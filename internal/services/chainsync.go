package services

import (
	"context"
	"delong/internal/models"
	"delong/pkg/contracts"
	"delong/pkg/schedule"
	"delong/pkg/tee"
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
	name          string
	notifier      *ws.Notifier
	ctrCaller     *contracts.ContractCaller
	db            *gorm.DB
	algoScheduler *schedule.AlgoScheduler
}

type ChainsyncServiceOptions struct {
	CtrCaller     *contracts.ContractCaller
	KeyVault      *tee.KeyVault
	Notifier      *ws.Notifier
	MysqlDb       *gorm.DB
	AlgoScheduler *schedule.AlgoScheduler
}

func NewChainsyncService(opts ChainsyncServiceOptions) *ChainsyncService {
	return &ChainsyncService{
		name:          "chainsync-service",
		notifier:      opts.Notifier,
		ctrCaller:     opts.CtrCaller,
		db:            opts.MysqlDb,
		algoScheduler: opts.AlgoScheduler,
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
	go s.listenAlgoSubmitted(ctx)
	go s.listenVoteCasted(ctx)
	go s.listenAlgoResolved(ctx)

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
				s.ctrCaller.DataContributionCtrtAddr(),
				s.ctrCaller.WsClient(),
			)
			if err != nil {
				return nil, err
			}
			return ctr.WatchDataRegistered(opts, ch, nil, nil)
		},
		func(evt *contracts.DataContributionDataRegistered) {
			txHash := evt.Raw.TxHash.Hex()
			log.Printf("Received event tx=%s", txHash)
			var status string
			defer func() {
				if txHash != "" && status != "" {
					if err := s.notifier.PushStatus(txHash, status); err != nil {
						log.Printf("PushStatus failed: tx=%s, err=%v", txHash, err)
					}
				}
			}()

			receipt, err := s.ctrCaller.HttpClient().TransactionReceipt(ctx, evt.Raw.TxHash)
			if err != nil {
				log.Printf("Receipt fetch failed: %v", err)
				return
			}
			// First get block info (needed for both success and failure cases)
			blockNumber := evt.Raw.BlockNumber

			// We need to fetch the block to get its timestamp
			block, err := s.ctrCaller.HttpClient().BlockByNumber(ctx, new(big.Int).SetUint64(blockNumber))
			if err != nil {
				log.Printf("Failed to fetch block details: %v", err)
				return
			}

			blockTime := time.Unix(int64(block.Time()), 0)

			// Determine transaction status based on receipt
			if receipt.Status != types.ReceiptStatusSuccessful {
				status = models.TX_STATUS_FAILED
			} else {
				status = models.TX_STATUS_CONFIRMED
			}

			// Update transaction status
			err = models.UpdateTransactionStatus(s.db, txHash, status, &blockNumber, &blockTime)
			if err != nil {
				log.Printf("Failed to update transaction status to %s: %v", status, err)
				return
			}
		},
	)
}

func (s *ChainsyncService) listenAlgoSubmitted(ctx context.Context) {
	log.Println("Watching AlgorithmSubmitted...")

	contracts.WatchEventLoop(
		ctx,
		func(opts *bind.WatchOpts, ch chan *contracts.AlgorithmReviewAlgorithmSubmitted) (ethereum.Subscription, error) {
			ctr, err := contracts.NewAlgorithmReview(
				s.ctrCaller.AlgoReviewCtrtAddr(),
				s.ctrCaller.WsClient(),
			)
			if err != nil {
				return nil, err
			}
			return ctr.WatchAlgorithmSubmitted(opts, ch, nil, nil)
		},
		func(evt *contracts.AlgorithmReviewAlgorithmSubmitted) {
			txHash := evt.Raw.TxHash.Hex()
			log.Printf("Received event tx=%s", txHash)
			var status string
			defer func() {
				if txHash != "" && status != "" {
					if err := s.notifier.PushStatus(txHash, status); err != nil {
						log.Printf("PushStatus failed: tx=%s, err=%v", txHash, err)
					}
				}
			}()

			receipt, err := s.ctrCaller.HttpClient().TransactionReceipt(ctx, evt.Raw.TxHash)
			if err != nil {
				log.Printf("Receipt fetch failed: %v", err)
				return
			}
			// First get block info (needed for both success and failure cases)
			blockNumber := evt.Raw.BlockNumber

			// We need to fetch the block to get its timestamp
			block, err := s.ctrCaller.HttpClient().BlockByNumber(ctx, new(big.Int).SetUint64(blockNumber))
			if err != nil {
				log.Printf("Failed to fetch block details: %v", err)
				return
			}

			blockTime := time.Unix(int64(block.Time()), 0)

			// Determine transaction status based on receipt
			if receipt.Status != types.ReceiptStatusSuccessful {
				status = models.TX_STATUS_FAILED
			} else {
				status = models.TX_STATUS_CONFIRMED
			}
			// Update transaction status
			err = models.UpdateTransactionStatus(s.db, txHash, status, &blockNumber, &blockTime)
			if err != nil {
				log.Printf("Failed to update transaction status to %s: %v", status, err)
				return
			}
		},
	)
}

