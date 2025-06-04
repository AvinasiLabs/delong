package tee

import (
	"context"
	"crypto/ecdsa"
	"crypto/sha256"
	"io"
	"log"
	"sync"

	"github.com/ethereum/go-ethereum/crypto"
	"golang.org/x/crypto/hkdf"
)

type EthereumAccount struct {
	PrivateKey *ecdsa.PrivateKey
	Address    string // hex string
}

type KeyVault struct {
	adapter      IDstackClient
	ethCache     sync.Map // map[string]*EthereumAccount
	symmKeyCache sync.Map // map[string][]byte
}

func NewKeyVault(adapter IDstackClient) *KeyVault {
	return &KeyVault{
		adapter:      adapter,
		ethCache:     sync.Map{},
		symmKeyCache: sync.Map{},
	}
}

// NewKeyVaultFromConfig create KeyVault via config
func NewKeyVaultFromConfig(appenv string) *KeyVault {
	var client IDstackClient
	switch appenv {
	case "local":
		client = NewDstackClientAdapter()
	case "staging":
		client = NewTappdClientAdapter()
	default:
		log.Fatalf("Unsupported app env: %v", appenv)
	}

	return NewKeyVault(client)
}

// func (t *KeyVault) getRawKey(ctx context.Context, kc *KeyContext) ([]byte, error) {
// 	cacheKey := kc.CacheKey()
// 	val, ok := t.cache.Load(cacheKey)
// 	if ok {
// 		return val.([]byte), nil
// 	}

// 	resp, err := t.client.DeriveKey(ctx, kc.Path())
// 	if err != nil {
// 		return nil, err
// 	}

// 	block, _ := pem.Decode([]byte(resp.Key))
// 	if block == nil {
// 		return nil, fmt.Errorf("Failed to decode PEM block")
// 	}

// 	privateKey, err := x509.ParsePKCS8PrivateKey(block.Bytes)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to parse PKCS8 private key: %w", err)
// 	}

// 	ecdsaKey, ok := privateKey.(*ecdsa.PrivateKey)
// 	if !ok {
// 		return nil, fmt.Errorf("private key is not an ECDSA key")
// 	}

// 	raw := ecdsaKey.D.Bytes()

// 	// raw, err := hex.DecodeString(keyStr)
// 	// if err != nil {
// 	// 	return nil, err
// 	// }

// 	t.cache.Store(cacheKey, raw)
// 	return raw, nil
// }

func (t *KeyVault) DeriveSymmetricKey(ctx context.Context, kc *KeyContext) ([]byte, error) {
	cacheKey := kc.CacheKey()
	if val, ok := t.symmKeyCache.Load(cacheKey); ok {
		return val.([]byte), nil
	}

	// raw, err := t.getRawKey(ctx, kc)
	raw, err := t.adapter.DeriveKey(ctx, kc)
	if err != nil {
		return nil, err
	}
	reader := hkdf.New(sha256.New, raw, kc.Salt(), kc.Info())
	key := make([]byte, 32)
	if _, err := io.ReadFull(reader, key); err != nil {
		return nil, err
	}
	t.symmKeyCache.Store(cacheKey, key)
	return key, nil
}

func (t *KeyVault) DeriveEthereumAccount(ctx context.Context, kc *KeyContext) (*EthereumAccount, error) {
	cacheKey := kc.CacheKey()
	if val, ok := t.ethCache.Load(cacheKey); ok {
		return val.(*EthereumAccount), nil
	}

	// raw, err := t.getRawKey(ctx, kc)
	// if err != nil {
	// 	return nil, err
	// }
	raw, err := t.adapter.DeriveKey(ctx, kc)
	if err != nil {
		return nil, err
	}

	priv, err := crypto.ToECDSA(raw)
	if err != nil {
		return nil, err
	}
	acc := &EthereumAccount{
		PrivateKey: priv,
		Address:    crypto.PubkeyToAddress(priv.PublicKey).Hex(),
	}
	t.ethCache.Store(cacheKey, acc)
	return acc, nil
}
