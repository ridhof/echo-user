package routes

import (
	"echo-user/constants"
	"echo-user/controllers"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// New returns echo object that contains all routes.
func New() *echo.Echo {
	e := echo.New()

	e.GET("/users", controllers.GetUsersController)
	e.GET("/users/:id", controllers.GetUserController)
	e.POST("/users", controllers.CreateUserController)

	// JWT Group
	r := e.Group("/jwt")
	r.Use(middleware.JWT([]byte(constants.SECRET_JWT)))
	r.DELETE("/users/:id", controllers.DeleteUserController)
	r.PUT("/users/:id", controllers.UpdateUserController)

	return e
}
