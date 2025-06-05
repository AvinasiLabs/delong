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
	go s.listenDataUsed(ctx)

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
		func(opts *bind.WatchOpts, ch chan *contracts.AlgorithmReviewExecutionSubmitted) (ethereum.Subscription, error) {
			ctr, err := contracts.NewAlgorithmReview(
				s.CtrCaller.AlgoReviewCtrtAddr(),
				s.CtrCaller.WsClient(),
			)
			if err != nil {
				return nil, err
			}
			return ctr.WatchExecutionSubmitted(opts, ch, nil)
		},
		func(evt *contracts.AlgorithmReviewExecutionSubmitted) {
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

			// Launch a timer to automatically resolve algorithm voting outcomes
			if status == models.TX_STATUS_CONFIRMED {
				startTime := time.Unix(evt.StartTime.Int64(), 0)
				endTime := time.Unix(evt.EndTime.Int64(), 0)
				exeId := uint(evt.ExecutionId.Uint64())
				cid := evt.Cid
				err = models.UpdateVoteDuration(s.Db, exeId, &startTime, &endTime)
				if err != nil {
					log.Printf("Failed to update algo vote duration: %v", err)
					s.Notifier.PushError(txHash, bizcode.MYSQL_WRITE_FAIL)
					return
				}
				s.AlgoScheduler.ScheduleResolve(exeId, cid, endTime)
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
			return ctr.WatchVoteCasted(opts, ch, nil)
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
			voteTime := time.Unix(evt.VoteTime.Int64(), 0)
			vote, err := models.CreateVote(dbtx, evt.Cid, evt.Member.Hex(), evt.Approved, voteTime)
			if err != nil {
				dbtx.Rollback()
				log.Printf("Failed to create vote: %v", err)
				s.Notifier.PushError(txHash, bizcode.MYSQL_WRITE_FAIL)
				return
			}

			transaction, err := models.CreateTransactionWithStatus(dbtx, txHash, vote.ID, models.ENTITY_TYPE_VOTE, status, &blockNumber, &blockTime)
			if err != nil {
				dbtx.Rollback()
				log.Printf("Failed to create transaction: %v", err)
				s.Notifier.PushError(txHash, bizcode.MYSQL_WRITE_FAIL)
				return
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
	log.Println("Watching AlgorithmResolved...")
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

			// algoCid := evt.Cid
			exeId := uint(evt.ExecutionId.Uint64())
			if evt.Approved {
				log.Printf("Execution task %v approved, notifying runtime service", exeId)
				err := models.UpdateReviewStatus(s.Db, exeId, models.ALGO_STATUS_APPROVED)
				if err != nil {
					log.Printf("Failed to update algo status to %s: %v", status, err)
					return
				}
				s.AlgoScheduler.ScheduleRun(exeId)
			} else {
				log.Printf("Execution task %v rejected", exeId)
				err := models.UpdateReviewStatus(s.Db, exeId, models.ALGO_STATUS_REJECTED)
				if err != nil {
					log.Printf("Failed to update algo status to %s: %v", status, err)
					return
				}
			}
		})
}

func (s *ChainsyncService) listenDataUsed(ctx context.Context) {
	log.Println("Watching DataUsed...")
	contracts.WatchEventLoop(ctx,
		func(opts *bind.WatchOpts, ch chan *contracts.DataContributionDataUsed) (ethereum.Subscription, error) {
			ctr, err := contracts.NewDataContribution(
				s.CtrCaller.DataContributionCtrtAddr(),
				s.CtrCaller.WsClient(),
			)
			if err != nil {
				return nil, err
			}
			return ctr.WatchDataUsed(opts, ch, nil, nil)
		},
		func(evt *contracts.DataContributionDataUsed) {
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
		})
}

func (s *ChainsyncService) recoverResolveTasks() error {
	log.Println("Recovering unresolved algos...")

	algoExes, err := models.GetReviewingAlgoExes(s.Db)
	if err != nil {
		return fmt.Errorf("failed to query reviewing algo exes: %w", err)
	}
	now := time.Now()

	for _, exe := range algoExes {
		// algo exes must have ent time
		if exe.VoteEndTime.After(now) {
			algo, err := models.GetAlgoByID(s.Db, exe.AlgoID)
			if err != nil {
				log.Printf("Failed to find algo by id: %v", err)
				continue
			}

			s.AlgoScheduler.ScheduleResolve(exe.ID, algo.Cid, *exe.VoteEndTime)
		}
	}

	return nil
}
