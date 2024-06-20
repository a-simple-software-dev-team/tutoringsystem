package services

import (
	"errors"
	"fmt"
	"net/smtp"
	"tutoringsystem/models"
	"tutoringsystem/utils"
)

type Notification struct {
	ID      uint   `json:"id"`
	UserID  uint   `json:"user_id"`
	Content string `json:"content"`
}

func GetNotificationsByUser(userID uint) ([]models.Notification, error) {
	var notifications []models.Notification
	if err := utils.DB.Where("user_id = ?", userID).Find(&notifications).Error; err != nil {
		return notifications, err
	}
	return notifications, nil
}

func GetEmailByUserID(userID uint) (string, error) {
	var user models.User
	if err := utils.DB.First(&user, userID).Error; err != nil {
		return "", err
	}
	return user.Email, nil
}

func SendNotificationsByEmail(userID uint, notifications []models.Notification) error {
	if len(notifications) == 0 {
		return errors.New("no notifications to send")
	}

	email, err := GetEmailByUserID(userID)
	if err != nil {
		return err
	}

	// 配置SMTP服务器信息
	smtpHost := "smtp.example.com"
	smtpPort := "587"
	smtpUser := "your-email@example.com"
	smtpPass := "your-email-password"

	// 设置邮件内容
	subject := "Your Notifications"
	body := "You have the following notifications:\n\n"
	for _, n := range notifications {
		body += n.Message + "\n"
	}

	msg := fmt.Sprintf("From: %s\nTo: %s\nSubject: %s\n\n%s", smtpUser, email, subject, body)

	// 发送邮件
	auth := smtp.PlainAuth("", smtpUser, smtpPass, smtpHost)
	err = smtp.SendMail(smtpHost+":"+smtpPort, auth, smtpUser, []string{email}, []byte(msg))
	return err
}
