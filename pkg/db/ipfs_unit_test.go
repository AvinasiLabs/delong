package db

import (
	"bytes"
	"context"
	"crypto/rand"
	"delong/pkg/aesgcm"
	"io"
	"strings"
	"testing"
)

// mockReader implements io.ReadCloser for testing
type mockReader struct {
	*bytes.Reader
	closed bool
}

func (m *mockReader) Close() error {
	m.closed = true
	return nil
}

func newMockReader(data []byte) *mockReader {
	return &mockReader{
		Reader: bytes.NewReader(data),
		closed: false,
	}
}

// mockIpfsStore simulates IPFS operations for unit testing
type mockIpfsStore struct {
	storage map[string][]byte
}

func newMockIpfsStore() *mockIpfsStore {
	return &mockIpfsStore{
		storage: make(map[string][]byte),
	}
}

func (m *mockIpfsStore) storeData(cid string, data []byte) {
	m.storage[cid] = data
}

func (m *mockIpfsStore) DownloadStream(ctx context.Context, cid string) (io.ReadCloser, error) {
	data, exists := m.storage[cid]
	if !exists {
		return nil, io.EOF
	}
	return newMockReader(data), nil
}

func TestDecryptReadCloser(t *testing.T) {
	testData := []byte("Hello, World!")
	mockReader := newMockReader(testData)
	
	drc := &decryptReadCloser{
		Reader: mockReader,
		closer: mockReader,
	}
	
	// Test reading
	buffer := make([]byte, len(testData))
	n, err := drc.Read(buffer)
	if err != nil {
		t.Fatalf("Failed to read: %v", err)
	}
	
	if n != len(testData) {
		t.Errorf("Expected to read %d bytes, got %d", len(testData), n)
	}
	
	if !bytes.Equal(buffer, testData) {
		t.Errorf("Read data doesn't match expected data")
	}
	
	// Test closing
	err = drc.Close()
	if err != nil {
		t.Fatalf("Failed to close: %v", err)
	}
	
	if !mockReader.closed {
		t.Error("Underlying reader was not closed")
	}
}

func TestDownloadDecryptedStreamLogic(t *testing.T) {
	// Generate test data and key
	originalData := []byte("This is test data for encryption/decryption testing!")
	key := make([]byte, 32)
	_, err := rand.Read(key)
	if err != nil {
		t.Fatalf("Failed to generate key: %v", err)
	}
	
	// Encrypt the data manually to simulate what would be stored in IPFS
	var encryptedBuffer bytes.Buffer
	encWriter, err := aesgcm.NewWriter(&encryptedBuffer, key)
	if err != nil {
		t.Fatalf("Failed to create encrypt writer: %v", err)
	}
	
	_, err = encWriter.Write(originalData)
	if err != nil {
		t.Fatalf("Failed to encrypt data: %v", err)
	}
	encWriter.Close()
	
	encryptedData := encryptedBuffer.Bytes()
	
	// Create mock IPFS store
	mockStore := newMockIpfsStore()
	mockStore.storeData("test-cid", encryptedData)
	
	// Test the decryption logic
	ctx := context.Background()
	encryptedStream, err := mockStore.DownloadStream(ctx, "test-cid")
	if err != nil {
		t.Fatalf("Failed to get encrypted stream: %v", err)
	}
	
	decryptReader, err := aesgcm.NewReader(encryptedStream, key)
	if err != nil {
		encryptedStream.Close()
		t.Fatalf("Failed to create decrypt reader: %v", err)
	}
	
	drc := &decryptReadCloser{
		Reader: decryptReader,
		closer: encryptedStream,
	}
	
	// Read and verify the decrypted data
	decryptedData, err := io.ReadAll(drc)
	if err != nil {
		t.Fatalf("Failed to read decrypted data: %v", err)
	}
	
	if !bytes.Equal(originalData, decryptedData) {
		t.Errorf("Decrypted data doesn't match original")
		t.Errorf("Original: %q", string(originalData))
		t.Errorf("Decrypted: %q", string(decryptedData))
	}
	
	// Test closing
	err = drc.Close()
	if err != nil {
		t.Fatalf("Failed to close decrypted stream: %v", err)
	}
	
	// Verify the underlying stream was closed
	if !encryptedStream.(*mockReader).closed {
		t.Error("Underlying encrypted stream was not closed")
	}
}

