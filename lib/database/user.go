package database

import (
	"echo-user/config"
	"echo-user/models"
)

// GetUsers return all users data
func GetUsers() (interface{}, error) {
	var users []models.User

	if err := config.DB.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

// GetUser return user data by given user id
func GetUser(id int) (interface{}, error) {
	var user models.User

	if err := config.DB.First(&user, id).Error; err != nil {
		return nil, err
	}
	return user, nil
}

// DeleteUser delete user data from database by given user id
func DeleteUser(id int) (interface{}, error) {
	var user models.User
	if err := config.DB.First(&user, id).Error; err != nil {
		return nil, err
	}

	if deleteErr := config.DB.Delete(&user).Error; deleteErr != nil {
		return nil, deleteErr
	}
	return nil, nil
}

// UpdateUser update user data by giving new user data and existing user id
func UpdateUser(id int, user *models.User) (interface{}, error) {
	var existingUser models.User
	if err := config.DB.First(&existingUser, id).Error; err != nil {
		return nil, err
	}

	existingUser.Name = user.Name
	existingUser.Email = user.Email
	if updateErr := config.DB.Save(&existingUser).Error; updateErr != nil {
		return nil, updateErr
	}
	
	return existingUser, nil
}

// CreateUser store new user data
func CreateUser(user *models.User) (interface{}, error) {
	if err := config.DB.Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}