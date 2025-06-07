/*
Dataset Caching System Design Documentation

This system manages two types of datasets with different caching strategies:

1. DYNAMIC DATASETS (from database)
2. STATIC DATASETS (from IPFS, encrypted)

═══════════════════════════════════════════════════════════════════════════════
DIRECTORY STRUCTURE:
═══════════════════════════════════════════════════════════════════════════════

/data/delong_dataset/                    <- storageRoot
├── current -> 20250106_143022/          <- symlink to latest dynamic version
├── 20250106_143022/                     <- dynamic dataset version (timestamp)
│   ├── dataset1.csv
│   ├── dataset2.csv
│   └── dataset3.csv
├── 20250106_140000/                     <- older dynamic version
│   └── ...
├── static_123_a1b2c3d4/                 <- static dataset cache
│   └── data/
│       └── blood_test.csv               <- decrypted static file
└── static_456_e5f6g7h8/                 <- another static dataset
    └── data/
        └── glucose_data.csv

═══════════════════════════════════════════════════════════════════════════════
VERSION MANAGEMENT:
═══════════════════════════════════════════════════════════════════════════════

DYNAMIC DATASETS:
┌─────────────────────────────────────────────────────────────────────────────┐
│ Timeline: Every 15 minutes, new version is created                         │
│                                                                             │
│ 14:00:00 ────────────► 14:15:00 ────────────► 14:30:00                    │
│    │                     │                      │                         │
│    ▼                     ▼                      ▼                         │
│ 20250106_140000/    20250106_141500/    20250106_143000/                  │
│ (old version)       (old version)       (current)                         │
│                                                                             │
│ current -> 20250106_143000/  (symlink always points to latest)            │
└─────────────────────────────────────────────────────────────────────────────┘

STATIC DATASETS:
┌─────────────────────────────────────────────────────────────────────────────┐
│ One-time download and cache, persistent until cleanup                      │
│                                                                             │
│ Dataset: __static__blood_test                                              │
│ Version: static_123_a1b2c3d4                                               │
│          │     │   │                                                       │
│          │     │   └── File hash (first 8 chars)                         │
│          │     └── Database ID                                             │
│          └── Static prefix                                                 │
│                                                                             │
│ Path: /data/delong_dataset/static_123_a1b2c3d4/data/blood_test.csv        │
└─────────────────────────────────────────────────────────────────────────────┘

═══════════════════════════════════════════════════════════════════════════════
REFERENCE COUNTING SYSTEM:
═══════════════════════════════════════════════════════════════════════════════

refCounts map[string]*MountVersionRef
│
├── "20250106_143000" -> {RefCount: 3}    <- 3 algorithms using this version
├── "20250106_140000" -> {RefCount: 0}    <- no algorithms, can be cleaned
├── "static_123_a1b2c3d4" -> {RefCount: 2} <- 2 algorithms using this static
└── "static_456_e5f6g7h8" -> {RefCount: 1} <- 1 algorithm using this static

ACQUIRE FLOW:
┌─────────────────┐
│ Algorithm wants │
│ dataset         │
└────────┬────────┘
         ▼
┌─────────────────┐     YES    ┌─────────────────┐
│ Static dataset? │ ──────────► │ Download from   │
│ (has prefix?)   │             │ IPFS & decrypt  │
└────────┬────────┘             └─────────────────┘
         │ NO
         ▼
┌─────────────────┐
│ Use current     │
│ dynamic version │
└─────────────────┘
         │
         ▼
┌─────────────────┐
│ Increment       │
│ RefCount++      │
└─────────────────┘

RELEASE FLOW:
┌─────────────────┐
│ Algorithm done  │
│ with dataset    │
└────────┬────────┘
         ▼
┌─────────────────┐
│ Decrement       │
│ RefCount--      │
└─────────────────┘

═══════════════════════════════════════════════════════════════════════════════
CLEANUP STRATEGY:
═══════════════════════════════════════════════════════════════════════════════

EVERY 5 MINUTES:
┌─────────────────────────────────────────────────────────────────────────────┐
│ For each version in refCounts:                                             │
│                                                                             │
│ if RefCount == 0 && version != currentVersion:                            │
│     ┌─────────────────┐                                                   │
│     │ DELETE          │                                                   │
│     │ - Directory     │                                                   │
│     │ - All files     │                                                   │
│     │ - RefCount entry│                                                   │
│     └─────────────────┘                                                   │
│                                                                             │
│ KEEP currentVersion (even if RefCount == 0)                               │
│ KEEP any version with RefCount > 0                                        │
└─────────────────────────────────────────────────────────────────────────────┘

═══════════════════════════════════════════════════════════════════════════════
STATIC DATASET PROCESSING FLOW:
═══════════════════════════════════════════════════════════════════════════════

Algorithm requests: "__static__blood_test"
         │
         ▼
┌─────────────────┐
│ Check database  │ ── SELECT * FROM static_datasets WHERE name = ?
│ for metadata    │
└────────┬────────┘
         ▼
┌─────────────────┐
│ Generate cache  │ ── static_123_a1b2c3d4
│ version ID      │    (ID_123 + hash_a1b2c3d4)
└────────┬────────┘
         ▼
┌─────────────────┐     EXISTS   ┌─────────────────┐
│ Check if file   │ ──────────► │ RefCount++      │
│ already cached  │              │ Return path     │
└────────┬────────┘              └─────────────────┘
         │ NOT EXISTS
         ▼
┌─────────────────┐
│ Derive key from │ ── TEE KeyVault
│ Author + purpose│
└────────┬────────┘
         ▼
┌─────────────────┐
│ Download from   │ ── IPFS CID -> encrypted data
│ IPFS            │
└────────┬────────┘
         ▼
┌─────────────────┐
│ Decrypt data    │ ── AES-GCM decryption
└────────┬────────┘
         ▼
┌─────────────────┐
│ Write to cache  │ ── /data/.../static_123_.../data/file.csv
│ directory       │
└────────┬────────┘
         ▼
┌─────────────────┐
│ RefCount = 1    │
│ Return path     │
└─────────────────┘

═══════════════════════════════════════════════════════════════════════════════
THREAD SAFETY:
═══════════════════════════════════════════════════════════════════════════════

All refCount operations are protected by sync.RWMutex:

AcquireCurrent():  dl.rm.Lock() -> modify refCounts -> dl.rm.Unlock()
Release():         dl.rm.Lock() -> modify refCounts -> dl.rm.Unlock()
Cleanup():         dl.rm.Lock() -> read/modify refCounts -> dl.rm.Unlock()

═══════════════════════════════════════════════════════════════════════════════
KEY ADVANTAGES:
═══════════════════════════════════════════════════════════════════════════════

1. MEMORY EFFICIENT: Only decrypt static datasets when needed
2. DISK EFFICIENT: Cache reuse across multiple algorithm runs
3. SECURE: Encrypted storage in IPFS, decrypted only in memory/local cache
4. CONCURRENT SAFE: Reference counting prevents deletion during use
5. AUTO CLEANUP: Removes unused versions automatically
6. FAST ACCESS: Local file system access after first download

*/

