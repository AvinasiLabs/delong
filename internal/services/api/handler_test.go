package api

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"testing"
	"time"
)

func TestUploadReport(t *testing.T) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	_ = writer.WriteField("userWallet", "0xabcabcabcabcabcabcabcabcabcabcabcabcabca")
	_ = writer.WriteField("dataset", "blood-basic-panel")
	_ = writer.WriteField("testTime", time.Now().UTC().Format("2006-01-02T15:04:05Z07:00"))

	filePath := "../../assets/diagnostic/blood_report.png"
	file, err := os.Open(filePath)
	if err != nil {
		t.Fatalf("failed to open file: %v", err)
	}
	defer file.Close()

	part, err := writer.CreateFormFile("file", filepath.Base(file.Name()))
	if err != nil {
		t.Fatalf("failed to create form file: %v", err)
	}
	_, err = io.Copy(part, file)
	if err != nil {
		t.Fatalf("failed to write file content: %v", err)
	}
	writer.Close()

	req, err := http.NewRequest("POST", "http://localhost:8080/api/report/upload", body)
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
	} else {
		t.Logf("Upload successful, response: %s", string(respBody))
	}
}
