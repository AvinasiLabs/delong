package models

import (
	"time"

	"gorm.io/gorm"
)

// ContractMeta is the schema for storing deployed contract addresses.
type ContractMeta struct {
	ID        uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	Name      string    `gorm:"type:varchar(255);not null;uniqueIndex;comment:contract identifier, e.g. 'data_contribution'" json:"name"`
	Address   string    `gorm:"type:varchar(42);not null;comment:contract address" json:"address"`
	CreatedAt time.Time `gorm:"type:datetime;autoCreateTime" json:"created_at"`
}

// GetContractAddress returns the contract address for a given name.
func GetContractAddress(db *gorm.DB, name string) (string, error) {
	var meta ContractMeta
	err := db.First(&meta, "name = ?", name).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return "", nil
		}
		return "", err
	}
	return meta.Address, nil
}

// SaveContractAddress stores a new contract address.
func SaveContractAddress(db *gorm.DB, name string, address string) error {
	meta := &ContractMeta{
		Name:    name,
		Address: address,
	}
	return db.Create(meta).Error
}
