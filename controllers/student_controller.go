package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"tutoringsystem/models"
	"tutoringsystem/services"
)

func GetStudent(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	student, err := services.GetStudentByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": student})
}

func CreateOrUpdateStudent(c *gin.Context) {
	var input services.StudentInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//fmt.Println("%+v", c)
	user := c.MustGet("user").(models.User)
	student, err := services.CreateOrUpdateStudent(user.ID, input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": student})
}
