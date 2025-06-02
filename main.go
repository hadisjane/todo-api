package main

import (
	"TodoApp/controllers"
	"TodoApp/controllers/gin_rest"
)
// "TodoApp/controllers"
// "TodoApp/controllers/console"


func main() {
	ginController := gin_rest.GinController{}
	controllers.Run(ginController)

	// consoleController := &console.ConsoleController{}
	// controllers.Run(consoleController)
}
