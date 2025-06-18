package service

import (
	"TodoApp/internal/errs"
	"TodoApp/internal/models"
	"TodoApp/internal/repository"
	"TodoApp/utils"
	"errors"
)

func CreateUser(u models.UserRegister) error {
	_, err := repository.GetUserByUsername(u.Username)
	if err != nil {
		if errors.Is(err, errs.ErrNotFound) {
			// User doesn't exist, we can proceed with creation
		} else {
			return err
		}
	} else {
		return errs.ErrUserAlreadyExists
	}

	hashedPassword, err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}
	u.Password = hashedPassword

	if err = repository.CreateUser(u); err != nil {
		return err
	}

	return nil
}

func GetUserByUsernameAndPassword(username string, password string) (models.User, error) {
	user, err := repository.GetUserByUsername(username)
	if err != nil {
		if errors.Is(err, errs.ErrNotFound) {
			return models.User{}, errs.ErrIncorrectUsernameOrPassword
		}
		return models.User{}, err
	}

	err = utils.VerifyPassword(user.Password, password)
	if err != nil {
		return models.User{}, errs.ErrIncorrectUsernameOrPassword
	}

	return user, nil
}
