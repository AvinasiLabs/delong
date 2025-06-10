//go:build integration
// +build integration

package db

import (
	"delong/internal"
	"testing"
)

func TestNewMysqlDb(t *testing.T) {
	config, err := internal.LoadConfigFromEnv()
	if err != nil {
		t.Fatalf("failed to load config from env: %v", err)
	}

	db, err := NewMysqlDb(config.MysqlDsn)
	if err != nil {
		t.Fatalf("failed to connect to MySQL: %v", err)
	}
	defer func() {
		sqlDb, _ := db.DB()
		sqlDb.Close()
	}()

	sqlDB, err := db.DB()
	if err != nil {
		t.Fatalf("failed to get sql.DB: %v", err)
	}

	if err := sqlDB.Ping(); err != nil {
		t.Fatalf("failed to ping MySQL: %v", err)
	}
}
