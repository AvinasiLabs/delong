package analysis

import (
	"context"
	"delong/internal"
	"delong/pkg/db"
	"fmt"
	"path/filepath"
	"testing"
)

func TestAnalyzeIntegration(t *testing.T) {
	config, err := internal.LoadConfigFromEnv()
	if err != nil {
		t.Fatalf("Failed to load config: %v", err)
	}

	store, err := db.NewMinioStore(config.MinioEndpoint, db.WithSecure(false), db.WithCredentials(config.MinioAccessKey, config.MinioSecretKey))
	if err != nil {
		t.Fatalf("Failed to create Minio store: %v", err)
	}

	ctx := context.Background()
	filePath := "../../assets/diagnostic/blood_report.png"
	fileName := filepath.Base(filePath)
	objectName := fmt.Sprintf("/v1/1/original/%s", fileName)

	err = store.Upload(ctx, "diagnostic", objectName, filePath, "image/png")
	if err != nil {
		t.Fatalf("Failed to upload file to Minio: %v", err)
	}

	t.Logf("Uploaded file to Minio: %s", objectName)

	client := NewReportAnalyzer(config.DiagnosticSrvEndpoint)

	raw, err := client.Analyze(ctx,
		"minio",
		"image/png",
		objectName,
		"1",
	)
	if err != nil {
		t.Fatalf("Analyze failed: %v", err)
	}

	t.Logf("Analyze returned JSON: %s", string(raw))
}
