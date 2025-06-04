package api

import (
	"bytes"
	"delong/internal/models"
	"delong/pkg/bizcode"
	"delong/pkg/responser"
	"encoding/json"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"strconv"
	"testing"
	"time"
)

const TEST_AUTHOR_WALLET = "0xa0Ee7A142d267C1f36714E4a8F75612F20a79720"

func TestStcDatasetCreate(t *testing.T) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	_ = writer.WriteField("name", "test-blood-dataset")
	_ = writer.WriteField("desc", "Test blood analysis dataset for integration testing")
	_ = writer.WriteField("author", "Test Author")
	_ = writer.WriteField("author_wallet", TEST_AUTHOR_WALLET)

	// Create a test CSV file content
	csvContent := `id,test_report_id,category,name,definition,result,reference_range,explanation,status,suggestions
1,101,blood,glucose,blood glucose level,95,70-100 mg/dL,normal glucose level,normal,maintain healthy diet
2,101,blood,cholesterol,total cholesterol,180,<200 mg/dL,acceptable cholesterol level,normal,continue current lifestyle
3,102,blood,hemoglobin,hemoglobin level,14.5,12-16 g/dL,normal hemoglobin,normal,no action needed`

	part, err := writer.CreateFormFile("file", "test_dataset.csv")
	if err != nil {
		t.Fatalf("failed to create form file: %v", err)
	}
	_, err = part.Write([]byte(csvContent))
	if err != nil {
		t.Fatalf("failed to write file content: %v", err)
	}
	writer.Close()

	req, err := http.NewRequest("POST", TEST_BASE_URL+"/static-datasets", body)
	if err != nil {
		t.Fatalf("failed to create request: %v", err)
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatalf("request failed: %v", err)
	}
	defer resp.Body.Close()

	respBody, _ := io.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK {
		t.Errorf("unexpected status %d: %s", resp.StatusCode, string(respBody))
	}

	apiResp := &responser.Response{}
	json.Unmarshal(respBody, apiResp)
	if apiResp.Code != bizcode.SUCCESS {
		t.Errorf("Failed to create static dataset, code=%v", apiResp.Code)
	}

	// Get transaction hash from response
	txHash, ok := apiResp.Data.(string)
	if !ok {
		t.Fatalf("Unexpected data format: %T", apiResp.Data)
	}

	t.Logf("Got transaction hash: %s", txHash)

	// Wait for WebSocket confirmation
	msg := waitForWsConfirmation(t, txHash, 15*time.Second)
	t.Logf("Received WebSocket message: %s", string(msg))

	wsResp := responser.ResponseRaw{}
	err = json.Unmarshal(msg, &wsResp)
	if err != nil {
		t.Fatalf("Failed to unmarshal WebSocket response: %v", err)
	}
	if wsResp.Code != bizcode.SUCCESS {
		t.Fatalf("Expected WebSocket SUCCESS, got %v", wsResp.Code)
	}

	var tx models.BlockchainTransaction
	err = json.Unmarshal(wsResp.Data, &tx)
	if err != nil {
		t.Fatalf("Failed to unmarshal transaction: %v", err)
	}

	datasetId := tx.EntityID
	t.Logf("Created dataset with ID: %d", datasetId)

	// Verify the created dataset by fetching it
	getResp, err := http.Get(TEST_BASE_URL + "/static-datasets/" + strconv.Itoa(int(datasetId)))
	if err != nil {
		t.Fatalf("Failed to GET created dataset: %v", err)
	}
	defer getResp.Body.Close()

	if getResp.StatusCode != http.StatusOK {
		t.Fatalf("GET created dataset failed, status: %d", getResp.StatusCode)
	}

	getBody, err := io.ReadAll(getResp.Body)
	if err != nil {
		t.Fatalf("Failed to read GET body: %v", err)
	}

	getRespJson := responser.Response{}
	err = json.Unmarshal(getBody, &getRespJson)
	if err != nil {
		t.Fatalf("Failed to unmarshal GET response: %v", err)
	}

	if getRespJson.Code != bizcode.SUCCESS {
		t.Fatalf("GET expected SUCCESS, got %v", getRespJson.Code)
	}

	// Verify dataset properties
	datasetBytes, err := json.Marshal(getRespJson.Data)
	if err != nil {
		t.Fatalf("Failed to marshal dataset data: %v", err)
	}

	var dataset models.StaticDataset
	err = json.Unmarshal(datasetBytes, &dataset)
	if err != nil {
		t.Fatalf("Failed to unmarshal dataset: %v", err)
	}

	if dataset.Name != "test-blood-dataset" {
		t.Errorf("Expected name 'test-blood-dataset', got '%s'", dataset.Name)
	}
	if dataset.Author != "Test Author" {
		t.Errorf("Expected author 'Test Author', got '%s'", dataset.Author)
	}
	if dataset.AuthorWallet != TEST_AUTHOR_WALLET {
		t.Errorf("Expected author wallet '%s', got '%s'", TEST_AUTHOR_WALLET, dataset.AuthorWallet)
	}
	if dataset.FileFormat != "csv" {
		t.Errorf("Expected file format 'csv', got '%s'", dataset.FileFormat)
	}
}

