package repositories

import (
	"TodoApp/errs"
	"TodoApp/models"
	"strings"
	"time"
)

var (
	lastID int
)


func init() {
	lastID = 0
	for _, t := range tasks {
		if t.ID > lastID {
			lastID = t.ID
		}
	}
}

func CreateTask(title string, done bool) (*models.Task, error) {
	for _, t := range tasks {
		if strings.EqualFold(t.Title, title) {
			return nil, errs.ErrTaskAlreadyExists
		}
	}

	lastID++

	task := models.Task{
		ID:        lastID,
		Title:     title,
		Done:      done,
		CreatedAt: time.Now(),
	}
	tasks = append(tasks, task)
	return &task, nil
}

func GetTask(id int) (*models.Task, error) {
	for _, task := range tasks {
		if task.ID == id {
			return &task, nil
		}
	}
	return nil, errs.ErrTaskNotFound
}

func DeleteTask(id int) error {
	for i, task := range tasks {
		if task.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			return nil
		}
	}
	return errs.ErrTaskNotFound
}

func CompleteTask(id int) (*models.Task, error) {
	for i := range tasks {
		if tasks[i].ID == id {
			if tasks[i].Done {
				return nil, errs.ErrTaskAlreadyCompleted
			}
			tasks[i].Done = true
			return &tasks[i], nil
		}
	}

	return nil, errs.ErrTaskNotFound
}

func ListTasks() []models.Task {
	return tasks
}
