package runtime

import (
	"archive/tar"
	"compress/gzip"
	"context"
	"delong/internal/models"
	"delong/pkg/contracts"
	"delong/pkg/db"
	"delong/pkg/schedule"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/docker/docker/api/types/mount"
	"github.com/ethereum/go-ethereum/common"
	"gorm.io/gorm"
)

type RuntimeService struct {
	name string
	RuntimeServiceOptions
}

type RuntimeServiceOptions struct {
	Db            *gorm.DB
	Loader        *DatasetLoader
	IpfsStore     *db.IpfsStore
	CtrCaller     *contracts.ContractCaller
	AlgoScheduler *schedule.AlgoScheduler
}

func NewService(opts RuntimeServiceOptions) *RuntimeService {
	return &RuntimeService{
		name:                  "runtime-service",
		RuntimeServiceOptions: opts,
	}
}

func (s *RuntimeService) Name() string {
	return s.name
}

func (s *RuntimeService) Init(ctx context.Context) error {
	var err error

	// Set algo scheduler handler
	s.AlgoScheduler.SetHandler(s)

	// Ensure dataset volume path exists
	if err = s.Loader.MustInit(); err != nil {
		return err
	}

	// Recover any pending algorithm executions
	err = s.recoverPendingExecutions(ctx)
	if err != nil {
		return err
	}

	// Start dataset management
	go s.runDatasetLifecycle(ctx)

	return nil
}

func (s *RuntimeService) Start(ctx context.Context) error {
	log.Println("Starting runtime service...")

	// Start listening for algorithm approval events
	// err := s.runEventLoop(ctx)
	// if err != nil {
	// 	return err
	// }

	go s.AlgoScheduler.Run(ctx)
	return nil
}

func (s *RuntimeService) Stop(ctx context.Context) error {
	log.Println("Stopping runtime service")
	return nil
}

func (s *RuntimeService) OnError(executionID uint, err error) {
	log.Printf("Failed to run algo, executionId=%d, err=%v", executionID, err.Error())
	_, err = models.UpdateExecutionStatus(s.Db, executionID, models.EXE_STATUS_FAILED)
	if err != nil {
		log.Printf("Failed to update exe status: %v", err)
	}
}

func (s *RuntimeService) OnCompleted(ctx context.Context, executionID uint, success bool, output []byte, errmsg []byte) {
	var err error
	var algoExe *models.AlgoExe
	if success {
		log.Printf("Execution %d completed successfully, algorithm ran without errors", executionID)
		algoExe, err = models.UpdateExecutionCompleted(s.Db, executionID, string(output))
	} else {
		log.Printf("Execution %d completed, but algorithm reported failure", executionID)
		algoExe, err = models.UpdateExecutionCompleted(s.Db, executionID, string(output), string(errmsg))
	}
	if err != nil {
		log.Printf("Failed to update execution result: %v", err)
		return
	}

	algo, err := models.GetAlgoByID(s.Db, algoExe.AlgoID)
	if err != nil {
		log.Printf("Failed to get algo by id: %v", err)
		return
	}
	usedAtUnix := time.Now().Unix()
	dbtx := s.Db.Begin()
	usage, err := models.CreateDataUsage(dbtx, algoExe.ScientistWallet, algo.Cid, algoExe.UsedDataset, usedAtUnix)
	if err != nil {
		dbtx.Rollback()
		log.Printf("Failed to create data usage record: %v", err)
		return
	}

	tx, err := s.CtrCaller.RecordUsage(ctx, common.HexToAddress(algoExe.ScientistWallet), algo.Cid, algoExe.UsedDataset, usedAtUnix)
	if err != nil {
		dbtx.Rollback()
		log.Printf("Failed to call RecordUsage: %v", err)
		return
	}

	_, err = models.CreateTransaction(dbtx, tx.Hash().Hex(), usage.ID, models.ENTITY_TYPE_DATAUSAGE)
	if err != nil {
		dbtx.Rollback()
		log.Printf("Failed to create block transaction: %v", err)
		return
	}

	if err := dbtx.Commit().Error; err != nil {
		log.Printf("Failed to commit transaction: %v", err)
	}
}

