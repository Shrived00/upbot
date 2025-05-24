package database

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func Connect() {
	DATABASE_URL := os.Getenv("DATABASE_URL")
	if DATABASE_URL == "" {
		log.Fatal("DATABASE_URL not set in environment")
	}
	log.Println("Trying to connect to database at:", DATABASE_URL)

	logger := logger.Default.LogMode((logger.Silent))
	connection, err := gorm.Open(postgres.Open(DATABASE_URL), &gorm.Config{
		Logger: logger,
	})
	if err != nil {
		log.Fatal("Failed to connect to database!")
	}

	DB = connection
	log.Println("Successfully connected to the database")

}

func AutoMigrate(models ...interface{}) error {
	return DB.AutoMigrate(models...)
}
