package main

import (
	"log"
	"os"
	"task-tracker/src/config"
	"task-tracker/src/models"
	"task-tracker/src/routers"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT not set defaulting to 9999")
		port = "9999"
	}
	config.InitMigration()
	config.DB.AutoMigrate(&models.User{})
	config.DB.AutoMigrate(&models.Task{})

	router := routers.InitRouter()
	router.Run(":" + port)
}
