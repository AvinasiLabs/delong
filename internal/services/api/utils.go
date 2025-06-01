package api

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"

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

// extractRepoName extracts the "owner/repo" portion from a github repository download URL.
func extractRepoName(link string) (string, error) {
	parsed, err := url.Parse(link)
	if err != nil {
		return "", fmt.Errorf("invalid URL: %w", err)
	}

	// Example path: /lilhammer111/algo-demo/tar.gz/c73e8d62...
	parts := strings.Split(strings.Trim(parsed.Path, "/"), "/")
	if len(parts) < 2 {
		return "", fmt.Errorf("unexpected path structure: %s", parsed.Path)
	}

	return parts[0] + "/" + parts[1], nil
}

func isAdmin(c *gin.Context) (bool, error) {
	role, exist := GetRole(c)
	if !exist {
		return false, fmt.Errorf("failed to get jwt payload of role")
	}

	return role == "admin", nil
}
