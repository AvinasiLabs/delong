package models

import "time"

// System-defined dataset registry
type DatasetRegistry struct {
	Name        string
	Title       string
	Description string
	CreatedAt   time.Time
}

// One uploaded report from user
type TestReport struct {
	ID           uint
	UserWallet   string
	FileHash     string // Hash of original file for deduplication
	RawReportCid string // CID of encrypted original file
	Dataset      string
	TestTime     time.Time

	Results []TestResult
}

// One individual result item under a report
type TestResult struct {
	ID             uint
	TestReportID   int
	Category       string
	Name           string
	Definition     string
	Result         string
	ReferenceRange string
	Explanation    string
	Status         string
	Suggestions    *string
}
