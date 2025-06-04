package api

import (
	"bytes"
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
	"mime/multipart"
	"net/http"

	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
)

type TestReportResource struct {
	ApiServiceOptions
}

// analyzeReportFile uploads file content to the analyzer endpoint and returns raw JSON bytes
func (r *TestReportResource) analyzeReportFile(c *gin.Context, fileName string, fileReader io.Reader, addr string) ([]byte, error) {
	var buf bytes.Buffer
	writer := multipart.NewWriter(&buf)

	part, err := writer.CreateFormFile("file", fileName)
	if err != nil {
		return nil, fmt.Errorf("creating form file: %w", err)
	}

	_, err = io.Copy(part, fileReader)
	if err != nil {
		return nil, fmt.Errorf("copying file content: %w", err)
	}

	err = writer.Close()
	if err != nil {
		return nil, fmt.Errorf("closing multipart writer: %w", err)
	}

	url := fmt.Sprintf("%s/api/analyzer/diagno_analyzer/analyze", addr)
	req, err := http.NewRequestWithContext(c.Request.Context(), http.MethodPost, url, &buf)
	if err != nil {
		return nil, fmt.Errorf("creating request: %w", err)
	}

	req.Header.Set("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("executing request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		b, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("status %d: %s", resp.StatusCode, string(b))
	}

	return io.ReadAll(resp.Body)
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
	ctx := c.Request.Context()
	kc := tee.NewKeyContext(tee.KEYKIND_ENC_KEY, req.UserWallet, "encrypt dynamic dataset")
	aesKey, err := r.KeyVault.DeriveSymmetricKey(ctx, kc)
	if err != nil {
		log.Printf("Failed to derive symmetric key: %v", err)
		responser.ResponseError(c, bizcode.KEY_DERIVE_FAIL)
		return
	}
	cid, err := r.IpfsStore.UploadEncrypted(ctx, reportFile.Data, aesKey)
	if err != nil {
		log.Printf("Failed to upload data: %v", err)
		responser.ResponseError(c, bizcode.IPFS_UPLOAD_FAIL)
		return
	}

	// Parse raw report file to structured data using direct file upload
	result, err := r.analyzeReportFile(c, reportFile.Filename, bytes.NewReader(reportFile.Data), r.DiagnosticSrvAddr)
	if err != nil {
		log.Printf("Failed to analyze raw report test: %v", err)
		responser.ResponseError(c, bizcode.REPORT_ANALYZE_FAIL)
		return
	}

	var anaResp types.AnalyzeResponse
	err = json.Unmarshal(result, &anaResp)
	if err != nil {
		log.Printf("Failed to deserialize wrapper response: %v", err)
		responser.ResponseError(c, bizcode.REPORT_DESERIALIZE_FAIL)
		return
	}
	if !anaResp.Success {
		log.Printf("Report analysis was not successful")
		responser.ResponseError(c, bizcode.REPORT_ANALYZE_FAIL)
		return
	}

	var raw types.RawReport
	err = json.Unmarshal(anaResp.Data.Report, &raw)
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

	tx, err := r.CtrCaller.RegisterData(ctx, userWallet, cid, req.Dataset)
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
