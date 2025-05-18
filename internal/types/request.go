package types

import "time"

type UploadReportReq struct {
	UserWallet string    `form:"userWallet" binding:"required,ethwallet"` // hex
	Dataset    string    `form:"dataset" binding:"required"`
	TestTime   time.Time `form:"testTime" binding:"required"`
}

type SubmitAlgoReq struct {
	ScientistWallet string `json:"scientist_wallet" binding:"required,ethwallet"` // hex
	Dataset         string `json:"dataset" binding:"required"`
	AlgoLink        string `json:"algo_link" binding:"required"`
}

type VoteReq struct {
	AlgoId uint   `json:"algo_id" binding:"required"`
	TxHash string `json:"tx_hash" binding:"required"`
}

type SetCommitteeMemberReq struct {
	MemberWallet string `json:"member_wallet" binding:"required,ethwallet"`
	IsApproved   bool   `json:"is_approved" binding:"required"`
}
