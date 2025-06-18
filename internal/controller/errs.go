package controller

import (
	"TodoApp/internal/errs"
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleError(c *gin.Context, err error) {
	if err == nil {
		return
	}

	// 404 Not Found
	if errors.Is(err, errs.ErrTaskNotFound) ||
		errors.Is(err, errs.ErrNotFound) {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	// 400 Bad Request
	if errors.Is(err, errs.ErrTaskAlreadyExists) ||
		errors.Is(err, errs.ErrTaskAlreadyCompleted) ||
		errors.Is(err, errs.ErrTaskTitleEmpty) ||
		errors.Is(err, errs.ErrInvalidId) ||
		errors.Is(err, errs.ErrUserAlreadyExists) ||
		errors.Is(err, errs.ErrIncorrectUsernameOrPassword) ||
		errors.Is(err, errs.ErrUnauthorized) {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// 401 Unauthorized
	if errors.Is(err, errs.ErrUnauthorized) ||
		errors.Is(err, errs.ErrIncorrectUsernameOrPassword) {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		return
	}

	// 403 Forbidden
	if errors.Is(err, errs.ErrUserAlreadyExists) ||
		errors.Is(err, errs.ErrUnauthorized) ||
		errors.Is(err, errs.ErrIncorrectUsernameOrPassword) {
		c.JSON(http.StatusForbidden, gin.H{
			"error": err.Error(),
		})
		return
	}

	// 500 Internal Server Error
	c.JSON(http.StatusInternalServerError, gin.H{
		"error": fmt.Sprintf("something went wrong: %s", err.Error()),
	})
}
