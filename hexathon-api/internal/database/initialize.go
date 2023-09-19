package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Global variable to access the database
var DB *gorm.DB

// Initialize Connection to the database
func Connect(debug string, dbUrl string) {
	var err error

	if debug == "true" {
		DB, err = gorm.Open(postgres.Open(dbUrl), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
		})
	} else {
		DB, err = gorm.Open(postgres.Open(dbUrl), &gorm.Config{})
	}

	if err != nil {
		log.Fatal("Failed to connect to database!")
	}

	log.Println("Connected to database!")
}
