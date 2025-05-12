package model

import (
	"time"

	"gorm.io/gorm"
)

const (
	EXESTATUS_QUEUED    = "QUEUED"
	EXESTATUS_RUNNING   = "RUNNING"
	EXESTATUS_COMPLETED = "COMPLETED"
	EXESTATUS_FAILED    = "FAILED"
)

// AlgoExecution tracks algorithm execution status and results
type AlgoExecution struct {
	ID        uint
	AlgoID    uint
	ExeStatus string // QUEUED, RUNNING, COMPLETED, FAILED
	StartTime *time.Time
	EndTime   *time.Time
	Result    string // Execution output
	ErrorMsg  string // Error message if failed
	CreatedAt time.Time
	UpdatedAt time.Time
}

// CreateAlgoExecution creates a new algorithm execution record
func CreateAlgoExecution(db *gorm.DB, algoID uint, status string) (*AlgoExecution, error) {
	now := time.Now()
	execution := &AlgoExecution{
		AlgoID:    algoID,
		ExeStatus: status,
		StartTime: &now,
	}

	if err := db.Create(execution).Error; err != nil {
		return nil, err
	}

	return execution, nil
}

// UpdateExecutionStatus updates the status of an execution
func UpdateExecutionStatus(db *gorm.DB, executionID uint, status string, errorMsg ...string) (*AlgoExecution, error) {
	now := time.Now()
	updates := map[string]any{
		"status":     status,
		"updated_at": now,
	}

	if status == EXESTATUS_COMPLETED || status == EXESTATUS_FAILED {
		updates["end_time"] = now
	}

	if len(errorMsg) > 0 && errorMsg[0] != "" {
		updates["error_msg"] = errorMsg[0]
	}

	var execution AlgoExecution
	err := db.Model(&AlgoExecution{}).Where("id = ?", executionID).Updates(updates).First(&execution).Error
	if err != nil {
		return nil, err
	}

	return &execution, nil
}

// UpdateExecutionResult updates the result of an execution
func UpdateExecutionResult(db *gorm.DB, executionID uint, result string) error {
	return db.Model(&AlgoExecution{}).Where("id = ?", executionID).Update("result", result).Error
}

// GetPendingExecutions retrieves all executions in QUEUED or RUNNING state
func GetPendingExecutions(db *gorm.DB) ([]AlgoExecution, error) {
	var executions []AlgoExecution
	err := db.Where("status IN ?", []string{EXESTATUS_QUEUED, EXESTATUS_RUNNING}).Find(&executions).Error
	return executions, err
}

// GetExecutionsByAlgoID gets all executions for a specific algorithm
func GetExecutionsByAlgoID(db *gorm.DB, algoID uint) ([]AlgoExecution, error) {
	var executions []AlgoExecution
	err := db.Where("algo_id = ?", algoID).Order("created_at DESC").Find(&executions).Error
	return executions, err
}
