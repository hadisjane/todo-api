package db

import (
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var db *sqlx.DB

func ConnDB() error {
	if err := godotenv.Load(); err != nil {
		return err
	}

	var err error

	dsn := os.Getenv("DB_DSN")
	if db, err = sqlx.Connect("postgres", dsn); err != nil {
		return err
	}

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