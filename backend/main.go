package main

import (
	"log"
	"os"

	"backend/database"
	"backend/models"
	"backend/routes"
	"backend/worker"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading env")
	}
	database.Connect()
	if err := database.AutoMigrate(&models.User{}, &models.Log{}, &models.Task{}); err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	if err != nil {
		log.Fatal("Error connecting to database")
	}
	// Initialize Gin router
	router := gin.Default()
	router.Use(cors.Default())
	routes.SetupRouter(router)
	go worker.NotiWorker()
	go worker.StartPingWorker()

	// Run the server
	PORT := os.Getenv("PORT")
	router.Run(":" + PORT)
}