func TestStcDatasetList(t *testing.T) {
	u, _ := url.Parse(TEST_BASE_URL + "/static-datasets")
	q := u.Query()
	q.Set("page", "1")
	q.Set("page_size", "10")
	u.RawQuery = q.Encode()

	resp, err := http.Get(u.String())
	if err != nil {
		t.Fatalf("Request failed: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status 200, got %d", resp.StatusCode)
	}

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("Failed to read response body: %v", err)
	}

	apiResp := responser.Response{}
	err = json.Unmarshal(respBody, &apiResp)
	if err != nil {
		t.Fatalf("Failed to unmarshal response body: %v", err)
	}

	t.Logf("List response: %v", apiResp)

	if apiResp.Code != bizcode.SUCCESS {
		t.Errorf("Expected CODE SUCCESS, got %v", apiResp.Code)
	}

	// Check if data is in expected format
	dataMap, ok := apiResp.Data.(map[string]interface{})
	if !ok {
		t.Fatalf("Expected data to be a map, got %T", apiResp.Data)
	}

	if _, exists := dataMap["items"]; !exists {
		t.Errorf("Expected 'list' field in response data")
	}
	if _, exists := dataMap["total"]; !exists {
		t.Errorf("Expected 'total' field in response data")
	}
}

func TestStcDatasetTake(t *testing.T) {
	// First create a dataset to get an ID
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	_ = writer.WriteField("name", "test-take-dataset")
	_ = writer.WriteField("desc", "Test dataset for take operation")
	_ = writer.WriteField("author", "Take Test Author")
	_ = writer.WriteField("author_wallet", TEST_AUTHOR_WALLET)

	csvContent := `test,value
sample1,123
sample2,456`

	part, err := writer.CreateFormFile("file", "take_test.csv")
	if err != nil {
		t.Fatalf("failed to create form file: %v", err)
	}
	_, err = part.Write([]byte(csvContent))
	if err != nil {
		t.Fatalf("failed to write file content: %v", err)
	}
	writer.Close()

	req, err := http.NewRequest("POST", TEST_BASE_URL+"/static-datasets", body)
	if err != nil {
		t.Fatalf("failed to create request: %v", err)
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatalf("create request failed: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		respBody, _ := io.ReadAll(resp.Body)
		t.Fatalf("Create failed with status %d: %s", resp.StatusCode, string(respBody))
	}

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("Failed to read create response body: %v", err)
	}

	createResp := responser.Response{}
	err = json.Unmarshal(respBody, &createResp)
	if err != nil {
		t.Fatalf("Failed to unmarshal create response: %v", err)
	}

	if createResp.Code != bizcode.SUCCESS {
		t.Fatalf("Expected CREATE SUCCESS, got %v", createResp.Code)
	}

	// Get transaction hash and wait for confirmation
	txHash, ok := createResp.Data.(string)
	if !ok {
		t.Fatalf("Unexpected data format: %T", createResp.Data)
	}

	msg := waitForWsConfirmation(t, txHash, 15*time.Second)

	wsResp := responser.ResponseRaw{}
	err = json.Unmarshal(msg, &wsResp)
	if err != nil {
		t.Fatalf("Failed to unmarshal WebSocket response: %v", err)
	}

	if wsResp.Code != bizcode.SUCCESS {
		t.Fatalf("Expected WebSocket SUCCESS, got %v", wsResp.Code)
	}

	var tx models.BlockchainTransaction
	err = json.Unmarshal(wsResp.Data, &tx)
	if err != nil {
		t.Fatalf("Failed to unmarshal transaction: %v", err)
	}

	datasetId := tx.EntityID
	t.Logf("Created dataset with ID: %d", datasetId)

	// Now test getting the dataset by ID
	getResp, err := http.Get(TEST_BASE_URL + "/static-datasets/" + strconv.Itoa(int(datasetId)))
	if err != nil {
		t.Fatalf("Failed to GET static dataset: %v", err)
	}
	defer getResp.Body.Close()

	if getResp.StatusCode != http.StatusOK {
		t.Fatalf("GET /static-datasets/{id} failed, status: %d", getResp.StatusCode)
	}

	getBody, err := io.ReadAll(getResp.Body)
	if err != nil {
		t.Fatalf("Failed to read GET body: %v", err)
	}

	getRespJson := responser.Response{}
	err = json.Unmarshal(getBody, &getRespJson)
	if err != nil {
		t.Fatalf("Failed to unmarshal GET response: %v", err)
	}

	if getRespJson.Code != bizcode.SUCCESS {
		t.Fatalf("GET expected SUCCESS, got %v", getRespJson.Code)
	}

	// Verify the returned dataset
	getDatasetBytes, err := json.Marshal(getRespJson.Data)
	if err != nil {
		t.Fatalf("Failed to marshal get dataset data: %v", err)
	}

	var retrievedDataset models.StaticDataset
	err = json.Unmarshal(getDatasetBytes, &retrievedDataset)
	if err != nil {
		t.Fatalf("Failed to unmarshal retrieved dataset: %v", err)
	}

	if retrievedDataset.ID != uint64(datasetId) {
		t.Errorf("Expected ID %d, got %d", datasetId, retrievedDataset.ID)
	}
	if retrievedDataset.Name != "test-take-dataset" {
		t.Errorf("Expected name 'test-take-dataset', got '%s'", retrievedDataset.Name)
	}
	if retrievedDataset.Author != "Take Test Author" {
		t.Errorf("Expected author 'Take Test Author', got '%s'", retrievedDataset.Author)
	}
	if retrievedDataset.AuthorWallet != TEST_AUTHOR_WALLET {
		t.Errorf("Expected author wallet '%s', got '%s'", TEST_AUTHOR_WALLET, retrievedDataset.AuthorWallet)
	}
}

