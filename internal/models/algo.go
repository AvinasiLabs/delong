package models

import (
	"time"

	"gorm.io/gorm"
)

const AlgoJoinConfirmedTx = `
JOIN blockchain_transactions bt
ON bt.entity_id = algos.id
   AND bt.status = ?
   AND bt.entity_type = ?
`

type Algo struct {
	ID       uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	Name     string `gorm:"type:varchar(255);not null" json:"name"`
	AlgoLink string `gorm:"type:varchar(255);not null;uniqueIndex:idx_algo_link" json:"algo_link"`
	Cid      string `gorm:"type:varchar(255);not null;uniqueIndex:idx_cid;comment:source code CID" json:"cid"`

	// UsedDataset     string
	// ScientistWallet string // 0x...

	// Status        string     // enum: PENDING, APPROVED, REJECTED
	// VoteStartTime *time.Time // vote duration
	// VoteEndTime   *time.Time // vote duration
	CreatedAt time.Time `gorm:"autoCreateTime;index:idx_created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

// CreateAlgo creates a new algorithm
// func CreateAlgo(db *gorm.DB, name, link, scientistWallet, cid, dataset string) (*Algo, error) {
func CreateAlgo(db *gorm.DB, name, link, cid string) (*Algo, error) {
	algo := &Algo{
		Name:     name,
		AlgoLink: link,
		Cid:      cid,
		// Status:   ALGO_STATUS_PENDING,
		// ScientistWallet: scientistWallet,
		// UsedDataset:     dataset,
	}

	err := db.Create(algo).Error
	if err != nil {
		return nil, err
	}

	return algo, nil
}

// GetAlgoByID retrieves algorithm by ID
func GetAlgoByID(db *gorm.DB, id uint) (*Algo, error) {
	var algo Algo
	err := db.First(&algo, id).Error
	if err != nil {
		return nil, err
	}
	return &algo, nil
}

// GetAlgoByLink returns the most recent algorithm record with the given AlgoLink, if any.
func GetAlgoByLink(db *gorm.DB, link string) (*Algo, error) {
	var algo Algo
	err := db.Where("algo_link", link).First(&algo).Error
	if err != nil {
		return nil, err
	}
	return &algo, nil
}