package runtime

import (
	"context"
	"delong/internal/consts"
	"delong/internal/models"
	"delong/pkg/db"
	"delong/pkg/tee"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
	"time"

	"gorm.io/gorm"
)

type Exporter interface {
	Suffix() string
	Export(outputDir string) error
}

type MountVersionRef struct {
	RefCount int
}

type DatasetLoader struct {
	storageRoot string // e.g. /data/delong_dataset
	db          *gorm.DB
	ipfsStore   *db.IpfsStore
	keyVault    *tee.KeyVault
	refCounts   map[string]*MountVersionRef // version -> ref count,  e.g. 20250514_120000 -> 42
	rm          sync.RWMutex
}

func NewDatasetLoader(storageRoot string, db *gorm.DB, ipfsStore *db.IpfsStore, keyVault *tee.KeyVault) *DatasetLoader {
	return &DatasetLoader{
		storageRoot: storageRoot,
		db:          db,
		ipfsStore:   ipfsStore,
		keyVault:    keyVault,
		refCounts:   make(map[string]*MountVersionRef),
		rm:          sync.RWMutex{},
	}
}

func (dl *DatasetLoader) MustInit() error {
	return os.MkdirAll(dl.storageRoot, 0755)
}

// AcquireCurrent returns the path to the current version of a dataset and the version string
func (dl *DatasetLoader) AcquireCurrent(dataset string) (string, string, error) {
	if strings.HasPrefix(dataset, consts.StaticDatasetPrefix) {
		return dl.acquireStaticDataset(dataset)
	}

	curPath := filepath.Join(dl.storageRoot, "current")
	resolvedPath, err := os.Readlink(curPath)
	if err != nil {
		return "", "", fmt.Errorf("current symlink not found for dataset %s", dataset)
	}
	version := filepath.Base(resolvedPath)
	versionPath := filepath.Join(dl.storageRoot, version)
	dl.rm.Lock()
	defer dl.rm.Unlock()
	_, ok := dl.refCounts[version]
	if !ok {
		dl.refCounts[version] = &MountVersionRef{RefCount: 1}
	} else {
		dl.refCounts[version].RefCount++
	}
	return versionPath, version, nil
}

