package main

import (
	"net/http"

	"github.com/labstack/echo"
)

// User struct contains User object with attribute 
// such as ID, Name, Email, and Password
type User struct {
	ID				int			`json:"id" form:"id"`
	Name			string	`json:"name" form:"name"`
	Email			string	`json:"email" form:"email"`
	Password	string	`json:"password" form:"password"`
}

var users []User

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