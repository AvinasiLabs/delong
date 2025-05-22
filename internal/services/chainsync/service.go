package chainsync

import (
	"context"
	"delong/internal/models"
	"delong/pkg/bizcode"
	"delong/pkg/contracts"
	"delong/pkg/schedule"
	"delong/pkg/ws"
	"fmt"
	"log"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
	"gorm.io/gorm"
)

type ChainsyncService struct {
	name string
	ChainsyncServiceOptions
}

type ChainsyncServiceOptions struct {
	CtrCaller     *contracts.ContractCaller
	Notifier      *ws.Notifier
	Db            *gorm.DB
	AlgoScheduler *schedule.AlgoScheduler
}

func NewService(opts ChainsyncServiceOptions) *ChainsyncService {
	return &ChainsyncService{
		name:                    "chainsync-service",
		ChainsyncServiceOptions: opts,
	}
}

func (s *ChainsyncService) Name() string {
	return s.name
}

func (s *ChainsyncService) Init(ctx context.Context) error {
	err := s.recoverResolveTasks()
	if err != nil {
		return err
	}

	return nil
}

func (s *ChainsyncService) Start(ctx context.Context) error {
	go s.listenDataRegistered(ctx)
	go s.listenAlgoSubmitted(ctx)
	go s.listenVoteCasted(ctx)
	go s.listenCommitteeMemberUpdated(ctx)
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
				s.CtrCaller.DataContributionCtrtAddr(),
				s.CtrCaller.WsClient(),
			)
			if err != nil {
				return nil, err
			}
			return ctr.WatchDataRegistered(opts, ch, nil, nil)
		},
		func(evt *contracts.DataContributionDataRegistered) {
			txHash := evt.Raw.TxHash.Hex()
			log.Printf("Received event tx=%s", txHash)

			receipt, err := s.CtrCaller.HttpClient().TransactionReceipt(ctx, evt.Raw.TxHash)
			if err != nil {
				log.Printf("Receipt fetch failed: %v", err)
				s.Notifier.PushError(txHash, bizcode.RECEIPT_QUERY_FAIL)
				return
			}
			// First get block info (needed for both success and failure cases)
			blockNumber := evt.Raw.BlockNumber

			// We need to fetch the block to get its timestamp
			block, err := s.CtrCaller.HttpClient().BlockByNumber(ctx, new(big.Int).SetUint64(blockNumber))
			if err != nil {
				log.Printf("Failed to fetch block details: %v", err)
				s.Notifier.PushError(txHash, bizcode.BLOCK_QUERY_FAIL)
				return
			}

			blockTime := time.Unix(int64(block.Time()), 0)

			// Determine transaction status based on receipt
			var status string
			if receipt.Status != types.ReceiptStatusSuccessful {
				status = models.TX_STATUS_FAILED
			} else {
				status = models.TX_STATUS_CONFIRMED
			}

			// Update transaction status
			transaction, err := models.UpdateTransactionStatus(s.Db, txHash, status, &blockNumber, &blockTime)
			if err != nil {
				log.Printf("Failed to update transaction status to %s: %v", status, err)
				s.Notifier.PushError(txHash, bizcode.MYSQL_WRITE_FAIL)
				return
			}
			s.Notifier.PushTxResult(txHash, transaction)
		},
	)
}

