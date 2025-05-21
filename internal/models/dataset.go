package models

import (
	"time"

	"gorm.io/gorm"
)

type DatasetRegistry struct {
	ID          int
	Name        string
	UiName      string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func CreateDataset(db *gorm.DB, name, uiName, desc string) (*DatasetRegistry, error) {
	dataset := DatasetRegistry{
		Name:        name,
		UiName:      uiName,
		Description: desc,
	}
	err := db.Create(&dataset).Error
	if err != nil {
		return nil, err
	}
	return &dataset, nil
}

func GetDatasets(db *gorm.DB, page, pageSize int) ([]DatasetRegistry, int64, error) {
	var datasets []DatasetRegistry
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

func GetDatasetByID(db *gorm.DB, id uint) (*DatasetRegistry, error) {
	var dataset DatasetRegistry
	err := db.First(&dataset, id).Error
	if err != nil {
		return nil, err
	}
	return &dataset, err
}

func UpdateDataset(db *gorm.DB, id uint, uiName, desc string) (*DatasetRegistry, error) {
	updates := map[string]any{
		"ui_name":     uiName,
		"description": desc,
	}
	err := db.Model(&DatasetRegistry{}).Where("id = ?", id).Updates(updates).Error
	if err != nil {
		return nil, err
	}

	dataset := DatasetRegistry{}
	err = db.First(&dataset, id).Error
	return &dataset, err
}

func DeleteDataset(db *gorm.DB, id uint) error {
	return db.Delete(&DatasetRegistry{}, id).Error
}
