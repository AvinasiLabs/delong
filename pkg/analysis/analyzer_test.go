package analysis

import (
	"context"
	"delong/internal"
	"testing"
)

func TestAnalyzeFileIntegration(t *testing.T) {
	config, err := internal.LoadConfigFromEnv()
	if err != nil {
		t.Fatalf("Failed to load config: %v", err)
	}
	t.Logf("api: %v", config.DiagnosticSrvEndpoint)
	ctx := context.Background()
	filePath := "../../assets/diagnostic/blood_report.png"

	client := NewReportAnalyzer(config.DiagnosticSrvEndpoint)

	raw, err := client.AnalyzeFile(ctx, filePath)
	if err != nil {
		t.Fatalf("AnalyzeFile failed: %v", err)
	}

	t.Logf("AnalyzeFile returned JSON: %s", string(raw))
}

func TestAnalyzeFileWithNonexistentFile(t *testing.T) {
	config, err := internal.LoadConfigFromEnv()
	if err != nil {
		t.Fatalf("Failed to load config: %v", err)
	}

	ctx := context.Background()
	filePath := "nonexistent_file.pdf"

	client := NewReportAnalyzer(config.DiagnosticSrvEndpoint)

	_, err = client.AnalyzeFile(ctx, filePath)
	if err == nil {
		t.Fatal("Expected error for nonexistent file, but got nil")
	}

	t.Logf("Expected error for nonexistent file: %v", err)
}
