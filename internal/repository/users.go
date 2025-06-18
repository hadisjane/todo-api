package repository

import (
	"TodoApp/internal/db"
	"TodoApp/internal/models"
)

func GetUserByUsernameAndPassword(username string, password string) (models.User, error) {
	var user models.User
	err := db.GetDB().Get(&user, `
		SELECT id, username, created_at
		FROM users 
		WHERE username = $1 AND password = $2`, 
		username, password)

	if err != nil {
		return models.User{}, translateError(err)
	}

	return user, nil
}

func GetUserByUsername(username string) (user models.User, err error) {
	err = db.GetDB().Get(&user, `SELECT id, 
					   username, 
					   password,
					   created_at
				FROM users WHERE username = $1`, username)
	if err != nil {
		return models.User{}, translateError(err)
	}

	return user, nil
}

func CreateUser(user models.UserRegister) error {
	_, err := db.GetDB().Exec(`
		INSERT INTO users (username, email, password)
		VALUES ($1, $2, $3)`, 
		user.Username, 
		user.Email, 
		user.Password)
	return err
}
