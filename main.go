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

// CreateUserController create new user by given form data
func CreateUserController(c echo.Context) error {
	// binding data
	user := User{}
	c.Bind(&user)

	if len(users) == 0 {
		user.ID = 1
	} else {
		newID := users[len(users) - 1].ID + 1
		user.ID = newID
	}

	users = append(users, user)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message"	:	"success create user",
		"user"		:	user,
	})
}

func main() {
	e := echo.New()
	e.GET("/users", GetUsersController)
	e.POST("/users", CreateUserController)

	e.Logger.Fatal(e.Start(":8000"))
}