package api

import (
	"context"
	"delong/pkg/analysis"
	"delong/pkg/contracts"
	"delong/pkg/db"
	"delong/pkg/tee"
	"delong/pkg/ws"
	"log"
	"net/http"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ApiService struct {
	name       string
	addr       string
	ctrAddr    map[string]common.Address
	engine     *gin.Engine
	httpserver *http.Server

	// ipfsStore      *db.IpfsStore
	// minioStore     *db.MinioStore
	// mysqlDb        *gorm.DB
	// ctrCaller      *contracts.ContractCaller
	// keyVault       *tee.KeyVault
	// notifier       *ws.Notifier
	// reportAnalyzer *analysis.ReportAnalyzer

	ApiServiceOptions
}

type ApiServiceOptions struct {
	Addr           string
	IpfsStore      *db.IpfsStore
	MinioStore     *db.MinioStore
	MysqlDb        *gorm.DB
	CtrCaller      *contracts.ContractCaller
	KeyVault       *tee.KeyVault
	Notifier       *ws.Notifier
	ReportAnalyzer *analysis.ReportAnalyzer
}

func NewService(opts ApiServiceOptions) *ApiService {
	return &ApiService{
		name:              "api-service",
		addr:              opts.Addr,
		engine:            gin.Default(),
		ctrAddr:           map[string]common.Address{},
		httpserver:        &http.Server{},
		ApiServiceOptions: opts,
	}
}

func (s *ApiService) Name() string {
	return s.name
}

func (s *ApiService) Init(ctx context.Context) error {
	// register routes
	s.engine.GET("/ws", ws.NewHandler(s.Notifier.Hub()))
	apiGroup := s.engine.Group("/api")
	apiGroup.POST("/report/upload", s.UploadReport)
	apiGroup.GET("/report/:id", s.GetReports)
	apiGroup.POST("/algo/submit", s.SubmitAlgo)
	apiGroup.POST("/algo/vote", s.Vote)
	return nil
}

func (s *ApiService) Start(ctx context.Context) error {
	s.httpserver = &http.Server{
		Addr:    s.addr,
		Handler: s.engine,
	}

	go func() {
		err := s.httpserver.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.Printf("Failed to listen: %v", err)
		}
	}()

	log.Println("Api service started")
	<-ctx.Done()
	log.Println("API service context cancelled, will shut down")
	return nil
}

func (s *ApiService) Stop(ctx context.Context) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	err := s.httpserver.Shutdown(ctx)
	if err != nil {
		log.Printf("Failed to shutdown gracefully: %v", err)
		return err
	}

	log.Println("Http server shutdown cleanly")
	return nil
}
