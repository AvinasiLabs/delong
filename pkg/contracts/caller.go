package contracts

import (
	"delong/pkg/tee"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

const (
	CTRKEY_DATA_CONTRIBUTION = "DataContribution"
	CTRKEY_ALGORITHM_REVIEW  = "AlgorithmReview"
)

type ContractCaller struct {
	httpUrl      string
	wsUrl        string
	chainId      *big.Int
	httpClient   *ethclient.Client
	wsClient     *ethclient.Client
	contractAddr map[string]common.Address
}

func NewContractCaller(httpUrl, wsUrl string, chainId int64, contractAddr map[string]common.Address) (*ContractCaller, error) {
	httpClient, err := ethclient.Dial(httpUrl)
	if err != nil {
		return nil, err
	}
	wsClient, err := ethclient.Dial(wsUrl)
	if err != nil {
		return nil, err
	}
	return &ContractCaller{
		httpUrl:      httpUrl,
		wsUrl:        wsUrl,
		chainId:      big.NewInt(chainId),
		httpClient:   httpClient,
		wsClient:     wsClient,
		contractAddr: contractAddr,
	}, nil
}

func (c *ContractCaller) ContractAddress(key string) common.Address {
	return c.contractAddr[key]
}

func (c *ContractCaller) HttpClient() *ethclient.Client {
	return c.httpClient
}

func (c *ContractCaller) WsClient() *ethclient.Client {
	return c.wsClient
}

func (c *ContractCaller) RegisterData(ethAcc tee.EthereumAccount, userAccount common.Address, cid string, dataset string) (*types.Transaction, error) {
	ctct, err := NewDataContribution(c.ContractAddress(CTRKEY_DATA_CONTRIBUTION), c.httpClient)
	if err != nil {
		return nil, err
	}

	txOpts, err := bind.NewKeyedTransactorWithChainID(ethAcc.PrivateKey, c.chainId)
	if err != nil {
		return nil, err
	}

	return ctct.RegisterData(txOpts, userAccount, cid, dataset)
}
