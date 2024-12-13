package database

import (
	"belajar-api/config"
	"belajar-api/models"
	"fmt"
	"log"

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

func Connect(cfg *config.Config) *gorm.DB {
	// Construct the SQL Server connection string
	dsn := fmt.Sprintf("sqlserver://%s:%s@%s:%s?database=%s", cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName)
	// Open a connection to the SQL Server database
	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to SQL Server: %v", err)
	}

	log.Println("Connected to SQL Server!")
	return db
}

func Migrate(db *gorm.DB) {
	// Automatically migrate your schema, to keep your schema up to date.
	err := db.AutoMigrate(
		models.MasterOrganization{},
		models.MasterNature{},
		models.MasterPriority{},
		models.UserBelajar{},
	)

	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	log.Println("Database migration completed!")
}
