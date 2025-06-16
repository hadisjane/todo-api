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
		// Return ErrNotFound for user-related operations
		// The service layer will translate it to the appropriate error
		return errs.ErrNotFound
	} else {
		return err
	}
}
