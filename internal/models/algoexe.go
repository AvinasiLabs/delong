package models

import (
	"time"

	"gorm.io/gorm"
)

const (
	ALGO_STATUS_REVIEWING string = "REVIEWING"
	ALGO_STATUS_APPROVED  string = "APPROVED"
	ALGO_STATUS_REJECTED  string = "REJECTED"
)

const (
	EXE_STATUS_QUEUED    = "QUEUED"
	EXE_STATUS_RUNNING   = "RUNNING"
	EXE_STATUS_COMPLETED = "COMPLETED"
	EXE_STATUS_FAILED    = "FAILED"
)

const AlgoExeJoinConfirmedTx = `
JOIN blockchain_transactions bt
ON bt.entity_id = algo_exes.id
   AND bt.status = ?
   AND bt.entity_type = ?
`

// AlgoExe tracks algorithm execution status and results
type AlgoExe struct {
	ID              uint       `gorm:"primaryKey;autoIncrement;type:bigint unsigned" json:"id"`
	AlgoID          uint       `gorm:"type:bigint unsigned;not null;index:idx_algo_id" json:"algo_id"`
	UsedDataset     string     `gorm:"type:varchar(255);not null;index:idx_dataset" json:"used_dataset"`
	ScientistWallet string     `gorm:"type:varchar(255);not null;index:idx_wallet;comment:Ethereum wallet address in hexadecimal format (0x...)" json:"scientist_wallet"`
	ReviewStatus    string     `gorm:"type:enum('REVIEWING','APPROVED','REJECTED');not null;default:'REVIEWING';index:idx_review_status" json:"review_status"`
	VoteStartTime   *time.Time `gorm:"type:timestamp null" json:"vote_start_time"`
	VoteEndTime     *time.Time `gorm:"type:timestamp null" json:"vote_end_time"`
	Status          string     `gorm:"type:enum('QUEUED','RUNNING','COMPLETED','FAILED');not null;default:'QUEUED';index:idx_status" json:"status"`
	StartTime       *time.Time `gorm:"type:timestamp null" json:"start_time"`
	EndTime         *time.Time `gorm:"type:timestamp null" json:"end_time"`
	Result          string     `gorm:"type:text" json:"result"`
	ErrorMsg        string     `gorm:"type:text" json:"error_msg"`
	CreatedAt       time.Time  `gorm:"autoCreateTime;index:idx_created_at" json:"created_at"`
	UpdatedAt       time.Time  `gorm:"autoUpdateTime" json:"updated_at"`
}

// CreateAlgoExecution creates a new algorithm execution record
func CreateAlgoExecution(db *gorm.DB, algoID uint, dataset, scientist string) (*AlgoExe, error) {
	execution := &AlgoExe{
		AlgoID:          algoID,
		Status:          EXE_STATUS_QUEUED,
		UsedDataset:     dataset,
		ScientistWallet: scientist,
		ReviewStatus:    ALGO_STATUS_REVIEWING,
	}

	if err := db.Create(execution).Error; err != nil {
		return nil, err
	}

	return execution, nil
}

// GetPendingExecutions retrieves all executions in QUEUED or RUNNING state
func GetPendingAlgoExesConfirmed(db *gorm.DB) ([]AlgoExe, error) {
	var executions []AlgoExe
	// EXE_STATUS_QUEUED indicated that algo execution taks is not resolved
	err := db.Table("algo_exes").
		Joins(AlgoExeJoinConfirmedTx, TX_STATUS_CONFIRMED, ENTITY_TYPE_EXECUTION).
		Where("status = ?", EXE_STATUS_RUNNING).Find(&executions).Error
	return executions, err
}

func GetAlgoExes(db *gorm.DB, page, pageSize int) ([]AlgoExe, int64, error) {
	var algoExes []AlgoExe
	var total int64

	tx := db.Model(&AlgoExe{}).Joins(AlgoExeJoinConfirmedTx, TX_STATUS_CONFIRMED, ENTITY_TYPE_EXECUTION)
	if err := tx.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * pageSize
	err := tx.Offset(offset).Limit(pageSize).Order("algo_exes.created_at DESC").Find(&algoExes).Error
	return algoExes, total, err
}

func GetAlgoExeById(db *gorm.DB, id uint) (*AlgoExe, error) {
	var algoExe AlgoExe
	err := db.Model(&AlgoExe{}).Joins(AlgoExeJoinConfirmedTx, TX_STATUS_CONFIRMED, ENTITY_TYPE_EXECUTION).First(&algoExe, id).Error
	if err != nil {
		return nil, err
	}
	return &algoExe, nil
}

func GetReviewingAlgoExes(db *gorm.DB) ([]AlgoExe, error) {
	var algoExes []AlgoExe
	err := db.Model(&AlgoExe{}).
		Joins(AlgoExeJoinConfirmedTx, TX_STATUS_CONFIRMED, ENTITY_TYPE_EXECUTION).
		Where("algo_exes.review_status = ?", ALGO_STATUS_REVIEWING).
		Find(&algoExes).Error
	if err != nil {
		return nil, err
	}
	return algoExes, nil
}

func UpdateReviewStatus(db *gorm.DB, id uint, reviewStatus string) error {
	return db.Model(&AlgoExe{ID: id}).Update("review_status", reviewStatus).Error
}

// UpdateVoteDuration updates start time and end time after onchain event log received
func UpdateVoteDuration(db *gorm.DB, exeId uint, startTime *time.Time, endTime *time.Time) error {
	updates := map[string]any{}
	if startTime != nil {
		updates["vote_start_time"] = startTime
	}
	if endTime != nil {
		updates["vote_end_time"] = endTime
	}
	algoExe := &AlgoExe{ID: exeId}
	err := db.Model(algoExe).Updates(updates).Error
	if err != nil {
		return err
	}
	return nil
}

// UpdateExecutionStatus updates the status of an execution
func UpdateExecutionStatus(db *gorm.DB, executionID uint, status string) (*AlgoExe, error) {
	now := time.Now()
	updates := map[string]any{
		"status":     status,
		"updated_at": now,
	}

	if status == EXE_STATUS_COMPLETED || status == EXE_STATUS_FAILED {
		updates["end_time"] = now
	}

	var execution AlgoExe
	err := db.Model(&AlgoExe{}).Where("id = ?", executionID).Updates(updates).First(&execution).Error
	if err != nil {
		return nil, err
	}

	return &execution, nil
}

// UpdateExecutionCompleted updates the status of an execution to completed and sets the result
func UpdateExecutionCompleted(db *gorm.DB, executionID uint, result string, errorMsg ...string) (*AlgoExe, error) {
	now := time.Now()
	updates := map[string]any{
		"status":     EXE_STATUS_COMPLETED,
		"updated_at": now,
		"end_time":   now,
		"result":     result,
	}
	if len(errorMsg) > 0 && errorMsg[0] != "" {
		updates["error_msg"] = errorMsg[0]
	}

	var execution AlgoExe
	err := db.Model(&AlgoExe{}).Where("id = ?", executionID).Updates(updates).First(&execution).Error
	if err != nil {
		return nil, err
	}

	return &execution, nil
}

// GetExecutionByAlgoID gets the first execution for a specific algorithm
// func GetExecutionByAlgoID(db *gorm.DB, algoID uint) (*AlgoExe, error) {
// 	executions := AlgoExe{}
// 	err := db.Where("algo_id = ?", algoID).Order("created_at DESC").First(&executions).Error
// 	return &executions, err
// }
