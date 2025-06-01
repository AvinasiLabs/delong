package types

import (
	"delong/internal/models"
	"encoding/json"
	"time"
)

type ReportFile struct {
	Data        []byte
	Filename    string
	ContentType string
}

// Define temporary wrapper structure to handle the new response format
type AnalyzeResponse struct {
	Success bool `json:"success"`
	Data    struct {
		Report         json.RawMessage `json:"report"`
		ProcessingTime float64         `json:"processingTime"`
		FileInfo       any             `json:"fileInfo"`
	} `json:"data"`
}

type RawReport map[string][]TestResultItem

type TestResultItem struct {
	Name           string  `json:"name"`
	Definition     string  `json:"definition"`
	Result         string  `json:"result"`
	ReferenceRange string  `json:"reference_range"`
	Explanation    string  `json:"explanation"`
	Status         string  `json:"status"`
	Suggestions    *string `json:"suggestions"`
}

// ConvertToModel converts RawReport to a TestReport with nested TestResults.
func (raw RawReport) ConvertToModel(userWallet, fileHash, cid, dataset string, testTime time.Time) models.TestReport {
	var results []models.TestResult

	for category, items := range raw {
		for _, item := range items {
			results = append(results, models.TestResult{
				Category:       category,
				Name:           item.Name,
				Definition:     item.Definition,
				Result:         item.Result,
				ReferenceRange: item.ReferenceRange,
				Explanation:    item.Explanation,
				Status:         item.Status,
				Suggestions:    item.Suggestions,
			})
		}
	}

	return models.TestReport{
		UserWallet:   userWallet,
		Dataset:      dataset,
		FileHash:     fileHash,
		RawReportCid: cid,
		TestTime:     testTime,
		Results:      results,
	}
}
