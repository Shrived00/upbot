package database

import (
	"context"
	"crypto/tls"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client
var DB *mongo.Database

func Connect() {
	MONGO_URI := os.Getenv("MONGO_URI")
	DB_NAME := os.Getenv("MONGO_DB_NAME")

	if MONGO_URI == "" || DB_NAME == "" {
		log.Fatal("MONGO_URI or MONGO_DB_NAME not set in environment")
	}

	// Context with more generous timeout
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	clientOptions := options.Client().
		ApplyURI(MONGO_URI).
		SetTLSConfig(&tls.Config{}) // Force TLS

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal("MongoDB connection error:", err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal("MongoDB ping error:", err)
	}

	log.Println("Connected to MongoDB!")
	Client = client
	DB = client.Database(DB_NAME)
}
