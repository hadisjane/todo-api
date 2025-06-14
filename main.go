package main

import (
	"TodoApp/controller"
	"TodoApp/db"
	"fmt"

	"github.com/joho/godotenv"
)


func main() {
	if err := godotenv.Load(); err != nil {
		fmt.Println("Error loading .env file:", err)
		return
	}

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
