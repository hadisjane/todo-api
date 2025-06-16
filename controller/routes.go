package controller

import (
	"fmt"
	"TodoApp/middleware"

	"github.com/gin-gonic/gin"
)

func RunServer() error {
	r := gin.Default()

	r.GET("/", Ping)

	authG := r.Group("/auth")
	{
		authG.POST("/register", Register)
		authG.POST("/login", Login)
	}

	apiG := r.Group("/api", middleware.CheckUserAuthentication)

	todosG := apiG.Group("/todos")
	{
		todosG.GET("", ListTasks)
		todosG.GET("/:id", GetTask)
		todosG.POST("", CreateTask)
		todosG.PUT("/:id", CompleteTask)
		todosG.DELETE("/:id", DeleteTask)
	}

	err := r.Run(":8989")
	if err != nil {
		fmt.Println("Error running server:", err)
	}
	return nil
}