package main

import (
	"task-tracker/src/config"
	"task-tracker/src/models"
	"task-tracker/src/routers"
)

func main() {
	config.InitMigration()
	config.DB.AutoMigrate(&models.User{})
	config.DB.AutoMigrate(&models.Task{})

	router := routers.InitRouter()
	router.Run(":9999")
}
