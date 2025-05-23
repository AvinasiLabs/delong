package api

import (
	"delong/internal/models"
	"delong/internal/types"
	"delong/pkg/bizcode"
	"delong/pkg/responser"
	"errors"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type DatasetResource struct{ ApiServiceOptions }

func (r *DatasetResource) CreateHandler(c *gin.Context) {
	req := types.DatasetCreateReq{}
	err := c.ShouldBind(&req)
	if err != nil {
		log.Printf("Failed to bind request: %v", err)
		responser.ResponseError(c, bizcode.BAD_REQUEST)
		return
	}

	dataset, err := models.CreateDataset(r.MysqlDb, req.Name, req.UiName, req.Description)
	if err != nil {
		log.Printf("Failed to create dataset in db: %v", err)
		responser.ResponseError(c, bizcode.MYSQL_WRITE_FAIL)
		return
	}

	responser.ResponseData(c, dataset)
}

func (r *DatasetResource) ListHandler(c *gin.Context) {
	page, pageSize := parsePageParams(c)
	datasets, total, err := models.GetDatasets(r.MysqlDb, page, pageSize)
	if err != nil {
		log.Printf("Failed to list datasets: %v", err)
		responser.ResponseError(c, bizcode.MYSQL_READ_FAIL)
		return
	}
	responser.ResponseList(c, page, pageSize, total, datasets)
}

func (r *DatasetResource) TakeHandler(c *gin.Context) {
	var id uint
	if err := parseUintParam(c.Param("id"), &id); err != nil {
		log.Printf("Failed to parse id param: %v", err)
		responser.ResponseError(c, bizcode.BAD_REQUEST)
		return
	}

	dataset, err := models.GetDatasetByID(r.MysqlDb, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Printf("Dataset %v not found", id)
			responser.ResponseError(c, bizcode.NOT_FOUND)
			return
		}
		log.Printf("Failed to get dataset: %v", err)
		responser.ResponseError(c, bizcode.MYSQL_READ_FAIL)
		return
	}

	responser.ResponseData(c, dataset)
}

func (r *DatasetResource) UpdateHandler(c *gin.Context) {
	var id uint
	if err := parseUintParam(c.Param("id"), &id); err != nil {
		responser.ResponseError(c, bizcode.BAD_REQUEST)
		return
	}

	req := types.DatasetUpdateReq{}
	if err := c.ShouldBind(&req); err != nil {
		responser.ResponseError(c, bizcode.BAD_REQUEST)
		return
	}

	updated, err := models.UpdateDataset(r.MysqlDb, id, req.UiName, req.Description)
	if err != nil {
		log.Printf("Failed to update dataset: %v", err)
		responser.ResponseError(c, bizcode.MYSQL_WRITE_FAIL)
		return
	}

	responser.ResponseData(c, updated)
}

func (r *DatasetResource) DeleteHandler(c *gin.Context) {
	var id uint
	if err := parseUintParam(c.Param("id"), &id); err != nil {
		responser.ResponseError(c, bizcode.BAD_REQUEST)
		return
	}

	err := models.DeleteDataset(r.MysqlDb, id)
	if err != nil {
		log.Printf("Failed to delete dataset: %v", err)
		responser.ResponseError(c, bizcode.MYSQL_WRITE_FAIL)
		return
	}
	responser.ResponseOk(c)
}
