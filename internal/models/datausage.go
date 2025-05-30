package models

import (
	"time"

	"gorm.io/gorm"
)

type DataUsage struct {
	ID              uint      `gorm:"primaryKey;autoIncrement;type:bigint unsigned" json:"id"`
	ScientistWallet string    `gorm:"type:varchar(255);not null;index:idx_wallet" json:"scientist_wallet"`
	Cid             string    `gorm:"type:varchar(255);not null;index:idx_cid" json:"cid"`
	Dataset         string    `gorm:"type:varchar(255);not null;index:idx_dataset;comment:Dataset name" json:"dataset"`
	UsedAt          time.Time `gorm:"type:timestamp;not null;index:idx_used_at" json:"used_at"`
	CreatedAt       time.Time `gorm:"type:timestamp;not null;autoCreateTime" json:"created_at"`
	UpdatedAt       time.Time `gorm:"type:timestamp;not null;autoUpdateTime" json:"updated_at"`
}

func CreateDataUsage(db *gorm.DB, scientist string, cid string, dataset string, usedAtUnix int64) (*DataUsage, error) {
	record := &DataUsage{
		ScientistWallet: scientist,
		Cid:             cid,
		Dataset:         dataset,
		UsedAt:          time.Unix(usedAtUnix, 0),
	}
	if err := db.Create(record).Error; err != nil {
		return nil, err
	}

	return record, nil
}
