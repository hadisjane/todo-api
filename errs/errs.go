package errs

import "errors"

var (
	ErrTaskNotFound = errors.New("task not found")
	ErrTaskAlreadyExists = errors.New("task already exists")
	ErrTaskInvalid = errors.New("task is invalid")
	ErrTaskAlreadyCompleted = errors.New("task already completed")
	ErrTaskTitleEmpty = errors.New("task title is empty")
	ErrInvalidId = errors.New("invalid id")
)