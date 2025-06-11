//go:build integration
// +build integration

package runtime

// TODO: No data record for dynamic dataset

// func TestDatasetLoader(t *testing.T) {
// 	config, err := internal.LoadConfigFromEnv()
// 	if err != nil {
// 		t.Fatalf("Failed to load config: %v", err)
// 	}

// 	dbConn, err := db.NewMysqlDb(config.MysqlDsn)
// 	if err != nil {
// 		t.Fatalf("Failed to connect to db: %v", err)
// 	}

// 	ipfsStore, err := db.NewIpfsStore(config.IpfsApiAddr)
// 	if err != nil {
// 		t.Fatalf("Failed to connect to ipfs: %v", err)
// 	}

// 	tempDir := t.TempDir()
// 	loader := NewDatasetLoader(tempDir, dbConn, ipfsStore, nil) // nil for build, TODO
// 	t.Log("Exporting datasets...")
// 	if err := loader.Export(); err != nil {
// 		t.Fatalf("Failed to export: %v", err)
// 	}

// 	linkPath := filepath.Join(tempDir, "current")
// 	resolvedPath, err := os.Readlink(linkPath)
// 	if err != nil {
// 		t.Fatalf("Failed to read current symlink: %v", err)
// 	}
// 	t.Logf("Link: %v", resolvedPath)

// 	if !filepath.IsAbs(resolvedPath) {
// 		resolvedPath = filepath.Join(tempDir, resolvedPath)
// 	}

// 	files, err := os.ReadDir(resolvedPath)
// 	if err != nil || len(files) == 0 {
// 		t.Fatalf("Exported dataset directory is empty: %v", err)
// 	}

// 	t.Log("Acquiring dataset...")
// 	fileName := files[0].Name()
// 	datasetName := strings.TrimSuffix(fileName, ".csv")
// 	fullPath, version, err := loader.AcquireCurrent(datasetName)
// 	if err != nil {
// 		t.Fatalf("Failed to acquire dataset: %v", err)
// 	}

// 	filePath := filepath.Join(fullPath, datasetName+".csv")
// 	if _, err := os.Stat(filePath); err != nil {
// 		t.Errorf("Expected dataset file missing: %s", filePath)
// 	}
// 	t.Logf("File: %v", filePath)

// 	// Read file header
// 	f, err := os.Open(filePath)
// 	if err != nil {
// 		t.Fatalf("Failed to open exported file: %v", err)
// 	}
// 	defer f.Close()

// 	r := csv.NewReader(f)
// 	for range 5 {
// 		record, err := r.Read()
// 		if err != nil {
// 			t.Fatalf("Failed to read record: %v", err)
// 		}
// 		t.Logf("Record: %v", record)
// 	}

// 	loader.Release(datasetName, version)
// 	t.Log("Released dataset, waiting before cleanup...")
// 	time.Sleep(1 * time.Second)

// 	t.Log("Cleaning up...")
// 	if err := loader.Cleanup(); err != nil {
// 		t.Errorf("Cleanup failed: %v", err)
// 	}

// 	if _, err := os.Stat(fullPath); err == nil {
// 		t.Errorf("Expected version directory %s to be removed", version)
// 	}
// }
