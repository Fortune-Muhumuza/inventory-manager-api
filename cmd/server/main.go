package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/yourusername/inventory-manager-api/internal/inventory"
	"github.com/yourusername/inventory-manager-api/pkg/db"
	"github.com/yourusername/inventory-manager-api/api"
)

func main() {
	// Initialize the database connection
	database, err := db.Connect(os.Getenv("DB_DSN"))
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}
	defer database.Close()

	// Set up the router
	router := gin.Default()

	// Initialize routes
	api.SetupRoutes(router, inventory.NewService(database))

	// Start the server
	if err := router.Run(":8080"); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
