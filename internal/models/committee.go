package models

import (
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type CommitteeMember struct {
	ID           int       `gorm:"primaryKey;autoIncrement" json:"id"`
	MemberWallet string    `gorm:"type:varchar(255);not null;uniqueIndex:idx_member_wallet" json:"member_wallet"`
	IsApproved   bool      `gorm:"not null;default:false;index:idx_is_approved" json:"is_approved"`
	CreatedAt    time.Time `gorm:"autoCreateTime;index:idx_created_at" json:"created_at"`
	UpdatedAt    time.Time `gorm:"autoUpdateTime" json:"updated_at"`
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

func GetCommitteeMemberByWallet(db *gorm.DB, wallet string) (*CommitteeMember, error) {
	var member CommitteeMember
	err := db.Table("committee_members").
		Joins("JOIN blockchain_transactions bt ON bt.entity_id = committee_members.id").
		Where("bt.status = ? AND bt.entity_type = ?", TX_STATUS_CONFIRMED, ENTITY_TYPE_COMMITTEE).
		Where("committee_members.member_wallet = ?", wallet).
		First(&member).Error
	return &member, err
}
