package main

import (
	"strconv"
	"net/http"

	"github.com/labstack/echo"
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
)

// User struct contains User object with attribute 
// such as ID, Name, Email, and Password
type User struct {
	gorm.Model
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
		if user.ID == uint(idPointer) {
			idPointer--
		}
		delete(users, int(user.ID))

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

// UpdateUserController update a user by given user ID and its form data
func UpdateUserController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"message": "failed to get a user, user with ID " + c.Param("id") + " is not found",
		})
	}

	user, isExist := users[id]
	if isExist {
		newUser := User{}
		c.Bind(&newUser)
		newUser.ID = user.ID

		users[int(newUser.ID)] = newUser
		return c.JSON(http.StatusOK, map[string]interface{}{
			"message"	: "success update a user",
			"user"		: users[int(newUser.ID)],
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

	user.ID = uint(idPointer) + 1
	idPointer++

	users[int(user.ID)] = user
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message"	:	"success create user",
		"user"		:	user,
	})
}

func main() {
	dsn := "root:mysql@tcp(127.0.0.1:3306)/echo_user"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&User{})

	e := echo.New()
	e.GET("/users", GetUsersController)
	e.GET("/users/:id", GetUserController)
	e.POST("/users", CreateUserController)
	e.DELETE("/users/:id", DeleteUserController)
	e.PUT("/users/:id", UpdateUserController)

	e.Logger.Fatal(e.Start(":8000"))
}