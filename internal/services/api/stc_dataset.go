package api

import (
	"bytes"
	"context"
	"delong/internal/consts"
	"delong/internal/models"
	"delong/internal/types"
	"delong/pkg/bizcode"
	"delong/pkg/responser"
	"delong/pkg/tee"
	"encoding/csv"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

const (
	defaultSampleSize = 5
)

type StaticDatasetResource struct{ ApiServiceOptions }

type FakeDataRequest struct {
	Data     []any `json:"data"`
	NSamples int   `json:"n_samples"`
}

type FakeDataResponse struct {
	Status string `json:"status"`
	Result []any  `json:"result"`
}

// ColumnType represents the detected type of a CSV column
type ColumnType int

const (
	ColumnTypeString ColumnType = iota
	ColumnTypeNumber
	ColumnTypeBoolean
)

// ColumnInfo holds information about a CSV column
type ColumnInfo struct {
	Type ColumnType
	Data []any
}

// generateSampleDataForColumn calls the fake data API to generate sample data for a column
func (r *StaticDatasetResource) generateSampleDataForColumn(ctx context.Context, columnData []any, nSamples int) ([]any, error) {
	// Limit data to 5000 samples for API call
	maxApiSamples := 5000
	if len(columnData) > maxApiSamples {
		columnData = columnData[:maxApiSamples]
	}

	reqData := FakeDataRequest{
		Data:     columnData,
		NSamples: nSamples,
	}

	jsonData, err := json.Marshal(reqData)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "POST", r.SampleSrvAddr+"/api/delong/v1/data_pipe/fake_data", bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to call fake data API: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("fake data API returned status %d", resp.StatusCode)
	}

	var response FakeDataResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	if response.Status != "success" {
		return nil, fmt.Errorf("fake data API returned error status: %s", response.Status)
	}

	return response.Result, nil
}

// detectColumnType analyzes a column's data to determine its type
func detectColumnType(values []string) ColumnType {
	numericCount := 0
	booleanCount := 0
	total := len(values)

	if total == 0 {
		return ColumnTypeString
	}

	for _, val := range values {
		val = strings.TrimSpace(strings.ToLower(val))

		// Check for boolean
		if val == "true" || val == "false" || val == "1" || val == "0" || val == "yes" || val == "no" {
			booleanCount++
		} else if _, err := fmt.Sscanf(val, "%f", new(float64)); err == nil {
			// Check for numeric
			numericCount++
		}
	}

	// If more than 80% are numeric, treat as number
	if float64(numericCount)/float64(total) > 0.8 {
		return ColumnTypeNumber
	}

	// If more than 80% are boolean, treat as boolean
	if float64(booleanCount)/float64(total) > 0.8 {
		return ColumnTypeBoolean
	}

	return ColumnTypeString
}

// parseColumnData converts string values to appropriate types based on column type
func parseColumnData(values []string, colType ColumnType) []any {
	result := make([]any, len(values))

	for i, val := range values {
		val = strings.TrimSpace(val)

		switch colType {
		case ColumnTypeNumber:
			if num, err := fmt.Sscanf(val, "%f", new(float64)); err == nil && num > 0 {
				var f float64
				fmt.Sscanf(val, "%f", &f)
				result[i] = f
			} else {
				result[i] = 0.0
			}
		case ColumnTypeBoolean:
			val = strings.ToLower(val)
			result[i] = val == "true" || val == "1" || val == "yes"
		default: // ColumnTypeString
			result[i] = val
		}
	}

	return result
}

// formatValue converts any back to string for CSV output
func formatValue(val any) string {
	switch v := val.(type) {
	case float64:
		return fmt.Sprintf("%.2f", v)
	case int:
		return fmt.Sprintf("%d", v)
	case bool:
		if v {
			return "true"
		}
		return "false"
	case string:
		return v
	default:
		return fmt.Sprintf("%v", v)
	}
}

