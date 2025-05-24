package main

import (
	"log"
	"os"

	"backend/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading env")
	}
	// Initialize Gin router
	router := gin.Default()
	router.Use(cors.Default())
	routes.SetupRouter(router)

	// Run the server
	PORT := os.Getenv("PORT")
	router.Run(":" + PORT)
}
