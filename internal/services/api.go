package services

import (
	"context"
	"crypto/sha256"
	"delong/internal/models"
	"delong/internal/types"
	"delong/pkg/analysis"
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

	ipfsStore      *db.IpfsStore
	minioStore     *db.MinioStore
	mysqlDb        *gorm.DB
	ctrCaller      *contracts.ContractCaller
	keyVault       *tee.KeyVault
	notifier       *ws.Notifier
	reportAnalyzer *analysis.ReportAnalyzer
}

type ApiServiceOptions struct {
	Addr           string
	IpfsStore      *db.IpfsStore
	MinioStore     *db.MinioStore
	MysqlDb        *gorm.DB
	CtrCaller      *contracts.ContractCaller
	KeyVault       *tee.KeyVault
	Notifier       *ws.Notifier
	ReportAnalyzer *analysis.ReportAnalyzer
}

func NewApiService(opts ApiServiceOptions) *ApiService {
	return &ApiService{
		name:           "api-service",
		addr:           opts.Addr,
		engine:         gin.Default(),
		ipfsStore:      opts.IpfsStore,
		minioStore:     opts.MinioStore,
		mysqlDb:        opts.MysqlDb,
		ctrCaller:      opts.CtrCaller,
		keyVault:       opts.KeyVault,
		notifier:       opts.Notifier,
		reportAnalyzer: opts.ReportAnalyzer,
	}
}

func (s *ApiService) Name() string {
	return s.name
}

