//go:build integration
// +build integration

package api

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"crypto/rand"
	"delong/internal"
	"delong/internal/models"
	"delong/internal/types"
	"delong/pkg/bizcode"
	"delong/pkg/contracts"
	"delong/pkg/db"
	"delong/pkg/responser"
	"delong/pkg/tee"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/big"
	"net/http"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/mr-tron/base58"
)

func TestVoteCreateAndList(t *testing.T) {
	// Generate unique CID for this test to avoid conflicts
	testAlgoCid := generateRandomCID()

	// Step 1: Create a vote
	tx := vote(t, testAlgoCid, true)
	msg := waitForWsConfirmation(t, tx.Hash().Hex(), 1*time.Minute)

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

	t.Logf("Vote created successfully for CID: %s", testAlgoCid)

	// Step 2: List votes for the algorithm
	reqUrl := TEST_BASE_URL + "/votes?algo_cid=" + testAlgoCid

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

	// Verify we have at least one vote (the one we just created)
	if len(votes) == 0 {
		t.Fatalf("Expected at least 1 vote for algoCid=%v, got %d", testAlgoCid, len(votes))
	}

	t.Logf("Found %d votes for algoCid=%v", len(votes), testAlgoCid)

	// Optional: Verify the vote details
	foundOurVote := false
	for _, vote := range votes {
		if vote.AlgoCid == testAlgoCid {
			foundOurVote = true
			t.Logf("Vote details - AlgoCid: %s, IsApproved: %v", vote.AlgoCid, vote.Approve)
			break
		}
	}

	if !foundOurVote {
		t.Errorf("Did not find our vote with AlgoCid=%s in the returned list", testAlgoCid)
	}
}

func TestVoteSetVotingDuration(t *testing.T) {
	duration := 3600

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

	t.Logf("DstackClientType: %v", config.DstackClientType)
	keyVault := tee.NewKeyVaultFromConfig(tee.ClientKind(config.DstackClientType))
	if keyVault == nil {
		t.Fatal("Failed to create key vault")
	}

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
		fundingPrivKey, 0.005, 0.01,
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
		t.Fatalf("Failed to fund wallet: %v", err)
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

// generateRandomCID generates a random IPFS CID-like string for testing
func generateRandomCID() string {
	// Generate 32 random bytes
	randomBytes := make([]byte, 32)
	_, err := rand.Read(randomBytes)
	if err != nil {
		panic(fmt.Sprintf("Failed to generate random bytes: %v", err))
	}

	// Encode to base58 and add "Qm" prefix to make it look like a real CID
	encoded := base58.Encode(randomBytes)
	return "Qm" + encoded // Use the full encoded string
}

func vote(t *testing.T, cid string, isApproved bool) *ethtypes.Transaction {
	caller, ctrt, config := voteTestSetup(t)
	// Ensure algo submitted
	scientistAddr := common.HexToAddress("0xa0Ee7A142d267C1f36714E4a8F75612F20a79720")
	executionId := uint(time.Now().UnixNano())
	dataset := "test-dataset"

	tx, err := caller.SubmitAlgorithm(t.Context(), executionId, scientistAddr, cid, dataset)
	if err != nil {
		t.Fatalf("Failed to submit algo: %v", err)
	}

	t.Logf("Submitted algorithm %s with tx: %s", cid, tx.Hash().Hex())

	tmpAccountPrivKey, tmpAccountAddress := createFundedTempAccount(t, caller)
	// Appoint the tee eth account as a member of the committee
	setCommitteeMember(t, tmpAccountAddress.Hex(), true)

	txOpts, err := bind.NewKeyedTransactorWithChainID(tmpAccountPrivKey, big.NewInt(config.ChainId))
	if err != nil {
		t.Fatal(err)
	}

	tx, err = ctrt.Vote(txOpts, cid, isApproved)
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
