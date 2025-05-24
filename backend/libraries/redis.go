package libraries

import (
	"context"
	"log"
	"os"
	"strconv"
	"sync"

	"github.com/go-redis/redis/v8"
	"github.com/joho/godotenv"
)

var lock = &sync.Mutex{}

type Singleton struct {
	client *redis.Client
}

var instance *Singleton

func GetInstance() *redis.Client {
	if instance == nil {
		lock.Lock()
		defer lock.Unlock()
		if instance == nil {
			instance = &Singleton{
				client: GetClient(),
			}
		}
	}
	return instance.client
}

func GetClient() *redis.Client {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: .env file not found, reading environment variables directly")
	}

	addr := os.Getenv("REDIS_ADDR")
	password := os.Getenv("REDIS_PASSWORD")
	username := os.Getenv("REDIS_USERNAME")
	dbStr := os.Getenv("REDIS_DB")

	db := 0
	if dbStr != "" {
		db, err = strconv.Atoi(dbStr)
		if err != nil {
			log.Printf("Invalid REDIS_DB value, defaulting to 0: %v", err)
			db = 0
		}
	}

	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		Username: username,
		DB:       db,
	})
	ctx := context.Background()
	_, err = client.Ping(ctx).Result()
	if err != nil {
		log.Fatal("SERVER - Error connecting to redis")
	}
	log.Print("SERVER - Connected to redis")
	return client
}
