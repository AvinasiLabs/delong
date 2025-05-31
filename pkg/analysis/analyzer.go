package analysis

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

type ReportAnalyzer struct {
	endpoint string
	httpCli  *http.Client
}

func NewReportAnalyzer(endpoint string) *ReportAnalyzer {
	return &ReportAnalyzer{endpoint: endpoint, httpCli: &http.Client{}}
}

// AnalyzeFile uploads a file directly to the analyzer endpoint and returns raw JSON bytes
func (c *ReportAnalyzer) AnalyzeFile(ctx context.Context, filePath string) ([]byte, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("opening file: %w", err)
	}
	defer file.Close()

	fileName := filepath.Base(filePath)
	return c.AnalyzeFileWithReader(ctx, fileName, file)
}

// AnalyzeFileWithReader uploads file content from a reader to the analyzer endpoint
func (c *ReportAnalyzer) AnalyzeFileWithReader(ctx context.Context, fileName string, fileReader io.Reader) ([]byte, error) {
	var buf bytes.Buffer
	writer := multipart.NewWriter(&buf)

	part, err := writer.CreateFormFile("file", fileName)
	if err != nil {
		return nil, fmt.Errorf("creating form file: %w", err)
	}

	_, err = io.Copy(part, fileReader)
	if err != nil {
		return nil, fmt.Errorf("copying file content: %w", err)
	}

	err = writer.Close()
	if err != nil {
		return nil, fmt.Errorf("closing multipart writer: %w", err)
	}

	url := fmt.Sprintf("%s/api/analyzer/diagno_analyzer/analyze", c.endpoint)
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, &buf)
	if err != nil {
		return nil, fmt.Errorf("creating request: %w", err)
	}

	req.Header.Set("Content-Type", writer.FormDataContentType())

	resp, err := c.httpCli.Do(req)
	if err != nil {
		return nil, fmt.Errorf("executing request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		b, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("status %d: %s", resp.StatusCode, string(b))
	}

	return io.ReadAll(resp.Body)
}
