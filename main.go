package main

import (
	"strconv"
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
		"users": users,
	})
}

// GetUserController get a user by given user ID
func GetUserController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"message": "failed to get a user, user with ID " + c.Param("id") + " is not found",
		})
	}

	for _, user := range users {
		if user.ID == id {
			return c.JSON(http.StatusOK, map[string]interface{}{
				"message"	: "success to get user data by given ID",
				"user"		: user,
			})
		}
	}

	return c.JSON(http.StatusBadRequest, map[string]interface{}{
		"message": "user with ID " + c.Param("id") + " is not found.",
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
	e.GET("/users/:id", GetUserController)
	e.POST("/users", CreateUserController)

	e.Logger.Fatal(e.Start(":8000"))
}