package tee

import (
	"context"
	"crypto/ecdsa"
	"crypto/x509"
	"encoding/hex"
	"encoding/pem"
	"fmt"
	"sync"

	"github.com/Dstack-TEE/dstack/sdk/go/dstack"
	"github.com/Dstack-TEE/dstack/sdk/go/tappd"
)

// IDstackClient defines a standard interface for key derivation from a TEE source.
type IDstackClient interface {
	DeriveKey(ctx context.Context, keyCtx *KeyContext) ([]byte, error)
}

// DstackClientAdapter wraps dstack.DstackClient to implement the KeyDerivationClient interface.
type DstackClientAdapter struct {
	client *dstack.DstackClient
	cache  sync.Map
}

// NewDstackClientAdapter creates a new DstackClientAdapter instance.
func NewDstackClientAdapter() *DstackClientAdapter {
	return &DstackClientAdapter{
		client: dstack.NewDstackClient(),
		cache:  sync.Map{},
	}
}

// DeriveKey derives a key using dstack.DstackClient and caches the result.
func (a *DstackClientAdapter) DeriveKey(ctx context.Context, keyCtx *KeyContext) ([]byte, error) {
	cacheKey := keyCtx.CacheKey()
	val, ok := a.cache.Load(cacheKey)
	if ok {
		return val.([]byte), nil
	}

	resp, err := a.client.GetKey(ctx, keyCtx.Path(), keyCtx.Purpose())
	if err != nil {
		return nil, err
	}
	raw, err := hex.DecodeString(resp.Key)
	if err != nil {
		return nil, err
	}

	a.cache.Store(cacheKey, raw)
	return raw, nil
}

// TappdClientAdapter wraps tappd.TappdClient to implement the KeyDerivationClient interface.
type TappdClientAdapter struct {
	client *tappd.TappdClient
	cache  sync.Map
}

// NewTappdClientAdapter creates a new TappdClientAdapter instance.
func NewTappdClientAdapter() *TappdClientAdapter {
	return &TappdClientAdapter{
		client: tappd.NewTappdClient(),
		cache:  sync.Map{},
	}
}

// DeriveKey derives a key using tappd.TappdClient, parses it from PEM, and caches the result.
func (a *TappdClientAdapter) DeriveKey(ctx context.Context, kc *KeyContext) ([]byte, error) {
	cacheKey := kc.CacheKey()
	val, ok := a.cache.Load(cacheKey)
	if ok {
		return val.([]byte), nil
	}

	resp, err := a.client.DeriveKey(ctx, kc.Path())
	if err != nil {
		return nil, err
	}

	block, _ := pem.Decode([]byte(resp.Key))
	if block == nil {
		return nil, fmt.Errorf("Failed to decode PEM block")
	}

	privateKey, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf("failed to parse PKCS8 private key: %w", err)
	}

	ecdsaKey, ok := privateKey.(*ecdsa.PrivateKey)
	if !ok {
		return nil, fmt.Errorf("private key is not an ECDSA key")
	}

	raw := ecdsaKey.D.Bytes()

	a.cache.Store(cacheKey, raw)
	return raw, nil
}
