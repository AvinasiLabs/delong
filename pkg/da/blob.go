package da

import (
	"context"
	"crypto/ecdsa"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/crypto/kzg4844"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/holiman/uint256"
)

func initWithKey(PrivateKeyInHex string) (key *ecdsa.PrivateKey, sender common.Address) {
	var err error
	key, err = crypto.HexToECDSA(PrivateKeyInHex)
	if err != nil {
		log.Println("Error generating ecdsa private key from hex", err)
		return
	}
	pubKey := key.Public()
	pubKeyECDSA, ok := pubKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("ECDSA")
		return
	}

	sender = crypto.PubkeyToAddress(*pubKeyECDSA)
	return
}

func newRPCProvider(url string) (client *ethclient.Client, err error) {
	client, err = ethclient.Dial(url)
	if err != nil {
		log.Println("Failed to setup client", "error", err)
		return
	}
	return
}

func createAndSendBlobTx(client *ethclient.Client, sender common.Address, recipient common.Address, key *ecdsa.PrivateKey) (signedTx *types.Transaction, err error) {
	onlyCheckBlobTxGasPrice := false
	sendBlobWithDataTx := true

	nonce, err := client.NonceAt(context.Background(), sender, nil)
	if err != nil {
		log.Println("Failed to get nonce", "err", err)
	}
	log.Println("Nonce is: ", nonce)

	chainId, err := client.ChainID(context.Background())
	if err != nil {
		log.Println("Failed to get chainId", "err", err)
		return
	}
	log.Println("chainId is: ", chainId)

	lastBlock, err := client.BlockByNumber(context.Background(), nil)
	if err != nil {
		log.Println("Failed to get last Block", "err", err)
		return
	}

	excessBlobGas := *lastBlock.ExcessBlobGas()
	// calcBlobFee := eip4844.CalcBlobFee(excessBlobGas)
	blobFeeCap := uint256.NewInt(excessBlobGas)
	log.Println("blobFeeCap is: ", blobFeeCap)
	blobFeeCap.Mul(blobFeeCap, uint256.NewInt(2))
	log.Println("blobFeeCap will be send is: ", blobFeeCap)

	tip := lastBlock.BaseFee().Mul(lastBlock.BaseFee(), big.NewInt(10))
	// In case of tx replacement you will need to pay penalties
	/*
		https://github.com/ethereum/go-ethereum/blob/767b00b0b514771a663f3362dd0310fc28d40c25/core/txpool/blobpool/blobpool.go#L1148
	*/
	// penalty := big.NewInt(120000000000)
	// tip := lastBlock.BaseFee().Add(lastBlock.BaseFee(), penalty)
	log.Println("lastBlock.BaseFee() is: ", lastBlock.BaseFee())
	log.Println("tip is:                 ", tip)

	maxFeePerGas := lastBlock.BaseFee().Add(lastBlock.BaseFee(), tip)
	log.Println("maxFeePerGas is:        ", maxFeePerGas)

	if onlyCheckBlobTxGasPrice {
		return nil, nil
	}

	var myBlob kzg4844.Blob
	if sendBlobWithDataTx {
		test := []byte("Obrigado, Leo Papais...")
		copy(myBlob[:], test)
	}

	myBlobCommit, err := kzg4844.BlobToCommitment(&myBlob)
	if err != nil {
		log.Println("Failed to create commitment", "err", err)
		return
	}
	myBlobProof, err := kzg4844.ComputeBlobProof(&myBlob, myBlobCommit)
	if err != nil {
		log.Println("Failed to create proof", "err", err)
		return
	}
	sidecar := types.BlobTxSidecar{
		Blobs:       []kzg4844.Blob{myBlob},
		Commitments: []kzg4844.Commitment{myBlobCommit},
		Proofs:      []kzg4844.Proof{myBlobProof},
	}
	// log.Printf("Blog generated: %+v\n", sidecar)

	tx := types.NewTx(&types.BlobTx{
		ChainID:    uint256.MustFromBig(chainId),
		Nonce:      nonce,
		GasTipCap:  uint256.MustFromBig(tip),
		GasFeeCap:  uint256.MustFromBig(maxFeePerGas),
		Gas:        250000,
		To:         recipient,
		Value:      uint256.NewInt(0),
		Data:       nil,
		BlobFeeCap: uint256.NewInt(3e10), // TODO: calculate BlobFeeCap
		BlobHashes: sidecar.BlobHashes(),
		Sidecar:    &sidecar,
	})
	log.Printf("Raw Tx is: %+v\n", tx)

	signedTx, err = types.SignTx(tx, types.NewCancunSigner(tx.ChainId()), key)
	if err != nil {
		log.Println("Error signing transaction", "error", err)
		return
	}

	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Println("Failed to send transaction", "error", err)
		return
	}
	return
}
