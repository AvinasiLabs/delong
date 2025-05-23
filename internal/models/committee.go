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

// UpsertCommitteeMember creates or updates a member based on wallet
func UpsertCommitteeMember(db *gorm.DB, memberWallet string, isApproved bool) (*CommitteeMember, error) {
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

// GetConfirmedCommitteeMembers returns confirmed committee members
func GetConfirmedCommitteeMembers(db *gorm.DB, page, pageSize int) ([]CommitteeMember, int64, error) {
	var members []CommitteeMember
	var total int64

	err := db.Table("committee_members").
		Joins("JOIN blockchain_transactions bt ON bt.entity_id = committee_members.id").
		Where("bt.status = ? AND bt.entity_type = ?", TX_STATUS_CONFIRMED, ENTITY_TYPE_COMMITTEE).
		Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = db.Table("committee_members").
		Select("committee_members.*").
		Joins("JOIN blockchain_transactions bt ON bt.entity_id = committee_members.id").
		Where("bt.status = ? AND bt.entity_type = ?", TX_STATUS_CONFIRMED, ENTITY_TYPE_COMMITTEE).
		Order("committee_members.created_at DESC").
		Offset((page - 1) * pageSize).
		Limit(pageSize).
		Find(&members).Error

	return members, total, err
}

// GetConfirmedCommitteeMemberByID returns one confirmed member
func GetConfirmedCommitteeMemberByID(db *gorm.DB, id uint) (*CommitteeMember, error) {
	var member CommitteeMember
	err := db.Table("committee_members").
		Joins("JOIN blockchain_transactions bt ON bt.entity_id = committee_members.id").
		Where("bt.status = ? AND bt.entity_type = ?", TX_STATUS_CONFIRMED, ENTITY_TYPE_COMMITTEE).
		Where("committee_members.id = ?", id).
		First(&member).Error
	return &member, err
}
