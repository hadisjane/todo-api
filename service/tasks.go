package service

import (
	"TodoApp/errs"
	"TodoApp/models"
	"TodoApp/repository"
)

func CreateTask(userID int, title string, done bool) (*models.Task, error) {
	if title == "" {
		return nil, errs.ErrTaskInvalid
	}

	return repository.CreateTask(userID, title, done)
}

func GetUserTask(userID, taskID int) (*models.Task, error) {
	return repository.GetUserTask(userID, taskID)
}

func DeleteUserTask(userID, taskID int) error {
	return repository.DeleteUserTask(userID, taskID)
}

func CompleteTask(userID, taskID int) (*models.Task, error) {
	return repository.CompleteTask(userID, taskID)
}

func ListUserTasks(userID int) ([]models.Task, error) {
	return repository.ListUserTasks(userID)
}
