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

type KeyContext struct {
	kind    KeyKind
	name    string
	purpose string
}

func NewKeyContext(kind KeyKind, name, purpose string) *KeyContext {
	return &KeyContext{
		kind:    kind,
		name:    name,
		purpose: purpose,
	}
}

func (k *KeyContext) Path() string {
	safeName := strings.ReplaceAll(k.name, "/", "_")
	return path.Join(string(k.kind), safeName)
}

func (k *KeyContext) Purpose() string {
	return k.purpose
}

func (k *KeyContext) Salt() []byte {
	salt := fmt.Sprintf("%s:%s:%s", k.kind, k.name, k.purpose)
	return []byte(salt)
}

func (k *KeyContext) Info() []byte {
	info := fmt.Sprintf(
		"purpose=%s,kind=%s,name=%s,version=1",
		k.purpose,
		k.kind,
		k.name,
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
