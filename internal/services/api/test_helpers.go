package api

import (
	"context"
	"delong/pkg/bizcode"
	"delong/pkg/responser"
	"encoding/json"
	"io"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/gorilla/websocket"
)

var (
	TEST_BASE_URL = getEnvOrDefault("TEST_BASE_URL", "http://localhost:8080/api")
	TEST_WS_URL   = getEnvOrDefault("TEST_WS_URL", "ws://localhost:8080/ws")
)

func getEnvOrDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func waitForWsConfirmation(t *testing.T, txHash string, timeout time.Duration) []byte {
	t.Helper()

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

func assertApiSuccess(t *testing.T, resp *http.Response) *responser.Response {
	t.Helper()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("failed to read response body: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("unexpected status %d: %s", resp.StatusCode, string(respBody))
	}

	apiResp := &responser.Response{}
	if err := json.Unmarshal(respBody, apiResp); err != nil {
		t.Fatalf("failed to unmarshal response: %v", err)
	}

	if apiResp.Code != bizcode.SUCCESS {
		t.Errorf("API returned error, code=%v", apiResp.Code)
	}

	return apiResp
}

func postAndAssertSuccess(t *testing.T, endpoint string, body io.Reader, contentType string) *responser.Response {
	t.Helper()

	url := TEST_BASE_URL + endpoint
	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		t.Fatalf("failed to create request: %v", err)
	}

	if contentType != "" {
		req.Header.Set("Content-Type", contentType)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatalf("request failed: %v", err)
	}
	defer resp.Body.Close()

	return assertApiSuccess(t, resp)
}

func assertPostSuccessAndWaitConfirm(t *testing.T, endpoint string, body io.Reader, contentType string, timeout time.Duration) []byte {
	t.Helper()

	respBody := postAndAssertSuccess(t, endpoint, body, contentType)
	txHash, ok := respBody.Data.(string)
	if !ok {
		t.Fatalf("Unexpected data format: %T", respBody.Data)
	}

	return waitForWsConfirmation(t, txHash, timeout)
}
