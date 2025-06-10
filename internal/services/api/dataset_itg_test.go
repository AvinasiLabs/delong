//go:build integration
// +build integration

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

func createDataset(t *testing.T) uint {
	uniqueName := "test" + "-" + strconv.FormatInt(time.Now().UnixNano(), 10)

	body := types.DatasetCreateReq{
		Name:        uniqueName,
		UiName:      "Test Title",
		Description: "Test Description",
	}
	jsonBody, err := json.Marshal(body)
	if err != nil {
		t.Fatalf("Failed to marshal create body: %v", err)
	}

	resp, err := http.Post(TEST_BASE_URL+"/datasets", "application/json", bytes.NewBuffer(jsonBody))
	if err != nil {
		t.Fatalf("POST /datasets failed: %v", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("Failed to read POST response: %v", err)
	}

	var apiResp responser.ResponseRaw
	if err := json.Unmarshal(respBody, &apiResp); err != nil {
		t.Fatalf("Failed to unmarshal POST response: %v", err)
	}
	if apiResp.Code != bizcode.SUCCESS {
		t.Fatalf("POST response code != SUCCESS: %v", apiResp.Code)
	}

	var created models.DynamicDataset
	if err := json.Unmarshal(apiResp.Data, &created); err != nil {
		t.Fatalf("Failed to decode created dataset: %v", err)
	}

	return uint(created.ID)
}

func TestDatasetCreateAndTake(t *testing.T) {
	id := createDataset(t)
	// Now take dataset by ID
	getResp, err := http.Get(TEST_BASE_URL + "/datasets/" + strconv.Itoa(int(id)))
	if err != nil {
		t.Fatalf("GET /datasets/:id failed: %v", err)
	}
	defer getResp.Body.Close()
	getBody, _ := io.ReadAll(getResp.Body)

	t.Logf("Dataset take response: %s", getBody)
	if getResp.StatusCode != http.StatusOK {
		t.Fatalf("GET failed, status: %d", getResp.StatusCode)
	}
}

func TestDatasetList(t *testing.T) {
	reqUrl := TEST_BASE_URL + "/datasets?page=1&page_size=10"

	resp, err := http.Get(reqUrl)
	if err != nil {
		t.Fatalf("GET /datasets failed: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("Failed to read response body: %v", err)
	}

	t.Logf("Dataset list response: %s", body)

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("Expected status 200, got %d", resp.StatusCode)
	}

	apiResp := responser.ResponseRaw{}
	if err := json.Unmarshal(body, &apiResp); err != nil {
		t.Fatalf("Failed to unmarshal response: %v", err)
	}
	if apiResp.Code != bizcode.SUCCESS {
		t.Fatalf("Expected SUCCESS, got %v", apiResp.Code)
	}

	var list struct {
		Items    []models.DynamicDataset `json:"items"`
		Total    int                     `json:"total"`
		Page     int                     `json:"page"`
		PageSize int                     `json:"page_size"`
	}
	if err := json.Unmarshal(apiResp.Data, &list); err != nil {
		t.Fatalf("Failed to parse dataset list: %v", err)
	}

	if list.Page != 1 || list.PageSize != 10 {
		t.Errorf("Unexpected pagination info: page=%d, size=%d", list.Page, list.PageSize)
	}
	t.Logf("Fetched %d datasets (total: %d)", len(list.Items), list.Total)
}

func TestDatasetUpdate(t *testing.T) {
	id := createDataset(t)
	test_update_desc := "Updated Description"
	updateBody := types.DatasetUpdateReq{
		Description: test_update_desc,
	}
	jsonBody, _ := json.Marshal(updateBody)
	req, err := http.NewRequest(http.MethodPut, TEST_BASE_URL+"/datasets/"+strconv.Itoa(int(id)), bytes.NewBuffer(jsonBody))
	if err != nil {
		t.Fatalf("Failed to create PUT request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatalf("PUT /datasets/:id failed: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("Update failed, status: %d", resp.StatusCode)
	}

	respBody, _ := io.ReadAll(resp.Body)
	apiResp := responser.ResponseRaw{}
	_ = json.Unmarshal(respBody, &apiResp)
	if apiResp.Code != bizcode.SUCCESS {
		t.Fatalf("Expected SUCCESS on update, got %v", apiResp.Code)
	}

	var updated models.DynamicDataset
	_ = json.Unmarshal(apiResp.Data, &updated)
	t.Logf("Update Response: %v", updated)

	if updated.Description != test_update_desc {
		t.Errorf("Expected title updated, got: %s", updated.Description)
	}
}

func TestDatasetDelete(t *testing.T) {
	id := createDataset(t)

	req, _ := http.NewRequest(http.MethodDelete, TEST_BASE_URL+"/datasets/"+strconv.Itoa(int(id)), nil)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatalf("DELETE /datasets/:id failed: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("DELETE failed with status: %d", resp.StatusCode)
	}
	body, _ := io.ReadAll(resp.Body)
	apiResp := responser.Response{}
	_ = json.Unmarshal(body, &apiResp)
	if apiResp.Code != bizcode.SUCCESS {
		t.Errorf("Expected SUCCESS on delete, got %v", apiResp.Code)
	}

	getResp, err := http.Get(TEST_BASE_URL + "/datasets/" + strconv.Itoa(int(id)))
	if err != nil {
		t.Fatalf("GET after delete failed: %v", err)
	}
	defer getResp.Body.Close()
	getBody, _ := io.ReadAll(getResp.Body)

	apiResp = responser.Response{}
	_ = json.Unmarshal(getBody, &apiResp)

	if apiResp.Code == bizcode.SUCCESS {
		t.Errorf("Expected NOT_FOUND or error after delete, got: %v", apiResp)
	}
}
