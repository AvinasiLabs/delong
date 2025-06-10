//go:build integration
// +build integration

package api

import (
	"delong/internal/models"
	"delong/pkg/bizcode"
	"delong/pkg/responser"
	"encoding/json"
	"io"
	"net/http"
	"testing"
)

func TestContractList(t *testing.T) {
	reqUrl := TEST_BASE_URL + "/contracts"

	resp, err := http.Get(reqUrl)
	if err != nil {
		t.Fatalf("GET /contracts failed: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("Failed to read response body: %v", err)
	}

	t.Logf("Contract list response: %s", body)

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("Expected status 200, got %d", resp.StatusCode)
	}

	var apiResp responser.ResponseRaw
	if err := json.Unmarshal(body, &apiResp); err != nil {
		t.Fatalf("Failed to unmarshal response: %v", err)
	}
	if apiResp.Code != bizcode.SUCCESS {
		t.Fatalf("Expected SUCCESS, got %v", apiResp.Code)
	}

	var contracts []models.ContractMeta
	if err := json.Unmarshal(apiResp.Data, &contracts); err != nil {
		t.Fatalf("Failed to parse contract list: %v", err)
	}

	t.Logf("Fetched %d contracts", len(contracts))

	// Verify contract structure if any contracts exist
	if len(contracts) > 0 {
		contract := contracts[0]
		if contract.Name == "" {
			t.Errorf("Contract name should not be empty")
		}
		if contract.Address == "" {
			t.Errorf("Contract address should not be empty")
		}
		if contract.ID == 0 {
			t.Errorf("Contract ID should not be zero")
		}
	}
}
