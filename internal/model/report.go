package model

import "time"

// System-defined dataset registry
type Dataset struct {
	Name        string
	Title       string
	Description string
	CreatedAt   time.Time
}

// One uploaded report from user
type TestReport struct {
	ID           int
	UserWallet   string
	RawReportCid string
	Dataset      string
	TestTime     time.Time

	Results []TestResult `gorm:"foreignKey:ReportID"` // Logical link only
}

// One individual result item under a report
type TestResult struct {
	ID             int
	ReportID       int
	Category       string
	Name           string
	Definition     string
	Result         string
	ReferenceRange string
	Explanation    string
	Status         string
	Suggestions    *string
}
