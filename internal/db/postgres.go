package db

import (
	"TodoApp/internal/configs"
	"fmt"
	"os"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var db *sqlx.DB

func ConnDB() error {
	var err error

	// Build DSN from config
	cfg := configs.AppSettings.PostgresParams
	
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.Host,
		cfg.Port,
		cfg.User,
		os.Getenv("DB_PASSWORD"), // Get password from environment
		cfg.Database,
	)

	db, err = sqlx.Connect("postgres", dsn)
	if err != nil {
		return fmt.Errorf("error connecting to database: %v", err)
	}

	// Set connection pool settings
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)

	return nil
}

func CloseDB() error {
	if db != nil {
		return db.Close()
	}
	return nil
}

func GetDB() *sqlx.DB {
	return db
}
