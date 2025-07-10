package internal

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Config struct {
	DstackClientType          string
	IpfsApiAddr               string
	EthHttpUrl                string
	EthWsUrl                  string
	ChainId                   int64
	DiagnosticSrvAddr         string
	SampleSrvAddr             string
	MysqlDsn                  string
	OfficialAccountPrivateKey string
	UseJwt                    bool
	JwtSecret                 string
}

const (
	ENVKEY_DSTACK_CLIENT_TYPE           = "DSTACK_CLIENT_TYPE"
	ENVKEY_IPFS_ADDR                    = "IPFS_ADDR"
	ENVKEY_CHAIN_ID                     = "CHAIN_ID"
	ENVKEY_ETH_HTTP_URL                 = "ETH_HTTP_URL"
	ENVKEY_ETH_WS_URL                   = "ETH_WS_URL"
	ENVKEY_DIAGNOSTIC_SRV_ADDR          = "DIAGNOSTIC_SRV_ADDR"
	ENVKEY_SAMPLE_SRV_ADDR              = "SAMPLE_SRV_ADDR"
	ENVKEY_MYSQL_DSN                    = "MYSQL_DSN"
	ENVKEY_OFFICIAL_ACCOUNT_PRIVATE_KEY = "OFFICIAL_ACCOUNT_PRIVATE_KEY"
	ENVKEY_USE_JWT                      = "USE_JWT"
	ENVKEY_JWT_SECRET                   = "JWT_SECRET"
)

func LoadConfigFromEnv() (*Config, error) {
	chainIdStr := os.Getenv(ENVKEY_CHAIN_ID)
	chainId, err := strconv.ParseInt(chainIdStr, 10, 64)
	if err != nil {
		return nil, err
	}

	useJwtStr := os.Getenv(ENVKEY_USE_JWT)
	useJwt, err := strconv.ParseBool(useJwtStr)
	if err != nil {
		return nil, err
	}

	config := &Config{
		DstackClientType:          os.Getenv(ENVKEY_DSTACK_CLIENT_TYPE),
		IpfsApiAddr:               os.Getenv(ENVKEY_IPFS_ADDR),
		EthHttpUrl:                os.Getenv(ENVKEY_ETH_HTTP_URL),
		EthWsUrl:                  os.Getenv(ENVKEY_ETH_WS_URL),
		ChainId:                   chainId,
		DiagnosticSrvAddr:         os.Getenv(ENVKEY_DIAGNOSTIC_SRV_ADDR),
		SampleSrvAddr:             os.Getenv(ENVKEY_SAMPLE_SRV_ADDR),
		MysqlDsn:                  os.Getenv(ENVKEY_MYSQL_DSN),
		OfficialAccountPrivateKey: os.Getenv(ENVKEY_OFFICIAL_ACCOUNT_PRIVATE_KEY),
		UseJwt:                    useJwt,
		JwtSecret:                 os.Getenv(ENVKEY_JWT_SECRET),
	}
	return config, nil
}

func (c *Config) String() string {
	var builder strings.Builder
	builder.WriteString("\nConfiguration:\n")
	builder.WriteString(fmt.Sprintf("\tDSTACK_CLIENT_TYPE: %s\n", c.DstackClientType))
	builder.WriteString(fmt.Sprintf("\tIPFS API Address: %s\n", c.IpfsApiAddr))
	builder.WriteString(fmt.Sprintf("\tEthereum RPC URL: %s\n", c.EthHttpUrl))
	builder.WriteString(fmt.Sprintf("\tEthereum WS URL: %s\n", c.EthWsUrl))
	builder.WriteString(fmt.Sprintf("\tChain ID: %d\n", c.ChainId))
	builder.WriteString(fmt.Sprintf("\tDiagnostic Service Addr: %s\n", c.DiagnosticSrvAddr))
	builder.WriteString(fmt.Sprintf("\tSample Service Addr: %s\n", c.SampleSrvAddr))
	builder.WriteString(fmt.Sprintf("\tMySQL DSN: %s\n", c.MysqlDsn))
	builder.WriteString(fmt.Sprintf("\tOfficial Account Private Key: %s\n", c.OfficialAccountPrivateKey))
	builder.WriteString(fmt.Sprintf("\tUse Jwt: %v\n", c.UseJwt))
	builder.WriteString(fmt.Sprintf("\tJwt Secret: %s\n", c.JwtSecret))
	return builder.String()
}
