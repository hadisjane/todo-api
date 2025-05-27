package services

import (
	"TodoApp/errs"
	"TodoApp/models"
	"TodoApp/repositories"
	
)

func CreateTask(title string, done bool) (*models.Task, error) {
	if title == "" {
		return nil, errs.ErrTaskInvalid
	}

	task := models.Task{
		Title:     title,
		Done:      done,
	}

	return repositories.CreateTask(task.Title, task.Done)
}

func GetTask(id int) (*models.Task, error) {
	return repositories.GetTask(id)
}

func DeleteTask(id int) error {
	return repositories.DeleteTask(id)
}

func CompleteTask(id int) (*models.Task, error) {
	return repositories.CompleteTask(id)
}

func ListTasks() []models.Task {
	return repositories.ListTasks()
}
