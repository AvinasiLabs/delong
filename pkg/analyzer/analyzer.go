package analyzer

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type ReportAnalyzer struct {
	endpoint string
	httpCli  *http.Client
}

func NewReportAnalyzer(endpoint string) *ReportAnalyzer {
	return &ReportAnalyzer{endpoint: endpoint, httpCli: &http.Client{}}
}

type RecognizeRequest struct {
	StorageType string `json:"storage_type"`
	DataType    string `json:"data_type"`
	DataPath    string `json:"data_path"`
	UserID      string `json:"user_id"`
}

type recognizeResponse struct {
	Result string `json:"result"`
}

type DiagnosticRequest struct {
	UserID string `json:"user_id"`
	SessID string `json:"sess_id"`
}

// Analyze performs recognize then diagnostic calls and returns raw JSON bytes
func (c *ReportAnalyzer) Analyze(ctx context.Context, storageType, dataType, dataPath, userID string) ([]byte, error) {
	req1 := RecognizeRequest{storageType, dataType, dataPath, userID}
	sessAPI := fmt.Sprintf("%s/api/analyzer/diagno_analyzer/recognize", c.endpoint)
	body1, err := c.doPost(ctx, sessAPI, req1)
	if err != nil {
		return nil, fmt.Errorf("recognize request failed: %w", err)
	}

	var resp1 recognizeResponse
	if err := json.Unmarshal(body1, &resp1); err != nil {
		return nil, fmt.Errorf("parsing recognize response: %w", err)
	}

	req2 := DiagnosticRequest{userID, resp1.Result}
	analAPI := fmt.Sprintf("%s/api/analyzer/diagno_analyzer/diagnostic", c.endpoint)
	body2, err := c.doPost(ctx, analAPI, req2)
	if err != nil {
		return nil, fmt.Errorf("diagnostic request failed: %w", err)
	}

	return body2, nil
}

// doPost sends a JSON payload and returns the response body
func (c *ReportAnalyzer) doPost(ctx context.Context, url string, payload any) ([]byte, error) {
	buf, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewReader(buf))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.httpCli.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		b, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("status %d: %s", resp.StatusCode, string(b))
	}

	return io.ReadAll(resp.Body)
}