func (s *ChainsyncService) listenAlgoSubmitted(ctx context.Context) {
	log.Println("Watching AlgorithmSubmitted...")

	contracts.WatchEventLoop(
		ctx,
		func(opts *bind.WatchOpts, ch chan *contracts.AlgorithmReviewAlgorithmSubmitted) (ethereum.Subscription, error) {
			ctr, err := contracts.NewAlgorithmReview(
				s.CtrCaller.AlgoReviewCtrtAddr(),
				s.CtrCaller.WsClient(),
			)
			if err != nil {
				return nil, err
			}
			return ctr.WatchAlgorithmSubmitted(opts, ch, nil, nil)
		},
		func(evt *contracts.AlgorithmReviewAlgorithmSubmitted) {
			txHash := evt.Raw.TxHash.Hex()
			log.Printf("Received event tx=%s", txHash)

			receipt, err := s.CtrCaller.HttpClient().TransactionReceipt(ctx, evt.Raw.TxHash)
			if err != nil {
				log.Printf("Receipt fetch failed: %v", err)
				s.Notifier.PushError(txHash, bizcode.RECEIPT_QUERY_FAIL)
				return
			}
			// First get block info (needed for both success and failure cases)
			blockNumber := evt.Raw.BlockNumber

			// We need to fetch the block to get its timestamp
			block, err := s.CtrCaller.HttpClient().BlockByNumber(ctx, new(big.Int).SetUint64(blockNumber))
			if err != nil {
				log.Printf("Failed to fetch block details: %v", err)
				s.Notifier.PushError(txHash, bizcode.BLOCK_QUERY_FAIL)
				return
			}

			blockTime := time.Unix(int64(block.Time()), 0)

			// Determine transaction status based on receipt
			var status string
			if receipt.Status != types.ReceiptStatusSuccessful {
				status = models.TX_STATUS_FAILED
			} else {
				status = models.TX_STATUS_CONFIRMED
			}

			dbtx := s.Db.Begin()
			defer func() {
				if r := recover(); r != nil {
					dbtx.Rollback()
					panic(r)
				}
			}()

			// Update transaction status
			transaction, err := models.UpdateTransactionStatus(dbtx, txHash, status, &blockNumber, &blockTime)
			if err != nil {
				dbtx.Rollback()
				log.Printf("Failed to update transaction status to %s: %v", status, err)
				s.Notifier.PushError(txHash, bizcode.MYSQL_WRITE_FAIL)
				return
			}

			// Launch a timer to automatically resolve algorithm voting outcomes
			if status == models.TX_STATUS_CONFIRMED {
				startTime := time.Unix(evt.StartTime.Int64(), 0)
				endTime := time.Unix(evt.EndTime.Int64(), 0)
				err := models.UpdateAlgoVoteDuration(dbtx, uint(evt.AlgoId.Uint64()), &startTime, &endTime)
				if err != nil {
					dbtx.Rollback()
					log.Printf("Failed to update algo vote duration: %v", err)
					s.Notifier.PushError(txHash, bizcode.MYSQL_WRITE_FAIL)
					return
				}
				s.AlgoScheduler.ScheduleResolve(uint(evt.AlgoId.Uint64()), endTime)
			}

			if err = dbtx.Commit().Error; err != nil {
				log.Printf("Failed to commit transaction status update: %v", err)
				s.Notifier.PushError(txHash, bizcode.MYSQL_WRITE_FAIL)
				return
			}

			s.Notifier.PushTxResult(txHash, transaction)
		},
	)
}

