package aesgcm

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
)

// Encrypt encrypts plaintext using AES-GCM, prepending a random nonce.
// Returns [nonce | ciphertext].
func Encrypt(plaintext, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	aead, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}
	nonce := make([]byte, aead.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}
	ct := aead.Seal(nil, nonce, plaintext, nil)
	return append(nonce, ct...), nil
}

// Decrypt undoes Encrypt: it splits out the nonce and decrypts.
func Decrypt(combined, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	aead, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}
	ns := aead.NonceSize()
	if len(combined) < ns {
		return nil, fmt.Errorf("combined data too short")
	}
	nonce, ct := combined[:ns], combined[ns:]
	pt, err := aead.Open(nil, nonce, ct, nil)
	if err != nil {
		return nil, err
	}
	return pt, nil
}

// EncryptHexKey is like Encrypt, but accepts a hex-encoded key.
func EncryptHexKey(plaintext []byte, hexKey string) ([]byte, error) {
	key, err := hex.DecodeString(hexKey)
	if err != nil {
		return nil, err
	}
	if l := len(key); l != 16 && l != 24 && l != 32 {
		return nil, err
	}
	return Encrypt(plaintext, key)
}

// DecryptHexKey is like Decrypt, but accepts a hex-encoded key.
func DecryptHexKey(combined []byte, hexKey string) ([]byte, error) {
	key, err := hex.DecodeString(hexKey)
	if err != nil {
		return nil, err
	}
	if l := len(key); l != 16 && l != 24 && l != 32 {
		return nil, err
	}
	return Decrypt(combined, key)
}