func (s *ApiService) Init(ctx context.Context) error {
	// register routes
	s.engine.GET("/ws", ws.NewHandler(s.notifier.Hub()))
	apiGroup := s.engine.Group("/api")
	apiGroup.POST("/report/upload", s.UploadReport)
	apiGroup.GET("/report/:id", s.GetReports)
	apiGroup.POST("/algo/submit", s.SubmitAlgo)
	apiGroup.POST("/algo/vote", s.Vote)
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

type ReportFile struct {
	data        []byte
	filename    string
	contentType string
}

type UploadReportReq struct {
	UserWallet string    `form:"userWallet" binding:"required,ethwallet"` // hex
	Dataset    string    `form:"dataset" binding:"required"`
	TestTime   time.Time `form:"testTime" binding:"required"`
}

func (s *ApiService) UploadReport(c *gin.Context) {
	req := UploadReportReq{}
	if err := c.ShouldBind(&req); err != nil {
		log.Printf("Failed to bind request for uploading report: %v", err)
		ResponseError(c, BAD_REQUEST)
		return
	}
	userWallet := common.HexToAddress(req.UserWallet)

	reportFile, err := s.readFile(c, "file")
	if err != nil {
		log.Printf("Failed to read file: %v", err)
		ResponseError(c, FILE_READ_FAIL)
		return
	}

	// Calculate hash of original file for deduplication
	originalHash := sha256.Sum256(reportFile.data)
	originalHashStr := hex.EncodeToString(originalHash[:])

	// Check if this file has already been uploaded (by any user)
	var existingReport models.TestReport
	if err := s.mysqlDb.Where("file_hash = ?", originalHashStr).First(&existingReport).Error; err == nil {
		// File already exists in the system
		log.Printf("File with hash %s already exists in the system", originalHashStr)
		ResponseError(c, REPORT_ALREADY_EXIST)
		return
	}

	// Upload encrypted raw report file
	aesKey, err := s.keyVault.DeriveSymmetricKey(c, tee.KeyCtxUploadReportEncrypt, 32)
	if err != nil {
		log.Printf("Failed to derive symmetric key: %v", err)
		ResponseError(c, KEY_DERIVE_FAIL)
		return
	}
	cid, err := s.ipfsStore.UploadEncrypted(c, reportFile.data, aesKey)
	if err != nil {
		log.Printf("Failed to upload data: %v", err)
		ResponseError(c, IPFS_UPLOAD_FAIL)
		return
	}

	// Parse raw report file to structured data
	objName := fmt.Sprintf("/v1/1/original/%s", reportFile.filename)
	err = s.minioStore.UploadBytes(c, "diagnostic", objName, reportFile.data, reportFile.contentType)
	if err != nil {
		log.Printf("Failed to upload file to minio: %v", err)
		ResponseError(c, MINIO_UPLOAD_FAIL)
		return
	}
	result, err := s.reportAnalyzer.Analyze(c, "minio", reportFile.contentType, objName, userWallet.Hex())
	if err != nil {
		log.Printf("Failed to analyze raw report test: %v", err)
		ResponseError(c, REPORT_ANALYZE_FAIL)
		return
	}
	var raw types.RawReport
	err = json.Unmarshal(result, &raw)
	if err != nil {
		log.Printf("Failed to deserialize raw report: %v", err)
		ResponseError(c, REPORT_DESERIALIZE_FAIL)
		return
	}

	testReport := raw.ConvertToModel(userWallet.Hex(), originalHashStr, cid, req.Dataset, req.TestTime)
	dbtx := s.mysqlDb.Begin()
	defer func() {
		if r := recover(); r != nil {
			dbtx.Rollback()
			panic(r)
		}
	}()

	if err = dbtx.Create(&testReport).Error; err != nil {
		dbtx.Rollback()
		log.Printf("Failed to write report to MySQL: %v", err)
		ResponseError(c, MYSQL_WRITE_FAIL)
		return
	}

	tx, err := s.ctrCaller.RegisterData(c, userWallet, cid, req.Dataset)
	if err != nil {
		dbtx.Rollback()
		log.Printf("Failed to register data on-chain: %v", err)
		ResponseError(c, ETHEREUM_CALL_FAIL)
		return
	}
	txHash := tx.Hash().Hex()

	// Create blockchain transaction record
	_, err = models.CreateTransaction(dbtx, txHash, testReport.ID)
	if err != nil {
		dbtx.Rollback()
		log.Printf("Failed to create blockchain transaction record: %v", err)
		ResponseError(c, MYSQL_WRITE_FAIL)
		return
	}

	if err := dbtx.Commit().Error; err != nil {
		log.Printf("Failed to commit db transaction: %v", err)
		ResponseError(c, MYSQL_WRITE_FAIL)
		return
	}

	ResponseData(c, txHash)
}

func (s *ApiService) GetReports(c *gin.Context) {}

type SubmitAlgoReq struct {
	ScientistWallet string `json:"scientist_wallet" binding:"required,ethwallet"` // hex
	Dataset         string `json:"dataset" binding:"required"`
	AlgoLink        string `json:"algo_link" binding:"required"`
}

func (s *ApiService) SubmitAlgo(c *gin.Context) {
	ctx := c.Request.Context()
	req := SubmitAlgoReq{}
	err := c.ShouldBind(&req)
	if err != nil {
		log.Printf("Failed to bind request for submitting algorithm: %v", err)
		ResponseError(c, BAD_REQUEST)
		return
	}

	resp, err := http.Get(req.AlgoLink)
	if err != nil {
		log.Printf("Failed to open algorithm link: %v", err)
		ResponseError(c, AlGO_LINK_INVALID)
		return
	}
	defer resp.Body.Close()

	cid, err := s.ipfsStore.UploadStream(ctx, resp.Body)
	if err != nil {
		log.Printf("Failed to upload algorithm to IPFS: %v", err)
		ResponseError(c, IPFS_UPLOAD_FAIL)
		return
	}

	dbtx := s.mysqlDb.Begin()
	defer func() {
		if r := recover(); r != nil {
			dbtx.Rollback()
			panic(r)
		}
	}()

	algo, err := models.CreateAlgo(dbtx, "", req.AlgoLink, req.ScientistWallet, cid, req.Dataset)
	if err != nil {
		dbtx.Rollback()
		log.Printf("Failed to create algorithm record: %v", err)
		ResponseError(c, MYSQL_WRITE_FAIL)
		return
	}

	tx, err := s.ctrCaller.SubmitAlgorithm(ctx, common.HexToAddress(req.ScientistWallet), cid, req.Dataset)
	if err != nil {
		dbtx.Rollback()
		log.Printf("Failed to submit algorithm to blockchain: %v", err)
		ResponseError(c, ETHEREUM_CALL_FAIL)
		return
	}

	txHash := tx.Hash().Hex()
	// Create blockchain transaction record
	_, err = models.CreateTransaction(dbtx, txHash, algo.ID)
	if err != nil {
		dbtx.Rollback()
		log.Printf("Failed to create blockchain transaction record: %v", err)
		ResponseError(c, MYSQL_WRITE_FAIL)
		return
	}

	if err := dbtx.Commit().Error; err != nil {
		log.Printf("Failed to commit db transaction: %v", err)
		ResponseError(c, MYSQL_WRITE_FAIL)
		return
	}

	ResponseData(c, txHash)
}

type VoteReq struct {
	AlgoId uint   `json:"algo_id" binding:"required"`
	TxHash string `json:"tx_hash" binding:"required"`
}

func (s *ApiService) Vote(c *gin.Context) {
	req := VoteReq{}
	err := c.ShouldBind(&req)
	if err != nil {
		log.Printf("Failed to bind vote request: %v", err)
		ResponseError(c, BAD_REQUEST)
		return
	}
	_, err = models.CreateTransaction(s.mysqlDb, req.TxHash, 0) // Using 0 as a placeholder; will be updated later by the off-chain listener
	if err != nil {
		log.Printf("Failed to create blockchain transaction record: %v", err)
		ResponseData(c, MYSQL_WRITE_FAIL)
		return
	}

	ResponseData(c, req.TxHash)
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
