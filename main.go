package main

import (
	"TodoApp/internal/configs"
	"TodoApp/internal/controller"
	"TodoApp/internal/db"
	"TodoApp/logger"
	"log"
)


func main() {
	// Load configurations
	if err := configs.ReadSettings(); err != nil {
		log.Fatalf("Failed to load configurations: %v", err)
	}

	// Initialize logger
	if err := logger.Init(); err != nil {
		log.Fatalf("Failed to initialize logger: %v", err)
	}

	// Initialize database connection
	if err := db.ConnDB(); err != nil {
		logger.Error.Fatalf("Error connecting to database: %v", err)
	}
	logger.Info.Println("Database connection established successfully")

	// Run migrations
	if err := db.InitMigrations(); err != nil {
		logger.Error.Fatalf("Error initializing migrations: %v", err)
	}
	logger.Info.Println("Migrations initialized successfully")

	// Start the server
	if err := controller.RunServer(); err != nil {
		logger.Error.Fatalf("Error running server: %v", err)
	}
}
