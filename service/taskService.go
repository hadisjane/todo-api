package service

import (
	"TodoApp/errs"
	"TodoApp/models"
	"TodoApp/repository"
)

func CreateTask(title string, done bool) (*models.Task, error) {
	if title == "" {
		return nil, errs.ErrTaskInvalid
	}

	task := models.Task{
		Title:     title,
		Done:      done,
	}

	return repository.CreateTask(task.Title, task.Done)
}

func GetTask(id int) (*models.Task, error) {
	return repository.GetTask(id)
}

func DeleteTask(id int) error {
	return repository	.DeleteTask(id)
}

func CompleteTask(id int) (*models.Task, error) {
	return repository.CompleteTask(id)
}

func ListTasks() ([]models.Task, error) {
	return repository.ListTasks()
}