func TestStcDatasetCreateDuplicateFile(t *testing.T) {
	csvContent := `duplicate,test
	value1,123
	value2,456`

	// Create first dataset
	body1 := &bytes.Buffer{}
	writer1 := multipart.NewWriter(body1)
	_ = writer1.WriteField("name", "first-duplicate-dataset")
	_ = writer1.WriteField("desc", "First dataset with specific content")
	_ = writer1.WriteField("author", "Duplicate Test Author")
	_ = writer1.WriteField("author_wallet", TEST_AUTHOR_WALLET)

	part1, _ := writer1.CreateFormFile("file", "duplicate1.csv")
	part1.Write([]byte(csvContent))
	writer1.Close()

	req1, _ := http.NewRequest("POST", TEST_BASE_URL+"/static-datasets", body1)
	req1.Header.Set("Content-Type", writer1.FormDataContentType())

	resp1, err := http.DefaultClient.Do(req1)
	if err != nil {
		t.Fatalf("first request failed: %v", err)
	}
	defer resp1.Body.Close()

	if resp1.StatusCode != http.StatusOK {
		respBody, _ := io.ReadAll(resp1.Body)
		t.Fatalf("first request failed with status %d: %s", resp1.StatusCode, string(respBody))
	}

	respBody1, _ := io.ReadAll(resp1.Body)
	apiResp1 := &responser.Response{}
	json.Unmarshal(respBody1, &apiResp1)

	if apiResp1.Code != bizcode.SUCCESS {
		t.Fatalf("First dataset creation failed: %v", apiResp1.Code)
	}

	// Wait for first dataset to be confirmed
	txHash1, _ := apiResp1.Data.(string)
	_ = waitForWsConfirmation(t, txHash1, 15*time.Second)

	// Try to create second dataset with same file content
	body2 := &bytes.Buffer{}
	writer2 := multipart.NewWriter(body2)
	_ = writer2.WriteField("name", "second-duplicate-dataset")
	_ = writer2.WriteField("desc", "Second dataset with same content")
	_ = writer2.WriteField("author", "Duplicate Test Author 2")
	_ = writer2.WriteField("author_wallet", TEST_AUTHOR_WALLET)

	part2, _ := writer2.CreateFormFile("file", "duplicate2.csv")
	part2.Write([]byte(csvContent))
	writer2.Close()

	req2, _ := http.NewRequest("POST", TEST_BASE_URL+"/static-datasets", body2)
	req2.Header.Set("Content-Type", writer2.FormDataContentType())

	resp2, err := http.DefaultClient.Do(req2)
	if err != nil {
		t.Fatalf("second request failed: %v", err)
	}
	defer resp2.Body.Close()

	respBody2, _ := io.ReadAll(resp2.Body)

	apiResp2 := &responser.Response{}
	json.Unmarshal(respBody2, &apiResp2)

	// Should get duplicate error
	if apiResp2.Code == bizcode.SUCCESS {
		t.Errorf("Expected duplicate error, but creation succeeded")
	} else if apiResp2.Code == bizcode.STATIC_DATASET_EXIST {
		t.Logf("Correctly detected duplicate file: %v", apiResp2.Code)
	} else {
		t.Logf("Got error code %v, response: %s", apiResp2.Code, string(respBody2))
	}
}
