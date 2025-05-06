package db

import (
	"context"
	"os"
	"testing"
)

func TestIpfsClient_AddAndGet(t *testing.T) {
	api_addr := os.Getenv("IPFS_ADDR")
	t.Logf("IPFS_ADDR: %v", api_addr)

	// addr, err := ma.NewMultiaddr(api_addr)
	// if err != nil {
	// 	t.Fatalf("Failed to new multi addr: %v", err)
	// }

	// api, err := rpc.NewApi(addr)
	// if err != nil {
	// 	t.Skip("Skipping test as local IPFS node is not available: ", err)
	// }

	// client := NewIpfsClient(api)

	client, err := NewIpfsStore(api_addr)
	if err != nil {
		t.Skip("Skipping test as local IPFS node is not available: ", err)
	}

	ctx := context.Background()

	originalData := []byte("This is a test file for IpfsClient unit test.")

	cidStr, err := client.Upload(ctx, originalData)
	if err != nil {
		t.Fatalf("Failed to add file to IPFS: %v", err)
	}

	t.Logf("cid: %v", cidStr)
	if cidStr == "" {
		t.Fatalf("Empty CID returned from Add")
	}

	retrievedData, err := client.Download(ctx, cidStr)
	if err != nil {
		t.Fatalf("Failed to get file from IPFS: %v", err)
	}

	if string(retrievedData) != string(originalData) {
		t.Fatalf("Retrieved data does not match. Expected: %q, got: %q", originalData, retrievedData)
	}
}
