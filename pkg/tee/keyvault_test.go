package tee

import (
	"context"
	"encoding/hex"
	"testing"
)

func TestDeriveSymmetricKey(t *testing.T) {
	ctx := context.Background()
	keyVault := NewKeyVaultFromConfig(VersionDstack)

	kc := NewKeyContext(KEYKIND_ENC_KEY, "testuser", "encrypt_data")

	key, err := keyVault.DeriveSymmetricKey(ctx, kc, 32)
	if err != nil {
		t.Fatalf("DeriveSymmetricKey failed: %v", err)
	}
	if len(key) != 32 {
		t.Fatalf("Expected key length 32, got %d", len(key))
	}

	t.Logf("Derived symmetric key (hex): %s", hex.EncodeToString(key))
}

func TestDeriveEthereumAccount(t *testing.T) {
	ctx := context.Background()
	keyVault := NewKeyVaultFromConfig(VersionDstack)

	kc := NewKeyContext(KEYKIND_ETH_ACC, "testuser", "eth_identity")

	acct, err := keyVault.DeriveEthereumAccount(ctx, kc)
	if err != nil {
		t.Fatalf("DeriveEthereumAccount failed: %v", err)
	}
	if acct == nil {
		t.Fatal("Expected non-nil EthereumAccount")
	}
	if len(acct.Address) == 0 {
		t.Fatal("Expected non-empty Ethereum address")
	}

	t.Logf("Derived Ethereum address: %s", acct.Address)

	if _, err := hex.DecodeString(acct.Address[2:]); err != nil {
		t.Fatalf("Invalid Ethereum address format: %v", err)
	}
}
