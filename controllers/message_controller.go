package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"tutoringsystem/models"
	"tutoringsystem/services"
)

func SendMessage(c *gin.Context) {
	var msg models.Message
	if err := c.ShouldBindJSON(&msg); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newMsg, err := services.CreateMessage(msg.SenderID, msg.ReceiverID, msg.Content)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, newMsg)
}

func GetMessages(c *gin.Context) {
	userID := c.Param("userid")
	msgs, err := services.GetMessagesByUser(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, msgs)
}
