package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/nsiebenaller/go-api-example/router"
	"github.com/nsiebenaller/go-api-example/services"
)

func main() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Connect to database
	log.Println("Connecting to database")
	services.ConnectDb()

	// Create server
	server := gin.Default()
	server.SetTrustedProxies(nil)

	// Configure router
	log.Println("Configuring router")
	router.ConfigureRouter(server)

	// Start server
	server.Run("localhost:8080")
}
