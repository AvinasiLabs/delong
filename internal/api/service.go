package api

import (
	"context"
	"crypto/sha256"
	"delong/internal/control"
	"delong/internal/model"
	"delong/internal/types"
	"delong/pkg/analyzer"
	"delong/pkg/contracts"
	"delong/pkg/db"
	"delong/pkg/tee"
	"delong/pkg/ws"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ApiService struct {
	name       string
	addr       string
	ctrAddr    map[string]common.Address
	engine     *gin.Engine
	httpserver *http.Server

	ipfsStore  *db.IpfsStore
	minioStore *db.MinioStore
	mysqlDb    *gorm.DB

	ethCaller *contracts.ContractCaller
	keyVault  *tee.KeyVault

	notifier *ws.Notifier

	reportAnalyzer *analyzer.ReportAnalyzer
}

type ApiServiceOptions struct {
	Addr         string
	Storage      *control.StorageDeps
	Blockchain   *control.BlockchainDeps
	Notification *control.NotificationDeps
	Analyzer     *control.AnalyzerDeps
}

const SERVICE_NAME = "api-service"

func NewApiService(opts ApiServiceOptions) *ApiService {
	return &ApiService{
		name:           SERVICE_NAME,
		addr:           opts.Addr,
		engine:         gin.Default(),
		ipfsStore:      opts.Storage.IpfsStore,
		minioStore:     opts.Storage.MinioStore,
		mysqlDb:        opts.Storage.MysqlDb,
		ethCaller:      opts.Blockchain.EthCaller,
		keyVault:       opts.Blockchain.KeyVault,
		notifier:       opts.Notification.Notifier,
		reportAnalyzer: opts.Analyzer.ReportAnalyzer,
	}
}

func (s *ApiService) Name() string {
	return s.name
}

func (s *ApiService) Init(ctx context.Context) error {
	s.registerRoutes()
	return nil
}

func (s *ApiService) Start(ctx context.Context) error {
	s.httpserver = &http.Server{
		Addr:    s.addr,
		Handler: s.engine,
	}

	go func() {
		err := s.httpserver.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.Printf("Failed to listen: %v", err)
		}
	}()

	log.Println("Api service started")
	<-ctx.Done()
	log.Println("API service context cancelled, will shut down")
	return nil
}

func (s *ApiService) Stop(ctx context.Context) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	err := s.httpserver.Shutdown(ctx)
	if err != nil {
		log.Printf("Failed to shutdown gracefully: %v", err)
		return err
	}

	log.Println("Http server shutdown cleanly")
	return nil
}

func (s *ApiService) registerRoutes() {
	s.engine.GET("/ws", ws.NewHandler(s.notifier.Hub()))

	apiGroup := s.engine.Group("/api")

	apiGroup.POST("/report/upload", s.UploadReport)
	apiGroup.GET("/report/:id", s.GetReports)
}

type UploadReportReq struct {
	UserWallet string    `form:"userWallet" binding:"required,ethwallet"`
	Dataset    string    `form:"dataset" binding:"required"`
	TestTime   time.Time `form:"testTime" binding:"required"`
}

func (s *ApiService) UploadReport(c *gin.Context) {
	req := UploadReportReq{}
	if err := c.ShouldBind(&req); err != nil {
		log.Printf("Failed to bind request: %v", err)
		c.JSON(400, gin.H{"error": "Invalid request"})
		return
	}
	userWallet := common.HexToAddress(req.UserWallet)

	reportFile, err := s.readFile(c, "file")
	if err != nil {
		log.Printf("Failed to read file: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Reading file failed"})
		return
	}

	// Calculate hash of original file for deduplication
	originalHash := sha256.Sum256(reportFile.data)
	originalHashStr := hex.EncodeToString(originalHash[:])

	// Check if this file has already been uploaded (by any user)
	var existingReport model.TestReport
	if err := s.mysqlDb.Where("file_hash = ?", originalHashStr).First(&existingReport).Error; err == nil {
		// File already exists in the system
		log.Printf("File with hash %s already exists in the system", originalHashStr)
		c.JSON(http.StatusConflict, gin.H{"error": "this file has already been uploaded to the system", "existing_cid": existingReport.RawReportCid})
		return
	}

	// Upload encrypted raw report file
	aesKey, err := s.keyVault.DeriveSymmetricKey(c, tee.KeyCtxUploadReportEncrypt, 32)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Deriving symmetric key failed"})
		return
	}
	cid, err := s.ipfsStore.UploadEncrypted(c, reportFile.data, aesKey)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Uploading data failed"})
		return
	}

	// Parse raw report file to structured data
	objName := fmt.Sprintf("/v1/1/original/%s", reportFile.filename)
	err = s.minioStore.UploadBytes(c, "diagnostic", objName, reportFile.data, reportFile.contentType)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Uploading file to MinIO failed"})
		return
	}
	result, err := s.reportAnalyzer.Analyze(c, "minio", reportFile.contentType, objName, userWallet.Hex())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Analyzing report failed"})
		return
	}
	var raw types.RawReport
	err = json.Unmarshal(result, &raw)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Analyzing report failed"})
		return
	}

	// Create report record
	testReport := raw.ConvertToModel(userWallet.Hex(), originalHashStr, cid, req.Dataset, req.TestTime)
	err = s.mysqlDb.Create(&testReport).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Writing report to MySQL failed"})
		return
	}

	// Call contract to submit on-chain transaction
	tx, err := s.ethCaller.RegisterData(c, userWallet, cid, req.Dataset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Registering data failed"})
		return
	}
	txHash := tx.Hash().Hex()
	log.Printf("Transaction hash: %s", txHash)

	// Create blockchain transaction record
	_, err = model.CreateTransaction(s.mysqlDb, txHash, uint(testReport.ID))
	if err != nil {
		log.Printf("Failed to create blockchain transaction record: %v", err)
		// Don't interrupt the operation, just log the error
	}

	c.JSON(http.StatusOK, gin.H{"msg": "ok", "data": txHash})
}

type ReportFile struct {
	data        []byte
	filename    string
	contentType string
}

// readFile reads a file from the request form data.
func (s *ApiService) readFile(c *gin.Context, fieldName string) (*ReportFile, error) {
	fh, err := c.FormFile(fieldName)
	if err != nil {
		return nil, err
	}

	f, err := fh.Open()
	if err != nil {
		return nil, err
	}
	defer f.Close()

	// detect content type
	head := make([]byte, 512)
	n, _ := f.Read(head)
	contentType := http.DetectContentType(head[:n])

	// restart seek point
	_, err = f.Seek(0, io.SeekStart)
	if err != nil {
		return nil, err
	}

	filedata, err := io.ReadAll(f)
	if err != nil {
		return nil, err
	}

	return &ReportFile{
		data:        filedata,
		filename:    fh.Filename,
		contentType: contentType,
	}, nil
}

func (s *ApiService) GetReports(c *gin.Context) {}