// generateSampleCSV generates a sample CSV file from the original CSV data
func (r *StaticDatasetResource) generateSampleCSV(ctx context.Context, originalFile io.Reader) (string, error) {
	// Parse the original CSV
	csvReader := csv.NewReader(originalFile)
	records, err := csvReader.ReadAll()
	if err != nil {
		return "", fmt.Errorf("failed to read CSV: %w", err)
	}

	if len(records) < 2 {
		return "", fmt.Errorf("CSV must have at least header and one data row")
	}

	headers := records[0]
	dataRows := records[1:]

	// Analyze each column to detect its type and prepare data
	columnInfos := make([]ColumnInfo, len(headers))
	for i := range headers {
		// Extract all values for this column
		columnValues := make([]string, len(dataRows))
		for rowIdx, row := range dataRows {
			if i < len(row) {
				columnValues[rowIdx] = row[i]
			} else {
				columnValues[rowIdx] = ""
			}
		}

		// Detect column type
		colType := detectColumnType(columnValues)

		// Parse data according to detected type
		columnInfos[i] = ColumnInfo{
			Type: colType,
			Data: parseColumnData(columnValues, colType),
		}
	}

	// Generate sample data for each column
	sampleSize := defaultSampleSize
	sampleData := make([][]any, len(headers))

	for i, colInfo := range columnInfos {
		if len(colInfo.Data) == 0 {
			// If no data, create default values based on type
			defaultData := make([]any, 3)
			switch colInfo.Type {
			case ColumnTypeNumber:
				defaultData = []any{1.0, 2.0, 3.0}
			case ColumnTypeBoolean:
				defaultData = []any{true, false, true}
			default: // ColumnTypeString
				defaultData = []any{"sample1", "sample2", "sample3"}
			}
			colInfo.Data = defaultData
		}

		samples, err := r.generateSampleDataForColumn(ctx, colInfo.Data, sampleSize)
		if err != nil {
			log.Printf("Failed to generate sample data for column %s: %v", headers[i], err)
			// Use some default values if API fails
			samples = make([]any, sampleSize)
			for j := range samples {
				if len(colInfo.Data) > 0 {
					samples[j] = colInfo.Data[j%len(colInfo.Data)]
				} else {
					switch colInfo.Type {
					case ColumnTypeNumber:
						samples[j] = float64(j + 1)
					case ColumnTypeBoolean:
						samples[j] = j%2 == 0
					default:
						samples[j] = fmt.Sprintf("sample%d", j+1)
					}
				}
			}
		}
		sampleData[i] = samples
	}

	// Create CSV content
	var csvBuffer bytes.Buffer
	csvWriter := csv.NewWriter(&csvBuffer)

	// Write headers
	if err := csvWriter.Write(headers); err != nil {
		return "", fmt.Errorf("failed to write CSV headers: %w", err)
	}

	// Write sample data rows
	for rowIdx := 0; rowIdx < sampleSize; rowIdx++ {
		row := make([]string, len(headers))
		for colIdx := 0; colIdx < len(headers); colIdx++ {
			if rowIdx < len(sampleData[colIdx]) {
				row[colIdx] = formatValue(sampleData[colIdx][rowIdx])
			} else {
				// Fallback value
				switch columnInfos[colIdx].Type {
				case ColumnTypeNumber:
					row[colIdx] = "0.00"
				case ColumnTypeBoolean:
					row[colIdx] = "false"
				default:
					row[colIdx] = "sample"
				}
			}
		}
		if err := csvWriter.Write(row); err != nil {
			return "", fmt.Errorf("failed to write CSV row: %w", err)
		}
	}

	csvWriter.Flush()
	if err := csvWriter.Error(); err != nil {
		return "", fmt.Errorf("failed to flush CSV writer: %w", err)
	}

	return csvBuffer.String(), nil
}

// addStaticPrefix adds the static dataset prefix to distinguish from dynamic datasets
func addStaticPrefix(name string) string {
	return consts.StaticDatasetPrefix + name
}

