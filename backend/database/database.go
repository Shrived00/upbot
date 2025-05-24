package database

import (
	"log"
	"os"
	"strings"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func Connect() {
	DATABASE_URL := strings.TrimSpace(os.Getenv("DATABASE_URL"))
	if DATABASE_URL == "" {
		log.Fatal("DATABASE_URL not set in environment")
	}
	log.Println("Trying to connect to database at:", DATABASE_URL)

	logger := logger.Default.LogMode((logger.Silent))
	connection, err := gorm.Open(postgres.Open(DATABASE_URL), &gorm.Config{
		Logger: logger,
	})
	if err != nil {
		log.Fatal("Failed to connect to database!", err)
	}

	DB = connection
	log.Println("Successfully connected to the database")

}

func AutoMigrate(models ...interface{}) error {
	return DB.AutoMigrate(models...)
}
