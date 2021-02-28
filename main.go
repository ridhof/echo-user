package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
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
func GetUsersController(c *gin.Context) {
	usersSlice := mapToSlice(users)

	c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success get all users",
		"users": usersSlice,
	})
}

// GetUserController get a user by given user ID
func GetUserController(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusNotFound, map[string]interface{}{
			"message": "failed to get a user, user with ID " + c.Param("id") + " is not found",
		})
		return
	}

	user, isExist := users[id]
	if isExist {
		c.JSON(http.StatusOK, map[string]interface{}{
			"message"	: "success to get user data by given ID",
			"user"		: user,
		})
		return
	} 

	c.JSON(http.StatusBadRequest, map[string]interface{}{
		"message": "user with ID " + c.Param("id") + " is not found.",
	})
}

// DeleteUserController delete a user by given user ID
func DeleteUserController(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusNotFound, map[string]interface{}{
			"message": "failed to get a user, user with ID " + c.Param("id") + " is not found",
		})
		return
	}

	user, isExist := users[id]
	if isExist {
		if user.ID == idPointer {
			idPointer--
		}
		delete(users, user.ID)

		usersSlice := mapToSlice(users)
		c.JSON(http.StatusOK, map[string]interface{}{
			"message"	: "success delete a user",
			"users"		:	usersSlice,
		})
		return
	}

	c.JSON(http.StatusBadRequest, map[string]interface{}{
		"message": "user with ID " + c.Param("id") + " is not found.",
	})
}

// UpdateUserController update a user by given user ID and its form data
func UpdateUserController(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusNotFound, map[string]interface{}{
			"message": "failed to get a user, user with ID " + c.Param("id") + " is not found",
		})
		return
	}

	user, isExist := users[id]
	if isExist {
		newUser := User{}
		c.Bind(&newUser)
		newUser.ID = user.ID

		users[newUser.ID] = newUser
		c.JSON(http.StatusOK, map[string]interface{}{
			"message"	: "success update a user",
			"user"		: users[newUser.ID],
		})
		return
	}

	c.JSON(http.StatusBadRequest, map[string]interface{}{
		"message": "user with ID " + c.Param("id") + " is not found.",
	})
}

// CreateUserController create new user by given form data
func CreateUserController(c *gin.Context) {
	// binding data
	user := User{}
	c.Bind(&user)

	user.ID = idPointer + 1
	idPointer++

	users[user.ID] = user
	c.JSON(http.StatusOK, map[string]interface{}{
		"message"	:	"success create user",
		"user"		:	user,
	})
}

func main() {
	router := gin.New()
	router.Use(gin.Logger())

	router.GET("/users", GetUsersController)
	router.GET("/users/:id", GetUserController)
	router.POST("/users", CreateUserController)
	router.DELETE("/users/:id", DeleteUserController)
	router.PUT("/users/:id", UpdateUserController)

	// e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", os.Getenv("PORT"))))
	router.Run(fmt.Sprintf(":%s", os.Getenv("PORT")))
}