package handlers

import (
	"net/http"
	"vanminton-be/api/models"

	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	users := models.User{
		// Code to retrieve resource from database or other source
	}
	c.JSON(http.StatusOK, resource)
}
