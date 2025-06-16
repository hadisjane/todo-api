package main

import (
	"TodoApp/controller"
	"TodoApp/db"
	"fmt"
	"log"

	"github.com/joho/godotenv"
)

func init() {
	// Сначала загружаем переменные окружения
	if err := godotenv.Load(); err != nil {
		log.Printf("Warning: Error loading .env file: %v", err)
	}
}

func main() {
	if err := db.ConnDB(); err != nil {
		fmt.Println("Error connecting to database:", err)
		return
	}
	fmt.Println("Database connection established successfully")

	if err := db.InitMigrations(); err != nil {
		fmt.Println("Error initializing migrations:", err)
		return
	}
	fmt.Println("Migrations initialized successfully")

	if err := controller.RunServer(); err != nil {
		fmt.Println("Error running server:", err)
		return
	}
}
