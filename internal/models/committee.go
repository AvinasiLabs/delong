package models

import (
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type CommitteeMember struct {
	ID           int
	MemberWallet string
	IsApproved   bool
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

// CreateOrUpdateCommitteeMember creates or updates a member based on wallet
func CreateOrUpdateCommitteeMember(db *gorm.DB, memberWallet string, isApproved bool) (*CommitteeMember, error) {
	cm := CommitteeMember{
		MemberWallet: memberWallet,
		IsApproved:   isApproved,
	}

	err := db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "member_wallet"}}, // UNIQUE constraint column
		DoUpdates: clause.AssignmentColumns([]string{"is_approved", "updated_at"}),
	}).Create(&cm).Error

	if err != nil {
		return nil, err
	}
	return &cm, nil
}