func (s *ChainsyncService) listenVoteCasted(ctx context.Context) {
	log.Println("Watching VoteCasted...")

	contracts.WatchEventLoop(
		ctx,
		func(opts *bind.WatchOpts, ch chan *contracts.AlgorithmReviewVoteCasted) (ethereum.Subscription, error) {
			ctr, err := contracts.NewAlgorithmReview(
				s.CtrCaller.AlgoReviewCtrtAddr(),
				s.CtrCaller.WsClient(),
			)
			if err != nil {
				return nil, err
			}
			return ctr.WatchVoteCasted(opts, ch, nil, nil)
		},
		func(evt *contracts.AlgorithmReviewVoteCasted) {
			txHash := evt.Raw.TxHash.Hex()
			log.Printf("Received event tx=%s", txHash)
			receipt, err := s.CtrCaller.HttpClient().TransactionReceipt(ctx, evt.Raw.TxHash)
			if err != nil {
				log.Printf("Receipt fetch failed: %v", err)
				s.Notifier.PushError(txHash, bizcode.RECEIPT_QUERY_FAIL)
				return
			}
			// First get block info (needed for both success and failure cases)
			blockNumber := evt.Raw.BlockNumber

			// We need to fetch the block to get its timestamp
			block, err := s.CtrCaller.HttpClient().BlockByNumber(ctx, new(big.Int).SetUint64(blockNumber))
			if err != nil {
				log.Printf("Failed to fetch block details: %v", err)
				s.Notifier.PushError(txHash, bizcode.BLOCK_QUERY_FAIL)
				return
			}

			blockTime := time.Unix(int64(block.Time()), 0)

			// Determine transaction status based on receipt
			var status string
			if receipt.Status != types.ReceiptStatusSuccessful {
				status = models.TX_STATUS_FAILED
			} else {
				status = models.TX_STATUS_CONFIRMED
			}

			dbtx := s.Db.Begin()
			defer func() {
				if r := recover(); r != nil {
					dbtx.Rollback()
					panic(r)
				}
			}()

			// var pending models.BlockchainTransaction
			// if err := dbtx.Where("tx_hash = ?", txHash).First(&pending).Error; err != nil {
			// 	dbtx.Rollback()
			// 	log.Printf("pending tx not found: %v", err)
			// 	s.Notifier.PushError(txHash, bizcode.MYSQL_READ_FAIL)
			// 	return
			// }

			vote, err := models.CreateVote(dbtx, uint(evt.AlgoId.Uint64()), evt.Voter.Hex(), evt.Approved, blockTime)
			if err != nil {
				dbtx.Rollback()
				log.Printf("Failed to create vote: %v", err)
				s.Notifier.PushError(txHash, bizcode.MYSQL_WRITE_FAIL)
				return
			}

			transaction, err := models.CreateTransactionWithStatus(dbtx, txHash, vote.ID, models.ENTITY_TYPE_VOTE, status)
			if err != nil {
				dbtx.Rollback()
				log.Printf("Failed to create transaction: %v", err)
				s.Notifier.PushError(txHash, bizcode.MYSQL_WRITE_FAIL)
				return
			}

			// transaction, err := models.UpdateTransactionEntity(dbtx, txHash, vote.ID, status, &blockNumber, &blockTime)
			// if err != nil {
			// 	dbtx.Rollback()
			// 	log.Printf("Failed to update transaction status to %s: %v", status, err)
			// 	s.Notifier.PushError(txHash, bizcode.MYSQL_WRITE_FAIL)
			// 	return
			// }

			if err = dbtx.Commit().Error; err != nil {
				log.Printf("Failed to commit transaction status update: %v", err)
				s.Notifier.PushError(txHash, bizcode.MYSQL_WRITE_FAIL)
				return
			}

			s.Notifier.PushTxResult(txHash, transaction)
		},
	)
}

func (s *ChainsyncService) listenCommitteeMemberUpdated(ctx context.Context) {
	log.Println("Watching CommitteeMemberUpdated...")

	contracts.WatchEventLoop(
		ctx,
		func(opts *bind.WatchOpts, ch chan *contracts.AlgorithmReviewCommitteeMemberUpdated) (ethereum.Subscription, error) {
			ctr, err := contracts.NewAlgorithmReview(
				s.CtrCaller.AlgoReviewCtrtAddr(),
				s.CtrCaller.WsClient(),
			)
			if err != nil {
				return nil, err
			}
			return ctr.WatchCommitteeMemberUpdated(opts, ch, nil)
		},
		func(evt *contracts.AlgorithmReviewCommitteeMemberUpdated) {
			txHash := evt.Raw.TxHash.Hex()
			log.Printf("Received event tx=%s", txHash)

			receipt, err := s.CtrCaller.HttpClient().TransactionReceipt(ctx, evt.Raw.TxHash)
			if err != nil {
				log.Printf("Receipt fetch failed: %v", err)
				s.Notifier.PushError(txHash, bizcode.RECEIPT_QUERY_FAIL)
				return
			}
			// First get block info (needed for both success and failure cases)
			blockNumber := evt.Raw.BlockNumber

			// We need to fetch the block to get its timestamp
			block, err := s.CtrCaller.HttpClient().BlockByNumber(ctx, new(big.Int).SetUint64(blockNumber))
			if err != nil {
				log.Printf("Failed to fetch block details: %v", err)
				s.Notifier.PushError(txHash, bizcode.BLOCK_QUERY_FAIL)
				return
			}

			blockTime := time.Unix(int64(block.Time()), 0)

			// Determine transaction status based on receipt
			var status string
			if receipt.Status != types.ReceiptStatusSuccessful {
				status = models.TX_STATUS_FAILED
			} else {
				status = models.TX_STATUS_CONFIRMED
			}

			// Update transaction status
			transaction, err := models.UpdateTransactionStatus(s.Db, txHash, status, &blockNumber, &blockTime)
			if err != nil {
				log.Printf("Failed to update transaction status to %s: %v", status, err)
				s.Notifier.PushError(txHash, bizcode.MYSQL_READ_FAIL)
				return
			}

			s.Notifier.PushTxResult(txHash, transaction)
		},
	)
}

