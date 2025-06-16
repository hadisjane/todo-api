package repository

import (
	"TodoApp/db"
	"TodoApp/errs"
	"TodoApp/models"
)


func CreateTask(userID int, title string, done bool) (*models.Task, error) {
	// Проверяем, существует ли уже задача с таким названием у этого пользователя
	existingTasks, err := ListUserTasks(userID)
	if err != nil {
		return nil, translateError(err)
	}

	for _, t := range existingTasks {
		if t.Title == title {
			return nil, errs.ErrTaskAlreadyExists
		}
	}

	// Начинаем транзакцию для атомарного создания задачи
	tx, err := db.GetDB().Beginx()
	if err != nil {
		return nil, translateError(err)
	}
	defer tx.Rollback()

	// Получаем следующий ID для задачи пользователя
	var taskID int
	err = tx.Get(&taskID, "SELECT next_task_id($1)", userID)
	if err != nil {
		return nil, translateError(err)
	}

	// Создаем задачу с указанным ID
	task := &models.Task{}
	err = tx.Get(task, `
		INSERT INTO tasks (id, user_id, title, done) 
		VALUES ($1, $2, $3, $4) 
		RETURNING id, user_id, title, done, created_at`,
		taskID, userID, title, done)

	if err != nil {
		return nil, translateError(err)
	}

	if err := tx.Commit(); err != nil {
		return nil, translateError(err)
	}

	return task, nil
}

func ListUserTasks(userID int) ([]models.Task, error) {
	var tasks []models.Task
	err := db.GetDB().Select(&tasks, 
		`SELECT id, user_id, title, done, created_at 
		FROM tasks 
		WHERE user_id = $1 
		ORDER BY id`,
		userID)
	if err != nil {
		return nil, translateError(err)
	}
	return tasks, nil
}

func GetUserTask(userID, taskID int) (*models.Task, error) {
	t := models.Task{}
	err := db.GetDB().Get(&t, 
		"SELECT id, user_id, title, done, created_at FROM tasks WHERE id = $1 AND user_id = $2", 
		taskID, userID)
	if err != nil {
		return nil, translateError(err)
	}
	return &t, nil
}

func CompleteTask(userID, taskID int) (*models.Task, error) {
	t := models.Task{}
	err := db.GetDB().Get(&t, 
		"UPDATE tasks SET done = true WHERE id = $1 AND user_id = $2 RETURNING id, user_id, title, done, created_at", 
		taskID, userID)
	if err != nil {
		return nil, translateError(err)
	}
	return &t, nil
}
			

func DeleteUserTask(userID, taskID int) error {
	result, err := db.GetDB().Exec(
		"DELETE FROM tasks WHERE id = $1 AND user_id = $2", 
		taskID, userID)
	if err != nil {
		return translateError(err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return translateError(err)
	}

	if rowsAffected == 0 {
		return errs.ErrNotFound
	}
	return nil
}
