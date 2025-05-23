package models

import (
	"time"

	"gorm.io/gorm"
)

type Vote struct {
	ID        uint
	AlgoCid   string
	Voter     string // 0x...
	Approve   bool   // vote yes or no
	VotedAt   time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
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
