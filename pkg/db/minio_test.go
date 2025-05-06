package db

import (
	"context"
	"delong/internal"
	"os"
	"testing"
)

var testBucket = "testbucket"

func setupStore(t *testing.T) (*MinioStore, error) {
	t.Helper()

	config, err := internal.LoadConfigFromEnv()
	if err != nil {
		t.Fatalf("failed to load config from env: %v", err)
	}

	store, err := NewMinioStore(config.MinioEndpoint,
		WithCredentials(config.MinioAccessKey, config.MinioSecretKey),
		WithSecure(false),
	)
	if err != nil {
		t.Fatalf("failed to create minio store: %v", err)
	}
	return store, nil
}

func TestMinioStore_List(t *testing.T) {
	ctx := context.Background()
	store, err := setupStore(t)
	if err != nil {
		t.Fatalf("failed to setup store: %v", err)
	}
	buckets, err := store.client.ListBuckets(ctx)
	if err != nil {
		t.Fatalf("failed to list buckets: %v", err)
	}
	for _, bucket := range buckets {
		t.Logf("bucket: %s", bucket.Name)
	}
}

func TestMinioStore_CreateBucket(t *testing.T) {
	ctx := context.Background()
	store, err := setupStore(t)
	if err != nil {
		t.Fatalf("failed to setup store: %v", err)
	}

	err = store.CreateBucket(ctx, testBucket)
	if err != nil {
		t.Fatalf("failed to create bucket: %v", err)
	}
}

func TestMinioStore_Upload(t *testing.T) {
	ctx := context.Background()
	store, err := setupStore(t)
	if err != nil {
		t.Fatalf("failed to setup store: %v", err)
	}

	localFilePath := "test_upload.txt"
	content := []byte("Hello MinIO Upload Test!")
	if err := os.WriteFile(localFilePath, content, 0644); err != nil {
		t.Fatalf("failed to write local file: %v", err)
	}
	defer os.Remove(localFilePath)

	objectName := "test/test_upload.txt"

	if err := store.Upload(ctx, testBucket, objectName, localFilePath, "text/plain"); err != nil {
		t.Fatalf("failed to upload file: %v", err)
	}
}

func TestMinioStore_Download(t *testing.T) {
	ctx := context.Background()
	store, err := setupStore(t)
	if err != nil {
		t.Fatalf("failed to setup store: %v", err)
	}

	objectName := "test/test_upload.txt"
	downloadedFilePath := "test_download.txt"
	defer os.Remove(downloadedFilePath)

	if err := store.Download(ctx, testBucket, objectName, downloadedFilePath); err != nil {
		t.Fatalf("failed to download file: %v", err)
	}

	info, err := os.Stat(downloadedFilePath)
	if os.IsNotExist(err) {
		t.Fatalf("downloaded file does not exist: %v", err)
	}

	t.Logf("downloaded file size: %d", info.Size())
}
