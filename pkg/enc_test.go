package pkg

import (
	"bytes"
	"encoding/hex"
	"os"
	"testing"
)

// getTestKey returns the encryption key from the environment variable.
// It fails the test if the key is not exactly 32 bytes long.
func getTestKey(t *testing.T) []byte {
	keyHexStr := os.Getenv("ENC_KEY_HEX")
	keyStr, err := hex.DecodeString(keyHexStr)
	if err != nil {
		t.Fatalf("Failed to get ENC_KEY_HEX envvar: %v", err)
	}

	if len(keyStr) != 32 {
		t.Fatalf("ENC_KEY_HEX environment variable must be 32 bytes long, got length %d", len(keyStr))
	}
	return []byte(keyStr)
}

// TestEncryptDecrypt_Simple tests encryption and decryption of a simple string.
func TestEncryptDecrypt_Simple(t *testing.T) {
	key := getTestKey(t)
	plaintext := []byte("Test simple string")
	combined, err := EncryptGCM(plaintext, key)
	if err != nil {
		t.Fatalf("EncryptGCM returned error: %v", err)
	}
	decrypted, err := DecryptGCM(combined, key)
	if err != nil {
		t.Fatalf("DecryptGCM returned error: %v", err)
	}
	if !bytes.Equal(plaintext, decrypted) {
		t.Fatalf("decrypted text does not match plaintext; expected %s, got %s", plaintext, decrypted)
	}
}

// TestEncryptDecrypt_Empty tests encryption and decryption of an empty string.
func TestEncryptDecrypt_Empty(t *testing.T) {
	key := getTestKey(t)
	plaintext := []byte("")
	combined, err := EncryptGCM(plaintext, key)
	if err != nil {
		t.Fatalf("EncryptGCM returned error: %v", err)
	}
	decrypted, err := DecryptGCM(combined, key)
	if err != nil {
		t.Fatalf("DecryptGCM returned error: %v", err)
	}
	if !bytes.Equal(plaintext, decrypted) {
		t.Fatalf("decrypted text does not match plaintext; expected empty string, got %s", decrypted)
	}
}

// TestEncryptDecrypt_Long tests encryption and decryption of a longer text.
func TestEncryptDecrypt_Long(t *testing.T) {
	key := getTestKey(t)
	plaintext := []byte("This is a longer test message for AES-GCM encryption and decryption to ensure the algorithm can handle more data.")
	combined, err := EncryptGCM(plaintext, key)
	if err != nil {
		t.Fatalf("EncryptGCM returned error: %v", err)
	}
	decrypted, err := DecryptGCM(combined, key)
	if err != nil {
		t.Fatalf("DecryptGCM returned error: %v", err)
	}
	if !bytes.Equal(plaintext, decrypted) {
		t.Fatalf("decrypted text does not match plaintext; expected %s, got %s", plaintext, decrypted)
	}
}

// TestEncryptDecrypt_NonASCII tests encryption and decryption of text containing non-ASCII characters.
func TestEncryptDecrypt_NonASCII(t *testing.T) {
	key := getTestKey(t)
	plaintext := []byte("测试含有中文的文本")
	combined, err := EncryptGCM(plaintext, key)
	if err != nil {
		t.Fatalf("EncryptGCM returned error: %v", err)
	}
	decrypted, err := DecryptGCM(combined, key)
	if err != nil {
		t.Fatalf("DecryptGCM returned error: %v", err)
	}
	if !bytes.Equal(plaintext, decrypted) {
		t.Fatalf("decrypted text does not match plaintext; expected %s, got %s", plaintext, decrypted)
	}
}

func TestEncryptDecryptWithHexKey(t *testing.T) {
	hexKey := os.Getenv("ENCRYPTION_KEY")
	if hexKey == "" {
		t.Skip("ENCRYPTION_KEY environment variable not set")
	}

	plaintext := []byte("Test with hex key")

	combined, err := EncryptGCMWithHexKey(plaintext, hexKey)
	if err != nil {
		t.Fatalf("EncryptGCMWithHexKey returned error: %v", err)
	}

	decrypted, err := DecryptGCMWithHexKey(combined, hexKey)
	if err != nil {
		t.Fatalf("DecryptGCMWithHexKey returned error: %v", err)
	}

	if !bytes.Equal(plaintext, decrypted) {
		t.Fatalf("decrypted text does not match plaintext; expected %s, got %s", plaintext, decrypted)
	}
}
