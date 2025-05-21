package api

import (
	"bytes"
	"context"
	"delong/internal/models"
	"delong/internal/types"
	"delong/pkg/bizcode"
	"delong/pkg/responser"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"testing"
	"time"

	"github.com/gorilla/websocket"
)

const TEST_BASE_URL = "http://localhost:8080/api"
const TEST_WS_URL = "ws://localhost:8080/ws"
const TEST_MEMEBR_WALLET = "0xa0Ee7A142d267C1f36714E4a8F75612F20a79720"

func waitForWsConfirmation(t *testing.T, txHash string, timeout time.Duration) []byte {
	wsURL := TEST_WS_URL + "?task_id=" + txHash
	dialer := websocket.Dialer{}
	conn, _, err := dialer.Dial(wsURL, nil)
	if err != nil {
		t.Fatalf("WebSocket connection failed: %v", err)
	}
	defer conn.Close()

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	ch := make(chan []byte, 1)
	go func() {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			t.Logf("WebSocket read error: %v", err)
			close(ch)
			return
		}
		ch <- msg
	}()

	select {
	case <-ctx.Done():
		t.Fatalf("Timeout waiting for WebSocket tx confirmation for %s", txHash)
		return nil
	case msg := <-ch:
		return msg
	}
}

func TestCommitteeMemberCreate(t *testing.T) {
	is_approved := true
	body := types.SetCommitteeMemberReq{
		MemberWallet: TEST_MEMEBR_WALLET,
		IsApproved:   &is_approved,
	}
	jsonBody, _ := json.Marshal(body)

	req, err := http.NewRequest("POST", TEST_BASE_URL+"/committee", bytes.NewBuffer(jsonBody))
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatalf("Request failed: %v", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("Failed to read response body: %v", err)
	}
	apiResp := responser.Response{}
	err = json.Unmarshal(respBody, &apiResp)
	if err != nil {
		t.Fatalf("Failed to unmarshal response body: %v", err)
	}
	t.Logf("Create response: %v", apiResp)

	if apiResp.Code != bizcode.SUCCESS {
		t.Errorf("Expected CODE SUCCESS, got %v", apiResp.Code)
	}

	txHash, ok := apiResp.Data.(string)
	if !ok {
		t.Fatalf("Unexpected data format: %T", apiResp.Data)
	}

	msg := waitForWsConfirmation(t, txHash, 10*time.Second)
	t.Logf("Received msg: %v", string(msg))
}

func TestCommitteeMembersList(t *testing.T) {
	u, _ := url.Parse(TEST_BASE_URL + "/committee")
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
}

func TestCommitteeMemberTake(t *testing.T) {
	is_approved := false
	body := types.SetCommitteeMemberReq{
		MemberWallet: TEST_MEMEBR_WALLET,
		IsApproved:   &is_approved,
	}
	jsonBody, err := json.Marshal(body)
	if err != nil {
		t.Fatalf("Failed to marshal body: %v", err)
	}

	resp, err := http.Post(TEST_BASE_URL+"/committee", "application/json", bytes.NewBuffer(jsonBody))
	if err != nil {
		t.Fatalf("Failed to create committee member: %v", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("Create failed with status: %d", resp.StatusCode)
	}
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("Failed to read response body: %v", err)
	}
	apiResp := responser.Response{}
	err = json.Unmarshal(respBody, &apiResp)
	if err != nil {
		t.Fatalf("Failed to unmarshal resp body: %v", err)
	}
	if apiResp.Code != bizcode.SUCCESS {
		t.Fatalf("Expected CODE SUCCESS, got %v", apiResp.Code)
	}

	// Send ws request by specific txHash
	txHash, ok := apiResp.Data.(string)
	if !ok {
		t.Fatalf("Unexpected data format: %T", apiResp.Data)
	}

	msg := waitForWsConfirmation(t, txHash, 10*time.Second)
	t.Logf("Received msg: %v", string(msg))

	wsResp := responser.ResponseRaw{}
	err = json.Unmarshal(msg, &wsResp)
	if err != nil {
		t.Fatalf("Failed to unmarshal: %v", err)
	}
	if wsResp.Code != bizcode.SUCCESS {
		t.Fatalf("Expected CODE SUCCESS, got %v", apiResp.Code)
	}

	var tx models.BlockchainTransaction
	err = json.Unmarshal(wsResp.Data, &tx)
	if err != nil {
		t.Fatalf("Failed to unmarshal: %v", err)
	}

	memberId := tx.EntityID
	getResp, err := http.Get(TEST_BASE_URL + "/committee/" + strconv.Itoa(int(memberId)))
	if err != nil {
		t.Fatalf("Failed to GET committee member: %v", err)
	}
	defer getResp.Body.Close()

	if getResp.StatusCode != http.StatusOK {
		t.Fatalf("GET /committee/{id} failed, status: %d", getResp.StatusCode)
	}

	getBody, err := io.ReadAll(getResp.Body)
	if err != nil {
		t.Fatalf("Failed to read GET body: %v", err)
	}

	getRespJson := responser.ResponseRaw{}
	err = json.Unmarshal(getBody, &getRespJson)
	if err != nil {
		t.Fatalf("Failed to unmarshal GET response: %v", err)
	}
	if getRespJson.Code != bizcode.SUCCESS {
		t.Fatalf("GET expected SUCCESS, got %v", getRespJson.Code)
	}

	var member models.CommitteeMember
	err = json.Unmarshal(getRespJson.Data, &member)
	if err != nil {
		t.Fatalf("Failed to decode committee member: %v", err)
	}

	if member.ID != int(memberId) {
		t.Errorf("Expected ID %d, got %d", memberId, member.ID)
	}
	if member.MemberWallet != TEST_MEMEBR_WALLET {
		t.Errorf("Expected wallet %s, got %s", TEST_MEMEBR_WALLET, member.MemberWallet)
	}
}
