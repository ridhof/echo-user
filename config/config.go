package config

import (
	"fmt"
	"echo-user/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// DB shared gorm.DB object accross the code to use
var DB *gorm.DB

// InitialMigration to initialize database migration
func InitialMigration() {
	DB.AutoMigrate(&models.User{})
}

// InitDB to initialize database connection
func InitDB() {
	dbUsername := "root"
	dbPassword := "mysql"
	dbHost := "127.0.0.1"
	dbPort := "3306"
	dbName := "echo_user"

	connectionString := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?parseTime=True",
		dbUsername,
		dbPassword,
		dbHost,
		dbPort,
		dbName,
	)

	var err error
	DB, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	InitialMigration()
}
