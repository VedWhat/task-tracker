package handlers

import (
	"net/http"
	"task-tracker/src/helper"
	"task-tracker/src/models"

	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {

	var users []models.User
	filters := helper.HandleFilters(c.Request.URL.Query())
	err := models.GetUsers(&users, filters)
	if err != nil {
		c.JSON(http.StatusNotFound, err)
	} else {
		c.JSON(http.StatusOK, users)
	}
}

func GetUser(c *gin.Context) {
	var user models.User
	id := c.Param("id")
	err := models.GetUser(&user, id)
	if err != nil {
		c.JSON(http.StatusNotFound, err)
	} else {
		c.JSON(http.StatusOK, user)
	}
}

func CreateUser(c *gin.Context) {
	var user models.User
	c.BindJSON(&user)
	err := models.CreateUser(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	} else {
		c.JSON(http.StatusOK, user)
	}
}

func UpdateUser(c *gin.Context) {
	var user models.User
	id := c.Query("id")
	err := models.GetUser(&user, id)
	if err != nil {
		c.JSON(http.StatusNotFound, err)
	}
	c.BindJSON(&user)
	err = models.UpdateUser(&user, id)
	if err != nil {
		c.JSON(http.StatusNotFound, err)
	} else {
		c.JSON(http.StatusOK, user)
	}
}

func DeleteUser(c *gin.Context) {
	var user models.User
	id := c.Param("id")
	err := models.GetUser(&user, id)
	if err != nil {
		c.JSON(http.StatusNotFound, err)
	}

	err = models.DeleteUser(&user, id)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	} else {
		c.JSON(http.StatusOK, user)
	}

}
