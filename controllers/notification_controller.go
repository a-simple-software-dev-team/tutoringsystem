package controllers

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"tutoringsystem/models"
	"tutoringsystem/services"
)

// GetNotifications handles the request to retrieve notifications for a user.
func GetNotifications(c *gin.Context) {
	user := c.MustGet("user").(models.User)

	notifications, err := services.GetNotificationsByUser(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	emailErr := services.SendNotificationsByEmail(user.ID, notifications)
	if emailErr != nil {
		// Log the email sending error, if necessary
		log.Println("Failed to send email:", emailErr)
	}

	c.JSON(http.StatusOK, gin.H{"data": notifications})
}
