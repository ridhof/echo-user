package main

import (
	"echo-user/controllers"
	"echo-user/config"

	"github.com/labstack/echo"
)

func main() {
	config.InitDB()

	e := echo.New()
	e.GET("/users", controllers.GetUsersController)
	e.GET("/users/:id", controllers.GetUserController)
	e.POST("/users", controllers.CreateUserController)
	e.DELETE("/users/:id", controllers.DeleteUserController)
	e.PUT("/users/:id", controllers.UpdateUserController)

	e.Logger.Fatal(e.Start(":8000"))
}