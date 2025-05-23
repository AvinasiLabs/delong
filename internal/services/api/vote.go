package api

import (
	"delong/internal/models"
	"delong/internal/types"
	"delong/pkg/bizcode"
	"delong/pkg/responser"
	"log"

	"github.com/gin-gonic/gin"
)

type VoteResource struct {
	ApiServiceOptions
}

func (r *VoteResource) ListHandler(c *gin.Context) {
	algoCid := c.Query("algo_cid")
	var votes []models.Vote
	votes, err := models.GetVotesByAlgoCid(r.MysqlDb, algoCid)
	if err != nil {
		log.Printf("Failed to get votes by algoId: %v", err)
		responser.ResponseError(c, bizcode.MYSQL_READ_FAIL)
		return
	}

	responser.ResponseData(c, votes)
}

func (r *VoteResource) SetVotingDuration(c *gin.Context) {
	req := types.SetVotingDurationReq{}
	err := c.ShouldBind(&req)
	if err != nil {
		log.Printf("Failed to bind set committee member request: %v", err)
		responser.ResponseError(c, bizcode.BAD_REQUEST)
		return
	}
	if req.Duration <= 0 {
		log.Printf("Vote duration should great than 0, duration=%d", req.Duration)
		responser.ResponseError(c, bizcode.BAD_REQUEST)
		return
	}

	tx, err := r.CtrCaller.SetVotingDuration(c.Request.Context(), req.Duration)
	if err != nil {
		log.Printf("Failed to call setVotingDuration on ethereum: %v", err)
		responser.ResponseError(c, bizcode.ETHEREUM_CALL_FAIL)
		return
	}

	responser.ResponseData(c, tx.Hash().Hex())
}
