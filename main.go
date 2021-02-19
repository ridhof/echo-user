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

var users = map[int]User{}
var idPointer int //unexported

func mapToSlice(usersMap map[int]User) (slice []User) {
	slice = []User{}
	for _, user := range users {
		slice = append(slice, user)
	}

	return
}

// GetUsersController get all users
func GetUsersController(c echo.Context) error {
	usersSlice := mapToSlice(users)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success get all users",
		"users": usersSlice,
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

	user, isExist := users[id]
	if isExist {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"message"	: "success to get user data by given ID",
			"user"		: user,
		})
	} 

	return c.JSON(http.StatusBadRequest, map[string]interface{}{
		"message": "user with ID " + c.Param("id") + " is not found.",
	})
}

// DeleteUserController delete a user by given user ID
func DeleteUserController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"message": "failed to get a user, user with ID " + c.Param("id") + " is not found",
		})
	}

	user, isExist := users[id]
	if isExist {
		if user.ID == idPointer {
			idPointer--
		}
		delete(users, user.ID)

		usersSlice := mapToSlice(users)
		return c.JSON(http.StatusOK, map[string]interface{}{
			"message"	: "success delete a user",
			"users"		:	usersSlice,
		})
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

	user.ID = idPointer + 1
	idPointer++

	users[user.ID] = user
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
	e.DELETE("/users/:id", DeleteUserController)

	e.Logger.Fatal(e.Start(":8000"))
}