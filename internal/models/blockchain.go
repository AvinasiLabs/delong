package models

import (
	"time"

	"gorm.io/gorm"
)

// Blockchain transaction status constants
const (
	TX_STATUS_PENDING   = "PENDING"   // Transaction submitted but not confirmed
	TX_STATUS_CONFIRMED = "CONFIRMED" // Transaction confirmed on chain
	TX_STATUS_FAILED    = "FAILED"    // Transaction failed
)

const (
	// ENTITY_TYPE_ALGO        string = "ALGO"
	ENTITY_TYPE_EXECUTION   string = "EXECUTION"
	ENTITY_TYPE_VOTE        string = "VOTE"
	ENTITY_TYPE_COMMITTEE   string = "COMMITTEE"
	ENTITY_TYPE_TEST_REPORT string = "TEST_REPORT"
	ENTITY_TYPE_DATAUSAGE   string = "DATAUSAGE"
)

// BlockchainTransaction records blockchain transactions and their status
type BlockchainTransaction struct {
	ID             uint
	TxHash         string  // Transaction hash
	EntityID       uint    // Associated entity ID
	EntityType     string  // values: "vote", "algo", "committee"
	Status         string  // Transaction status (PENDING, CONFIRMED, FAILED)
	BlockNumber    *uint64 // Confirmation block number
	BlockTimestamp *time.Time
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

// GetTransactionByHash retrieves a transaction record by its hash
func GetTransactionByHash(db *gorm.DB, txHash string) (*BlockchainTransaction, error) {
	var tx BlockchainTransaction
	err := db.Where("tx_hash = ?", txHash).First(&tx).Error
	if err != nil {
		return nil, err
	}
	return &tx, nil
}

// CreateTransaction creates a new blockchain transaction record
func CreateTransaction(db *gorm.DB, txHash string, entityID uint, entityType string) (*BlockchainTransaction, error) {
	return CreateTransactionWithStatus(db, txHash, entityID, entityType, TX_STATUS_PENDING, nil, nil)
}

func CreateTransactionWithStatus(db *gorm.DB, txHash string, entityID uint, entityType string, status string, blockNumber *uint64, blockTimestamp *time.Time) (*BlockchainTransaction, error) {
	tx := &BlockchainTransaction{
		TxHash:         txHash,
		EntityID:       entityID,
		EntityType:     entityType,
		Status:         status,
		BlockNumber:    blockNumber,
		BlockTimestamp: blockTimestamp,
	}

	err := db.Create(tx).Error
	if err != nil {
		return nil, err
	}
	return tx, nil
}

// UpdateTransactionStatus updates the status of a transaction
func UpdateTransactionStatus(db *gorm.DB, txHash string, status string, blockNumber *uint64, blockTime *time.Time) (*BlockchainTransaction, error) {
	updates := map[string]any{
		"status": status,
	}

	if blockNumber != nil {
		updates["block_number"] = blockNumber
	}

	if blockTime != nil {
		updates["block_timestamp"] = blockTime
	}

	err := db.Model(&BlockchainTransaction{}).Where("tx_hash = ?", txHash).Updates(updates).Error
	if err != nil {
		return nil, err
	}

	var tx BlockchainTransaction
	if err := db.Where("tx_hash = ?", txHash).First(&tx).Error; err != nil {
		return nil, err
	}
	return &tx, nil
}
