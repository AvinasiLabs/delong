package models

import (
	"time"

	"gorm.io/gorm"
)

type DataUsage struct {
	ID              uint
	ScientistWallet string
	Cid             string
	Dataset         string
	UsedAt          time.Time
	CreatedAt       time.Time
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