func (s *RuntimeService) OnResolve(ctx context.Context, exeId uint, algoCid string, resolveAt time.Time) {
	delay := time.Until(resolveAt) + 60 // extra 60 seconds
	if delay <= 0 {
		delay = time.Second
	}
	select {
	case <-time.After(delay):
		log.Printf("Should resolve at %v, actually resolve at %v", resolveAt, time.Now().Unix())
		tx, err := s.CtrCaller.Resolve(ctx, algoCid, exeId)
		if err != nil {
			log.Printf("Resolve failed for algo %v: %v", algoCid, err)
		} else {
			txHash := tx.Hash().Hex()
			log.Printf("Resolved algo %v, tx=%v", algoCid, txHash)
		}
	case <-ctx.Done():
		log.Printf("Resolve cancelled for algo %v", algoCid)
	}
}

func (s *RuntimeService) OnRun(ctx context.Context, exeId uint) {
	var (
		success bool
		output  []byte
		errmsg  []byte
		err     error
	)
	defer func() {
		if err != nil {
			s.OnError(exeId, err)
		} else {
			s.OnCompleted(ctx, exeId, success, output, errmsg)
		}
	}()

	// TODO: minitor hardware resource
	execution, err := models.UpdateExecutionStatus(s.Db, exeId, models.EXE_STATUS_RUNNING)
	if err != nil {
		return
	}

	// Get algorithm cid and download algorithm from IPFS and create temporary working directory
	algo, err := models.GetAlgoByID(s.Db, execution.AlgoID)
	if err != nil {
		return
	}
	workDir, err := s.fetchAndExtractWorkDir(ctx, algo.Cid)
	if err != nil {
		return
	}
	defer os.RemoveAll(workDir)

	// Verify Dockerfile exists
	if _, err := os.Stat(filepath.Join(workDir, "Dockerfile")); os.IsNotExist(err) {
		return
	}

	// Build Docker image
	imageName := fmt.Sprintf("delong-algorithm-%d", algo.ID)
	log.Printf("Building image %s", imageName)
	err = s.AlgoScheduler.BuildImage(ctx, workDir, imageName)
	if err != nil {
		// log.Printf("Failed to build image: %v", err)
		// models.UpdateExecutionStatus(s.Db, execution.ID, models.EXE_STATUS_FAILED, "Image build failed")
		return
	}

	// Acquire current version of dataset used by algorithm
	path, version, err := s.Loader.AcquireCurrent(execution.UsedDataset)
	if err != nil {
		// log.Printf("Failed to acquire current dataset:%v", err)
		// models.UpdateExecutionStatus(s.Db, execution.ID, models.EXE_STATUS_FAILED, "Failed to acquire dataset")
		return
	}
	defer s.Loader.Release(execution.UsedDataset, version)

	// Build docker environment variables and mounts
	env := map[string]string{
		"DATASET_PATH": path,
	}
	mounts := []mount.Mount{
		{
			Type:     mount.TypeBind,
			Source:   path,
			Target:   "/data",
			ReadOnly: true,
		},
	}

	// Run algorithm container
	log.Printf("Running algorithm container for dataset %s", execution.UsedDataset)
	output, errmsg, success, err = s.AlgoScheduler.RunContainer(ctx, imageName, env, mounts)
	if err != nil {
		// log.Printf("Failed to run algorithm container: %v", err)
		// models.UpdateExecutionStatus(s.Db, execution.ID, models.EXE_STATUS_FAILED, fmt.Sprintf("Execution failed: %v", err))
		return
	}

	// Otherwise save results
	// resultStr := string(output)
	// log.Printf("Algorithm execution completed with output length: %d bytes", len(resultStr))
	// models.UpdateExecutionCompleted(s.Db, execution.ID, resultStr)
}

