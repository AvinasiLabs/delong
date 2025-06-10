//go:build integration
// +build integration

package api

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"delong/internal"
	"delong/internal/models"
	"delong/internal/types"
	"delong/pkg/bizcode"
	"delong/pkg/contracts"
	"delong/pkg/db"
	"delong/pkg/responser"
	"delong/pkg/tee"
	"encoding/json"
	"io"
	"log"
	"math/big"
	"math/rand"
	"net/http"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
)

const TEST_ALGO_CID = "Qmc2q2fMPuss3tLWVJgqEB7p98rB1Bnz7he6x5hXHRdUYg"

func TestVoteCreate(t *testing.T) {
	tx := vote(t, TEST_ALGO_CID, true)
	msg := waitForWsConfirmation(t, tx.Hash().Hex(), 10*time.Second)

	var wsResp responser.ResponseRaw
	if err := json.Unmarshal(msg, &wsResp); err != nil {
		t.Fatalf("Failed to unmarshal WS message: %v", err)
	}
	if wsResp.Code != bizcode.SUCCESS {
		t.Fatalf("Vote WS confirmation failed: %v", wsResp.Code)
	}

	var transaction models.BlockchainTransaction
	_ = json.Unmarshal(wsResp.Data, &transaction)
	if transaction.Status != models.TX_STATUS_CONFIRMED {
		t.Errorf("expected tx status CONFIRMED, got %v", transaction.Status)
	}
}

func TestVoteList(t *testing.T) {
	reqUrl := TEST_BASE_URL + "/votes?algo_cid=" + TEST_ALGO_CID

	resp, err := http.Get(reqUrl)
	if err != nil {
		t.Fatalf("GET /votes failed: %v", err)
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	var apiResp responser.ResponseRaw
	if err := json.Unmarshal(body, &apiResp); err != nil {
		t.Fatalf("Failed to unmarshal list response: %v", err)
	}
	if apiResp.Code != bizcode.SUCCESS {
		t.Fatalf("Expected list SUCCESS, got: %v", apiResp.Code)
	}

	var votes []models.Vote
	if err := json.Unmarshal(apiResp.Data, &votes); err != nil {
		t.Fatalf("Failed to decode vote list: %v", err)
	}
	t.Logf("Found %d votes for algoCid=%v", len(votes), TEST_ALGO_CID)
}

func TestVoteSetVotingDuration(t *testing.T) {
	duration := rand.Intn(290) + 10

	req := types.SetVotingDurationReq{
		Duration: int64(duration),
	}
	body, _ := json.Marshal(req)

	resp, err := http.Post(TEST_BASE_URL+"/set-voting-duration", "application/json", bytes.NewBuffer(body))
	if err != nil {
		t.Fatalf("POST /set-voting-duration failed: %v", err)
	}
	defer resp.Body.Close()

	respBody, _ := io.ReadAll(resp.Body)

	var apiResp responser.Response
	if err := json.Unmarshal(respBody, &apiResp); err != nil {
		t.Fatalf("Failed to parse response: %v", err)
	}

	if apiResp.Code != bizcode.SUCCESS {
		t.Fatalf("Expected SUCCESS, got: %v", apiResp.Code)
	}

	txHash, ok := apiResp.Data.(string)
	if !ok {
		t.Fatalf("Expected txHash string, got %T", apiResp.Data)
	}
	t.Logf("Tx hash: %v", txHash)

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()

	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			t.Fatalf("Timeout waiting for votingDuration to become %d", int64(duration))
		case <-ticker.C:
			actual := votingDuration(t)
			if actual.Int64() == int64(duration) {
				t.Logf("VotingDuration matched: %d", actual.Int64())
				return
			}
		}
	}
}

func voteTestSetup(t *testing.T) (*contracts.ContractCaller, *contracts.AlgorithmReview, *internal.Config) {
	ctx := t.Context()
	config, err := internal.LoadConfigFromEnv()
	if err != nil {
		t.Fatalf("Failed to load config: %v", err)
	}

	keyVault := tee.NewKeyVaultFromConfig(tee.ClientKind(config.DstackClientType))
	ethAcc, err := keyVault.DeriveEthereumAccount(ctx, tee.KeyCtxTEEContractOwner)
	if err != nil {
		t.Fatal(err)
	}
	fundingPrivKey, err := crypto.HexToECDSA(config.OfficialAccountPrivateKey)
	if err != nil {
		t.Fatalf("Failed to create funding private key: %v", err)
	}
	caller, err := contracts.NewContractCaller(
		config.EthHttpUrl, config.EthWsUrl, config.ChainId,
		keyVault,
		fundingPrivKey, 0.005, 0.1,
	)
	mysqlDb, err := db.NewMysqlDb(config.MysqlDsn)
	if err != nil {
		log.Fatalf("Failed to create mysql client: %v", err)
	}
	err = caller.EnsureContractsDeployed(ctx, mysqlDb)
	if err != nil {
		log.Fatalf("Failed to ensure contracts deployed: %v", err)
	}

	err = caller.EnsureWalletFunded(ctx, ethAcc.Address)
	if err != nil {
		t.Fatal(err)
	}

	ctrt, err := contracts.NewAlgorithmReview(caller.AlgoReviewCtrtAddr(), caller.HttpClient())
	if err != nil {
		t.Fatal(err)
	}
	return caller, ctrt, config
}

func createFundedTempAccount(t *testing.T, ctrt *contracts.ContractCaller) (*ecdsa.PrivateKey, common.Address) {
	privKey, err := crypto.GenerateKey()
	if err != nil {
		t.Fatal(err)
	}
	address := crypto.PubkeyToAddress(privKey.PublicKey)
	if err = ctrt.EnsureWalletFunded(t.Context(), address.Hex()); err != nil {
		t.Fatal(err)
	}

	return privKey, address
}

func vote(t *testing.T, cid string, isApproved bool) *ethtypes.Transaction {
	caller, ctrt, config := voteTestSetup(t)
	tmpAccountPrivKey, tmpAccountAddress := createFundedTempAccount(t, caller)
	// Appoint the tee eth account as a member of the committee
	setCommitteeMember(t, tmpAccountAddress.Hex(), true)

	txOpts, err := bind.NewKeyedTransactorWithChainID(tmpAccountPrivKey, big.NewInt(config.ChainId))
	if err != nil {
		t.Fatal(err)
	}

	tx, err := ctrt.Vote(txOpts, cid, isApproved)
	if err != nil {
		t.Fatal(err)
	}
	return tx
}

func votingDuration(t *testing.T) *big.Int {
	_, ctrt, _ := voteTestSetup(t)
	tx, err := ctrt.VotingDuration(nil)
	if err != nil {
		t.Fatal(err)
	}
	return tx
}
