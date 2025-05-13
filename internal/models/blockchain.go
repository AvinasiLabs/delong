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

// No transaction type constants needed

// BlockchainTransaction records blockchain transactions and their status
type BlockchainTransaction struct {
	ID             uint
	TxHash         string  // Transaction hash
	EntityID       uint    // Associated entity ID
	Status         string  // Transaction status (pending, confirmed, failed)
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
func CreateTransaction(db *gorm.DB, txHash string, entityID uint) (*BlockchainTransaction, error) {
	tx := &BlockchainTransaction{
		TxHash:   txHash,
		EntityID: entityID,
		Status:   TX_STATUS_PENDING,
	}

	err := db.Create(tx).Error
	if err != nil {
		return nil, err
	}
	return tx, nil
}

// UpdateTransactionEntity updates both the EntityID and Status
func UpdateTransactionEntity(db *gorm.DB, txHash string, entityID uint, status string, blockNumber *uint64, blockTime *time.Time) error {
	updates := map[string]any{
		"entity_id": entityID,
		"status":    status,
	}
	if blockNumber != nil {
		updates["block_number"] = blockNumber
	}
	if blockTime != nil {
		updates["block_timestamp"] = blockTime
	}
	return db.Model(&BlockchainTransaction{}).
		Where("tx_hash = ?", txHash).
		Updates(updates).
		Error
}

// UpdateTransactionStatus updates the status of a transaction
func UpdateTransactionStatus(db *gorm.DB, txHash string, status string, blockNumber *uint64, blockTime *time.Time) error {
	updates := map[string]any{
		"status": status,
	}

	if blockNumber != nil {
		updates["block_number"] = blockNumber
	}

	if blockTime != nil {
		updates["block_timestamp"] = blockTime
	}

	return db.Model(&BlockchainTransaction{}).Where("tx_hash = ?", txHash).Updates(updates).Error
}

// GetTransactionStatusByEntity retrieves blockchain transaction status by entity ID
func GetTransactionStatusByEntity(db *gorm.DB, entityID uint) (*BlockchainTransaction, error) {
	var tx BlockchainTransaction
	err := db.Where("entity_id = ?", entityID).Order("created_at DESC").First(&tx).Error
	if err != nil {
		return nil, err
	}
	return &tx, nil
}