// recoverPendingExecutions retries any algorithms that were in progress when service stopped
func (s *RuntimeService) recoverPendingExecutions(ctx context.Context) error {
	algoExes, err := models.GetPendingAlgoExesConfirmed(s.Db)
	if err != nil {
		return err
	}

	for _, exe := range algoExes {
		log.Printf("Recovering execution %d", exe.ID)
		s.AlgoScheduler.ScheduleRun(exe.ID)
	}

	return nil
}

// fetchAndExtractWorkDir fetches and extracts a work directory from an IPFS CID
func (s *RuntimeService) fetchAndExtractWorkDir(ctx context.Context, cidStr string) (string, error) {
	// Download tar.gz from IPFS
	r, err := s.IpfsStore.DownloadStream(ctx, cidStr)
	if err != nil {
		return "", fmt.Errorf("failed to download %s: %w", cidStr, err)
	}

	// Open gzip reader
	gzr, err := gzip.NewReader(r)
	if err != nil {
		return "", fmt.Errorf("gzip open failed: %w", err)
	}
	defer gzr.Close()

	// Read tar archive
	tr := tar.NewReader(gzr)

	// Create a temporary directory
	workDir, err := os.MkdirTemp("", fmt.Sprintf("algo-%s-", cidStr))
	if err != nil {
		return "", fmt.Errorf("mkdir temp failed: %w", err)
	}
	log.Printf("Downloading algorithm %s to %s", cidStr, workDir)

	// Extract entries one by one
	count := 0
	for {
		hdr, err := tr.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			return "", fmt.Errorf("tar read failed: %w", err)
		}

		target := filepath.Join(workDir, hdr.Name)
		if hdr.FileInfo().IsDir() {
			if err := os.MkdirAll(target, 0755); err != nil {
				return "", err
			}
		} else {
			// Ensure parent directory exists
			if err := os.MkdirAll(filepath.Dir(target), 0755); err != nil {
				return "", err
			}
			f, err := os.OpenFile(target, os.O_CREATE|os.O_WRONLY, hdr.FileInfo().Mode())
			if err != nil {
				return "", err
			}
			if _, err := io.Copy(f, tr); err != nil {
				f.Close()
				return "", err
			}
			f.Close()
		}
		count++
		log.Printf("Extracting #%d: %s (%d bytes)", count, hdr.Name, hdr.Size)
	}

	log.Printf("Work dir: %v", workDir)
	// After extraction, locate the single root subdirectory
	entries, err := os.ReadDir(workDir)
	if err != nil {
		return "", fmt.Errorf("failed to read workDir %s: %w", workDir, err)
	}
	// Filter for directories only
	var dirs []os.DirEntry
	for _, e := range entries {
		if e.IsDir() {
			dirs = append(dirs, e)
		}
	}
	if len(dirs) != 1 {
		return "", fmt.Errorf("expected single root directory after extraction, found %d", len(dirs))
	}
	root := filepath.Join(workDir, dirs[0].Name())
	log.Printf("Root: %s", root)
	return root, nil
}

// runDatasetLifecycle runs the dataset management tasks
func (s *RuntimeService) runDatasetLifecycle(ctx context.Context) {
	if err := s.Loader.Export(); err != nil {
		log.Printf("Initial dataset export failed: %v", err)
	}

	updateTicker := time.NewTicker(15 * time.Minute) // Update datasets every 15 minutes
	cleanupTicker := time.NewTicker(5 * time.Minute) // Clean unused datasets versions every 5 minutes
	defer updateTicker.Stop()
	defer cleanupTicker.Stop()
	for {
		select {
		case <-ctx.Done():
			log.Println("Stopping dataset management tasks...")
			return
		case <-updateTicker.C:
			if err := s.Loader.Export(); err != nil {
				log.Printf("Failed to update datasets: %v", err)
			}
		case <-cleanupTicker.C:
			if err := s.Loader.Cleanup(); err != nil {
				log.Printf("Failed to clean unused dataset versions: %v", err)
			}
		}
	}
}
