//go:build integration
// +build integration

package db

import (
	"bytes"
	"context"
	"crypto/rand"
	"delong/internal"
	"io"
	"testing"
)

const (
	testDataSize = 1024 * 1024 // 1MB test data
	testCSVData  = `name,age,city,salary
John Doe,30,New York,50000
Jane Smith,25,Los Angeles,60000
Bob Johnson,35,Chicago,45000
Alice Brown,28,Houston,55000
Charlie Wilson,32,Phoenix,48000`
)

func TestIPFSEncryptDecryptIntegration(t *testing.T) {
	// Skip if not running integration tests
	if testing.Short() {
		t.Skip("Skipping integration test in short mode")
	}

	// Load config from environment variables like in main.go
	config, err := internal.LoadConfigFromEnv()
	if err != nil {
		t.Skipf("Failed to load config from environment: %v (make sure IPFS_API_ADDR is set)", err)
	}

	// Initialize IPFS store
	ipfsStore, err := NewIpfsStore(config.IpfsApiAddr)
	if err != nil {
		t.Skipf("Failed to create IPFS store: %v (make sure IPFS node is running)", err)
	}

	ctx := context.Background()

	t.Run("EncryptDecryptSmallData", func(t *testing.T) {
		testEncryptDecryptRoundTrip(t, ctx, ipfsStore, []byte(testCSVData))
	})

	t.Run("EncryptDecryptLargeData", func(t *testing.T) {
		// Generate large test data
		largeData := generateTestData(testDataSize)
		testEncryptDecryptRoundTrip(t, ctx, ipfsStore, largeData)
	})

	t.Run("EncryptDecryptStream", func(t *testing.T) {
		testEncryptDecryptStreamRoundTrip(t, ctx, ipfsStore, []byte(testCSVData))
	})

	t.Run("EncryptDecryptEmptyData", func(t *testing.T) {
		testEncryptDecryptRoundTrip(t, ctx, ipfsStore, []byte(""))
	})

	t.Run("EncryptDecryptBinaryData", func(t *testing.T) {
		// Generate random binary data
		binaryData := make([]byte, 1024)
		_, err := rand.Read(binaryData)
		if err != nil {
			t.Fatalf("Failed to generate random data: %v", err)
		}
		testEncryptDecryptRoundTrip(t, ctx, ipfsStore, binaryData)
	})
}

func testEncryptDecryptRoundTrip(t *testing.T, ctx context.Context, ipfsStore *IpfsStore, originalData []byte) {
	// Generate a random encryption key
	key := make([]byte, 32) // 256-bit key for AES
	_, err := rand.Read(key)
	if err != nil {
		t.Fatalf("Failed to generate encryption key: %v", err)
	}

	// Test UploadEncryptedStream + DownloadDecrypted
	t.Run("StreamUpload_ByteDownload", func(t *testing.T) {
		// Upload encrypted data
		reader := bytes.NewReader(originalData)
		cid, err := ipfsStore.UploadEncryptedStream(ctx, reader, key)
		if err != nil {
			t.Fatalf("Failed to upload encrypted data: %v", err)
		}

		if cid == "" {
			t.Fatal("Expected non-empty CID")
		}

		// Download and decrypt data
		decryptedData, err := ipfsStore.DownloadDecrypted(ctx, cid, key)
		if err != nil {
			t.Fatalf("Failed to download and decrypt data: %v", err)
		}

		// Verify the data matches
		if !bytes.Equal(originalData, decryptedData) {
			t.Errorf("Decrypted data doesn't match original data")
			t.Errorf("Original length: %d, Decrypted length: %d", len(originalData), len(decryptedData))
			if len(originalData) < 200 && len(decryptedData) < 200 {
				t.Errorf("Original: %q", string(originalData))
				t.Errorf("Decrypted: %q", string(decryptedData))
			}
		}
	})

	// Test UploadEncryptedStream + DownloadDecryptedStream
	t.Run("StreamUpload_StreamDownload", func(t *testing.T) {
		// Upload encrypted data
		reader := bytes.NewReader(originalData)
		cid, err := ipfsStore.UploadEncryptedStream(ctx, reader, key)
		if err != nil {
			t.Fatalf("Failed to upload encrypted data: %v", err)
		}

		if cid == "" {
			t.Fatal("Expected non-empty CID")
		}

		// Download and decrypt data as stream
		decryptedStream, err := ipfsStore.DownloadDecryptedStream(ctx, cid, key)
		if err != nil {
			t.Fatalf("Failed to download and decrypt stream: %v", err)
		}
		defer decryptedStream.Close()

		// Read all data from stream
		decryptedData, err := io.ReadAll(decryptedStream)
		if err != nil {
			t.Fatalf("Failed to read from decrypted stream: %v", err)
		}

		// Verify the data matches
		if !bytes.Equal(originalData, decryptedData) {
			t.Errorf("Decrypted stream data doesn't match original data")
			t.Errorf("Original length: %d, Decrypted length: %d", len(originalData), len(decryptedData))
			if len(originalData) < 200 && len(decryptedData) < 200 {
				t.Errorf("Original: %q", string(originalData))
				t.Errorf("Decrypted: %q", string(decryptedData))
			}
		}
	})
}

