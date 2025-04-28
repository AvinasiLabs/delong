package api

import (
	"context"
	"delong/internal/control"
	"delong/pkg"
	"delong/pkg/contracts"
	"delong/pkg/db"
	"delong/pkg/tee"
	"delong/pkg/ws"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
)

type ApiService struct {
	name       string
	addr       string
	ctrAddr    map[string]common.Address
	engine     *gin.Engine
	httpserver *http.Server

	state *ApiServiceState
}

type ApiServiceState struct {
	ipfsStore *db.IpfsStore
	ethCaller *contracts.ContractCaller
	KeyVault  *tee.KeyVault
	Notifier  *ws.Notifier
}

const SERVICE_NAME = "api-service"

func NewApiService(addr string, ctrAddr map[string]common.Address) *ApiService {
	return &ApiService{
		name:    SERVICE_NAME,
		addr:    addr,
		ctrAddr: ctrAddr,
		engine:  gin.Default(),
		state:   &ApiServiceState{},
	}
}

func (s *ApiService) Name() string {
	return s.name
}

func (s *ApiService) Init(ctx context.Context, gs *control.ServiceState) error {
	s.state.ipfsStore = gs.IpfsStore
	s.state.ethCaller = gs.EthCaller
	s.state.KeyVault = gs.KeyVault
	s.state.Notifier = gs.Notifier
	s.registerRoutes()
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

func (s *ApiService) registerRoutes() {
	s.engine.GET("/ws", ws.NewHandler(s.state.Notifier.Hub()))

	apiGroup := s.engine.Group("/api")

	apiGroup.POST("/report/upload", s.UploadReport)
	apiGroup.GET("/report/:id", s.GetReports)
}

type UploadReportReq struct {
	UserWallet string `form:"userWallet" binding:"required,ethwallet"`
	Dataset    string `form:"dataset" binding:"required"`
}

func (s *ApiService) UploadReport(c *gin.Context) {
	req := UploadReportReq{}
	err := c.ShouldBind(&req)
	if err != nil {
		log.Printf("Failed to bind request: %v", err)
		c.JSON(400, gin.H{"error": "Invalid request"})
		return
	}

	userWallet := common.HexToAddress(req.UserWallet)
	dataset := req.Dataset

	fh, err := c.FormFile("file")
	if err != nil {
		log.Printf("Failed to get uploaded file: %v", err)
		c.JSON(400, gin.H{"error": "Uploading file failed"})
		return
	}

	f, err := fh.Open()
	if err != nil {
		log.Printf("Failed to open uploaded file: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Opening file failed"})
	}
	defer f.Close()

	fd, err := io.ReadAll(f)
	if err != nil {
		log.Printf("Failed to read uploaded file: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Reading file failed"})
		return
	}

	aesKc := tee.NewKeyContext(tee.KindEncryptionKey, "upload-report-key", "encrypt user test report")
	aesKey, err := s.state.KeyVault.DeriveSymmetricKey(c, aesKc, 32)

	combined, err := pkg.EncryptGCM(fd, aesKey)
	if err != nil {
		log.Printf("Failed to encrypt file: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Encrypting file failed"})
		return
	}

	cid, err := s.state.ipfsStore.Add(c, combined)
	if err != nil {
		log.Printf("Failed to add file to IPFS: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Adding file to IPFS failed"})
		return
	}

	ethAccKc := tee.NewKeyContext(tee.KindEthAccount, "tee-eth-account", "call as owner")
	ethAcc, err := s.state.KeyVault.DeriveEthereumAccount(c, ethAccKc)
	if err != nil {
		log.Printf("Failed to derive Ethereum account: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Deriving Ethereum account failed"})
		return
	}

	tx, err := s.state.ethCaller.RegisterData(*ethAcc, userWallet, cid, dataset)
	if err != nil {
		log.Printf("Failed to register data: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Registering data failed"})
		return
	}

	log.Printf("Transaction hash: %s", tx.Hash().Hex())

	c.JSON(http.StatusOK, gin.H{"msg": "ok", "data": tx.Hash().Hex()})
}

func (s *ApiService) GetReports(c *gin.Context) {}
