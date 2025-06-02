package api

import (
	"delong/internal/models"
	"delong/pkg/bizcode"
	"delong/pkg/responser"
	"log"

	"github.com/gin-gonic/gin"
)

type ContractMetaResource struct{ ApiServiceOptions }

func (r *ContractMetaResource) ListHandler(c *gin.Context) {
	metas, err := models.GetContracts(r.MysqlDb)
	if err != nil {
		log.Printf("Failed to list all contracts: %v", err)
		responser.ResponseError(c, bizcode.MYSQL_READ_FAIL)
		return
	}
	responser.ResponseData(c, metas)
}
