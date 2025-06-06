package api

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
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

// buildGitHubDownloadUrl constructs a GitHub archive download URL and extracts the repo name.
// Returns the download URL and repo name in "owner/repo" format.
// Example repoUrl: https://github.com/lilhammer111/algo-demo
// Example sha: c73e8d62a0ae5d68040cabb461c7b51b7630020c
// Example result: https://codeload.github.com/lilhammer111/algo-demo/tar.gz/c73e8d62a0ae5d68040cabb461c7b51b7630020c
func buildGitHubDownloadUrl(repoUrl string, commitSha string) (string, string, error) {
	// Parse the repository URL to extract owner/repo
	parsed, err := url.Parse(repoUrl)
	if err != nil {
		return "", "", err
	}

	// Extract owner/repo from path (e.g., "/lilhammer111/algo-demo")
	path := strings.Trim(parsed.Path, "/")
	parts := strings.Split(path, "/")
	if len(parts) < 2 {
		return "", "", fmt.Errorf("invalid repository url")
	}

	owner := parts[0]
	repo := parts[1]
	repoName := fmt.Sprintf("%s/%s", owner, repo)

	// Build the download URL
	downloadUrl := fmt.Sprintf("https://codeload.github.com/%s/%s/tar.gz/%s", owner, repo, commitSha)

	return downloadUrl, repoName, nil
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

// hashSha256 reads all data from rs, computes its SHA-256 hash,
// resets rs back to the start, and returns the hex-encoded digest.
// rs must implement io.ReadSeeker so that it can be rewound.
func hashSha256(rs io.ReadSeeker) (string, error) {
	// Create a new SHA-256 hasher
	hasher := sha256.New()

	// Read entire content into the hasher
	if _, err := io.Copy(hasher, rs); err != nil {
		return "", err
	}

	// Convert the hash sum to a hex string
	hashBytes := hasher.Sum(nil)
	hashStr := hex.EncodeToString(hashBytes)

	// Rewind the reader so it can be read again later
	if _, err := rs.Seek(0, io.SeekStart); err != nil {
		return "", err
	}

	return hashStr, nil
}

// filePath converts a file name to a standardized format: lowercase with underscores and .csv extension.
// Removes all spaces, converts to lowercase, and replaces spaces with underscores.
// Example: "BLOOD TEST" -> "blood_test.csv"
func filePath(fileName string) (string, error) {
	// Trim spaces first
	cleaned := strings.TrimSpace(fileName)
	if cleaned == "" {
		return "", fmt.Errorf("fileName cannot be empty")
	}
	
	// Convert to lowercase
	result := strings.ToLower(cleaned)
	
	// Replace multiple consecutive spaces with single underscore
	for strings.Contains(result, "  ") {
		result = strings.ReplaceAll(result, "  ", " ")
	}
	result = strings.ReplaceAll(result, " ", "_")
	
	// Add .csv extension
	return result + ".csv", nil
}
