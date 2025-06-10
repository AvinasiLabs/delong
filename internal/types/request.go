package types

import "time"

type UploadReportReq struct {
	UserWallet string    `form:"user_wallet" binding:"required,ethwallet"` // hex
	Dataset    string    `form:"dataset" binding:"required"`
	TestTime   time.Time `form:"test_time" binding:"required"`
}

type SubmitAlgoExeReq struct {
	ScientistWallet string `json:"scientist_wallet" binding:"required,ethwallet"` // hex
	Dataset         string `json:"dataset" binding:"required"`
	GithubRepo      string `json:"github_repo" binding:"required"` // github code url
	CommitHash      string `json:"commit_hash" binding:"required"`
	AlgoLink        string `json:"algo_link"`
}

type VoteReq struct {
	AlgoId uint   `json:"algo_id" binding:"required"`
	TxHash string `json:"tx_hash" binding:"required"`
}

type SetCommitteeMemberReq struct {
	MemberWallet string `json:"member_wallet" binding:"required,ethwallet"`
	IsApproved   *bool  `json:"is_approved" binding:"required"`
}

type SetVotingDurationReq struct {
	Duration int64 `json:"duration" binding:"required"` // seconds
}

// Dataset
type DatasetCreateReq struct {
	Name        string `json:"name" binding:"required"`
	UiName      string `json:"ui_name" binding:"required"`
	Description string `json:"description" binding:"required"`
}

type DatasetUpdateReq struct {
	Description string `json:"description" binding:"required"`
}

type StcDatasetCreateReq struct {
	Name         string `form:"name" binding:"required"`
	Desc         string `form:"desc"`
	Author       string `form:"author"`
	AuthorWallet string `form:"author_wallet"`
}

type StcDatasetUpdateReq struct {
	Desc   string `json:"desc"`
	Author string `json:"author"`
}