func (s *ChainsyncService) listenAlgoResolved(ctx context.Context) {
	log.Println("Watching ")
	contracts.WatchEventLoop(ctx,
		func(opts *bind.WatchOpts, ch chan *contracts.AlgorithmReviewAlgorithmResolved) (ethereum.Subscription, error) {
			ctr, err := contracts.NewAlgorithmReview(
				s.CtrCaller.AlgoReviewCtrtAddr(),
				s.CtrCaller.WsClient(),
			)
			if err != nil {
				return nil, err
			}
			return ctr.WatchAlgorithmResolved(opts, ch, nil)
		},
		func(evt *contracts.AlgorithmReviewAlgorithmResolved) {
			txHash := evt.Raw.TxHash.Hex()
			log.Printf("Received event tx=%s", txHash)

			receipt, err := s.CtrCaller.HttpClient().TransactionReceipt(ctx, evt.Raw.TxHash)
			if err != nil {
				log.Printf("Receipt fetch failed: %v", err)
				s.Notifier.PushError(txHash, bizcode.RECEIPT_QUERY_FAIL)
				return
			}
			// First get block info (needed for both success and failure cases)
			blockNumber := evt.Raw.BlockNumber

			// We need to fetch the block to get its timestamp
			block, err := s.CtrCaller.HttpClient().BlockByNumber(ctx, new(big.Int).SetUint64(blockNumber))
			if err != nil {
				log.Printf("Failed to fetch block details: %v", err)
				s.Notifier.PushError(txHash, bizcode.BLOCK_QUERY_FAIL)
				return
			}

			blockTime := time.Unix(int64(block.Time()), 0)

			// Determine transaction status based on receipt
			var status string
			if receipt.Status != types.ReceiptStatusSuccessful {
				status = models.TX_STATUS_FAILED
			} else {
				status = models.TX_STATUS_CONFIRMED
			}

			// Update transaction status
			transaction, err := models.UpdateTransactionStatus(s.Db, txHash, status, &blockNumber, &blockTime)
			if err != nil {
				log.Printf("Failed to update transaction status to %s: %v", status, err)
				s.Notifier.PushError(txHash, bizcode.MYSQL_READ_FAIL)
				return
			}

			s.Notifier.PushTxResult(txHash, transaction)

			algoId := evt.AlgoId.Uint64()
			if evt.Approved {
				log.Printf("Algo %d approved, notifying runtime service", algoId)
				err := models.UpdateAlgoStatus(s.Db, uint(algoId), models.ALGO_STATUS_APPROVED)
				if err != nil {
					log.Printf("Failed to update algo status to %s: %v", status, err)
					return
				}

				s.AlgoScheduler.ScheduleRun(uint(algoId))
			} else {
				log.Printf("Algo %d rejected", algoId)
				err := models.UpdateAlgoStatus(s.Db, uint(algoId), models.ALGO_STATUS_REJECTED)
				if err != nil {
					log.Printf("Failed to update algo status to %s: %v", status, err)
					return
				}
			}
		})
}

func (s *ChainsyncService) recoverResolveTasks() error {
	log.Println("Recovering unresolved algos...")

	algos, err := models.GetPendingConfirmedAlgos(s.Db)
	if err != nil {
		return fmt.Errorf("failed to query unresolved algos: %w", err)
	}

	now := time.Now()
	for _, algo := range algos {
		if algo.EndTime == nil {
			log.Printf("Skip unresolved algo %d: missing end time", algo.ID)
			continue
		}

		if algo.EndTime.After(now) {
			log.Printf("Recovered resolve task for algo %d (endTime = %s)", algo.ID, algo.EndTime)
			s.AlgoScheduler.ScheduleResolve(algo.ID, *algo.EndTime)
		} else {
			// TODO resolve immediately
			log.Printf("Skipped expired resolve for algo %d (endTime = %s)", algo.ID, algo.EndTime)
		}
	}
	return nil
}
