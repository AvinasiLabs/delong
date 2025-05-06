package main

import (
	"context"
	"delong/internal"
	"delong/internal/api"
	"delong/internal/chainsync"
	"delong/internal/control"
	"delong/pkg/analyzer"
	"delong/pkg/contracts"
	"delong/pkg/db"
	"delong/pkg/tee"
	"delong/pkg/ws"
	"log"

	"github.com/ethereum/go-ethereum/crypto"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func main() {
	ctx := context.Background()
	config, err := internal.LoadConfigFromEnv()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	log.Println(config)

	ipfsStore, err := db.NewIpfsStore(config.IpfsApiAddr)
	if err != nil {
		log.Fatalf("Failed to create ipfs client: %v", err)
	}

	minioStore, err := db.NewMinioStore(
		config.MinioEndpoint,
		db.WithSecure(false),
		db.WithCredentials(config.MinioAccessKey, config.MinioSecretKey),
	)
	if err != nil {
		log.Fatalf("Failed to create minio client: %v", err)
	}

	mysqlDb, err := db.NewMysqlDb(config.MysqlDsn)
	if err != nil {
		log.Fatalf("Failed to create mysql client: %v", err)
	}

	reportAnalyzer := analyzer.NewReportAnalyzer(config.DiagnosticSrvEndpoint)

	keyVault := tee.NewKeyVault()

	fundingPrivKey, err := crypto.HexToECDSA(config.OfficialAccountPrivateKey)
	if err != nil {
		log.Fatalf("Failed to create funding private key: %v", err)
	}

	ethCaller, err := contracts.NewContractCaller(
		config.EthHttpUrl, config.EthWsUrl, config.ChainId,
		keyVault,
		fundingPrivKey, 0.005, 0.1,
	)
	if err != nil {
		log.Fatalf("Failed to create contract caller: %v", err)
	}

	err = ethCaller.EnsureContractsDeployed(ctx, mysqlDb)
	if err != nil {
		log.Fatalf("Failed to ensure contracts deployed: %v", err)
	}

	hub := ws.NewHub()
	notifier := ws.NewNotifier(hub)

	deps := control.NewDependencies(ipfsStore, minioStore, mysqlDb, reportAnalyzer, ethCaller, keyVault, notifier)

	apiService := api.NewApiService(api.ApiServiceOptions{
		Addr:         ":8080",
		Storage:      deps.Storage,
		Blockchain:   deps.Blockchain,
		Notification: deps.Notification,
		Analyzer:     deps.Analyzer,
	})

	chainsyncService := chainsync.NewChainsyncService(chainsync.ChainsyncServiceOptions{
		Blockchain:   deps.Blockchain,
		Notification: deps.Notification,
	})

	srvMgr := control.NewServiceManager(deps, apiService, chainsyncService)

	srvMgr.Run(context.Background())
}
