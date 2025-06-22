package main

import (
	"book-api/internal/book"
	"book-api/internal/jobs"
	"book-api/internal/routes"
	"book-api/pkg/database"
	"context"
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	database.Connect()
	database.DB.AutoMigrate(&book.Book{})

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	jobs.StartDispatcher(ctx, 3)

	r := gin.Default()
	routes.SetupRoutes(r)

	go func() {
		if err := r.Run(":8080"); err != nil {
			log.Fatalf("Server failed: %v", err)
		}
	}()

	// Wait for shutdown signal
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")
	cancel()
}
