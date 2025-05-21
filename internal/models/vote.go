package models

import (
	"time"

	"gorm.io/gorm"
)

type Vote struct {
	ID        uint
	AlgoID    uint
	Voter     string // 0x...
	Approve   bool   // vote yes or no
	VoteTime  time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
}

func CreateVote(db *gorm.DB, algoID uint, voter string, approve bool, voteTime time.Time) (*Vote, error) {
	vote := &Vote{
		AlgoID:   algoID,
		Voter:    voter,
		Approve:  approve,
		VoteTime: voteTime,
	}
	if err := db.Create(vote).Error; err != nil {
		return nil, err
	}
	return vote, nil
}

func GetVotesByAlgoID(db *gorm.DB, algoID uint) ([]Vote, error) {
	var votes []Vote
	err := db.Table("votes").
		Joins("JOIN blockchain_transactions bt ON bt.entity_id = votes.id").
		Where("votes.algo_id = ? AND bt.status = ? AND bt.entity_type = ?", algoID, TX_STATUS_CONFIRMED, ENTITY_TYPE_VOTE).
		Select("votes.*").
		Find(&votes).Error
	return votes, err
}