// removeStaticPrefix removes the static dataset prefix for display purposes
func removeStaticPrefix(name string) string {
	if strings.HasPrefix(name, consts.StaticDatasetPrefix) {
		return name[len(consts.StaticDatasetPrefix):]
	}
	return name
}

// isStaticDataset checks if a dataset name has the static prefix
func isStaticDataset(name string) bool {
	return strings.HasPrefix(name, consts.StaticDatasetPrefix)
}

func (r *StaticDatasetResource) CreateHandler(c *gin.Context) {
	req := types.StcDatasetCreateReq{}
	if err := c.ShouldBind(&req); err != nil {
		log.Printf("Failed to bind request: %v", err)
		responser.ResponseError(c, bizcode.BAD_REQUEST)
		return
	}

	file, header, err := c.Request.FormFile("file")
	if err != nil {
		log.Printf("Failed to get uploaded static dataset file: %v", err)
		responser.ResponseError(c, bizcode.BAD_REQUEST)
		return
	}
	defer file.Close()

	// Calculate hash of original file for deduplication
	hash, err := hashSha256(file)
	if err != nil {
		log.Printf("Failed to calculate file hash: %v", err)
		responser.ResponseError(c, bizcode.INTERNAL_SERVER_ERROR)
		return
	}
	_, err = models.GetStcDatasetByHash(r.MysqlDb, hash)
	if err == nil { // file exists
		log.Printf("File with hash %v already exists", hash)
		responser.ResponseError(c, bizcode.STATIC_DATASET_EXIST)
		return
	}

	// Upload to ipfs
	ctx := c.Request.Context()
	kc := tee.NewKeyContext(tee.KEYKIND_ENC_KEY, req.Author, consts.PurposeEncStaticDataset)
	key, err := r.KeyVault.DeriveSymmetricKey(ctx, kc)
	if err != nil {
		log.Printf("Failed to derive symmetric key: %v", err)
		responser.ResponseError(c, bizcode.KEY_DERIVE_FAIL)
		return
	}

	cid, err := r.IpfsStore.UploadEncryptedStream(ctx, file, key)
	if err != nil {
		log.Printf("Failed to upload to IPFS: %v", err)
		responser.ResponseError(c, bizcode.INTERNAL_SERVER_ERROR)
		return
	}

	// Get file format
	fileFormat := filepath.Ext(header.Filename)
	if fileFormat != "" {
		fileFormat = fileFormat[1:] // remove the dot
	}

	// Generate sample dataset and upload to IPFS
	var sampleUrl string
	if fileFormat == "csv" {
		// Reset file reader for sample generation
		if _, err := file.Seek(0, 0); err != nil {
			log.Printf("Failed to reset file reader: %v", err)
			responser.ResponseError(c, bizcode.INTERNAL_SERVER_ERROR)
			return
		}

		sampleCSV, err := r.generateSampleCSV(ctx, file)
		if err != nil {
			log.Printf("Failed to generate sample CSV: %v", err)
			// Don't fail the entire operation if sample generation fails
			sampleUrl = ""
		} else {
			// Upload sample CSV to IPFS without encryption
			sampleCid, err := r.IpfsStore.Upload(ctx, []byte(sampleCSV))
			if err != nil {
				log.Printf("Failed to upload sample CSV to IPFS: %v", err)
				sampleUrl = ""
			} else {
				// Construct API URL for accessing sample data
				sampleUrl = fmt.Sprintf("/api/sample/%s", sampleCid)
			}
		}

		// Reset file reader again for the main upload
		if _, err := file.Seek(0, 0); err != nil {
			log.Printf("Failed to reset file reader for main upload: %v", err)
			responser.ResponseError(c, bizcode.INTERNAL_SERVER_ERROR)
			return
		}
	}

	dbtx := r.MysqlDb.Begin()
	if dbtx.Error != nil {
		log.Printf("Failed to begin transaction: %v", dbtx.Error)
		responser.ResponseError(c, bizcode.MYSQL_WRITE_FAIL)
		return
	}

	// Create static dataset
	normName, err := normalizeFileName(req.Name)
	if err != nil {
		log.Printf("Failed to normalize name: %v", err)
		responser.ResponseError(c, bizcode.INTERNAL_SERVER_ERROR)
		return
	}
	stcDatasetName := fmt.Sprintf("%s%s", consts.StaticDatasetPrefix, normName)
	createReq := models.CreateStcDatasetReq{
		Name:         stcDatasetName, // Store with prefix
		UiName:       req.Name,
		Desc:         req.Desc,
		FileHash:     hash,
		IpfsCid:      cid,
		FileSize:     header.Size,
		FileFormat:   fileFormat,
		Author:       req.Author,
		AuthorWallet: req.AuthorWallet,
		SampleUrl:    sampleUrl,
		FilePath:     fmt.Sprintf("/data/%s.csv", stcDatasetName),
	}

	dataset, err := models.CreateStcDataset(dbtx, createReq)
	if err != nil {
		dbtx.Rollback()
		log.Printf("Failed to create data asset: %v", err)
		responser.ResponseError(c, bizcode.MYSQL_WRITE_FAIL)
		return
	}

	tx, err := r.CtrCaller.RegisterData(ctx, common.HexToAddress(dataset.AuthorWallet), dataset.IpfsCid, dataset.Name)
	if err != nil {
		dbtx.Rollback()
		log.Printf("Failed to register static dataset: %v", err)
		responser.ResponseError(c, bizcode.MYSQL_WRITE_FAIL)
		return
	}
	txHash := tx.Hash().Hex()
	_, err = models.CreateTransaction(dbtx, txHash, uint(dataset.ID), models.ENTITY_TYPE_STATIC_DATASET)
	if err != nil {
		dbtx.Rollback()
		log.Printf("Failed to create transaction: %v", err)
		responser.ResponseError(c, bizcode.MYSQL_WRITE_FAIL)
		return
	}

	if err := dbtx.Commit().Error; err != nil {
		log.Printf("Failed to commit transaction: %v", err)
		responser.ResponseError(c, bizcode.MYSQL_WRITE_FAIL)
		return
	}

	responser.ResponseData(c, txHash)
}

