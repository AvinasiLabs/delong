package models

import (
	"time"

	"gorm.io/gorm"
)

const (
	ALGO_STATUS_PENDING  string = "PENDING"
	ALGO_STATUS_APPROVED string = "APPROVED"
	ALGO_STATUS_REJECTED string = "REJECTED"
)

type Algo struct {
	ID              uint
	Name            string
	ScientistWallet string // 0x...
	AlgoLink        string
	Cid             string
	UsedDataset     string
	Status          string    // enum: PENDING, APPROVED, REJECTED
	StartTime       time.Time // vote duration
	EndTime         time.Time // vote duration
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

type AlgoDetails struct {
	Algo
	Votes      []Vote          `json:"votes"`
	Executions []AlgoExecution `json:"executions"`
}

// CreateAlgo creates a new algorithm
func CreateAlgo(db *gorm.DB, name, link, scientistWallet, cid, dataset string) (*Algo, error) {
	algo := &Algo{
		Name:            name,
		ScientistWallet: scientistWallet,
		AlgoLink:        link,
		Cid:             cid,
		UsedDataset:     dataset,
		Status:          ALGO_STATUS_PENDING,
	}

	err := db.Create(algo).Error
	if err != nil {
		return nil, err
	}

	return algo, nil
}

// GetConfirmedAlgos return all algos of which tx status are CONFIRMED
func GetConfirmedAlgos(db *gorm.DB, page, pageSize int) ([]Algo, int64, error) {
	var algos []Algo
	var total int64

	tx := db.Table("algos").
		Select("algos.*").
		Joins("JOIN blockchain_transactions bt ON bt.entity_id = algos.id").
		Where("bt.status = ? AND bt.entity_type = ?", TX_STATUS_CONFIRMED, ENTITY_TYPE_ALGO)

	if err := tx.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * pageSize
	err := tx.Offset(offset).Limit(pageSize).Order("algos.created_at DESC").Find(&algos).Error
	return algos, total, err
}

// UpdateAlgoVoteDuration updates start time and end time after onchain event log received
func UpdateAlgoVoteDuration(db *gorm.DB, id uint, startTime time.Time, endTime time.Time) error {
	algo := &Algo{ID: id}
	updates := map[string]time.Time{
		"start_time": startTime,
		"end_time":   endTime,
	}
	err := db.Model(algo).Updates(updates).Error
	if err != nil {
		return err
	}
	return nil
}

// UpdateAlgoStatus updates algorithm status
func UpdateAlgoStatus(db *gorm.DB, id uint, status string) error {
	algo := &Algo{ID: id}
	err := db.Model(algo).Update("status", status).Error
	if err != nil {
		return err
	}

	return nil
}

// GetConfirmedAlgoByID retrieves algos of which tx status is CONFIRMED
func GetConfirmedAlgoByID(db *gorm.DB, id uint) (*Algo, error) {
	var algo Algo
	err := db.Table("algos").
		Select("algos.*").
		Joins("JOIN blockchain_transactions bt ON bt.entity_id = algos.id").
		Where("bt.status = ? AND bt.entity_type = ?", TX_STATUS_CONFIRMED, ENTITY_TYPE_ALGO).
		Where("algos.id = ?", id).
		First(&algo).Error

	if err != nil {
		return nil, err
	}

	return &algo, nil
}

// GetAlgoByID retrieves algorithm by ID
func GetAlgoByID(db *gorm.DB, id uint) (*Algo, error) {
	var algo Algo
	err := db.Where("id = ?", id).First(&algo).Error
	if err != nil {
		return nil, err
	}
	return &algo, nil
}

// GetAlgoByCid retrieves algorithm by CID
func GetAlgoByCid(db *gorm.DB, cid string) (*Algo, error) {
	var algo Algo
	err := db.Where("cid = ?", cid).First(&algo).Error
	if err != nil {
		return nil, err
	}
	return &algo, nil
}

// GetPendingConfirmedAlgos returns all PENDING algos of which tx is CONFIRMED
func GetPendingConfirmedAlgos(db *gorm.DB) ([]Algo, error) {
	var algos []Algo
	err := db.
		Table("algos").
		Joins("JOIN blockchain_transactions bt ON bt.entity_id = algos.id").
		Where("algos.status = ? AND bt.status = ? AND bt.entity_type = ?", ALGO_STATUS_PENDING, TX_STATUS_CONFIRMED, ENTITY_TYPE_ALGO).
		Select("algos.*").
		Find(&algos).Error
	return algos, err
}

// GetAlgoDetailsByID retrieves algorithm details by ID
func GetAlgoDetailsByID(db *gorm.DB, id uint) (*AlgoDetails, error) {
	var algo Algo
	if err := db.
		Table("algos").
		Select("algos.*").
		Joins("JOIN blockchain_transactions bt ON bt.entity_id = algos.id").
		Where("bt.status = ? AND bt.entity_type = ? AND algos.id = ?", TX_STATUS_CONFIRMED, ENTITY_TYPE_ALGO, id).
		First(&algo).Error; err != nil {
		return nil, err
	}

	var votes []Vote
	if err := db.
		Table("votes").
		Select("votes.*").
		Joins("JOIN blockchain_transactions bt ON bt.entity_id = votes.id").
		Where("bt.status = ? AND bt.entity_type = ? AND votes.algo_id = ?", TX_STATUS_CONFIRMED, ENTITY_TYPE_VOTE, id).
		Order("vote_time ASC").
		Find(&votes).Error; err != nil {
		return nil, err
	}

	var executions []AlgoExecution
	if err := db.
		Where("algo_id = ?", id).
		Order("created_at DESC").
		Find(&executions).Error; err != nil {
		return nil, err
	}

	return &AlgoDetails{
		Algo:       algo,
		Votes:      votes,
		Executions: executions,
	}, nil
}
