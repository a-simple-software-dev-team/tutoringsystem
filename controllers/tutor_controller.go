package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"tutoringsystem/models"
	"tutoringsystem/services"
)

func GetTutor(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	tutor, err := services.GetTutorByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": tutor})
}

func CreateOrUpdateTutor(c *gin.Context) {
	var input services.TutorInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := c.MustGet("user").(models.User)
	tutor, err := services.CreateOrUpdateTutor(user.ID, input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": tutor})
}
