package da

import (
	"crypto/ecdsa"
	"log"
	"testing"

	"github.com/ethereum/go-ethereum/common"
)

const (
	PrivateKeyInHex = "ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80"
	ProviderRpcUrl  = "http://127.0.0.1:8545"
)

var (
	Key       *ecdsa.PrivateKey
	Sender    common.Address
	Recipient = common.HexToAddress("")
)

func TestSendBlobTransaction(t *testing.T) {

	t.Run("Send Blob Transaction", func(t *testing.T) {
		log.Println("Connecting to Ethereum RPC Provider...")
		client, err := newRPCProvider(ProviderRpcUrl)
		if err != nil {
			t.Fatalf("Error connecting to Ethereum RPC Provider %v", err)
		}
		log.Println("Connected.")
		key, sender := initWithKey(PrivateKeyInHex)
		tx, err := createAndSendBlobTx(client, sender, Recipient, key)
		if err != nil {
			t.Fatalf("Error creating and sending blob transaction %v", err)
		}
		if tx == nil {
			t.Fatal("Blob transaction not sent")
			return
		}
		log.Println("Transaction sent", "hash", tx.Hash().Hex())
	})
}
