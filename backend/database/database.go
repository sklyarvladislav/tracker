package database

import (
	"log"

	"github.com/sklyarvladislav/tracker/backend/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

// Initialize sets up the database connection and runs migrations
func Initialize(dbPath string) error {
	var err error
	DB, err = gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		return err
	}

	log.Println("Database connection established")
	return nil
}

// Migrate runs database migrations
func Migrate() error {
	log.Println("Running database migrations...")
	err := DB.AutoMigrate(&models.Habit{}, &models.HabitEntry{})
	if err != nil {
		return err
	}
	log.Println("Migrations completed successfully")
	return nil
}