func testEncryptDecryptStreamRoundTrip(t *testing.T, ctx context.Context, ipfsStore *IpfsStore, originalData []byte) {
	// Generate a random encryption key
	key := make([]byte, 32) // 256-bit key for AES
	_, err := rand.Read(key)
	if err != nil {
		t.Fatalf("Failed to generate encryption key: %v", err)
	}

	// Upload encrypted data
	reader := bytes.NewReader(originalData)
	cid, err := ipfsStore.UploadEncryptedStream(ctx, reader, key)
	if err != nil {
		t.Fatalf("Failed to upload encrypted data: %v", err)
	}

	// Download and decrypt data as stream
	decryptedStream, err := ipfsStore.DownloadDecryptedStream(ctx, cid, key)
	if err != nil {
		t.Fatalf("Failed to download and decrypt stream: %v", err)
	}
	defer decryptedStream.Close()

	// Read data in chunks to test streaming behavior
	var decryptedData []byte
	buffer := make([]byte, 256) // Small buffer to test multiple reads
	for {
		n, err := decryptedStream.Read(buffer)
		if n > 0 {
			decryptedData = append(decryptedData, buffer[:n]...)
		}
		if err == io.EOF {
			break
		}
		if err != nil {
			t.Fatalf("Failed to read from decrypted stream: %v", err)
		}
	}

	// Verify the data matches
	if !bytes.Equal(originalData, decryptedData) {
		t.Errorf("Decrypted stream data doesn't match original data")
		t.Errorf("Original length: %d, Decrypted length: %d", len(originalData), len(decryptedData))
		if len(originalData) < 200 && len(decryptedData) < 200 {
			t.Errorf("Original: %q", string(originalData))
			t.Errorf("Decrypted: %q", string(decryptedData))
		}
	}
}

func TestIPFSEncryptDecryptWrongKey(t *testing.T) {
	// Skip if not running integration tests
	if testing.Short() {
		t.Skip("Skipping integration test in short mode")
	}

	// Load config from environment variables
	config, err := internal.LoadConfigFromEnv()
	if err != nil {
		t.Skipf("Failed to load config from environment: %v", err)
	}

	// Initialize IPFS store
	ipfsStore, err := NewIpfsStore(config.IpfsApiAddr)
	if err != nil {
		t.Skipf("Failed to create IPFS store: %v", err)
	}

	ctx := context.Background()
	originalData := []byte(testCSVData)

	// Generate encryption key
	key := make([]byte, 32)
	_, err = rand.Read(key)
	if err != nil {
		t.Fatalf("Failed to generate encryption key: %v", err)
	}

	// Upload encrypted data
	reader := bytes.NewReader(originalData)
	cid, err := ipfsStore.UploadEncryptedStream(ctx, reader, key)
	if err != nil {
		t.Fatalf("Failed to upload encrypted data: %v", err)
	}

	// Try to decrypt with wrong key
	wrongKey := make([]byte, 32)
	_, err = rand.Read(wrongKey)
	if err != nil {
		t.Fatalf("Failed to generate wrong key: %v", err)
	}

	// This should fail
	_, err = ipfsStore.DownloadDecrypted(ctx, cid, wrongKey)
	if err == nil {
		t.Error("Expected decryption to fail with wrong key, but it succeeded")
	}

	// Test with stream version too
	_, err = ipfsStore.DownloadDecryptedStream(ctx, cid, wrongKey)
	if err == nil {
		t.Error("Expected stream decryption to fail with wrong key, but it succeeded")
	}
}

