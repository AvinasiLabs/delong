package control

import (
	"delong/pkg/analyzer"
	"delong/pkg/contracts"
	"delong/pkg/db"
	"delong/pkg/tee"
	"delong/pkg/ws"

	"gorm.io/gorm"
)

type StorageDeps struct {
	IpfsStore  *db.IpfsStore
	MinioStore *db.MinioStore
	MysqlDb    *gorm.DB
}

type BlockchainDeps struct {
	EthCaller *contracts.ContractCaller
	KeyVault  *tee.KeyVault
}

type NotificationDeps struct {
	Notifier *ws.Notifier
}

type AnalyzerDeps struct {
	ReportAnalyzer *analyzer.ReportAnalyzer
}

type Dependencies struct {
	Storage      *StorageDeps
	Blockchain   *BlockchainDeps
	Notification *NotificationDeps
	Analyzer     *AnalyzerDeps
}

func NewDependencies(
	ipfsStore *db.IpfsStore,
	minioStore *db.MinioStore,
	msyqlDb *gorm.DB,
	reportAnalyzer *analyzer.ReportAnalyzer,
	ethCaller *contracts.ContractCaller,
	keyVault *tee.KeyVault,
	notifier *ws.Notifier,
) *Dependencies {
	return &Dependencies{
		Storage: &StorageDeps{
			IpfsStore:  ipfsStore,
			MinioStore: minioStore,
			MysqlDb:    msyqlDb,
		},
		Blockchain: &BlockchainDeps{
			EthCaller: ethCaller,
			KeyVault:  keyVault,
		},
		Notification: &NotificationDeps{
			Notifier: notifier,
		},
		Analyzer: &AnalyzerDeps{
			ReportAnalyzer: reportAnalyzer,
		},
	}
}
