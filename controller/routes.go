package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func RunServer() error {
	r := gin.Default()

	r.GET("/", Ping)
	r.GET("/todos", ListTasks)
	r.GET("/todos/:id", GetTask)
	r.POST("/todos", CreateTask)
	r.PUT("/todos/:id", CompleteTask)
	r.DELETE("/todos/:id", DeleteTask)

	err := r.Run(":8989")
	if err != nil {
		fmt.Println("Error running server:", err)
	}
	return nil
}