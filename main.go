package main

import (
	"net/http"

	"github.com/labstack/echo"
)

// GetUsersController get all users
func GetUsersController(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success get all users",
		"users": []string{},
	})
}

func main() {
	e := echo.New()
	e.GET("/users", GetUsersController)

	e.Logger.Fatal(e.Start(":8000"))
}