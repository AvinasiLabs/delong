package api

import "testing"

func TestExtractRepoName(t *testing.T) {
	repo, err := extractRepoName(TEST_ALGO_LINK)
	if err != nil {
		t.Errorf("Failed to extract repo name: %v", err)
	}
	expected := "lilhammer111/algo-demo"
	if repo != expected {
		t.Errorf("expected %v, got %v", expected, repo)
	}
}
