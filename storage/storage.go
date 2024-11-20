package storage

import (
	"fmt"
	"os"

	"github.com/aelpxy/fresh/models"
	_ "github.com/mattn/go-sqlite3"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type StorageSystem struct {
	db       *gorm.DB
	rootPath string
}

type ObjectSystem struct {
	bucket string
	// db     *gorm.DB
}

func NewFreshStorageSystem(rootPath string) (*StorageSystem, error) {
	if err := os.MkdirAll(rootPath, 0755); err != nil {
		return nil, fmt.Errorf("failed to create root directory: %v", err)
	}

	dsn := fmt.Sprintf("file:./%s/storage.db?cache=shared&mode=rwc", rootPath)
	sqlStatement := "PRAGMA journal_mode=WAL;"

	db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %v", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("failed to get SQL DB: %v", err)
	}

	if _, err := sqlDB.Exec(sqlStatement); err != nil {
		return nil, fmt.Errorf("failed to set WAL mode: %v", err)
	}

	if err := db.AutoMigrate(&models.Bucket{}, &models.Object{}); err != nil {
		return nil, fmt.Errorf("failed to migrate database: %v", err)
	}

	return &StorageSystem{
		db:       db,
		rootPath: rootPath,
	}, nil
}
