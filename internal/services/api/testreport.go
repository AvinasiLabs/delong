package api

import (
	"crypto/sha256"
	"delong/internal/models"
	"delong/internal/types"
	"delong/pkg/bizcode"
	"delong/pkg/responser"
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

type TestReportResource struct {
	ApiServiceOptions
}

func (r *TestReportResource) CreateHandler(c *gin.Context) {
	req := types.UploadReportReq{}
	if err := c.ShouldBind(&req); err != nil {
		log.Printf("Failed to bind request for uploading report: %v", err)
		responser.ResponseError(c, bizcode.BAD_REQUEST)
		return
	}
	userWallet := common.HexToAddress(req.UserWallet)

	reportFile, err := readFile(c, "file")
	if err != nil {
		log.Printf("Failed to read file: %v", err)
		responser.ResponseError(c, bizcode.FILE_READ_FAIL)
		return
	}

	// Calculate hash of original file for deduplication
	originalHash := sha256.Sum256(reportFile.Data)
	originalHashStr := hex.EncodeToString(originalHash[:])

	// Check if this file has already been uploaded (by any user)
	var existingReport models.TestReport
	if err := r.MysqlDb.Where("file_hash = ?", originalHashStr).First(&existingReport).Error; err == nil {
		// File already exists in the system
		log.Printf("File with hash %s already exists in the system", originalHashStr)
		responser.ResponseError(c, bizcode.REPORT_ALREADY_EXIST)
		return
	}

	// Upload encrypted raw report file
	aesKey, err := r.KeyVault.DeriveSymmetricKey(c, tee.KeyCtxUploadReportEncrypt, 32)
	if err != nil {
		log.Printf("Failed to derive symmetric key: %v", err)
		responser.ResponseError(c, bizcode.KEY_DERIVE_FAIL)
		return
	}
	cid, err := r.IpfsStore.UploadEncrypted(c, reportFile.Data, aesKey)
	if err != nil {
		log.Printf("Failed to upload data: %v", err)
		responser.ResponseError(c, bizcode.IPFS_UPLOAD_FAIL)
		return
	}

	// Parse raw report file to structured data
	objName := fmt.Sprintf("/v1/1/original/%s", reportFile.Filename)
	err = r.MinioStore.UploadBytes(c, "diagnostic", objName, reportFile.Data, reportFile.ContentType)
	if err != nil {
		log.Printf("Failed to upload file to minio: %v", err)
		responser.ResponseError(c, bizcode.MINIO_UPLOAD_FAIL)
		return
	}
	result, err := r.ReportAnalyzer.Analyze(c, "minio", reportFile.ContentType, objName, userWallet.Hex())
	if err != nil {
		log.Printf("Failed to analyze raw report test: %v", err)
		responser.ResponseError(c, bizcode.REPORT_ANALYZE_FAIL)
		return
	}
	var raw types.RawReport
	err = json.Unmarshal(result, &raw)
	if err != nil {
		log.Printf("Failed to deserialize raw report: %v", err)
		responser.ResponseError(c, bizcode.REPORT_DESERIALIZE_FAIL)
		return
	}

	testReport := raw.ConvertToModel(userWallet.Hex(), originalHashStr, cid, req.Dataset, req.TestTime)
	dbtx := r.MysqlDb.Begin()
	defer func() {
		if r := recover(); r != nil {
			dbtx.Rollback()
			panic(r)
		}
	}()

	if err = dbtx.Create(&testReport).Error; err != nil {
		dbtx.Rollback()
		log.Printf("Failed to write report to MySQL: %v", err)
		responser.ResponseError(c, bizcode.MYSQL_WRITE_FAIL)
		return
	}

	tx, err := r.CtrCaller.RegisterData(c, userWallet, cid, req.Dataset)
	if err != nil {
		dbtx.Rollback()
		log.Printf("Failed to register data on-chain: %v", err)
		responser.ResponseError(c, bizcode.ETHEREUM_CALL_FAIL)
		return
	}
	txHash := tx.Hash().Hex()

	// Create blockchain transaction record
	_, err = models.CreateTransaction(dbtx, txHash, testReport.ID, models.ENTITY_TYPE_TEST_REPORT)
	if err != nil {
		dbtx.Rollback()
		log.Printf("Failed to create blockchain transaction record: %v", err)
		responser.ResponseError(c, bizcode.MYSQL_WRITE_FAIL)
		return
	}

	if err := dbtx.Commit().Error; err != nil {
		log.Printf("Failed to commit db transaction: %v", err)
		responser.ResponseError(c, bizcode.MYSQL_WRITE_FAIL)
		return
	}

	responser.ResponseData(c, txHash)
}

// func (r *TestReportResource) DeleteHandler(c *gin.Context) {

// }

// readFile reads a file from the request form data.
func readFile(c *gin.Context, fieldName string) (*types.ReportFile, error) {
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
