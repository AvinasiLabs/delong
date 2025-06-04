package models

import (
	"log"

	"gorm.io/gorm"
)

// AutoMigrateDatabase performs database migration using GORM AutoMigrate
func AutoMigrateDatabase(db *gorm.DB) error {
	log.Println("Starting database migration...")

	// Define all models to migrate
	err := db.AutoMigrate(
		&Algo{},
		&AlgoExe{},
		&BlockchainTransaction{},
		&CommitteeMember{},
		&ContractMeta{},
		&DatasetRegistry{},
		&DataUsage{},
		&TestReport{},
		&TestResult{},
		&Vote{},
		&StaticDataset{},
	)

	if err != nil {
		log.Printf("Database migration failed: %v", err)
		return err
	}

	// Insert default dataset data if not exists
	// err = insertDefaultDatasets(db)
	// if err != nil {
	// 	log.Printf("Failed to insert default datasets: %v", err)
	// 	return err
	// }

	log.Println("Database migration completed successfully")
	return nil
}