func (s *ChainsyncService) listenVoteCasted(ctx context.Context) {
	log.Println("Watching VoteCasted...")

	contracts.WatchEventLoop(
		ctx,
		func(opts *bind.WatchOpts, ch chan *contracts.AlgorithmReviewVoteCasted) (ethereum.Subscription, error) {
			ctr, err := contracts.NewAlgorithmReview(
				s.ctrCaller.AlgoReviewCtrtAddr(),
				s.ctrCaller.WsClient(),
			)
			if err != nil {
				return nil, err
			}
			return ctr.WatchVoteCasted(opts, ch, nil, nil)
		},
		func(evt *contracts.AlgorithmReviewVoteCasted) {
			txHash := evt.Raw.TxHash.Hex()
			log.Printf("Received event tx=%s", txHash)
			var status string
			defer func() {
				if txHash != "" && status != "" {
					if err := s.notifier.PushStatus(txHash, status); err != nil {
						log.Printf("PushStatus failed: tx=%s, err=%v", txHash, err)
					}
				}
			}()

			receipt, err := s.ctrCaller.HttpClient().TransactionReceipt(ctx, evt.Raw.TxHash)
			if err != nil {
				log.Printf("Receipt fetch failed: %v", err)
				return
			}
			// First get block info (needed for both success and failure cases)
			blockNumber := evt.Raw.BlockNumber

			// We need to fetch the block to get its timestamp
			block, err := s.ctrCaller.HttpClient().BlockByNumber(ctx, new(big.Int).SetUint64(blockNumber))
			if err != nil {
				log.Printf("Failed to fetch block details: %v", err)
				return
			}

			blockTime := time.Unix(int64(block.Time()), 0)

			// Determine transaction status based on receipt
			if receipt.Status != types.ReceiptStatusSuccessful {
				status = models.TX_STATUS_FAILED
			} else {
				status = models.TX_STATUS_CONFIRMED
			}

			dbtx := s.db.Begin()
			defer func() {
				if r := recover(); r != nil {
					dbtx.Rollback()
					panic(r)
				}
			}()

			var pending models.BlockchainTransaction
			if err := dbtx.Where("tx_hash = ?", txHash).First(&pending).Error; err != nil {
				dbtx.Rollback()
				log.Printf("pending tx not found: %v", err)
				return
			}

			vote, err := models.CreateVote(dbtx, uint(evt.AlgoId.Uint64()), evt.Voter.Hex(), evt.Approved, blockTime)
			if err != nil {
				dbtx.Rollback()
				log.Printf("Failed to create vote: %v", err)
				return
			}

			err = models.UpdateTransactionEntity(dbtx, txHash, vote.ID, status, &blockNumber, &blockTime)
			if err != nil {
				dbtx.Rollback()
				log.Printf("Failed to update transaction status to %s: %v", status, err)
				return
			}

			if err = dbtx.Commit().Error; err != nil {
				log.Printf("Failed to commit transaction status update: %v", err)
				return
			}
		},
	)
}

func (s *ChainsyncService) listenAlgoResolved(ctx context.Context) {
	log.Println("Watching ")
	contracts.WatchEventLoop(ctx,
		func(opts *bind.WatchOpts, ch chan *contracts.AlgorithmReviewAlgorithmResolved) (ethereum.Subscription, error) {
			ctr, err := contracts.NewAlgorithmReview(
				s.ctrCaller.AlgoReviewCtrtAddr(),
				s.ctrCaller.WsClient(),
			)
			if err != nil {
				return nil, err
			}
			return ctr.WatchAlgorithmResolved(opts, ch, nil)
		},
		func(evt *contracts.AlgorithmReviewAlgorithmResolved) {
			txHash := evt.Raw.TxHash.Hex()
			log.Printf("Received event tx=%s", txHash)
			var status string
			defer func() {
				if txHash != "" && status != "" {
					if err := s.notifier.PushStatus(txHash, status); err != nil {
						log.Printf("PushStatus failed: tx=%s, err=%v", txHash, err)
					}
				}
			}()

			receipt, err := s.ctrCaller.HttpClient().TransactionReceipt(ctx, evt.Raw.TxHash)
			if err != nil {
				log.Printf("Receipt fetch failed: %v", err)
				return
			}
			// First get block info (needed for both success and failure cases)
			blockNumber := evt.Raw.BlockNumber

			// We need to fetch the block to get its timestamp
			block, err := s.ctrCaller.HttpClient().BlockByNumber(ctx, new(big.Int).SetUint64(blockNumber))
			if err != nil {
				log.Printf("Failed to fetch block details: %v", err)
				return
			}

			blockTime := time.Unix(int64(block.Time()), 0)

			// Determine transaction status based on receipt
			if receipt.Status != types.ReceiptStatusSuccessful {
				status = models.TX_STATUS_FAILED
			} else {
				status = models.TX_STATUS_CONFIRMED
			}

			// Update transaction status
			err = models.UpdateTransactionStatus(s.db, txHash, status, &blockNumber, &blockTime)
			if err != nil {
				log.Printf("Failed to update transaction status to %s: %v", status, err)
				return
			}

			algoId := evt.AlgoId.Uint64()
			if evt.Approved {
				log.Printf("Algo %d approved, notifying runtime service", algoId)
				err := models.UpdateAlgoStatus(s.db, uint(algoId), models.ALGO_STATUS_APPROVED)
				if err != nil {
					log.Printf("Failed to update algo status to %s: %v", status, err)
					return
				}

				select {
				case s.algoScheduler.AlgoIdCh <- uint(algoId):
					log.Printf("Notification sent for algo %d", algoId)
				default:
					log.Printf("Algorithm event channel is full, skipping notification for algo %d", algoId)
				}
			} else {
				log.Printf("Algo %d rejected", algoId)
				err := models.UpdateAlgoStatus(s.db, uint(algoId), models.ALGO_STATUS_REJECTED)
				if err != nil {
					log.Printf("Failed to update algo status to %s: %v", status, err)
					return
				}
			}
		})
}
