package main

import (
	"context"
	"delong/internal"
	"delong/internal/api"
	"delong/internal/chainsync"
	"delong/internal/control"
	"delong/pkg/contracts"
	"delong/pkg/db"
	"delong/pkg/tee"
	"delong/pkg/ws"
	"log"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func main() {
	config, err := internal.NewConfigFromEnv()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	log.Println(config)

	ipfsStore, err := db.NewIpfsStore(config.IpfsApiAddr)
	if err != nil {
		log.Fatalf("Failed to create ipfs client: %v", err)
	}

	ethCaller, err := contracts.NewContractCaller(config.EthHttpUrl, config.EthWsUrl, config.ChainId, config.CtrAddr)
	if err != nil {
		log.Fatalf("Failed to create contract caller: %v", err)
	}

	keyVault := tee.NewKeyVault()

	hub := ws.NewHub()
	notifier := ws.NewNotifier(hub)

	state := control.NewServiceState(ipfsStore, ethCaller, keyVault, notifier)

	apiService := api.NewApiService(":8080", config.CtrAddr)
	chainsyncService := chainsync.NewChainsyncService()

	srvMgr := control.NewServiceManager(state, apiService, chainsyncService)

	srvMgr.Run(context.Background())
}
