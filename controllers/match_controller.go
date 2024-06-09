package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"tutoringsystem/models"
	"tutoringsystem/services"
)

func GetMatches(c *gin.Context) {
	user := c.MustGet("user").(models.User)

	matches, err := services.GetMatchesByUser(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": matches})
}

func MatchStudentsAndTutors(c *gin.Context) {
	user := c.MustGet("user").(models.User)

	err := services.MatchStudentsAndTutors(user.ID, user.Role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Matching completed successfully"})
}
