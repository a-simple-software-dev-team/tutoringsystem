package services

import (
	"errors"
	"tutoringsystem/models"
	"tutoringsystem/utils"
)

type ScheduleInput struct {
	Date      string `json:"date"`
	Time      string `json:"time"`
	StudentID uint   `json:"student_id"`
}

func GetSchedulesByUser(userID uint, role string) ([]models.Schedule, error) {
	var schedules []models.Schedule
	query := utils.DB
	if role == "tutor" {
		query = query.Where("tutor_id = ?", userID)
	} else {
		query = query.Where("student_id = ?", userID)
	}

	if err := query.Find(&schedules).Error; err != nil {
		return schedules, err
	}
	return schedules, nil
}

func UpdateSchedules(userID uint, role string, input []ScheduleInput) error {
	if role != "tutor" {
		return errors.New("only tutors can update schedules")
	}

	var schedules []models.Schedule
	for _, schedInput := range input {
		schedule := models.Schedule{
			TutorID:   userID,
			StudentID: schedInput.StudentID,
			Date:      schedInput.Date,
			Time:      schedInput.Time,
		}
		schedules = append(schedules, schedule)
	}

	tx := utils.DB.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	if err := tx.Where("tutor_id = ?", userID).Delete(&models.Schedule{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Create(&schedules).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}
