//go:build integration
// +build integration

package runtime

import (
	"context"
	"delong/internal"
	"delong/internal/models"
	"delong/pkg/contracts"
	"delong/pkg/db"
	"delong/pkg/schedule"
	"net/http"
	"testing"
	"time"
)

func TestRuntimeService(t *testing.T) {
	config, err := internal.LoadConfigFromEnv()
	if err != nil {
		t.Fatal(err)
	}

	mysqlConn, err := db.NewMysqlDb(config.MysqlDsn)
	if err != nil {
		t.Fatal(err)
	}

	ipfsStore, err := db.NewIpfsStore(config.IpfsApiAddr)
	if err != nil {
		t.Fatal(err)
	}

	algoScheduler, err := schedule.NewAlgoScheduler()
	if err != nil {
		t.Fatal(err)
	}
	loader := NewDatasetLoader(t.TempDir(), mysqlConn, ipfsStore, nil) // nil for build, TODO
	rt := NewService(RuntimeServiceOptions{
		Db:            mysqlConn,
		Loader:        loader,
		IpfsStore:     ipfsStore,
		CtrCaller:     &contracts.ContractCaller{},
		AlgoScheduler: algoScheduler,
	})

	ctx, cancel := context.WithTimeout(t.Context(), 5*time.Minute)
	defer cancel()

	if err := rt.Init(ctx); err != nil {
		t.Fatalf("init failed: %v", err)
	}

	go rt.Start(ctx)

	// Build a test algo
	testAlgoLink := "https://codeload.github.com/lilhammer111/algo-demo/tar.gz/c73e8d62a0ae5d68040cabb461c7b51b7630020c"
	resp, err := http.Get(testAlgoLink)
	if err != nil {
		t.Fatalf("Failed to download algorithm: %v", err)
	}
	defer resp.Body.Close()

	cid, err := rt.IpfsStore.UploadStream(ctx, resp.Body)
	if err != nil {
		t.Fatalf("Failed to upload algorithm to IPFS: %v", err)
	}

	testRepo := "lilhammer111/algo-demo"

	algo, err := models.CreateAlgo(
		mysqlConn,
		testRepo,
		testAlgoLink,
		cid,
	)

	algoExe, err := models.CreateAlgoExecution(mysqlConn, algo.ID, "TEST_DATASET", "TEST_SCIENTIST")
	rt.AlgoScheduler.ScheduleRun(algoExe.ID)

	time.Sleep(3 * time.Second)
	for range 300 {
		exe, err := models.GetAlgoExeById(mysqlConn, algoExe.ID)
		if err == nil && exe.Status == models.EXE_STATUS_COMPLETED {
			t.Logf("Execution completed! Result:\n%s", exe.Result)
			return
		}
		time.Sleep(1 * time.Second)
	}

	t.Fatal("Execution did not complete in time")
}
