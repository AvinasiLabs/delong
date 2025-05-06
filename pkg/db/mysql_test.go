package db

import (
	"testing"
)

func TestNewMysqlDb(t *testing.T) {
	dsn := "root:root123@tcp(localhost:3306)/delong?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := NewMysqlDb(dsn)
	if err != nil {
		t.Fatalf("failed to connect to MySQL: %v", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		t.Fatalf("failed to get sql.DB: %v", err)
	}

	if err := sqlDB.Ping(); err != nil {
		t.Fatalf("failed to ping MySQL: %v", err)
	}
}
