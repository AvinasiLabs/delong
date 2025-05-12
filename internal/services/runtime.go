package services

import (
	"archive/tar"
	"bytes"
	"compress/gzip"
	"context"
	"delong/internal/model"
	"delong/pkg/contracts"
	"delong/pkg/db"
	"delong/pkg/scheduler"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"time"

	"gorm.io/gorm"
)

type RuntimeService struct {
	name          string
	db            *gorm.DB
	ipfsStore     *db.IpfsStore
	ctrCaller     *contracts.ContractCaller
	algoScheduler *scheduler.AlgoScheduler
}

type RuntimeServiceOptions struct {
	Db            *gorm.DB
	IpfsStore     *db.IpfsStore
	CtrCaller     *contracts.ContractCaller
	AlgoScheduler *scheduler.AlgoScheduler
}

func NewRuntimeService(opts RuntimeServiceOptions) *RuntimeService {
	return &RuntimeService{
		name:          "runtime-service",
		db:            opts.Db,
		ipfsStore:     opts.IpfsStore,
		ctrCaller:     opts.CtrCaller,
		algoScheduler: opts.AlgoScheduler,
	}
}

func (s *RuntimeService) Name() string {
	return s.name
}

func (s *RuntimeService) Init(ctx context.Context) error {
	// Recover any pending algorithm executions
	return s.recoverPendingExecutions(ctx)
}

func (s *RuntimeService) Start(ctx context.Context) error {
	log.Println("Starting runtime service...")

	// Start listening for algorithm approval events
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case algoID := <-s.algoScheduler.AlgoIdCh:
				log.Printf("Received approved algorithm ID: %d", algoID)
				go func(algoID uint) {
					execution, err := model.CreateAlgoExecution(s.db, algoID, model.EXESTATUS_QUEUED)
					if err != nil {
						log.Printf("Failed to create execution record: %v", err)
						return
					}

					// Execute algorithm in a new goroutine
					go func() {
						execCtx, cancel := context.WithTimeout(context.Background(), 1*time.Hour)
						defer cancel()

						err := s.executeAlgorithm(execCtx, execution.ID)
						if err != nil {
							log.Printf("Algorithm execution %d failed: %v", execution.ID, err)
						}
					}()
				}(algoID)
			}
		}
	}()

	return nil
}

func (s *RuntimeService) Stop(ctx context.Context) error {
	log.Println("Stopping runtime service")
	return nil
}

// executeAlgorithm handles the full algorithm execution process
func (s *RuntimeService) executeAlgorithm(ctx context.Context, executionID uint) error {
	// Update execution status to RUNNING
	execution, err := model.UpdateExecutionStatus(s.db, executionID, "RUNNING")
	if err != nil {
		return err
	}

	// Get algorithm details
	algo, err := model.GetAlgoByID(s.db, execution.AlgoID)
	if err != nil {
		model.UpdateExecutionStatus(s.db, executionID, "FAILED", "Algorithm not found")
		return err
	}

	// Download algorithm from IPFS and create temporary working directory
	workDir, err := s.fetchAndExtractWorkDir(ctx, algo.Cid)
	if err != nil {
		model.UpdateExecutionStatus(s.db, executionID, "FAILED", "Failed to prepare work dir")
		return err
	}
	defer os.RemoveAll(workDir)

	// Verify Dockerfile exists
	if _, err := os.Stat(filepath.Join(workDir, "Dockerfile")); os.IsNotExist(err) {
		model.UpdateExecutionStatus(s.db, executionID, "FAILED", "Dockerfile not found")
		return fmt.Errorf("dockerfile not found in algorithm package")
	}

	// Build Docker image
	imageName := fmt.Sprintf("delong-algorithm-%d", algo.ID)
	log.Printf("Building image %s", imageName)
	err = s.algoScheduler.BuildImage(ctx, workDir, imageName)
	if err != nil {
		model.UpdateExecutionStatus(s.db, executionID, "FAILED", "Image build failed")
		return err
	}

	// Run algorithm container
	log.Printf("Running algorithm container for dataset %s", algo.UsedDataset)
	output, err := s.algoScheduler.RunContainer(ctx, imageName, nil)
	if err != nil {
		model.UpdateExecutionStatus(s.db, executionID, "FAILED", fmt.Sprintf("Execution failed: %v", err))
		return err
	}

	// Save results
	resultStr := string(output)
	log.Printf("Algorithm execution completed with output length: %d bytes", len(resultStr))
	model.UpdateExecutionResult(s.db, executionID, resultStr)

	// Update execution status to COMPLETED
	_, err = model.UpdateExecutionStatus(s.db, executionID, "COMPLETED")

	return err
}

// recoverPendingExecutions retries any algorithms that were in progress when service stopped
func (s *RuntimeService) recoverPendingExecutions(ctx context.Context) error {
	// executions, err := model.GetPendingExecutions(s.db)
	// if err != nil {
	// 	return err
	// }

	// for _, exec := range executions {
	// 	log.Printf("Recovering execution %d for algorithm %d", exec.ID, exec.AlgoID)
	// 	s.scheduleAlgorithm(ctx, exec.AlgoID)
	// }

	return nil
}

// fetchAndExtractWorkDir
func (s *RuntimeService) fetchAndExtractWorkDir(ctx context.Context, cidStr string) (string, error) {
	// Download tar.gz from IPFS
	data, err := s.ipfsStore.Download(ctx, cidStr)
	if err != nil {
		return "", fmt.Errorf("failed to download %s: %w", cidStr, err)
	}

	// Open gzip reader
	gzr, err := gzip.NewReader(bytes.NewReader(data))
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
	}

	return workDir, nil
}
