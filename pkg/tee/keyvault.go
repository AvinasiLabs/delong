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
	Address    string
}

type KeyVault struct {
	client *dstack.DstackClient
	cache  sync.Map // map[string][]byte
}

func NewKeyVault(opts ...dstack.DstackClientOption) *KeyVault {
	return &KeyVault{
		client: dstack.NewDstackClient(opts...),
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
	raw, err := t.getRawKey(ctx, kc)
	if err != nil {
		return nil, err
	}
	reader := hkdf.New(sha256.New, raw, kc.Salt(), kc.Info())
	key := make([]byte, length)
	if _, err := io.ReadFull(reader, key); err != nil {
		return nil, err
	}
	return key, nil
}

func (t *KeyVault) DeriveEthereumAccount(ctx context.Context, kc *KeyContext) (*EthereumAccount, error) {
	raw, err := t.getRawKey(ctx, kc)
	if err != nil {
		return nil, err
	}
	priv, err := crypto.ToECDSA(raw)
	if err != nil {
		return nil, err
	}
	addr := crypto.PubkeyToAddress(priv.PublicKey)
	return &EthereumAccount{
		PrivateKey: priv,
		Address:    addr.Hex(),
	}, nil
}
