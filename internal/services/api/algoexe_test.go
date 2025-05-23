package api

import (
	"bytes"
	"delong/internal/models"
	"delong/internal/types"
	"delong/pkg/bizcode"
	"delong/pkg/responser"
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"testing"
	"time"
)

const (
	TEST_ALGO_LINK        = "https://codeload.github.com/lilhammer111/algo-demo/tar.gz/c73e8d62a0ae5d68040cabb461c7b51b7630020c"
	TEST_SCIENTIST_WALLET = "0xa0Ee7A142d267C1f36714E4a8F75612F20a79720"
	TEST_DATASET_NAME     = "test-1747851424415680000"
)

func TestAlgoExeCreateAndTake(t *testing.T) {
	req := types.SubmitAlgoExeReq{
		AlgoLink:        TEST_ALGO_LINK,
		ScientistWallet: TEST_SCIENTIST_WALLET,
		Dataset:         TEST_DATASET_NAME,
	}
	body, _ := json.Marshal(req)

	resp, err := http.Post(TEST_BASE_URL+"/algoexes", "application/json", bytes.NewBuffer(body))
	if err != nil {
		t.Fatalf("POST /algoexes failed: %v", err)
	}
	defer resp.Body.Close()

	respBody, _ := io.ReadAll(resp.Body)
	var apiResp responser.Response
	_ = json.Unmarshal(respBody, &apiResp)
	if apiResp.Code != bizcode.SUCCESS {
		t.Fatalf("Algo create failed: %v", apiResp.Code)
	}

	txHash, ok := apiResp.Data.(string)
	if !ok {
		t.Fatalf("Expected txHash string, got %T", apiResp.Data)
	}

	msg := waitForWsConfirmation(t, txHash, 20*time.Second)

	var wsResp responser.ResponseRaw
	_ = json.Unmarshal(msg, &wsResp)
	if wsResp.Code != bizcode.SUCCESS {
		t.Fatalf("Algo TX failed on-chain: %v", wsResp.Code)
	}

	var tx models.BlockchainTransaction
	_ = json.Unmarshal(wsResp.Data, &tx)

	getResp, err := http.Get(TEST_BASE_URL + "/algoexes/" + strconv.Itoa(int(tx.EntityID)))
	if err != nil {
		t.Fatalf("GET /algoexes/:id failed: %v", err)
	}
	defer getResp.Body.Close()
	body, _ = io.ReadAll(getResp.Body)

	var takeResp responser.ResponseRaw
	_ = json.Unmarshal(body, &takeResp)
	if takeResp.Code != bizcode.SUCCESS {
		t.Fatalf("Take algo failed: %v", takeResp.Code)
	}

	t.Logf("Algo Take response: %s", body)
}

func TestAlgoExeList(t *testing.T) {
	resp, err := http.Get(TEST_BASE_URL + "/algoexes?page=1&page_size=10")
	if err != nil {
		t.Fatalf("GET /algoexes failed: %v", err)
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	t.Logf("List response: %s", body)

	var apiResp responser.ResponseRaw
	_ = json.Unmarshal(body, &apiResp)
	if apiResp.Code != bizcode.SUCCESS {
		t.Fatalf("List failed with code: %v", apiResp.Code)
	}

	var result struct {
		Items    []models.Algo `json:"items"`
		Total    int           `json:"total"`
		Page     int           `json:"page"`
		PageSize int           `json:"page_size"`
	}
	_ = json.Unmarshal(apiResp.Data, &result)

	if result.Page != 1 || result.PageSize != 10 {
		t.Errorf("Unexpected page info: %+v", result)
	}
	t.Logf("Listed %d algos", len(result.Items))
}
