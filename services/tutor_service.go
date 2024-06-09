package services

import (
	"strings"
	"tutoringsystem/models"
	"tutoringsystem/utils"
)

type TutorInput struct {
	Subjects      []string `json:"subjects"`
	Grades        []string `json:"grades"`
	AvailableTime []struct {
		Day       string `json:"day"`
		TimeRange string `json:"time_range"`
	} `json:"available_time"`
}

func GetTutorByID(id uint) (models.Tutor, error) {
	var tutor models.Tutor
	if err := utils.DB.Preload("AvailableTimes").First(&tutor, id).Error; err != nil {
		return tutor, err
	}
	return tutor, nil
}

func CreateOrUpdateTutor(userID uint, input TutorInput) (models.Tutor, error) {
	var tutor models.Tutor
	if err := utils.DB.Where("user_id = ?", userID).First(&tutor).Error; err != nil {
		if err.Error() == "record not found" {
			tutor.UserID = userID
		} else {
			return tutor, err
		}
	}

	tutor.Subjects = strings.Join(input.Subjects, ",")
	tutor.Grades = strings.Join(input.Grades, ",")

	if err := utils.DB.Save(&tutor).Error; err != nil {
		return tutor, err
	}

	utils.DB.Model(&tutor).Association("AvailableTimes").Clear()
	for _, at := range input.AvailableTime {
		availableTime := models.AvailableTime{
			Day:       at.Day,
			TimeRange: at.TimeRange,
			TutorID:   tutor.ID,
		}
		utils.DB.Create(&availableTime)
	}

	utils.DB.Preload("AvailableTimes").Find(&tutor)
	return tutor, nil
}