func TestDownloadDecryptedStreamChunkedReading(t *testing.T) {
	// Test reading in small chunks to verify streaming behavior
	originalData := []byte(strings.Repeat("Test data for chunked reading! ", 100)) // ~3KB
	key := make([]byte, 32)
	_, err := rand.Read(key)
	if err != nil {
		t.Fatalf("Failed to generate key: %v", err)
	}
	
	// Encrypt the data
	var encryptedBuffer bytes.Buffer
	encWriter, err := aesgcm.NewWriter(&encryptedBuffer, key)
	if err != nil {
		t.Fatalf("Failed to create encrypt writer: %v", err)
	}
	
	_, err = encWriter.Write(originalData)
	if err != nil {
		t.Fatalf("Failed to encrypt data: %v", err)
	}
	encWriter.Close()
	
	encryptedData := encryptedBuffer.Bytes()
	
	// Create encrypted stream
	encryptedStream := newMockReader(encryptedData)
	
	// Create decrypt reader
	decryptReader, err := aesgcm.NewReader(encryptedStream, key)
	if err != nil {
		t.Fatalf("Failed to create decrypt reader: %v", err)
	}
	
	drc := &decryptReadCloser{
		Reader: decryptReader,
		closer: encryptedStream,
	}
	
	// Read in small chunks
	var decryptedData []byte
	buffer := make([]byte, 64) // Small buffer
	
	for {
		n, err := drc.Read(buffer)
		if n > 0 {
			decryptedData = append(decryptedData, buffer[:n]...)
		}
		if err == io.EOF {
			break
		}
		if err != nil {
			t.Fatalf("Failed to read chunk: %v", err)
		}
	}
	
	if !bytes.Equal(originalData, decryptedData) {
		t.Errorf("Chunked decrypted data doesn't match original")
		t.Errorf("Original length: %d, Decrypted length: %d", len(originalData), len(decryptedData))
	}
	
	drc.Close()
}

func TestDownloadDecryptedErrorHandling(t *testing.T) {
	key := make([]byte, 32)
	rand.Read(key)
	
	// Test with invalid encrypted data
	invalidData := []byte("This is not properly encrypted data")
	encryptedStream := newMockReader(invalidData)
	
	// This should fail when trying to decrypt
	_, err := aesgcm.NewReader(encryptedStream, key)
	// The error might occur when creating the reader or when reading
	// Either case is acceptable for this test
	if err != nil {
		// Expected - invalid data should cause an error
		return
	}
	
	// If NewReader succeeded, the error should occur when reading
	decryptReader, _ := aesgcm.NewReader(encryptedStream, key)
	drc := &decryptReadCloser{
		Reader: decryptReader,
		closer: encryptedStream,
	}
	
	buffer := make([]byte, 100)
	_, err = drc.Read(buffer)
	if err == nil {
		t.Error("Expected error when reading invalid encrypted data, but got none")
	}
	
	drc.Close()
}

func TestDownloadDecryptedEmptyData(t *testing.T) {
	key := make([]byte, 32)
	rand.Read(key)
	
	// Encrypt empty data
	var encryptedBuffer bytes.Buffer
	encWriter, err := aesgcm.NewWriter(&encryptedBuffer, key)
	if err != nil {
		t.Fatalf("Failed to create encrypt writer: %v", err)
	}
	
	_, err = encWriter.Write([]byte{})
	if err != nil {
		t.Fatalf("Failed to encrypt empty data: %v", err)
	}
	encWriter.Close()
	
	encryptedData := encryptedBuffer.Bytes()
	encryptedStream := newMockReader(encryptedData)
	
	// Decrypt
	decryptReader, err := aesgcm.NewReader(encryptedStream, key)
	if err != nil {
		t.Fatalf("Failed to create decrypt reader for empty data: %v", err)
	}
	
	drc := &decryptReadCloser{
		Reader: decryptReader,
		closer: encryptedStream,
	}
	
	// Read should return EOF immediately for empty data
	buffer := make([]byte, 100)
	n, err := drc.Read(buffer)
	
	if n != 0 {
		t.Errorf("Expected 0 bytes read for empty data, got %d", n)
	}
	
	if err != io.EOF {
		t.Errorf("Expected EOF for empty data, got: %v", err)
	}
	
	drc.Close()
}