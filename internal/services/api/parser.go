package api

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

// parseUintParam parses a string into a uint and stores the result in the provided pointer.
func parseUintParam(s string, out *uint) error {
	id64, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		return err
	}
	*out = uint(id64)
	return nil
}

// parsePageParams extracts pagination parameters from the query string,
// defaulting to page=1 and page_size=10 if not provided or invalid.
func parsePageParams(c *gin.Context) (int, int) {
	page := 1
	pageSize := 10
	if p := c.Query("page"); p != "" {
		if val, err := strconv.Atoi(p); err == nil && val > 0 {
			page = val
		}
	}
	if ps := c.Query("page_size"); ps != "" {
		if val, err := strconv.Atoi(ps); err == nil && val > 0 {
			pageSize = val
		}
	}
	return page, pageSize
}
