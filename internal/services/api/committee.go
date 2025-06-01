package api

import (
	"delong/internal/models"
	"delong/internal/types"
	"delong/pkg/bizcode"
	"delong/pkg/responser"
	"errors"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CommitteeResource struct{ ApiServiceOptions }

func (r *CommitteeResource) CreateHandler(c *gin.Context) {
	req := types.SetCommitteeMemberReq{}
	err := c.ShouldBind(&req)
	if err != nil {
		log.Printf("Failed to bind set committee member request: %v", err)
		responser.ResponseError(c, bizcode.BAD_REQUEST)
		return
	}
	memberWallet := common.HexToAddress(req.MemberWallet)

	sure, err := isAdmin(c)
	if err != nil {
		log.Printf("Failed to check admin status: %v", err)
		responser.ResponseError(c, bizcode.INTERNAL_SERVER_ERROR)
		return
	}
	if !sure {
		log.Printf("Not admin")
		responser.ResponseError(c, bizcode.FORBIDDEN)
		return
	}

	dbtx := r.MysqlDb.Begin()
	defer func() {
		if r := recover(); r != nil {
			dbtx.Rollback()
			panic(r)
		}
	}()

	cm, err := models.UpsertCommitteeMember(dbtx, req.MemberWallet, *req.IsApproved)
	if err != nil {
		dbtx.Rollback()
		log.Printf("Failed to create algorithm record: %v", err)
		responser.ResponseError(c, bizcode.MYSQL_WRITE_FAIL)
		return
	}

	tx, err := r.CtrCaller.SetCommitteeMember(c.Request.Context(), memberWallet, *req.IsApproved)
	if err != nil {
		dbtx.Rollback()
		log.Printf("Failed to set committee member: %v", err)
		responser.ResponseError(c, bizcode.ETHEREUM_CALL_FAIL)
		return
	}

	_, err = models.CreateTransaction(dbtx, tx.Hash().Hex(), uint(cm.ID), models.ENTITY_TYPE_COMMITTEE)
	if err != nil {
		dbtx.Rollback()
		log.Printf("Failed to create blockchain transaction record: %v", err)
		responser.ResponseError(c, bizcode.MYSQL_WRITE_FAIL)
		return
	}

	if err := dbtx.Commit().Error; err != nil {
		log.Printf("Failed to commit db transaction: %v", err)
		responser.ResponseError(c, bizcode.MYSQL_WRITE_FAIL)
		return
	}

	responser.ResponseData(c, tx.Hash().Hex())
}

func (r *CommitteeResource) ListHandler(c *gin.Context) {
	page, pageSize := parsePageParams(c)
	members, total, err := models.GetConfirmedCommitteeMembers(r.MysqlDb, page, pageSize)
	if err != nil {
		log.Printf("Failed to get committee members: %v", err)
		responser.ResponseError(c, bizcode.MYSQL_READ_FAIL)
		return
	}
	responser.ResponseList(c, page, pageSize, total, members)
}

func (r *CommitteeResource) TakeHandler(c *gin.Context) {
	var id uint
	if err := parseUintParam(c.Param("id"), &id); err != nil {
		log.Printf("Failed to parse id param: %v", err)
		responser.ResponseError(c, bizcode.BAD_REQUEST)
		return
	}

	member, err := models.GetConfirmedCommitteeMemberByID(r.MysqlDb, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Printf("Committee member %v not found", id)
			responser.ResponseError(c, bizcode.NOT_FOUND)
			return
		}
		log.Printf("Failed to get committee member: %v", err)
		responser.ResponseError(c, bizcode.MYSQL_READ_FAIL)
		return
	}

	responser.ResponseData(c, member)
}
