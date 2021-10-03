package handlers

import (
	"net/http"
	"task-tracker/src/helper"
	"task-tracker/src/models"

	"github.com/gin-gonic/gin"
)

func GetTask(c *gin.Context) {
	id := c.Param("id")
	var task models.Task
	err := models.GetTask(&task, id)
	if err != nil {
		c.JSON(http.StatusNotFound, "Id Not found")
		return

	} else {
		c.JSON(http.StatusOK, task)
	}
}

func GetTasks(c *gin.Context) {
	query := helper.HandleFilters(c.Request.URL.Query())
	var tasks []models.Task
	err := models.GetTasks(&tasks, query)
	if err != nil {
		c.JSON(http.StatusNotFound, "No tasks found")
	} else {
		c.JSON(http.StatusOK, tasks)
	}

}

func CreateTask(c *gin.Context) {
	var task models.Task
	c.ShouldBindJSON(&task)
	err := models.CreateTask(&task)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	} else {
		c.JSON(http.StatusOK, task)
	}
}

func UpdateTask(c *gin.Context) {
	var task models.Task
	id := c.Query("id")
	err := models.GetTask(&task, id)
	if err != nil {
		c.JSON(http.StatusNotFound, err)
		return

	}
	c.BindJSON(&task)
	err = models.UpdateTask(&task, id)
	if err != nil {
		c.JSON(http.StatusNotFound, err)
		return

	} else {
		c.JSON(http.StatusOK, task)
	}
}

func DeleteTask(c *gin.Context) {
	var task models.Task
	id := c.Param("id")
	err := models.GetTask(&task, id)
	if err != nil {
		c.JSON(http.StatusNotFound, err)
		return

	}

	err = models.DeleteTask(&task, id)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return

	} else {
		c.JSON(http.StatusOK, task)
	}
}
