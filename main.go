package main

import (
	"TodoApp/controllers"
	"TodoApp/controllers/gin_rest"
)

func main() {
	ginController := gin_rest.GinController{}
	controllers.Run(ginController)
}
