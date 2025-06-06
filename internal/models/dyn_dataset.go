package models

import (
	"time"

	"gorm.io/gorm"
)

type DynamicDataset struct {
	ID          int       `gorm:"primaryKey;autoIncrement" json:"id"`
	Name        string    `gorm:"type:varchar(255);not null;uniqueIndex:idx_name" json:"name"`
	Description string    `gorm:"type:text" json:"description"`
	FilePath    string    `gorm:"type:varchar(255)" json:"file_path"`
	CreatedAt   time.Time `gorm:"autoCreateTime;index:idx_created_at" json:"created_at"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

func CreateDataset(db *gorm.DB, name, uiName, desc, filePath string) (*DynamicDataset, error) {
	dataset := DynamicDataset{
		Name:        name,
		Description: desc,
		FilePath:    filePath,
	}
	err := db.Create(&dataset).Error
	if err != nil {
		return nil, err
	}
	return &dataset, nil
}

func GetDatasets(db *gorm.DB, page, pageSize int) ([]DynamicDataset, int64, error) {
	var datasets []DynamicDataset
	var total int64

	tx := db.Model(&datasets)
	if err := tx.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	err := tx.Limit(pageSize).
		Offset((page - 1) * pageSize).
		Order("created_at DESC").
		Find(&datasets).Error
	return datasets, total, err
}

func GetDatasetByID(db *gorm.DB, id uint) (*DynamicDataset, error) {
	var dataset DynamicDataset
	err := db.First(&dataset, id).Error
	if err != nil {
		return nil, err
	}
	return &dataset, err
}

func UpdateDataset(db *gorm.DB, id uint, desc string) (*DynamicDataset, error) {
	updates := map[string]any{
		"description": desc,
	}
	err := db.Model(&DynamicDataset{}).Where("id = ?", id).Updates(updates).Error
	if err != nil {
		return nil, err
	}

	dataset := DynamicDataset{}
	err = db.First(&dataset, id).Error
	return &dataset, err
}

func DeleteDataset(db *gorm.DB, id uint) error {
	return db.Delete(&DynamicDataset{}, id).Error
}
