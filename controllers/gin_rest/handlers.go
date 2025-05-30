package gin_rest

import (
	"TodoApp/errs"
	"TodoApp/services"
	"TodoApp/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "TodoApp server up and running",
	})
}

func CreateTask(c *gin.Context) {
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

	task, err := services.CreateTask(req.Title, req.Done)
	if err != nil {
		HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, task)
}

func GetTask(c *gin.Context) {
	id, err := utils.ExtractID(c.Request)
	if err != nil {
		HandleError(c, err)
		return
	}

	task, err := services.GetTask(id)
	if err != nil {
		HandleError(c, err)
		return
	}

	if task == nil {
		HandleError(c, errs.ErrTaskNotFound)
		return
	}

	c.JSON(http.StatusOK, task)
}

func DeleteTask(c *gin.Context) {
	id, err := utils.ExtractID(c.Request)
	if err != nil {
		HandleError(c, err)
		return
	}

	err = services.DeleteTask(id)
	if err != nil {
		HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Task deleted successfully",
	})
}

func CompleteTask(c *gin.Context) {
	id, err := utils.ExtractID(c.Request)
	if err != nil {
		HandleError(c, err)
		return
	}

	_, err = services.CompleteTask(id)
	if err != nil {
		HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Task completed successfully",
	})
}

func ListTasks(c *gin.Context) {
	tasks := services.ListTasks()
	c.JSON(http.StatusOK, tasks)
}
