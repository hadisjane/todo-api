package controller

import (
	"TodoApp/internal/errs"
	"TodoApp/internal/middleware"
	"TodoApp/internal/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "TodoApp server up and running",
	})
}

func CreateTask(c *gin.Context) {
	userID := c.GetInt(middleware.UserIDCtx)
	if userID == 0 {
		HandleError(c, errs.ErrUnauthorized)
		return
	}

	var req struct {
		Title string `json:"title"`
		Done  bool   `json:"done"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		HandleError(c, err)
		return
	}

	if req.Title == "" {
		HandleError(c, errs.ErrTaskTitleEmpty)
		return
	}

	task, err := service.CreateTask(userID, req.Title, req.Done)
	if err != nil {
		HandleError(c, err)
		return
	}

	c.JSON(http.StatusCreated, task)
}

func GetTask(c *gin.Context) {
	userID := c.GetInt(middleware.UserIDCtx)
	if userID == 0 {
		HandleError(c, errs.ErrUnauthorized)
		return
	}

	// Get ID from URL parameter
	idStr := c.Param("id")
	taskID, err := strconv.Atoi(idStr)
	if err != nil {
		HandleError(c, errs.ErrInvalidId)
		return
	}

	task, err := service.GetUserTask(userID, taskID)
	if err != nil {
		HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, task)
}

func DeleteTask(c *gin.Context) {
	userID := c.GetInt(middleware.UserIDCtx)
	if userID == 0 {
		HandleError(c, errs.ErrUnauthorized)
		return
	}

	// Get ID from URL parameter
	idStr := c.Param("id")
	taskID, err := strconv.Atoi(idStr)
	if err != nil {
		HandleError(c, errs.ErrInvalidId)
		return
	}

	err = service.DeleteUserTask(userID, taskID)
	if err != nil {
		HandleError(c, err)
		return
	}

	c.Status(http.StatusNoContent)
}

func CompleteTask(c *gin.Context) {
	userID := c.GetInt(middleware.UserIDCtx)
	if userID == 0 {
		HandleError(c, errs.ErrUnauthorized)
		return
	}

	// Get ID from URL parameter
	idStr := c.Param("id")
	taskID, err := strconv.Atoi(idStr)
	if err != nil {
		HandleError(c, errs.ErrInvalidId)
		return
	}

	task, err := service.CompleteTask(userID, taskID)
	if err != nil {
		HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, task)
}

func ListTasks(c *gin.Context) {
	userID := c.GetInt(middleware.UserIDCtx)
	if userID == 0 {
		HandleError(c, errs.ErrUnauthorized)
		return
	}

	tasks, err := service.ListUserTasks(userID)
	if err != nil {
		HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, tasks)
}
