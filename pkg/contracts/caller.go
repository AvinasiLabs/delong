package contracts

import (
	"context"
	"crypto/ecdsa"
	"delong/internal/model"
	"delong/pkg/tee"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"gorm.io/gorm"
)

const (
	CTRKEY_DATA_CONTRIBUTION = "data-contribution"
	CTRKEY_ALGORITHM_REVIEW  = "algorithm-review"
)

type ContractAddr struct {
	DataContribution common.Address
	AlgorithmReview  common.Address
}

type ContractCaller struct {
	httpUrl      string
	wsUrl        string
	chainId      *big.Int
	httpClient   *ethclient.Client
	wsClient     *ethclient.Client
	contractAddr *ContractAddr

	keyVault *tee.KeyVault

	fundingPrivKey *ecdsa.PrivateKey
	thresholdEth   float64
	topUpEth       float64
}

func NewContractCaller(httpUrl, wsUrl string, chainId int64, keyVault *tee.KeyVault,
	fundingPrivKey *ecdsa.PrivateKey, thresholdEth, topUpEth float64) (*ContractCaller, error) {
	httpClient, err := ethclient.Dial(httpUrl)
	if err != nil {
		return nil, err
	}
	wsClient, err := ethclient.Dial(wsUrl)
	if err != nil {
		return nil, err
	}
	return &ContractCaller{
		httpUrl:        httpUrl,
		wsUrl:          wsUrl,
		chainId:        big.NewInt(chainId),
		httpClient:     httpClient,
		wsClient:       wsClient,
		contractAddr:   &ContractAddr{},
		keyVault:       keyVault,
		fundingPrivKey: fundingPrivKey,
		thresholdEth:   thresholdEth,
		topUpEth:       topUpEth,
	}, nil
}

// func (c *ContractCaller) ContractAddress(key string) common.Address {
// 	return c.contractAddr[key]
// }

func (c *ContractCaller) HttpClient() *ethclient.Client {
	return c.httpClient
}

func (c *ContractCaller) WsClient() *ethclient.Client {
	return c.wsClient
}

func (c *ContractCaller) DataContributionCtrtAddr() common.Address {
	return c.contractAddr.DataContribution
}

// EthToWei converts an ETH amount to a wei amount.
func EthToWei(eth float64) *big.Int {
	f := new(big.Float).Mul(big.NewFloat(eth), big.NewFloat(1e18))
	result := new(big.Int)
	f.Int(result)
	return result
}

// WeiToEthString converts a wei amount to a string representation of ETH.
func WeiToEthString(wei *big.Int) string {
	f := new(big.Float).Quo(new(big.Float).SetInt(wei), big.NewFloat(1e18))
	return f.Text('f', 6)
}

func (c *ContractCaller) EnsureTeeAccountFunded(ctx context.Context, acc *tee.EthereumAccount) error {
	toAddr := common.HexToAddress(acc.Address)

	balanceWei, err := c.httpClient.BalanceAt(ctx, toAddr, nil)
	if err != nil {
		return err
	}
	log.Printf("Tee eth account balance: %s", WeiToEthString(balanceWei))

	thresholdWei := EthToWei(c.thresholdEth)
	if balanceWei.Cmp(thresholdWei) >= 0 {
		return nil // balance is sufficient
	}

	topUpWei := EthToWei(c.topUpEth)
	fromAddr := crypto.PubkeyToAddress(c.fundingPrivKey.PublicKey)
	nonce, err := c.httpClient.PendingNonceAt(ctx, fromAddr)
	if err != nil {
		return err
	}
	gasPrice, err := c.httpClient.SuggestGasPrice(ctx)
	if err != nil {
		return err
	}

	tx := types.NewTransaction(
		nonce,
		toAddr,
		topUpWei,
		21000, // normal transfer fixed gas
		gasPrice,
		nil,
	)

	signer := types.LatestSignerForChainID(c.chainId)
	signedTx, err := types.SignTx(tx, signer, c.fundingPrivKey)
	if err != nil {
		return err
	}

	err = c.httpClient.SendTransaction(ctx, signedTx)
	if err != nil {
		return err
	}

	finalBalance, err := c.httpClient.BalanceAt(ctx, toAddr, nil)
	if err != nil {
		log.Printf("Warning: funded but failed to query final balance for %s: %v", toAddr.Hex(), err)
	} else {
		log.Printf(
			"Auto-funded TEE account %s (tx: %s), new balance: %s ETH",
			toAddr.Hex(), signedTx.Hash().Hex(), WeiToEthString(finalBalance),
		)
	}

	return nil
}

func (c *ContractCaller) EnsureContractsDeployed(ctx context.Context, db *gorm.DB) error {
	ethAcc, err := c.keyVault.DeriveEthereumAccount(ctx, tee.KeyCtxTEEContractOwner)
	if err != nil {
		return err
	}

	err = c.EnsureTeeAccountFunded(ctx, ethAcc)
	if err != nil {
		return err
	}

	auth, err := bind.NewKeyedTransactorWithChainID(ethAcc.PrivateKey, c.chainId)
	if err != nil {
		return err
	}

	// DataContribution
	addrStr, err := model.GetContractAddress(db, CTRKEY_DATA_CONTRIBUTION)
	if err != nil {
		return err
	}
	if addrStr == "" {
		addr, tx, _, err := DeployDataContribution(auth, c.httpClient)
		if err != nil {
			return err
		}
		log.Printf("Deployed DataContribution at %s (tx: %s)", addr.Hex(), tx.Hash().Hex())
		if err := model.SaveContractAddress(db, CTRKEY_DATA_CONTRIBUTION, addr.Hex()); err != nil {
			return err
		}
		c.contractAddr.DataContribution = addr
	} else {
		c.contractAddr.DataContribution = common.HexToAddress(addrStr)
	}

	// AlgorithmReview
	// addrStr, _ = model.GetContractAddress(db, CTRKEY_ALGORITHM_REVIEW)
	// if addrStr == "" {
	// 	addr, tx, _, err := DeployAlgorithmReview(auth, c.httpClient)
	// 	if err != nil {
	// 		return err
	// 	}
	// 	log.Printf("Deployed AlgorithmReview at %s (tx: %s)", addr.Hex(), tx.Hash().Hex())
	// 	if err := model.SaveContractAddress(db, CTRKEY_ALGORITHM_REVIEW, addr.Hex()); err != nil {
	// 		return err
	// 	}
	// 	c.contractAddr.AlgorithmReview = addr
	// } else {
	// 	c.contractAddr.AlgorithmReview = common.HexToAddress(addrStr)
	// }

	return nil
}

func (c *ContractCaller) RegisterData(ctx context.Context, userAccount common.Address, cid string, dataset string) (*types.Transaction, error) {
	ethAcc, err := c.keyVault.DeriveEthereumAccount(ctx, tee.KeyCtxTEEContractOwner)
	if err != nil {
		return nil, err
	}
	log.Printf("Tee eth account: %v", ethAcc.Address)

	err = c.EnsureTeeAccountFunded(ctx, ethAcc)
	if err != nil {
		return nil, err
	}

	ctct, err := NewDataContribution(c.contractAddr.DataContribution, c.httpClient)
	if err != nil {
		return nil, err
	}

	txOpts, err := bind.NewKeyedTransactorWithChainID(ethAcc.PrivateKey, c.chainId)
	if err != nil {
		return nil, err
	}

	return ctct.RegisterData(txOpts, userAccount, cid, dataset)
}
