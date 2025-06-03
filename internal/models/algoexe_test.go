package models

import (
	"database/sql"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func setupMockDB(t *testing.T) (*gorm.DB, sqlmock.Sqlmock) {
	sqlDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to create mock: %v", err)
	}

	gormDB, err := gorm.Open(mysql.New(mysql.Config{
		Conn:                      sqlDB,
		SkipInitializeWithVersion: true,
	}), &gorm.Config{})
	if err != nil {
		t.Fatalf("Failed to create gorm db: %v", err)
	}

	return gormDB, mock
}

// Common column definitions
var algoExeSelectColumns = []string{
	"id", "algo_id", "used_dataset", "scientist_wallet",
	"review_status", "vote_start_time", "vote_end_time", "status",
	"start_time", "end_time", "result", "error_msg", "created_at", "updated_at",
	"algo_name", "algo_link", "cid",
}

// COUNT query regex (only match the JOIN on blockchain_transactions and algos)
func countQueryRegex() string {
	return regexp.QuoteMeta("SELECT count(*) FROM `algo_exes`") + `.*JOIN.*blockchain_transactions.*JOIN.*algos.*`
}

// SELECT query regex for page=1 (OFFSET=0 is omitted)
func selectQueryRegexPage1() string {
	return regexp.QuoteMeta("SELECT") +
		`.*algo_exes\.\*.*algos\.name as algo_name.*algos\.algo_link.*algos\.cid.*` +
		`FROM.*algo_exes.*JOIN.*blockchain_transactions.*JOIN.*algos.*` +
		`ORDER BY.*created_at.*LIMIT \?`
}

// SELECT query regex for page>1 (includes LIMIT and OFFSET)
func selectQueryRegexPageN() string {
	return regexp.QuoteMeta("SELECT") +
		`.*algo_exes\.\*.*algos\.name as algo_name.*algos\.algo_link.*algos\.cid.*` +
		`FROM.*algo_exes.*JOIN.*blockchain_transactions.*JOIN.*algos.*` +
		`ORDER BY.*created_at.*LIMIT \? OFFSET \?`
}

func TestGetAlgoExesWithAlgoInfo_Unit(t *testing.T) {
	tests := []struct {
		name         string
		page         int
		pageSize     int
		setupMock    func(mock sqlmock.Sqlmock)
		expectCount  int
		expectTotal  int64
		expectError  bool
		expectErrMsg string
	}{
		{
			name:     "successful three-table join query (page=1, offset=0)",
			page:     1,
			pageSize: 10,
			setupMock: func(mock sqlmock.Sqlmock) {
				// 1) mock COUNT query
				mock.ExpectQuery(countQueryRegex()).
					WithArgs("CONFIRMED", "EXECUTION").
					WillReturnRows(sqlmock.NewRows([]string{"count(*)"}).AddRow(3))

				// 2) mock SELECT query for page=1 (only LIMIT)
				now := time.Now()
				rows := sqlmock.NewRows(algoExeSelectColumns).
					AddRow(1, 1, "dataset1", "0x123456789abcdef", "REVIEWING", nil, nil, "COMPLETED",
						nil, &now, "success result", "", now.Add(-2*time.Hour), now,
						"Machine Learning Algorithm", "https://github.com/user/ml-algo", "QmHash123").
					AddRow(2, 2, "dataset2", "0x987654321fedcba", "APPROVED", nil, nil, "COMPLETED",
						nil, &now, "another result", "", now.Add(-1*time.Hour), now,
						"Deep Learning Model", "https://github.com/user/dl-model", "QmHash456")

				mock.ExpectQuery(selectQueryRegexPage1()).
					WithArgs("CONFIRMED", "EXECUTION", 10). // only 3 args: two JOIN args + LIMIT
					WillReturnRows(rows)
			},
			expectCount: 2,
			expectTotal: 3,
			expectError: false,
		},
		{
			name:     "database error on count query",
			page:     1,
			pageSize: 10,
			setupMock: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery(countQueryRegex()).
					WithArgs("CONFIRMED", "EXECUTION").
					WillReturnError(sql.ErrConnDone)
			},
			expectError:  true,
			expectErrMsg: "connection is already closed",
		},
		{
			name:     "no confirmed transactions found (page=1)",
			page:     1,
			pageSize: 10,
			setupMock: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery(countQueryRegex()).
					WithArgs("CONFIRMED", "EXECUTION").
					WillReturnRows(sqlmock.NewRows([]string{"count(*)"}).AddRow(0))

				emptyRows := sqlmock.NewRows(algoExeSelectColumns)
				mock.ExpectQuery(selectQueryRegexPage1()).
					WithArgs("CONFIRMED", "EXECUTION", 10).
					WillReturnRows(emptyRows)
			},
			expectCount: 0,
			expectTotal: 0,
			expectError: false,
		},
		{
			name:     "pagination test - second page (page=2, offset>0)",
			page:     2,
			pageSize: 5,
			setupMock: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery(countQueryRegex()).
					WithArgs("CONFIRMED", "EXECUTION").
					WillReturnRows(sqlmock.NewRows([]string{"count(*)"}).AddRow(12))

				now := time.Now()
				rows := sqlmock.NewRows(algoExeSelectColumns).
					AddRow(6, 3, "dataset6", "0xabc123", "APPROVED", nil, nil, "COMPLETED",
						nil, &now, "result6", "", now.Add(-6*time.Hour), now,
						"Algorithm 6", "https://github.com/user/algo6", "QmHash6")

				mock.ExpectQuery(selectQueryRegexPageN()).
					WithArgs("CONFIRMED", "EXECUTION", 5, 5). // 4 args: two JOIN args + LIMIT + OFFSET
					WillReturnRows(rows)
			},
			expectCount: 1,
			expectTotal: 12,
			expectError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, mock := setupMockDB(t)
			tt.setupMock(mock)

			results, total, err := GetAlgoExesWithAlgoInfo(db, tt.page, tt.pageSize)

			if tt.expectError {
				assert.Error(t, err)
				if tt.expectErrMsg != "" {
					assert.Contains(t, err.Error(), tt.expectErrMsg)
				}
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectCount, len(results))
				assert.Equal(t, tt.expectTotal, total)

				if len(results) > 0 {
					first := results[0]
					assert.NotZero(t, first.ID)
					assert.NotEmpty(t, first.AlgoName)
					assert.NotEmpty(t, first.Cid)
					assert.NotEmpty(t, first.UsedDataset)
				}
			}

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("Mock expectations were not met: %v", err)
			}
		})
	}
}
