//go:build integration
// +build integration

package api

// go:embed testdata/blood_report.png
// var testReportFile []byte

// func TestUploadReport(t *testing.T) {
// 	body := &bytes.Buffer{}
// 	writer := multipart.NewWriter(body)
// 	_ = writer.WriteField("user_wallet", "0xabcabcabcabcabcabcabcabcabcabcabcabcabca")
// 	_ = writer.WriteField("dataset", "blood-basic-panel")
// 	_ = writer.WriteField("test_time", time.Now().UTC().Format("2006-01-02T15:04:05Z07:00"))

// 	part, err := writer.CreateFormFile("file", "blood_report.png")
// 	if err != nil {
// 		t.Fatalf("failed to create form file: %v", err)
// 	}
// 	_, err = part.Write(testReportFile)
// 	if err != nil {
// 		t.Fatalf("failed to write file content: %v", err)
// 	}
// 	writer.Close()

// 	_ = assertApiSuccessAndWaitConfirm(t, "/reports", body, writer.FormDataContentType())
// }
