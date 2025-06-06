package runtime

import (
	"delong/internal/models"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
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
	refCounts   map[string]*MountVersionRef // version -> ref count,  e.g. 20250514_120000 -> 42
	rm          sync.RWMutex
}

func NewDatasetLoader(storageRoot string, db *gorm.DB) *DatasetLoader {
	return &DatasetLoader{
		storageRoot: storageRoot,
		db:          db,
		refCounts:   make(map[string]*MountVersionRef),
		rm:          sync.RWMutex{},
	}
}

func (dl *DatasetLoader) MustInit() error {
	return os.MkdirAll(dl.storageRoot, 0755)
}

// AcquireCurrent returns the path to the current version of a dataset and the version string
func (dl *DatasetLoader) AcquireCurrent(dataset string) (string, string, error) {
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

// Release releases a dataset version, decrementing its reference count
func (dl *DatasetLoader) Release(dataset, version string) {
	dl.rm.Lock()
	defer dl.rm.Unlock()
	if ref, ok := dl.refCounts[version]; ok {
		ref.RefCount--
	}
}

// CleanupExpired removes expired dataset versions
// func (m *Loader) CleanupExpired() error {
// 	cutoff := time.Now().Add(-ttl)
// 	entries, err := os.ReadDir(m.storageRoot)
// 	if err != nil {
// 		return err
// 	}
// 	for _, entry := range entries {
// 		if !entry.IsDir() {
// 			continue
// 		}
// 		dataset := entry.Name()
// 		versionRoot := filepath.Join(m.storageRoot, dataset)
// 		versions, err := os.ReadDir(versionRoot)
// 		if err != nil {
// 			continue
// 		}
// 		for _, v := range versions {
// 			if !v.IsDir() || v.Name() == "current" {
// 				continue
// 			}
// 			version := v.Name()
// 			verPath := filepath.Join(versionRoot, version)
// 			info, err := os.Stat(verPath)
// 			if err != nil || info.ModTime().After(cutoff) {
// 				continue
// 			}
// 			key := fmt.Sprintf("%s:%s", dataset, version)
// 			m.rm.RLock()
// 			ref, ok := m.refCounts[key]
// 			m.rm.RUnlock()
// 			if !ok || ref.RefCount == 0 {
// 				os.RemoveAll(verPath)
// 			}
// 		}
// 	}
// 	return nil
// }

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
		if ref.RefCount == 0 {
			datasetDir := filepath.Join(dl.storageRoot, version)
			os.RemoveAll(datasetDir)
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
