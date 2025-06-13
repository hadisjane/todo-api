package repository

import (
	"TodoApp/db"
	"TodoApp/errs"
	"TodoApp/models"
)


func CreateTask(title string, done bool) (*models.Task, error) {
	// First check if task with same title already exists
	existingTasks, err := ListTasks()
	if err != nil {
		return nil, translateError(err)
	}

	for _, t := range existingTasks {
		if t.Title == title {
			return nil, errs.ErrTaskAlreadyExists
		}
	}

	task := &models.Task{}
	err = db.GetDB().Get(task, 
		"INSERT INTO tasks (title, done) VALUES ($1, $2) RETURNING id, title, done, created_at", 
		title, done)
	if err != nil {
		return nil, translateError(err)
	}
	return task, nil
}

func ListTasks() ([]models.Task, error) {
	var tasks []models.Task
	err := db.GetDB().Select(&tasks, "SELECT id, title, done, created_at FROM tasks")
	if err != nil {
		return nil, translateError(err)
	}
	return tasks, nil

}

func GetTask(id int) (*models.Task, error) {
	t := models.Task{}
	err := db.GetDB().Get(&t, "SELECT id, title, done, created_at FROM tasks WHERE id = $1", id)
	if err != nil {
		return nil, translateError(err)
	}
	return &t, nil
}

func CompleteTask(id int) (*models.Task, error) {
	task := &models.Task{}
	err := db.GetDB().Get(task, "UPDATE tasks SET done = true WHERE id = $1 RETURNING id, title, done, created_at", id)
	if err != nil {
		return nil, translateError(err)
	}
	return task, nil
}
			

func DeleteTask(id int) error {
	task, err := GetTask(id)
	if err != nil {
		return translateError(err)
	}
	if task == nil {
		return errs.ErrTaskNotFound
	}

	_, err = db.GetDB().Exec("DELETE FROM tasks WHERE id = $1", id)
	if err != nil {
		return translateError(err)
	}
	return nil
}
