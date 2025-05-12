package model

import (
	"time"

	"gorm.io/gorm"
)

const (
	ALGO_STATUS_APPROVED string = "APPROVED"
	ALGO_STATUS_REJECTED string = "REJECTED"
)

type Algo struct {
	ID              uint
	Name            string
	ScientistWallet string // 0x...
	AlgoLink        string
	Cid             string
	UsedDataset     string
	Status          string // enum: PENDING, APPROVED, REJECTED
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

// CreateAlgo creates a new algorithm
func CreateAlgo(db *gorm.DB, name, link, scientistWallet, cid, dataset string) (*Algo, error) {
	algo := &Algo{
		Name:            name,
		ScientistWallet: scientistWallet,
		AlgoLink:        link,
		Cid:             cid,
		UsedDataset:     dataset,
	}

	err := db.Create(algo).Error
	if err != nil {
		return nil, err
	}

	return algo, nil
}

// UpdateAlgoStatus updates algorithm status
func UpdateAlgoStatus(db *gorm.DB, id uint, status string) error {
	algo := &Algo{ID: id}
	err := db.Model(algo).Update("status", status).Error
	if err != nil {
		return err
	}

	return nil
}

// GetAlgoByID retrieves algorithm by ID
func GetAlgoByID(db *gorm.DB, id uint) (*Algo, error) {
	var algo Algo
	err := db.Where("id = ?", id).First(&algo).Error
	if err != nil {
		return nil, err
	}
	return &algo, nil
}

// GetAlgoByCid retrieves algorithm by CID
func GetAlgoByCid(db *gorm.DB, cid string) (*Algo, error) {
	var algo Algo
	err := db.Where("cid = ?", cid).First(&algo).Error
	if err != nil {
		return nil, err
	}
	return &algo, nil
}