// acquireStaticDataset downloads and decrypts a static dataset from IPFS
func (dl *DatasetLoader) acquireStaticDataset(dataset string) (string, string, error) {
	ctx := context.Background()

	// Get static dataset info from database
	staticDataset, err := models.GetStcDatasetByName(dl.db, dataset)
	if err != nil {
		return "", "", fmt.Errorf("failed to find static dataset %s: %w", dataset, err)
	}

	// Create a version identifier for this static dataset
	version := fmt.Sprintf("static_%d_%s", staticDataset.ID, staticDataset.FileHash[:8])

	// Check if already downloaded and cached
	versionPath := filepath.Join(dl.storageRoot, version)
	filePath := filepath.Join(versionPath, fmt.Sprintf("%v.csv", staticDataset.Name)) // Remove leading "/"

	// Check if directory and file both exist
	if _, err := os.Stat(versionPath); err == nil {
		if _, err := os.Stat(filePath); err == nil {
			log.Printf("Dataset already cached at %s", versionPath)
			// Both directory and file exist
			dl.rm.Lock()
			defer dl.rm.Unlock()
			_, ok := dl.refCounts[version]
			if !ok {
				dl.refCounts[version] = &MountVersionRef{RefCount: 1}
			} else {
				dl.refCounts[version].RefCount++
			}
			return versionPath, version, nil
		}
	}

	// Download and decrypt from IPFS
	kc := tee.NewKeyContext(tee.KEYKIND_ENC_KEY, staticDataset.Author, consts.PurposeEncStaticDataset)
	key, err := dl.keyVault.DeriveSymmetricKey(ctx, kc)
	if err != nil {
		return "", "", fmt.Errorf("failed to derive decryption key: %w", err)
	}

	decryptedData, err := dl.ipfsStore.DownloadDecrypted(ctx, staticDataset.IpfsCid, key)
	if err != nil {
		return "", "", fmt.Errorf("failed to download and decrypt static dataset: %w", err)
	}

	// Create directory structure
	if err := os.MkdirAll(versionPath, 0755); err != nil {
		return "", "", fmt.Errorf("failed to create directory structure: %w", err)
	}

	// Write decrypted data to file
	if err := os.WriteFile(filePath, decryptedData, 0644); err != nil {
		return "", "", fmt.Errorf("failed to write decrypted data: %w", err)
	}

	// Update reference count
	dl.rm.Lock()
	defer dl.rm.Unlock()
	dl.refCounts[version] = &MountVersionRef{RefCount: 1}

	log.Printf("Successfully downloaded and cached static dataset %s to %s", dataset, filePath)
	return versionPath, version, nil
}

// Release releases a dataset version, decrementing its reference count
func (dl *DatasetLoader) Release(dataset, version string) {
	dl.rm.Lock()
	defer dl.rm.Unlock()
	if ref, ok := dl.refCounts[version]; ok {
		ref.RefCount--
	}
}

