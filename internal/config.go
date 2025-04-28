package internal

import (
	"delong/pkg/contracts"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/common"
)

type Config struct {
	IpfsApiAddr string
	EthHttpUrl  string
	EthWsUrl    string
	ChainId     int64
	CtrAddr     map[string]common.Address
}

func NewConfig(ipfsApiAddr, ethHttpUrl, ethWsUrl string, chainId int64, dataContributionCtrAddr, algorithmCtrAddr common.Address) *Config {
	return &Config{
		IpfsApiAddr: ipfsApiAddr,
		EthHttpUrl:  ethHttpUrl,
		EthWsUrl:    ethWsUrl,
		ChainId:     chainId,
		CtrAddr: map[string]common.Address{
			contracts.CTRKEY_DATA_CONTRIBUTION: dataContributionCtrAddr,
			contracts.CTRKEY_ALGORITHM_REVIEW:  algorithmCtrAddr,
		},
	}
}

const (
	ENVKEY_IPFS_ADDR              = "IPFS_ADDR"
	ENVKEY_CHAIN_ID               = "CHAIN_ID"
	ENVKEY_ETH_HTTP_URL           = "ETH_HTTP_URL"
	ENVKEY_ETH_WS_URL             = "ETH_WS_URL"
	ENVKEY_CONTRIBUTION_CTRT_ADDR = "CONTRIBUTION_CTRT_ADDR"
	ENVKEY_ALGORITHM_CTRT_ADDR    = "ALGORITHM_CTRT_ADDR"

	// ENVKEY_ENC_KEY_HEX            = "ENC_KEY_HEX"
)

func NewConfigFromEnv() (*Config, error) {
	ethHttpUrl := os.Getenv(ENVKEY_ETH_HTTP_URL)
	ipfsApiAddr := os.Getenv(ENVKEY_IPFS_ADDR)
	ethWsUrl := os.Getenv(ENVKEY_ETH_WS_URL)
	chainIdStr := os.Getenv(ENVKEY_CHAIN_ID)
	algorithmCtrtAddr := os.Getenv(ENVKEY_ALGORITHM_CTRT_ADDR)
	contributionCtrtAddr := os.Getenv(ENVKEY_CONTRIBUTION_CTRT_ADDR)
	chainId, err := strconv.ParseInt(chainIdStr, 10, 64)
	if err != nil {
		return nil, err
	}
	return NewConfig(ipfsApiAddr, ethHttpUrl, ethWsUrl, chainId, common.HexToAddress(contributionCtrtAddr), common.HexToAddress(algorithmCtrtAddr)), nil
}

func (c *Config) String() string {
	var builder strings.Builder

	builder.WriteString("\nConfiguration:\n")
	builder.WriteString(fmt.Sprintf("\tIPFS API Address: %s\n", c.IpfsApiAddr))
	builder.WriteString(fmt.Sprintf("\tEthereum RPC URL: %s\n", c.EthHttpUrl))
	builder.WriteString(fmt.Sprintf("\tEthereum WS URL: %s\n", c.EthWsUrl))
	builder.WriteString(fmt.Sprintf("\tChain ID: %d\n", c.ChainId))
	builder.WriteString("\tSmart Contract Addresses:\n")
	for key, addr := range c.CtrAddr {
		builder.WriteString(fmt.Sprintf("\t\t- %s: %s\n", key, addr.Hex()))
	}

	return builder.String()
}
