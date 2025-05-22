package tee

import (
	"context"
	"crypto/ecdsa"
	"crypto/sha256"
	"encoding/hex"
	"io"
	"sync"

	"github.com/Dstack-TEE/dstack/sdk/go/dstack"
	"github.com/ethereum/go-ethereum/crypto"
	"golang.org/x/crypto/hkdf"
)

type EthereumAccount struct {
	PrivateKey *ecdsa.PrivateKey
	Address    string // hex string
}

type KeyVault struct {
	client       *dstack.DstackClient
	cache        sync.Map // map[string][]byte
	ethCache     sync.Map // map[string]*EthereumAccount
	symmKeyCache sync.Map // map[string][]byte
}

func NewKeyVault(opts ...dstack.DstackClientOption) *KeyVault {
	return &KeyVault{
		client:       dstack.NewDstackClient(opts...),
		cache:        sync.Map{},
		ethCache:     sync.Map{},
		symmKeyCache: sync.Map{},
	}
}

func (t *KeyVault) getRawKey(ctx context.Context, kc *KeyContext) ([]byte, error) {
	cacheKey := kc.CacheKey()
	val, ok := t.cache.Load(cacheKey)
	if ok {
		return val.([]byte), nil
	}

	resp, err := t.client.GetKey(ctx, kc.Path(), kc.Purpose())
	if err != nil {
		return nil, err
	}
	raw, err := hex.DecodeString(resp.Key)
	if err != nil {
		return nil, err
	}
	t.cache.Store(cacheKey, raw)
	return raw, nil
}

func (t *KeyVault) DeriveSymmetricKey(ctx context.Context, kc *KeyContext, length int) ([]byte, error) {
	cacheKey := kc.CacheKey()
	if val, ok := t.symmKeyCache.Load(cacheKey); ok {
		return val.([]byte), nil
	}

	raw, err := t.getRawKey(ctx, kc)
	if err != nil {
		return nil, err
	}
	reader := hkdf.New(sha256.New, raw, kc.Salt(), kc.Info())
	key := make([]byte, length)
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

	raw, err := t.getRawKey(ctx, kc)
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
