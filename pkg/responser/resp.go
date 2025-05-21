package responser

import (
	"delong/pkg/bizcode"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code bizcode.Code `json:"code,omitempty"`
	Data any          `json:"data,omitempty"`
}

type ResponseRaw struct {
	Code bizcode.Code    `json:"code"`
	Data json.RawMessage `json:"data,omitempty"` // delay parse
}

func ResponseOk(c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		Code: bizcode.SUCCESS,
	})
}

func ResponseData(c *gin.Context, data any) {
	c.JSON(http.StatusOK, Response{
		Code: bizcode.SUCCESS,
		Data: data,
	})
}

func ResponseError(c *gin.Context, code bizcode.Code) {
	c.JSON(http.StatusOK, Response{
		Code: code,
	})
}

func ResponseList(c *gin.Context, page, pageSize int, total int64, items any) {
	ResponseData(c, gin.H{
		"items":     items,
		"total":     total,
		"page":      page,
		"page_size": pageSize,
	})
}
