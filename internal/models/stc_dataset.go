package models

import (
	"time"

	"gorm.io/gorm"
)

type StaticDataset struct {
	ID           uint64    `gorm:"primaryKey;autoIncrement" json:"id"`
	Name         string    `gorm:"type:varchar(255);not null;index:idx_name" json:"name"`
	Desc         string    `gorm:"type:text" json:"desc"`
	FileHash     string    `gorm:"type:varchar(64);not null;uniqueIndex:idx_file_hash" json:"file_hash"` // SHA-256 hash of original file for deduplication
	IpfsCid      string    `gorm:"type:varchar(255);not null" json:"ipfs_cid"`
	FileSize     int64     `gorm:"bigint;not null" json:"file_size"`
	FileFormat   string    `gorm:"type:varchar(50);not null" json:"file_format"` // csv, json, parquet...
	Author       string    `gorm:"type:varchar(255)" json:"author"`
	AuthorWallet string    `gorm:"type:varchar(255);not null;index:idx_author_wallet" json:"author_wallet"`
	SampleUrl    string    `gorm:"type:varchar(255)" json:"sample_url"`
	CreatedAt    time.Time `gorm:"autoCreateTime;index:idx_created_at" json:"created_at"`
	UpdatedAt    time.Time `gorm:"autoUpdateTime" json:"updated_at"`
	// Tags    string `gorm:"type:text" json:"tags"`            // separated by commas
	// Version string `gorm:"type:varchar(50)" json:"version"`
	// License string `gorm:"type:varchar(100)" json:"license"`
}

type CreateStcDatasetReq struct {
	Name         string `json:"name"`
	Desc         string `json:"desc"`
	FileHash     string `json:"file_hash"`
	IpfsCid      string `json:"ipfs_cid"`
	FileSize     int64  `json:"file_size"`
	FileFormat   string `json:"file_format"`
	Author       string `json:"author"`
	AuthorWallet string `json:"author_wallet"`
	SampleUrl    string `json:"sample_url"`
}

func CreateStcDataset(db *gorm.DB, req CreateStcDatasetReq) (*StaticDataset, error) {
	asset := StaticDataset{
		Name:         req.Name,
		Desc:         req.Desc,
		FileHash:     req.FileHash,
		IpfsCid:      req.IpfsCid,
		FileSize:     req.FileSize,
		FileFormat:   req.FileFormat,
		Author:       req.Author,
		AuthorWallet: req.AuthorWallet,
		SampleUrl:    req.SampleUrl,
		// Tags:       req.Tags,
		// Version:    req.Version,
		// License:    req.License,
	}
	err := db.Create(&asset).Error
	if err != nil {
		return nil, err
	}
	return &asset, nil
}

func GetStcDataset(db *gorm.DB, page, pageSize int) ([]StaticDataset, int64, error) {
	var assets []StaticDataset
	var total int64

	tx := db.Model(&assets)
	if err := tx.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	err := tx.Limit(pageSize).
		Offset((page - 1) * pageSize).
		Order("created_at DESC").
		Find(&assets).Error
	return assets, total, err
}

func GetStcDatasetByID(db *gorm.DB, id uint) (*StaticDataset, error) {
	var asset StaticDataset
	err := db.First(&asset, id).Error
	if err != nil {
		return nil, err
	}
	return &asset, err
}

func GetStcDatasetByHash(db *gorm.DB, hash string) (*StaticDataset, error) {
	var asset StaticDataset
	err := db.Where("file_hash = ?", hash).First(&asset).Error
	if err != nil {
		return nil, err
	}
	return &asset, err
}

type UpdateStcDatasetReq struct {
	Description string `json:"description"`
	Author      string `json:"author"`
}

func UpdateStcDataset(db *gorm.DB, id uint, req UpdateStcDatasetReq) (*StaticDataset, error) {
	updates := map[string]any{
		"description": req.Description,
		"author":      req.Author,
	}
	err := db.Model(&StaticDataset{}).Where("id = ?", id).Updates(updates).Error
	if err != nil {
		return nil, err
	}

	asset := StaticDataset{}
	err = db.First(&asset, id).Error
	return &asset, err
}

func DeleteStcDataset(db *gorm.DB, id uint) error {
	return db.Delete(&StaticDataset{}, id).Error
}
