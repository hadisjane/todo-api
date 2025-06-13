package models

import "time"

type Task struct {
	ID        int       `json:"id" db:"id"`
	Title     string    `json:"title" binding:"required" db:"title"`
	Done      bool      `json:"done" db:"done"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}