func TestIPFSPlaintextUploadDownload(t *testing.T) {
	// Skip if not running integration tests
	if testing.Short() {
		t.Skip("Skipping integration test in short mode")
	}

	// Load config from environment variables
	config, err := internal.LoadConfigFromEnv()
	if err != nil {
		t.Skipf("Failed to load config from environment: %v", err)
	}

	// Initialize IPFS store
	ipfsStore, err := NewIpfsStore(config.IpfsApiAddr)
	if err != nil {
		t.Skipf("Failed to create IPFS store: %v", err)
	}

	ctx := context.Background()
	originalData := []byte(testCSVData)

	// Test normal upload and download (without encryption)
	cid, err := ipfsStore.Upload(ctx, originalData)
	if err != nil {
		t.Fatalf("Failed to upload data: %v", err)
	}

	downloadedData, err := ipfsStore.Download(ctx, cid)
	if err != nil {
		t.Fatalf("Failed to download data: %v", err)
	}

	if !bytes.Equal(originalData, downloadedData) {
		t.Errorf("Downloaded data doesn't match original data")
	}

	// Test stream upload and download
	reader := bytes.NewReader(originalData)
	cid2, err := ipfsStore.UploadStream(ctx, reader)
	if err != nil {
		t.Fatalf("Failed to upload stream: %v", err)
	}

	stream, err := ipfsStore.DownloadStream(ctx, cid2)
	if err != nil {
		t.Fatalf("Failed to download stream: %v", err)
	}
	defer stream.Close()

	streamData, err := io.ReadAll(stream)
	if err != nil {
		t.Fatalf("Failed to read stream: %v", err)
	}

	if !bytes.Equal(originalData, streamData) {
		t.Errorf("Stream data doesn't match original data")
	}
}

// generateTestData creates test data of the specified size
func generateTestData(size int) []byte {
	data := make([]byte, size)
	// Fill with a repeating pattern for better compression testing
	pattern := "This is test data for IPFS encryption/decryption testing. "
	patternBytes := []byte(pattern)

	for i := 0; i < size; i++ {
		data[i] = patternBytes[i%len(patternBytes)]
	}

	return data
}

// Benchmark tests
// func BenchmarkIPFSEncryptDecrypt(b *testing.B) {
// 	if testing.Short() {
// 		b.Skip("Skipping benchmark in short mode")
// 	}

// 	config, err := internal.LoadConfigFromEnv()
// 	if err != nil {
// 		b.Skip("Failed to load config from environment")
// 	}

// 	ipfsStore, err := NewIpfsStore(config.IpfsApiAddr)
// 	if err != nil {
// 		b.Skip("Failed to create IPFS store")
// 	}

// 	ctx := context.Background()
// 	testData := generateTestData(1024) // 1KB test data
// 	key := make([]byte, 32)
// 	rand.Read(key)

// 	b.ResetTimer()

// 	for i := 0; i < b.N; i++ {
// 		// Upload
// 		reader := bytes.NewReader(testData)
// 		cid, err := ipfsStore.UploadEncryptedStream(ctx, reader, key)
// 		if err != nil {
// 			b.Fatalf("Upload failed: %v", err)
// 		}

// 		// Download
// 		_, err = ipfsStore.DownloadDecrypted(ctx, cid, key)
// 		if err != nil {
// 			b.Fatalf("Download failed: %v", err)
// 		}
// 	}
// }
