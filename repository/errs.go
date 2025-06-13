package repository

import (
	"TodoApp/errs"
	"database/sql"
	"errors"
)

func translateError(err error) error {
	if err == nil {
		return nil
	} else if errors.Is(err, sql.ErrNoRows) {
		return errs.ErrTaskNotFound
	} else {
		return err
	}
}
