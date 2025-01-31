package common

import (
	"fmt"
	"log"
	"time"

	"github.com/hiepnguyen223/int3306-project/configs"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB

// InitDB initializes the database connection
func InitDB() *gorm.DB {
	db = connectPostgresDb()
	return db
}

// connectPostgresDb establishes a connection to PostgreSQL database
func connectPostgresDb() *gorm.DB {
	DB_USERNAME := configs.EnvDbUserName()
	DB_PASSWORD := configs.EnvDbPassword()
	DB_NAME := configs.EnvDbName()
	DB_HOST := configs.EnvDbHost()
	DB_PORT := configs.EnvDbPort()
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		DB_HOST, DB_PORT, DB_USERNAME, DB_PASSWORD, DB_NAME,
	)

	gormConfig := &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info), // Enable detailed logs
	}

	var err error
	for i := 0; i < 5; i++ { // Retry mechanism
		db, err = gorm.Open(postgres.Open(dsn), gormConfig)
		if err == nil {
			log.Println("Successfully connected to PostgreSQL database.")
			return db
		}
		log.Printf("Database connection failed. Retrying in 2 seconds... (Attempt %d)", i+1)
		time.Sleep(2 * time.Second)
	}

	log.Fatalf("Error connecting to PostgreSQL database: %v", err)
	return nil
}

// GetDB returns the database instance
func GetDB() *gorm.DB {
	if db == nil {
		log.Fatal("Database not initialized. Call InitDB() first.")
	}
	return db
}
