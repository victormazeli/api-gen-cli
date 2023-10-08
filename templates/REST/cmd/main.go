package main

import (
	"github.com/gin-gonic/gin"
	"goApiStartetProject/api/middlewares"
	"goApiStartetProject/api/routes"
	"goApiStartetProject/internal/config"
	"goApiStartetProject/internal/database"
	"log"
)

func main() {
	// Load config
	env := config.LoadEnvironmentConfig()

	// Initialize Database Connection
	Db := database.ConnectDB(env)

	// Migrate Tables
	err := Db.AutoMigrate()

	if err != nil {
		log.Fatalf(err.Error())
	}


	// Set Gin Mode
	gin.SetMode(gin.DebugMode)
	// Initialize Gin router
	r := gin.New()

	r.HandleMethodNotAllowed = true
	//Enable logger
	r.Use(gin.Logger())
	// Enable Recovery
	r.Use(gin.Recovery())

	// Apply middleware
	r.Use(middlewares.AuthMiddleware())

	// Setup Route
	rootPath := r.Group("")
	routes.SetupRoute(env, Db, rootPath)


	// Start the server
	r.Run(":8080")
}
