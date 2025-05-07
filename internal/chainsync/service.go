package chainsync

import (
	"context"
	"delong/internal/control"
	"delong/internal/model"
	"delong/pkg/contracts"
	"delong/pkg/ws"
	"log"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
	"gorm.io/gorm"
)

type ChainsyncService struct {
	name      string
	notifier  *ws.Notifier
	ethCaller *contracts.ContractCaller
	db        *gorm.DB
}

type ChainsyncServiceOptions struct {
	Blockchain   *control.BlockchainDeps
	Notification *control.NotificationDeps
	Storage      *control.StorageDeps
}

const SERVICE_NAME = "chainsync-service"

func NewChainsyncService(opts ChainsyncServiceOptions) *ChainsyncService {
	return &ChainsyncService{
		name:      SERVICE_NAME,
		notifier:  opts.Notification.Notifier,
		ethCaller: opts.Blockchain.EthCaller,
		db:        opts.Storage.MysqlDb,
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
			// First get block info (needed for both success and failure cases)
			blockNumber := evt.Raw.BlockNumber

			// We need to fetch the block to get its timestamp
			block, err := s.ethCaller.HttpClient().BlockByNumber(ctx, new(big.Int).SetUint64(blockNumber))
			if err != nil {
				log.Printf("Failed to fetch block details: %v", err)
				return
			}

			blockTime := time.Unix(int64(block.Time()), 0)

			// Determine transaction status based on receipt
			var status string
			if receipt.Status != types.ReceiptStatusSuccessful {
				status = model.TxStatusFailed
			} else {
				status = model.TxStatusConfirmed
			}

			// Update transaction status
			err = model.UpdateTransactionStatus(s.db, txHash, status, &blockNumber, &blockTime)
			if err != nil {
				log.Printf("Failed to update transaction status to %s: %v", status, err)
				return
			}

			s.notifier.PushStatus(txHash, status)
		},
	)
}
