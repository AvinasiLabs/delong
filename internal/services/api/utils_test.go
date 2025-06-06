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

func TestBuildGitHubDownloadUrl(t *testing.T) {
	tests := []struct {
		name        string
		repoUrl     string
		commitSha   string
		expectedUrl string
		expectedRepo string
		expectError bool
	}{
		{
			name:        "valid github url",
			repoUrl:     "https://github.com/lilhammer111/algo-demo",
			commitSha:   "c73e8d62a0ae5d68040cabb461c7b51b7630020c",
			expectedUrl: "https://codeload.github.com/lilhammer111/algo-demo/tar.gz/c73e8d62a0ae5d68040cabb461c7b51b7630020c",
			expectedRepo: "lilhammer111/algo-demo",
			expectError: false,
		},
		{
			name:        "another valid github url",
			repoUrl:     "https://github.com/owner/repo-name",
			commitSha:   "abc123def456",
			expectedUrl: "https://codeload.github.com/owner/repo-name/tar.gz/abc123def456",
			expectedRepo: "owner/repo-name",
			expectError: false,
		},
		{
			name:        "invalid url format",
			repoUrl:     "not-a-valid-url",
			commitSha:   "abc123",
			expectedUrl: "",
			expectedRepo: "",
			expectError: true,
		},
		{
			name:        "url with insufficient path parts",
			repoUrl:     "https://github.com/onlyowner",
			commitSha:   "abc123",
			expectedUrl: "",
			expectedRepo: "",
			expectError: true,
		},
		{
			name:        "empty repo url",
			repoUrl:     "",
			commitSha:   "abc123",
			expectedUrl: "",
			expectedRepo: "",
			expectError: true,
		},
		{
			name:        "url with trailing slash",
			repoUrl:     "https://github.com/owner/repo/",
			commitSha:   "abc123",
			expectedUrl: "https://codeload.github.com/owner/repo/tar.gz/abc123",
			expectedRepo: "owner/repo",
			expectError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			downloadUrl, repoName, err := buildGitHubDownloadUrl(tt.repoUrl, tt.commitSha)
			
			if tt.expectError {
				if err == nil {
					t.Errorf("expected error but got none")
				}
				return
			}
			
			if err != nil {
				t.Errorf("unexpected error: %v", err)
				return
			}
			
			if downloadUrl != tt.expectedUrl {
				t.Errorf("expected download URL %q, got %q", tt.expectedUrl, downloadUrl)
			}
			
			if repoName != tt.expectedRepo {
				t.Errorf("expected repo name %q, got %q", tt.expectedRepo, repoName)
			}
		})
	}
}

func TestNormalizeFileName(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		expected    string
		expectError bool
	}{
		{
			name:        "uppercase with space",
			input:       "BLOOD TEST",
			expected:    "blood_test",
			expectError: false,
		},
		{
			name:        "multiple spaces",
			input:       "BLOOD   TEST   REPORT",
			expected:    "blood_test_report",
			expectError: false,
		},
		{
			name:        "already lowercase",
			input:       "blood test",
			expected:    "blood_test",
			expectError: false,
		},
		{
			name:        "single word",
			input:       "GLUCOSE",
			expected:    "glucose",
			expectError: false,
		},
		{
			name:        "with leading/trailing spaces",
			input:       "  BLOOD TEST  ",
			expected:    "blood_test",
			expectError: false,
		},
		{
			name:        "empty string",
			input:       "",
			expected:    "",
			expectError: true,
		},
		{
			name:        "only spaces",
			input:       "   ",
			expected:    "",
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := normalizeFileName(tt.input)
			
			if tt.expectError {
				if err == nil {
					t.Errorf("expected error but got none")
				}
				return
			}
			
			if err != nil {
				t.Errorf("unexpected error: %v", err)
				return
			}
			
			if result != tt.expected {
				t.Errorf("expected %q, got %q", tt.expected, result)
			}
		})
	}
}
