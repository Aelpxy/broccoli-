package main

import (
	"github.com/aelpxy/fresh/cmd"
	"github.com/aelpxy/fresh/models"
	"github.com/charmbracelet/log"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	dsn := "file:./dev.db?cache=shared&mode=rwc"
	sqlStatement := "PRAGMA journal_mode=WAL;"

	database, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database", err)
	}

	sqlDB, err := database.DB()
	if err != nil {
		log.Fatal("failed to get SQL DB", err)
	}

	_, err = sqlDB.Exec(sqlStatement)
	if err != nil {
		log.Fatal("failed to set WAL mode", err)
	}

	err = database.AutoMigrate(&models.Bucket{}, &models.Object{})
	if err != nil {
		log.Fatal("failed to migrate database", err)
	}

	if err != nil {
		log.Fatal(err)
	}

	cmd.Execute()
}
