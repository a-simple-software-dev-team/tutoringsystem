package services

import (
	"tutoringsystem/models"
	"tutoringsystem/utils"
)

func GetNotificationsByUser(userID uint) ([]models.Notification, error) {
	var notifications []models.Notification
	if err := utils.DB.Where("user_id = ?", userID).Find(&notifications).Error; err != nil {
		return notifications, err
	}
	return notifications, nil
}
