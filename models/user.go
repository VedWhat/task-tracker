package models

import (
	"task-tracker/src/config"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}

func GetUser(user *User, id string) error {
	if err := config.DB.First(&user, id).Error; err != nil {
		return err
	} else {
		return nil
	}
}

func CreateUser(user *User) error {
	if err := config.DB.Create(&user).Error; err != nil {
		return err
	} else {
		return nil
	}
}

func UpdateUser(user *User, id string) error {
	var temp User
	if err := config.DB.First(&temp, id).Error; err != nil {
		return err
	} else {
		config.DB.Save(&user)
		return nil
	}
}

func DeleteUser(user *User, id string) error {

	if err := config.DB.Delete(&user, id).Error; err != nil {
		return err
	} else {
		return nil
	}
}

func GetUsers(users *[]User, filter map[string]interface{}) error {
	if err := config.DB.Where(filter).Find(&users).Error; err != nil {
		return err
	} else {
		return nil
	}
}
