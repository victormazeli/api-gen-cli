package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"goApiStartetProject/api/middlewares"
	"goApiStartetProject/api/routes"
	"goApiStartetProject/internal/config"
	"goApiStartetProject/internal/database/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const defaultPort = "8080"

func main() {
	// Load config
	cfg := config.LoadEnvironmentConfig()

	// Initialize Database Connection
	db, err := gorm.Open(postgres.Open(cfg.Database.URL), &gorm.Config{})

	if err != nil {
		log.Fatal("Cannot connect to Database")
	}

	// Migrate Tables
	err = db.AutoMigrate(&models.User{})

	if err != nil {
		log.Fatalf(err.Error())
	}

	port := cfg.Server.Port
	if port == "" {
		port = defaultPort
	}

	// Initialize Gin router
	r := gin.Default()
	r.HandleMethodNotAllowed = true
	r.Use(middlewares.CORS())
	// Health
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Server is Running!")
	})
	// Not Found
	r.NoRoute(middlewares.NotFound())
	// Method Not Allowed
	r.NoMethod(middlewares.MethodNotAllowed())
	// Setup Route
	rootPath := r.Group("")
	routes.SetupRoute(cfg, nil, rootPath)

	// Setup server
	srv := &http.Server{
		Addr: fmt.Sprintf(":%v", port),
		// Good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      r, // Pass our instance of gin in.
	}

	// Run our server in a goroutine so that it doesn't block.
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)

		}
	}()

	quit := make(chan os.Signal, 1)
	// Accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	// Block until we receive our signal.
	<-quit

	log.Println("Shutting down server...")

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown: ", err)
	}

	log.Println("Server exiting")
}
