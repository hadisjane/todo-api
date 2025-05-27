package errs

import "errors"

var (
	ErrTaskNotFound = errors.New("task not found")
	ErrTaskAlreadyExists = errors.New("task already exists")
	ErrTaskInvalid = errors.New("task is invalid")
	ErrTaskAlreadyCompleted = errors.New("task already completed")
)