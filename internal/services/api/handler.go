package api

import (
	"crypto/sha256"
	"delong/internal/models"
	"delong/internal/types"
	"delong/pkg/tee"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
)

func (s *ApiService) UploadReport(c *gin.Context) {
	req := types.UploadReportReq{}
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
	originalHash := sha256.Sum256(reportFile.Data)
	originalHashStr := hex.EncodeToString(originalHash[:])

	// Check if this file has already been uploaded (by any user)
	var existingReport models.TestReport
	if err := s.MysqlDb.Where("file_hash = ?", originalHashStr).First(&existingReport).Error; err == nil {
		// File already exists in the system
		log.Printf("File with hash %s already exists in the system", originalHashStr)
		ResponseError(c, REPORT_ALREADY_EXIST)
		return
	}

	// Upload encrypted raw report file
	aesKey, err := s.KeyVault.DeriveSymmetricKey(c, tee.KeyCtxUploadReportEncrypt, 32)
	if err != nil {
		log.Printf("Failed to derive symmetric key: %v", err)
		ResponseError(c, KEY_DERIVE_FAIL)
		return
	}
	cid, err := s.IpfsStore.UploadEncrypted(c, reportFile.Data, aesKey)
	if err != nil {
		log.Printf("Failed to upload data: %v", err)
		ResponseError(c, IPFS_UPLOAD_FAIL)
		return
	}

	// Parse raw report file to structured data
	objName := fmt.Sprintf("/v1/1/original/%s", reportFile.Filename)
	err = s.MinioStore.UploadBytes(c, "diagnostic", objName, reportFile.Data, reportFile.ContentType)
	if err != nil {
		log.Printf("Failed to upload file to minio: %v", err)
		ResponseError(c, MINIO_UPLOAD_FAIL)
		return
	}
	result, err := s.ReportAnalyzer.Analyze(c, "minio", reportFile.ContentType, objName, userWallet.Hex())
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
	dbtx := s.MysqlDb.Begin()
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

	tx, err := s.CtrCaller.RegisterData(c, userWallet, cid, req.Dataset)
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

func (s *ApiService) SubmitAlgo(c *gin.Context) {
	ctx := c.Request.Context()
	req := types.SubmitAlgoReq{}
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

	cid, err := s.IpfsStore.UploadStream(ctx, resp.Body)
	if err != nil {
		log.Printf("Failed to upload algorithm to IPFS: %v", err)
		ResponseError(c, IPFS_UPLOAD_FAIL)
		return
	}

	dbtx := s.MysqlDb.Begin()
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

	tx, err := s.CtrCaller.SubmitAlgorithm(ctx, common.HexToAddress(req.ScientistWallet), cid, req.Dataset)
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

func (s *ApiService) Vote(c *gin.Context) {
	req := types.VoteReq{}
	err := c.ShouldBind(&req)
	if err != nil {
		log.Printf("Failed to bind vote request: %v", err)
		ResponseError(c, BAD_REQUEST)
		return
	}
	_, err = models.CreateTransaction(s.MysqlDb, req.TxHash, 0) // Using 0 as a placeholder; will be updated later by the off-chain listener
	if err != nil {
		log.Printf("Failed to create blockchain transaction record: %v", err)
		ResponseData(c, MYSQL_WRITE_FAIL)
		return
	}

	ResponseData(c, req.TxHash)
}

func (s *ApiService) SetCommitteeMember(c *gin.Context) {
	req := types.SetCommitteeMemberReq{}
	err := c.ShouldBind(&req)
	if err != nil {
		log.Printf("Failed to bind set committee member request: %v", err)
		ResponseError(c, BAD_REQUEST)
		return
	}
	memberWallet := common.HexToAddress(req.MemberWallet)

	dbtx := s.MysqlDb.Begin()
	defer func() {
		if r := recover(); r != nil {
			dbtx.Rollback()
			panic(r)
		}
	}()

	cm, err := models.CreateOrUpdateCommitteeMember(dbtx, req.MemberWallet, req.IsApproved)
	if err != nil {
		dbtx.Rollback()
		log.Printf("Failed to create algorithm record: %v", err)
		ResponseError(c, MYSQL_WRITE_FAIL)
		return
	}

	tx, err := s.CtrCaller.SetCommitteeMember(c.Request.Context(), memberWallet, req.IsApproved)
	if err != nil {
		dbtx.Rollback()
		log.Printf("Failed to set committee member: %v", err)
		ResponseError(c, ETHEREUM_CALL_FAIL)
		return
	}

	_, err = models.CreateTransaction(dbtx, tx.Hash().Hex(), uint(cm.ID))
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

	ResponseData(c, tx)
}

// readFile reads a file from the request form data.
func (s *ApiService) readFile(c *gin.Context, fieldName string) (*types.ReportFile, error) {
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

	return &types.ReportFile{
		Data:        filedata,
		Filename:    fh.Filename,
		ContentType: contentType,
	}, nil
}