func (r *StaticDatasetResource) ListHandler(c *gin.Context) {
	page, pageSize := parsePageParams(c)
	assets, total, err := models.GetStcDataset(r.MysqlDb, page, pageSize)
	if err != nil {
		log.Printf("Failed to list data assets: %v", err)
		responser.ResponseError(c, bizcode.MYSQL_READ_FAIL)
		return
	}
	responser.ResponseList(c, page, pageSize, total, assets)
}

func (r *StaticDatasetResource) TakeHandler(c *gin.Context) {
	var id uint
	if err := parseUintParam(c.Param("id"), &id); err != nil {
		responser.ResponseError(c, bizcode.BAD_REQUEST)
		return
	}

	asset, err := models.GetStcDatasetByID(r.MysqlDb, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			responser.ResponseError(c, bizcode.NOT_FOUND)
			return
		}
		responser.ResponseError(c, bizcode.MYSQL_READ_FAIL)
		return
	}

	responser.ResponseData(c, asset)
}

func (r *StaticDatasetResource) SampleHandler(c *gin.Context) {
	cid := c.Param("cid")
	if cid == "" {
		responser.ResponseError(c, bizcode.BAD_REQUEST)
		return
	}

	ctx := c.Request.Context()
	data, err := r.IpfsStore.Download(ctx, cid)
	if err != nil {
		log.Printf("Failed to download sample data from IPFS: %v", err)
		responser.ResponseError(c, bizcode.NOT_FOUND)
		return
	}

	// Set headers for CSV file download
	c.Header("Content-Type", "text/csv")
	c.Header("Content-Disposition", fmt.Sprintf("attachment; filename=\"sample_%s.csv\"", cid[:8]))
	c.Data(200, "text/csv", data)
}
