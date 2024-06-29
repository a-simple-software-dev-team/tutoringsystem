package services

import (
	"tutoringsystem/models"
	"tutoringsystem/utils"
)

// CreateMessage creates a new message in the database
func CreateMessage(senderID, receiverID, content string) (*models.Message, error) {
	msg := &models.Message{
		SenderID:   senderID,
		ReceiverID: receiverID,
		Content:    content,
	}
	result := utils.DB.Create(msg)
	return msg, result.Error
}

// GetMessagesByUser retrieves messagres for a given user (either sent or received)
func GetMessagesByUser(userID string) ([]models.Message, error) {
	var messages []models.Message
	result := utils.DB.Where("sender_id = ? OR receiver_id = ?", userID, userID).Find(&messages)
	return messages, result.Error
}
