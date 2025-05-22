package api

import (
	"delong/internal/models"
	"delong/internal/types"
	"delong/pkg/bizcode"
	"delong/pkg/responser"
	"errors"
	"log"
	"net/http"

	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AlgoResource struct {
	ApiServiceOptions
}

func (r *AlgoResource) CreateHandler(c *gin.Context) {
	ctx := c.Request.Context()
	req := types.SubmitAlgoReq{}
	err := c.ShouldBind(&req)
	if err != nil {
		log.Printf("Failed to bind request for submitting algorithm: %v", err)
		responser.ResponseError(c, bizcode.BAD_REQUEST)
		return
	}

	repo, err := extractRepoName(req.AlgoLink)
	if err != nil {
		log.Printf("Failed to extract repo name from algo link: %v", err)
		responser.ResponseError(c, bizcode.BAD_REQUEST)
		return
	}

	resp, err := http.Get(req.AlgoLink)
	if err != nil {
		log.Printf("Failed to open algorithm link: %v", err)
		responser.ResponseError(c, bizcode.AlGO_LINK_INVALID)
		return
	}
	defer resp.Body.Close()

	cid, err := r.IpfsStore.UploadStream(ctx, resp.Body)
	if err != nil {
		log.Printf("Failed to upload algorithm to IPFS: %v", err)
		responser.ResponseError(c, bizcode.IPFS_UPLOAD_FAIL)
		return
	}

	dbtx := r.MysqlDb.Begin()
	defer func() {
		if r := recover(); r != nil {
			dbtx.Rollback()
			panic(r)
		}
	}()

	algo, err := models.CreateAlgo(dbtx, repo, req.AlgoLink, req.ScientistWallet, cid, req.Dataset)
	if err != nil {
		dbtx.Rollback()
		log.Printf("Failed to create algorithm record: %v", err)
		responser.ResponseError(c, bizcode.MYSQL_WRITE_FAIL)
		return
	}

	tx, err := r.CtrCaller.SubmitAlgorithm(ctx, common.HexToAddress(req.ScientistWallet), cid, req.Dataset)
	if err != nil {
		dbtx.Rollback()
		log.Printf("Failed to submit algorithm to blockchain: %v", err)
		responser.ResponseError(c, bizcode.ETHEREUM_CALL_FAIL)
		return
	}

	txHash := tx.Hash().Hex()
	// Create blockchain transaction record
	_, err = models.CreateTransaction(dbtx, txHash, algo.ID, models.ENTITY_TYPE_ALGO)
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

	responser.ResponseData(c, txHash)
}

func (r *AlgoResource) ListHandler(c *gin.Context) {
	page, pageSize := parsePageParams(c)
	algos, total, err := models.GetConfirmedAlgos(r.MysqlDb, page, pageSize)
	if err != nil {
		log.Printf("Failed to list confirmed algos: %v", err)
		responser.ResponseError(c, bizcode.MYSQL_READ_FAIL)
		return
	}
	responser.ResponseList(c, page, pageSize, total, algos)
}

func (r *AlgoResource) TakeHandler(c *gin.Context) {
	var id uint
	if err := parseUintParam(c.Param("id"), &id); err != nil {
		log.Printf("Failed to parse id param: %v", err)
		responser.ResponseError(c, bizcode.BAD_REQUEST)
		return
	}

	algoDetails, err := models.GetAlgoDetailsByID(r.MysqlDb, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Printf("Algo %v not found", id)
			responser.ResponseError(c, bizcode.NOT_FOUND)
			return
		}
		log.Printf("Failed to get algo %d: %v", id, err)
		responser.ResponseError(c, bizcode.MYSQL_READ_FAIL)
		return
	}

	responser.ResponseData(c, algoDetails)
}

// func (r *AlgoResource) DeleteHandler(c *gin.Context) {
// 	id := c.Param("id")
// 	c.JSON(http.StatusOK, gin.H{"message": "Delete", "id": id})
// }
