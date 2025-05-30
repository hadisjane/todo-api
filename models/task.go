package models

import "time"

type Task struct {
	ID        int       `json:"id"`
	Title     string    `json:"title" binding:"required"`
	Done      bool      `json:"done"`
	CreatedAt time.Time `json:"created_at"`
}
