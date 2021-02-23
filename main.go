package main

import (
	"fmt"
	"strconv"
	"net/http"

	"echo-user/config"
	"echo-user/models"

	"github.com/labstack/echo"
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
)

// DB shared gorm.DB object accross the code to use
var DB *gorm.DB

// GetUsersController get all users
func GetUsersController(c echo.Context) error {
	var usersDB []models.User
	if err := config.DB.Find(&usersDB).Error; err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success get all users",
		"users": usersDB,
	})
}

// GetUserController get a user by given user ID
func GetUserController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "failed to get a user, user with ID " + c.Param("id") + " is not found",
		})
	}

	var user models.User
	if err := config.DB.First(&user, id).Error; err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message"	: "success to get user data by given ID",
		"user"		: user,
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

	var user models.User
	if err := config.DB.First(&user, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"message": err.Error(),
		})
	}

	if err := config.DB.Delete(&models.User{}, user.ID).Error; err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete a user",
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

	var user models.User
	if err := config.DB.First(&user, id).Error; err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
		})
	}

	var updateUser models.User
	c.Bind(&updateUser)

	user.Name = updateUser.Name
	user.Email = updateUser.Email
	if err := config.DB.Save(&user).Error; err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update a user",
		"user" : user,
	})
}

// CreateUserController create new user by given form data
func CreateUserController(c echo.Context) error {
	// binding data
	user := models.User{}
	c.Bind(&user)

	if err := config.DB.Create(&user).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message"	:	"success create user",
		"user"		:	user,
	})
}

// InitDB to initialize database connection
func InitDB() {
	dbUsername := "root"
	dbPassword := "mysql"
	dbAddress := "127.0.0.1"
	dbPort := "3306"
	dbName := "echo_user"

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?parseTime=True",
		dbUsername,
		dbPassword,
		dbAddress,
		dbPort,
		dbName,
	)

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}
}

func main() {
	config.InitDB()

	e := echo.New()
	e.GET("/users", GetUsersController)
	e.GET("/users/:id", GetUserController)
	e.POST("/users", CreateUserController)
	e.DELETE("/users/:id", DeleteUserController)
	e.PUT("/users/:id", UpdateUserController)

	e.Logger.Fatal(e.Start(":8000"))
}