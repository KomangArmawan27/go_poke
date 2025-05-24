package config

import (
	"fmt"
	"log"

	"go-api/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// ConnectDatabase initializes the PostgreSQL connection
func ConnectDatabase() {
	// Load environment variables
	LoadEnv()

	// Get database credentials from .env
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		GetEnv("PGHOST"),
		GetEnv("PGUSER"),
		GetEnv("PGPASSWORD"),
		GetEnv("PGDATABASE"),
		GetEnv("PGPORT"),
	)

	// Open connection
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	DB = db

	// Auto Migrate Models
	DB.AutoMigrate(&models.User{})
	DB.AutoMigrate(&models.Product{})
	DB.AutoMigrate(&models.Log{})

	fmt.Println("✅ Successfully connected to the database!")
}
