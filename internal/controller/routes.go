package controller

import (
	"TodoApp/internal/configs"
	"TodoApp/internal/middleware"
	"TodoApp/logger"
	"fmt"

	"github.com/gin-gonic/gin"
)

func RunServer() error {
	// Set Gin mode based on configuration
	if configs.AppSettings.AppParams.GinMode == "release" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	r := gin.Default()

	// Add logging middleware
	r.Use(gin.LoggerWithWriter(logger.Info.Writer()))
	r.Use(gin.Recovery())

	// Health check endpoint
	r.GET("/", Ping)

	// Auth routes
	authG := r.Group("/auth")
	{
		authG.POST("/register", Register)
		authG.POST("/login", Login)
	}

	// API routes with authentication middleware
	apiG := r.Group("/api", middleware.CheckUserAuthentication)

	// Todo routes
	todosG := apiG.Group("/todos")
	{
		todosG.GET("", ListTasks)
		todosG.GET("/:id", GetTask)
		todosG.POST("", CreateTask)
		todosG.PUT("/:id", CompleteTask)
		todosG.DELETE("/:id", DeleteTask)
	}

	// Get server address from config
	serverAddr := ":" + configs.AppSettings.AppParams.PortRun
	if configs.AppSettings.AppParams.PortRun[0] == ':' {
		serverAddr = configs.AppSettings.AppParams.PortRun
	}

	logger.Info.Printf("Starting server on %s", serverAddr)
	
	// Start the server
	if err := r.Run(serverAddr); err != nil {
		logger.Error.Printf("Error running server: %v", err)
		return fmt.Errorf("failed to start server: %v", err)
	}

	return nil
}
