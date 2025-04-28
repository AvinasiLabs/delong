package tee

import (
	"fmt"
	"path"
	"strings"
)

type KeyKind string

const (
	KindEthAccount    KeyKind = "eth_account"
	KindEncryptionKey KeyKind = "encryption"
)

type KeyIdentity struct {
	kind KeyKind
	name string
}

type KeyContext struct {
	path    *KeyIdentity
	purpose string
}

func NewKeyContext(kind KeyKind, name, purpose string) *KeyContext {
	return &KeyContext{
		path: &KeyIdentity{
			kind: kind,
			name: name,
		},
		purpose: purpose,
	}
}

func (k *KeyContext) Path() string {
	safeName := strings.ReplaceAll(k.path.name, "/", "_")
	return path.Join(string(k.path.kind), safeName)
}

func (k *KeyContext) Purpose() string {
	return k.purpose
}

func (k *KeyContext) Salt() []byte {
	salt := fmt.Sprintf("%s:%s:%s", k.path.kind, k.path.name, k.purpose)
	return []byte(salt)
}

func (k *KeyContext) Info() []byte {
	info := fmt.Sprintf(
		"purpose=%s,kind=%s,name=%s,version=1",
		k.purpose,
		k.path.kind,
		k.path.name,
	)
	return []byte(info)
}

func (k *KeyContext) CacheKey() string {
	cacheKey := fmt.Sprintf(
		"%s:%s",
		k.Path(),
		k.Purpose(),
	)
	return cacheKey
}
