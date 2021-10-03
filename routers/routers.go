package routers

import (
	"task-tracker/src/handlers"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/api/v1/users", handlers.GetUsers) //Filters
	r.GET("/api/v1/users/:id", handlers.GetUser)
	r.POST("/api/v1/users", handlers.CreateUser)
	r.PUT("/api/v1/users/:id", handlers.UpdateUser)
	r.DELETE("/api/v1/users/:id", handlers.DeleteUser)

	r.GET("/api/v1/tasks", handlers.GetTasks) //Filters
	r.GET("/api/v1/tasks/:id", handlers.GetTask)
	r.POST("/api/v1/tasks", handlers.CreateTask)
	r.PUT("/api/v1/tasks/:id", handlers.UpdateTask)
	r.DELETE("/api/v1/tasks/:id", handlers.DeleteTask)

	return r
}
