package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"tutoringsystem/models"
	"tutoringsystem/services"
)

func GetNotifications(c *gin.Context) {
	user := c.MustGet("user").(models.User)

	notifications, err := services.GetNotificationsByUser(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": notifications})
}