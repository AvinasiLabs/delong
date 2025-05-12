package resp

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// ErrCode is a custom string type representing application error codes.
type ErrCode string

// General/common errors
const (
	BAD_REQUEST           ErrCode = "BAD_REQUEST"           // Invalid request parameters
	UNAUTHORIZED          ErrCode = "UNAUTHORIZED"          // Authentication required
	FORBIDDEN             ErrCode = "FORBIDDEN"             // Permission denied
	NOT_FOUND             ErrCode = "NOT_FOUND"             // Resource not found
	TOO_MANY_REQUESTS     ErrCode = "TOO_MANY_REQUESTS"     // Rate limit exceeded
	INTERNAL_SERVER_ERROR ErrCode = "INTERNAL_SERVER_ERROR" // Unhandled server error
)

// Database errors
const (
	DB_WRITE_FAIL ErrCode = "DB_WRITE_FAIL"      // Failed to write to database
	DB_READ_FAIL  ErrCode = "DB_READ_FAIL"       // Failed to query database
	DB_DUPLICATE  ErrCode = "DB_DUPLICATE_ENTRY" // Duplicate record found
)

var ErrMsg = map[ErrCode]string{}

type ApiResp struct {
	ErrCode ErrCode `json:"err_code"`
	Message string  `json:"message"`
	Data    any     `json:"data,omitempty"`
}

func Success(c *gin.Context, data any) {
	c.JSON(http.StatusOK, ApiResp{
		Data:    data,
		Message: "OK",
	})
}

func Fail(c *gin.Context, errCode ErrCode) {
	c.JSON(http.StatusOK, ApiResp{
		ErrCode: errCode,
		Message: ErrMsg[errCode],
	})
}

func FailWithMsg(c *gin.Context, errCode ErrCode, message string) {
	c.JSON(http.StatusOK, ApiResp{
		ErrCode: errCode,
		Message: message,
	})
}
