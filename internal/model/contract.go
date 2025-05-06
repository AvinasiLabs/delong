package model

import (
	"time"

	"gorm.io/gorm"
)

// ContractMeta is the schema for storing deployed contract addresses.
type ContractMeta struct {
	ID        uint
	Name      string // Contract identifier, e.g. "data_contribution"
	Address   string // Deployed contract address
	CreatedAt time.Time
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
