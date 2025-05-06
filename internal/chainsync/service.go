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
	name      string
	notifier  *ws.Notifier
	ethCaller *contracts.ContractCaller
}

type ChainsyncServiceOptions struct {
	Blockchain   *control.BlockchainDeps
	Notification *control.NotificationDeps
}

const SERVICE_NAME = "chainsync-service"

func NewChainsyncService(opts ChainsyncServiceOptions) *ChainsyncService {
	return &ChainsyncService{
		name:      SERVICE_NAME,
		notifier:  opts.Notification.Notifier,
		ethCaller: opts.Blockchain.EthCaller,
	}
}

func (s *ChainsyncService) Name() string {
	return s.name
}

func (s *ChainsyncService) Init(ctx context.Context) error {
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
				s.ethCaller.DataContributionCtrtAddr(),
				s.ethCaller.WsClient(),
			)
			if err != nil {
				return nil, err
			}
			return ctr.WatchDataRegistered(opts, ch, nil, nil)
		},
		func(evt *contracts.DataContributionDataRegistered) {
			txHash := evt.Raw.TxHash.Hex()
			log.Printf("Received event tx=%s", txHash)

			receipt, err := s.ethCaller.HttpClient().TransactionReceipt(ctx, evt.Raw.TxHash)
			if err != nil {
				log.Printf("Receipt fetch failed: %v", err)
				return
			}
			if receipt.Status != 1 {
				s.notifier.PushStatus(txHash, ws.StatusFailed)
				return
			}

			// TODO: persist cid, dataset, contributor

			s.notifier.PushStatus(txHash, ws.StatusConfirmed)
		},
	)
}
