package types

import (
	"delong/internal/model"
	"time"
)

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
func (raw RawReport) ConvertToModel(userWallet, dataset string, testTime time.Time) model.TestReport {
	var results []model.TestResult

	for category, items := range raw {
		for _, item := range items {
			results = append(results, model.TestResult{
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

	return model.TestReport{
		UserWallet: userWallet,
		Dataset:    dataset,
		TestTime:   testTime,
		Results:    results,
	}
}
