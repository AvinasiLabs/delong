package pkg

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
)

// EncryptGCM encrypts the plaintext using AES-GCM.
// It generates a random nonce and prepends it to the ciphertext.
// The function returns the combined data: [nonce | ciphertext].
func EncryptGCM(plaintext, key []byte) ([]byte, error) {
	// Create a new AES cipher using the provided key.
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, fmt.Errorf("error creating AES cipher: %v", err)
	}

	// Create a GCM cipher mode instance.
	aead, err := cipher.NewGCM(block)
	if err != nil {
		return nil, fmt.Errorf("error creating GCM: %v", err)
	}

	// Generate a random nonce.
	nonce := make([]byte, aead.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, fmt.Errorf("error generating nonce: %v", err)
	}

	// Encrypt the plaintext and obtain the ciphertext.
	ciphertext := aead.Seal(nil, nonce, plaintext, nil)
	// Prepend the nonce to the ciphertext.
	combined := append(nonce, ciphertext...)

	return combined, nil
}

// DecryptGCM decrypts data that has the nonce prepended to the ciphertext.
// It extracts the nonce and then decrypts the remaining ciphertext.
func DecryptGCM(combined, key []byte) ([]byte, error) {
	// Create a new AES cipher using the provided key.
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, fmt.Errorf("error creating AES cipher: %v", err)
	}

	// Create a GCM cipher mode instance.
	aead, err := cipher.NewGCM(block)
	if err != nil {
		return nil, fmt.Errorf("error creating GCM: %v", err)
	}

	nonceSize := aead.NonceSize()
	// Check that the combined data length is at least as long as the nonce.
	if len(combined) < nonceSize {
		return nil, fmt.Errorf("combined data too short")
	}

	// Extract the nonce and the ciphertext.
	nonce := combined[:nonceSize]
	ciphertext := combined[nonceSize:]

	// Decrypt the ciphertext.
	plaintext, err := aead.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return nil, fmt.Errorf("error decrypting data: %v", err)
	}

	return plaintext, nil
}

// EncryptGCMWithHexKey encrypts data using a key provided as a hex string.
// It handles the hex decoding of the key before performing encryption.
func EncryptGCMWithHexKey(plaintext []byte, hexKey string) ([]byte, error) {
	key, err := hex.DecodeString(hexKey)
	if err != nil {
		return nil, fmt.Errorf("error decoding hex key: %v", err)
	}

	// Validate key length (AES-256 requires 32 bytes)
	if len(key) != 32 {
		return nil, fmt.Errorf("invalid key length: expected 32 bytes, got %d", len(key))
	}

	return EncryptGCM(plaintext, key)
}

// DecryptGCMWithHexKey decrypts data using a key provided as a hex string.
// It handles the hex decoding of the key before performing decryption.
func DecryptGCMWithHexKey(combined []byte, hexKey string) ([]byte, error) {
	key, err := hex.DecodeString(hexKey)
	if err != nil {
		return nil, fmt.Errorf("error decoding hex key: %v", err)
	}

	// Validate key length (AES-256 requires 32 bytes)
	if len(key) != 32 {
		return nil, fmt.Errorf("invalid key length: expected 32 bytes, got %d", len(key))
	}

	return DecryptGCM(combined, key)
}
