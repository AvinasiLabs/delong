package control

import (
	"context"
	"delong/pkg/contracts"
	"delong/pkg/db"
	"delong/pkg/tee"
	"delong/pkg/ws"
	"log"

	"golang.org/x/sync/errgroup"
)

type Service interface {
	Name() string
	Init(context.Context, *ServiceState) error
	Start(context.Context) error
	Stop(context.Context) error
}

type ServiceState struct {
	IpfsStore *db.IpfsStore
	EthCaller *contracts.ContractCaller
	KeyVault  *tee.KeyVault
	Notifier  *ws.Notifier
}

func NewServiceState(
	ipfsClient *db.IpfsStore,
	ethCaller *contracts.ContractCaller,
	keyVault *tee.KeyVault,
	notifier *ws.Notifier,
) *ServiceState {
	return &ServiceState{
		IpfsStore: ipfsClient,
		EthCaller: ethCaller,
		KeyVault:  keyVault,
		Notifier:  notifier,
	}
}

type ServiceManager struct {
	services []Service
	state    *ServiceState
}

func NewServiceManager(state *ServiceState, srvs ...Service) *ServiceManager {
	return &ServiceManager{
		state:    state,
		services: srvs,
	}
}

func (sm *ServiceManager) Run(ctx context.Context) error {
	g, ctx := errgroup.WithContext(ctx)
	for _, srv := range sm.services {
		g.Go(func() error {
			srvname := srv.Name()
			log.Printf("Initializing %v ...", srvname)
			srv.Init(ctx, sm.state)
			log.Printf("Starting %v ...", srvname)
			return srv.Start(ctx)
		})
	}

	err := g.Wait()

	for _, srv := range sm.services {
		_ = srv.Stop(ctx)
	}

	return err
}
