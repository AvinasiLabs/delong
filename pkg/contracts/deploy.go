package contracts

import (
	"context"
	"delong/internal/model"
	"delong/pkg/tee"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"
	"gorm.io/gorm"
)

// EnsureDataContributionDeployed deploys the DataContribution contract if it does not already exist.
// It checks the MySQL record first, and stores the deployed address if newly created.
func EnsureDataContributionDeployed(
	ctx context.Context,
	db *gorm.DB,
	client *ethclient.Client,
	chainID *big.Int,
	vault *tee.KeyVault,
) (string, error) {
	addr, err := model.GetContractAddress(db, CTRKEY_DATA_CONTRIBUTION)
	if err == nil && addr != "" {
		log.Printf("Found existing %s contract at %s", CTRKEY_DATA_CONTRIBUTION, addr)
		return addr, nil
	}

	ethAcc, err := vault.DeriveEthereumAccount(ctx, tee.KeyCtxTEEContractOwner)
	if err != nil {
		return "", fmt.Errorf("failed to derive TEE Ethereum account: %w", err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(ethAcc.PrivateKey, chainID)
	if err != nil {
		return "", fmt.Errorf("failed to create transactor: %w", err)
	}

	contractAddr, tx, _, err := DeployDataContribution(auth, client)
	if err != nil {
		return "", fmt.Errorf("contract deployment failed: %w", err)
	}
	log.Printf("Deployed %s contract at %s (tx: %s)", CTRKEY_DATA_CONTRIBUTION, contractAddr.Hex(), tx.Hash().Hex())

	if err := model.SaveContractAddress(db, CTRKEY_DATA_CONTRIBUTION, contractAddr.Hex()); err != nil {
		return "", fmt.Errorf("failed to save contract address to DB: %w", err)
	}

	return contractAddr.Hex(), nil
}
