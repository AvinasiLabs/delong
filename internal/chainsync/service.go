package chainsync

import (
	"context"
	"delong/internal/control"
	"delong/pkg/contracts"
	"delong/pkg/ws"
	"log"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

type ChainsyncService struct {
	name  string
	state *ChainsyncServiceState
}

type ChainsyncServiceState struct {
	notifier  *ws.Notifier
	ethCaller *contracts.ContractCaller
}

const SERVICE_NAME = "chainsync-service"

func NewChainsyncService() *ChainsyncService {
	return &ChainsyncService{
		name:  SERVICE_NAME,
		state: &ChainsyncServiceState{},
	}
}

func (s *ChainsyncService) Name() string {
	return s.name
}

func (s *ChainsyncService) Init(ctx context.Context, gs *control.ServiceState) error {
	s.state.notifier = gs.Notifier
	s.state.ethCaller = gs.EthCaller
	return nil
}

func (s *ChainsyncService) Start(ctx context.Context) error {
	go s.listenDataRegistered(ctx)

	// todo listen other events

	log.Println("Chainsync service started")

	return nil
}

func (s *ChainsyncService) Stop(ctx context.Context) error {
	log.Println("Stopping...")
	return nil
}

func (s *ChainsyncService) listenDataRegistered(ctx context.Context) {
	log.Println("Watching DataRegistered...")

	contracts.WatchEventLoop(
		ctx,
		func(opts *bind.WatchOpts, ch chan *contracts.DataContributionDataRegistered) (ethereum.Subscription, error) {
			ctr, err := contracts.NewDataContribution(
				s.state.ethCaller.ContractAddress(contracts.CTRKEY_DATA_CONTRIBUTION),
				s.state.ethCaller.WsClient(),
			)
			if err != nil {
				return nil, err
			}
			return ctr.WatchDataRegistered(opts, ch, nil, nil)
		},
		func(evt *contracts.DataContributionDataRegistered) {
			txHash := evt.Raw.TxHash.Hex()
			log.Printf("Received event tx=%s", txHash)

			receipt, err := s.state.ethCaller.HttpClient().TransactionReceipt(ctx, evt.Raw.TxHash)
			if err != nil {
				log.Printf("Receipt fetch failed: %v", err)
				return
			}
			if receipt.Status != 1 {
				s.state.notifier.PushStatus(txHash, ws.StatusFailed)
				return
			}

			// TODO: persist cid, dataset, contributor

			s.state.notifier.PushStatus(txHash, ws.StatusConfirmed)
		},
	)
}
