package models

import (
	"time"

	"gorm.io/gorm"
)

type Vote struct {
	ID        uint      `gorm:"primaryKey;autoIncrement;type:bigint unsigned" json:"id"`
	AlgoCid   string    `gorm:"type:varchar(255);not null;index:idx_algo_cid" json:"algo_cid"`
	Voter     string    `gorm:"type:varchar(255);not null;index:idx_voter;comment:0x..." json:"voter"`
	Approve   bool      `gorm:"not null" json:"approve"`
	VotedAt   time.Time `gorm:"type:datetime;not null;autoCreateTime;index:idx_voted_at" json:"voted_at"`
	CreatedAt time.Time `gorm:"type:timestamp;not null;autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"type:timestamp;not null;autoUpdateTime" json:"updated_at"`
}

func CreateVote(db *gorm.DB, algoCid string, voter string, approve bool, voteTime time.Time) (*Vote, error) {
	vote := &Vote{
		AlgoCid: algoCid,
		Voter:   voter,
		Approve: approve,
		VotedAt: voteTime,
	}
	if err := db.Create(vote).Error; err != nil {
		return nil, err
	}
	return vote, nil
}

func GetVotesByAlgoCid(db *gorm.DB, algoCid string) ([]Vote, error) {
	var votes []Vote
	err := db.Table("votes").
		Joins("JOIN blockchain_transactions bt ON bt.entity_id = votes.id").
		Where("votes.algo_cid = ? AND bt.status = ? AND bt.entity_type = ?", algoCid, TX_STATUS_CONFIRMED, ENTITY_TYPE_VOTE).
		Select("votes.*").
		Find(&votes).Error
	return votes, err
}
