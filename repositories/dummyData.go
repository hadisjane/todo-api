package repositories

import (
	"TodoApp/models"
	"time"
)

var tasks = []models.Task{
	{ID: 1, Title: "Анжуманя", Done: true, CreatedAt: time.Now()},
	{ID: 2, Title: "Прескачат", Done: false, CreatedAt: time.Now()},
	{ID: 3, Title: "Бегит", Done: false, CreatedAt: time.Now()},
}