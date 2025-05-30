package models

import "time"

// One uploaded report from user
type TestReport struct {
	ID           uint         `gorm:"primaryKey;autoIncrement" json:"id"`
	UserWallet   string       `gorm:"type:varchar(255);not null;index:idx_user_wallet" json:"user_wallet"`
	FileHash     string       `gorm:"type:varchar(64);not null;uniqueIndex:idx_file_hash;comment:SHA-256 hash of original file for deduplication" json:"file_hash"`
	RawReportCid string       `gorm:"type:varchar(255);not null;uniqueIndex:idx_raw_report_cid;comment:CID of encrypted original file" json:"raw_report_cid"`
	Dataset      string       `gorm:"type:varchar(255);not null;index:idx_dataset;comment:dataset_registry.name" json:"dataset"`
	TestTime     time.Time    `gorm:"type:timestamp;not null;autoCreateTime;index:idx_test_time" json:"test_time"`

	Results []TestResult `gorm:"foreignKey:TestReportID" json:"results"`
}

// One individual result item under a report
type TestResult struct {
	ID             uint    `gorm:"primaryKey;autoIncrement" json:"id"`
	TestReportID   int     `gorm:"not null;index:idx_test_report_id" json:"test_report_id"`
	Category       string  `gorm:"type:varchar(255);not null;index:idx_category" json:"category"`
	Name           string  `gorm:"type:text;not null" json:"name"`
	Definition     string  `gorm:"type:text;not null" json:"definition"`
	Result         string  `gorm:"type:text;not null" json:"result"`
	ReferenceRange string  `gorm:"type:text" json:"reference_range"`
	Explanation    string  `gorm:"type:text;not null" json:"explanation"`
	Status         string  `gorm:"type:enum('above_range','below_range','within_range','unknown');not null" json:"status"`
	Suggestions    *string `gorm:"type:text" json:"suggestions"`
}
