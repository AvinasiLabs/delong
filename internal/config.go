package internal

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Config struct {
	IpfsApiAddr string
	EthHttpUrl  string
	EthWsUrl    string
	ChainId     int64

	MinioEndpoint  string
	MinioAccessKey string
	MinioSecretKey string

	DiagnosticSrvEndpoint string

	MysqlDsn                  string
	OfficialAccountPrivateKey string

	JwtSecret string
}

func NewConfig(
	ipfsApiAddr string,
	ethHttpUrl, ethWsUrl string, chainId int64,
	minioEndpoint, minioAccessKey, minioSecretKey string,
	diagnosticSrvEndpoint string,
	mysqlDsn string,
	officialAccountPrivateKey string,
	jwtSecret string,
) *Config {
	return &Config{
		IpfsApiAddr:               ipfsApiAddr,
		EthHttpUrl:                ethHttpUrl,
		EthWsUrl:                  ethWsUrl,
		ChainId:                   chainId,
		MinioEndpoint:             minioEndpoint,
		MinioAccessKey:            minioAccessKey,
		MinioSecretKey:            minioSecretKey,
		DiagnosticSrvEndpoint:     diagnosticSrvEndpoint,
		MysqlDsn:                  mysqlDsn,
		OfficialAccountPrivateKey: officialAccountPrivateKey,
		JwtSecret:                 jwtSecret,
	}
}

const (
	ENVKEY_IPFS_ADDR    = "IPFS_ADDR"
	ENVKEY_CHAIN_ID     = "CHAIN_ID"
	ENVKEY_ETH_HTTP_URL = "ETH_HTTP_URL"
	ENVKEY_ETH_WS_URL   = "ETH_WS_URL"

	ENVKEY_MINIO_ENDPOINT = "MINIO_ENDPOINT"
	ENVKEY_MINIO_AK       = "MINIO_AK"
	ENVKEY_MINIO_SK       = "MINIO_SK"

	ENVKEY_DIAGNOSTIC_SRV_ENDPOINT = "DIAGNOSTIC_SRV_ENDPOINT"

	ENVKEY_MYSQL_DSN = "MYSQL_DSN"

	ENVKEY_OFFICIAL_ACCOUNT_PRIVATE_KEY = "OFFICIAL_ACCOUNT_PRIVATE_KEY"

	ENVKEY_JWT_SECRET = "JWT_SECRET"
)

func LoadConfigFromEnv() (*Config, error) {
	ethHttpUrl := os.Getenv(ENVKEY_ETH_HTTP_URL)
	ipfsApiAddr := os.Getenv(ENVKEY_IPFS_ADDR)
	ethWsUrl := os.Getenv(ENVKEY_ETH_WS_URL)
	chainIdStr := os.Getenv(ENVKEY_CHAIN_ID)
	chainId, err := strconv.ParseInt(chainIdStr, 10, 64)
	if err != nil {
		return nil, err
	}

	minioEndpoint := os.Getenv(ENVKEY_MINIO_ENDPOINT)
	minioAccessKey := os.Getenv(ENVKEY_MINIO_AK)
	minioSecretKey := os.Getenv(ENVKEY_MINIO_SK)
	diagnosticSrvEndpoint := os.Getenv(ENVKEY_DIAGNOSTIC_SRV_ENDPOINT)
	mysqlDsn := os.Getenv(ENVKEY_MYSQL_DSN)
	officialAccountPk := os.Getenv(ENVKEY_OFFICIAL_ACCOUNT_PRIVATE_KEY)

	jwtSecret := os.Getenv(ENVKEY_JWT_SECRET)
	return NewConfig(
		ipfsApiAddr,
		ethHttpUrl, ethWsUrl, chainId,
		minioEndpoint, minioAccessKey, minioSecretKey,
		diagnosticSrvEndpoint,
		mysqlDsn,
		officialAccountPk,
		jwtSecret,
	), nil
}

func (c *Config) String() string {
	var builder strings.Builder

	builder.WriteString("\nConfiguration:\n")
	builder.WriteString(fmt.Sprintf("\tIPFS API Address: %s\n", c.IpfsApiAddr))
	builder.WriteString(fmt.Sprintf("\tEthereum RPC URL: %s\n", c.EthHttpUrl))
	builder.WriteString(fmt.Sprintf("\tEthereum WS URL: %s\n", c.EthWsUrl))
	builder.WriteString(fmt.Sprintf("\tChain ID: %d\n", c.ChainId))
	builder.WriteString(fmt.Sprintf("\tMinio Endpoint: %s\n", c.MinioEndpoint))
	builder.WriteString(fmt.Sprintf("\tMinio Access Key: %s\n", c.MinioAccessKey))
	builder.WriteString(fmt.Sprintf("\tMinio Secret Key: %s\n", c.MinioSecretKey))
	builder.WriteString(fmt.Sprintf("\tDiagnostic Service Endpoint: %s\n", c.DiagnosticSrvEndpoint))
	builder.WriteString(fmt.Sprintf("\tMySQL DSN: %s\n", c.MysqlDsn))
	builder.WriteString(fmt.Sprintf("\tOfficial Account Private Key: %s\n", c.OfficialAccountPrivateKey))
	builder.WriteString(fmt.Sprintf("\tJWT Secret: [HIDDEN]\n"))
	return builder.String()
}
