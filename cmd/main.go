package main

import (
	"context"
	"delong/internal"
	"delong/internal/services/api"
	"delong/internal/services/chainsync"
	"delong/internal/services/runtime"
	"delong/pkg/analysis"
	"delong/pkg/contracts"
	"delong/pkg/db"
	"delong/pkg/schedule"
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

	reportAnalyzer := analysis.NewReportAnalyzer(config.DiagnosticSrvEndpoint)

	keyVault := tee.NewKeyVault()

	fundingPrivKey, err := crypto.HexToECDSA(config.OfficialAccountPrivateKey)
	if err != nil {
		log.Fatalf("Failed to create funding private key: %v", err)
	}

	ctrCaller, err := contracts.NewContractCaller(
		config.EthHttpUrl, config.EthWsUrl, config.ChainId,
		keyVault,
		fundingPrivKey, 0.005, 0.1,
	)
	if err != nil {
		log.Fatalf("Failed to create contract caller: %v", err)
	}

	err = ctrCaller.EnsureContractsDeployed(ctx, mysqlDb)
	if err != nil {
		log.Fatalf("Failed to ensure contracts deployed: %v", err)
	}

	hub := ws.NewHub()
	notifier := ws.NewNotifier(hub)

	algoScheduler, err := schedule.NewAlgoScheduler()
	if err != nil {
		log.Fatalf("Failed to create algo scheduler: %v", err)
	}

	apiService := api.NewService(api.ApiServiceOptions{
		Addr:           ":8080",
		IpfsStore:      ipfsStore,
		MinioStore:     minioStore,
		MysqlDb:        mysqlDb,
		CtrCaller:      ctrCaller,
		KeyVault:       keyVault,
		Notifier:       notifier,
		ReportAnalyzer: reportAnalyzer,
	})

	chainsyncService := chainsync.NewService(chainsync.ChainsyncServiceOptions{
		CtrCaller:     ctrCaller,
		KeyVault:      keyVault,
		Notifier:      notifier,
		MysqlDb:       mysqlDb,
		AlgoScheduler: algoScheduler,
	})

	runtimeService := runtime.NewService(runtime.RuntimeServiceOptions{
		Loader:        runtime.NewDatasetLoader("/data/delong_dataset", mysqlDb),
		Db:            mysqlDb,
		IpfsStore:     ipfsStore,
		CtrCaller:     ctrCaller,
		AlgoScheduler: algoScheduler,
	})
	if err != nil {
		log.Fatalf("Failed to create runtime service: %v", err)
	}

	srvMgr := internal.NewServiceManager(apiService, chainsyncService, runtimeService)

	srvMgr.Run(context.Background())
}
