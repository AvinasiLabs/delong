package api

import (
	"delong/internal/models"
	"delong/internal/types"
	"delong/pkg/bizcode"
	"delong/pkg/responser"
	"delong/pkg/tee"
	"errors"
	"log"
	"path/filepath"

	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type StaticDatasetResource struct{ ApiServiceOptions }

func (r *StaticDatasetResource) CreateHandler(c *gin.Context) {
	req := types.StcDatasetCreateReq{}
	if err := c.ShouldBind(&req); err != nil {
		log.Printf("Failed to bind request: %v", err)
		responser.ResponseError(c, bizcode.BAD_REQUEST)
		return
	}

	file, header, err := c.Request.FormFile("file")
	if err != nil {
		log.Printf("Failed to get uploaded static dataset file: %v", err)
		responser.ResponseError(c, bizcode.BAD_REQUEST)
		return
	}
	defer file.Close()

	// Calculate hash of original file for deduplication
	hash, err := hashSha256(file)
	if err != nil {
		log.Printf("Failed to calculate file hash: %v", err)
		responser.ResponseError(c, bizcode.INTERNAL_SERVER_ERROR)
		return
	}
	_, err = models.GetStcDatasetByHash(r.MysqlDb, hash)
	if err == nil { // file exists
		log.Printf("File with hash %v already exists", hash)
		responser.ResponseError(c, bizcode.STATIC_DATASET_EXIST)
		return
	}

	// Upload to ipfs
	ctx := c.Request.Context()
	kc := tee.NewKeyContext(tee.KEYKIND_ENC_KEY, req.Author, "encrypt static dataset")
	key, err := r.KeyVault.DeriveSymmetricKey(ctx, kc)
	if err != nil {
		log.Printf("Failed to derive symmetric key: %v", err)
		responser.ResponseError(c, bizcode.KEY_DERIVE_FAIL)
		return
	}

	cid, err := r.IpfsStore.UploadEncryptedStream(ctx, file, key)
	if err != nil {
		log.Printf("Failed to upload to IPFS: %v", err)
		responser.ResponseError(c, bizcode.INTERNAL_SERVER_ERROR)
		return
	}

	// Get file format
	fileFormat := filepath.Ext(header.Filename)
	if fileFormat != "" {
		fileFormat = fileFormat[1:] // remove the dot
	}

	// TODO: Get dataset sample url

	dbtx := r.MysqlDb.Begin()
	if dbtx.Error != nil {
		log.Printf("Failed to begin transaction: %v", dbtx.Error)
		responser.ResponseError(c, bizcode.MYSQL_WRITE_FAIL)
		return
	}

	// Create static dataset
	createReq := models.CreateStcDatasetReq{
		Name:         req.Name,
		Desc:         req.Desc,
		FileHash:     hash,
		IpfsCid:      cid,
		FileSize:     header.Size,
		FileFormat:   fileFormat,
		Author:       req.Author,
		AuthorWallet: req.AuthorWallet,
		SampleUrl:    "",
	}

	dataset, err := models.CreateStcDataset(dbtx, createReq)
	if err != nil {
		dbtx.Rollback()
		log.Printf("Failed to create data asset: %v", err)
		responser.ResponseError(c, bizcode.MYSQL_WRITE_FAIL)
		return
	}

	tx, err := r.CtrCaller.RegisterData(ctx, common.HexToAddress(dataset.AuthorWallet), dataset.IpfsCid, dataset.Name)
	if err != nil {
		dbtx.Rollback()
		log.Printf("Failed to register static dataset: %v", err)
		responser.ResponseError(c, bizcode.MYSQL_WRITE_FAIL)
		return
	}
	txHash := tx.Hash().Hex()
	_, err = models.CreateTransaction(dbtx, txHash, uint(dataset.ID), models.ENTITY_TYPE_STATIC_DATASET)
	if err != nil {
		dbtx.Rollback()
		log.Printf("Failed to create transaction: %v", err)
		responser.ResponseError(c, bizcode.MYSQL_WRITE_FAIL)
		return
	}

	if err := dbtx.Commit().Error; err != nil {
		log.Printf("Failed to commit transaction: %v", err)
		responser.ResponseError(c, bizcode.MYSQL_WRITE_FAIL)
		return
	}

	responser.ResponseData(c, txHash)
}

func (r *StaticDatasetResource) ListHandler(c *gin.Context) {
	page, pageSize := parsePageParams(c)
	assets, total, err := models.GetStcDataset(r.MysqlDb, page, pageSize)
	if err != nil {
		log.Printf("Failed to list data assets: %v", err)
		responser.ResponseError(c, bizcode.MYSQL_READ_FAIL)
		return
	}
	responser.ResponseList(c, page, pageSize, total, assets)
}

func (r *StaticDatasetResource) TakeHandler(c *gin.Context) {
	var id uint
	if err := parseUintParam(c.Param("id"), &id); err != nil {
		responser.ResponseError(c, bizcode.BAD_REQUEST)
		return
	}

	asset, err := models.GetStcDatasetByID(r.MysqlDb, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			responser.ResponseError(c, bizcode.NOT_FOUND)
			return
		}
		responser.ResponseError(c, bizcode.MYSQL_READ_FAIL)
		return
	}

	responser.ResponseData(c, asset)
}
