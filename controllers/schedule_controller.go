package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"tutoringsystem/models"
	"tutoringsystem/services"
)

func GetSchedules(c *gin.Context) {
	user := c.MustGet("user").(models.User)

	schedules, err := services.GetSchedulesByUser(user.ID, user.Role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": schedules})
}

func UpdateSchedules(c *gin.Context) {
	var input []services.ScheduleInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := c.MustGet("user").(models.User)
	if err := services.UpdateSchedules(user.ID, user.Role, input); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Schedules updated successfully"})
}