// Cleanup removes older version dataset with no references
func (dl *DatasetLoader) Cleanup() error {
	currentPath := filepath.Join(dl.storageRoot, "current")
	target, err := os.Readlink(currentPath)
	if err != nil {
		log.Println("No symbolic link existed while cleaning up")
		return nil
	}
	currentVersion := filepath.Base(target)

	dl.rm.Lock()
	defer dl.rm.Unlock()

	for version, ref := range dl.refCounts {
		if version == currentVersion {
			continue
		}
		if ref.RefCount <= 0 {
			datasetDir := filepath.Join(dl.storageRoot, version)
			// check out: ensure no running containers use this directory
			if _, err := os.Stat(datasetDir); err == nil {
				log.Printf("Cleaning up unused dataset directory: %s", datasetDir)
				os.RemoveAll(datasetDir)
			}
			delete(dl.refCounts, version)
		}
	}
	return nil
}

// Suffix returns the file suffix for the dataset
func (dl *DatasetLoader) Suffix() string {
	return ".csv"
}

// Export export the specified dataset to CSV in the given output directory
func (dl *DatasetLoader) Export() error {
	version := time.Now().Format("20060102_150405")
	datasetDir := filepath.Join(dl.storageRoot, version)

	var datasets []models.DynamicDataset
	if err := dl.db.Find(&datasets).Error; err != nil {
		return fmt.Errorf("failed to list datasets: %w", err)
	}

	if err := os.MkdirAll(datasetDir, 0755); err != nil {
		return fmt.Errorf("failed to create dataset dir: %w", err)
	}

	for _, dataset := range datasets {
		fileName := dataset.Name + dl.Suffix()
		f, err := os.Create(filepath.Join(datasetDir, fileName))
		if err != nil {
			return fmt.Errorf("failed to create file: %w", err)
		}
		defer f.Close()

		w := csv.NewWriter(f)
		defer w.Flush()

		// Write csv header
		header := []string{
			"id", "test_report_id", "category", "name", "definition", "result",
			"reference_range", "explanation", "status", "suggestions",
		}
		w.Write(header)

		// Fetch test report Ids
		var testReportIds []uint
		err = dl.db.Model(&models.TestReport{}).
			Where("dataset = ?", dataset.Name).
			Pluck("id", &testReportIds).Error
		if err != nil {
			return fmt.Errorf("failed to fetch test report IDs: %w", err)
		}

		// Paging query
		pageSize := 10000
		for offset := 0; offset < len(testReportIds); offset += pageSize {
			end := offset + pageSize
			end = min(end, len(testReportIds))
			batchIds := testReportIds[offset:end]
			var results []models.TestResult
			err := dl.db.Where("test_report_id IN ?", batchIds).Find(&results).Error
			if err != nil {
				return fmt.Errorf("failed to query test results: %w", err)
			}
			for _, result := range results {
				row := []string{
					strconv.FormatUint(uint64(result.ID), 10),
					strconv.Itoa(result.TestReportID),
					result.Category,
					result.Name,
					result.Definition,
					result.Result,
					result.ReferenceRange,
					result.Explanation,
					result.Status,
				}
				if result.Suggestions != nil {
					row = append(row, *result.Suggestions)
				} else {
					row = append(row, "")
				}
				w.Write(row)
			}
			w.Flush()
			if err := w.Error(); err != nil {
				return fmt.Errorf("failed to write CSV: %w", err)
			}
		}
	}

	// Atomically move to current
	currentLink := filepath.Join(dl.storageRoot, "current")
	tmpLink := filepath.Join(dl.storageRoot, fmt.Sprintf(".tmp_current_%d", time.Now().UnixNano()))
	if err := os.Symlink(datasetDir, tmpLink); err != nil {
		return fmt.Errorf("failed to create tmp symlink: %w", err)
	}
	if err := os.Rename(tmpLink, currentLink); err != nil {
		return fmt.Errorf("failed to replace current symlink: %w", err)
	}

	return nil
}
