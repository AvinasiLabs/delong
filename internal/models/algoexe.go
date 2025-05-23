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
	ID              uint
	AlgoID          uint
	UsedDataset     string
	ScientistWallet string     // 0x...
	ReviewStatus    string     // enum: REVIEWING, APPROVED, REJECTED
	VoteStartTime   *time.Time // vote duration
	VoteEndTime     *time.Time // vote duration
	Status          string     // QUEUED, RUNNING, COMPLETED, FAILED
	StartTime       *time.Time // Container execution duration
	EndTime         *time.Time // Container execution duration
	Result          string     // Execution output
	ErrorMsg        string     // Error message if failed
	CreatedAt       time.Time
	UpdatedAt       time.Time
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
	// wrong query condition: Where("status IN ?", []string{EXE_STATUS_QUEUED, EXE_STATUS_RUNNING}).Find(&executions).Error
	// because EXE_STATUS_QUEUED indicated that algo execution taks is not resolved
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
