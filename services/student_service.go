package services

import (
	"strings"
	"tutoringsystem/models"
	"tutoringsystem/utils"
)

type StudentInput struct {
	SubjectsNeeded []string `json:"subjects_needed"`
	Grades         []string `json:"grades"`
	AvailableTime  []struct {
		Day       string `json:"day"`
		TimeRange string `json:"time_range"`
	} `json:"available_time"`
}

func GetStudentByID(id uint) (models.Student, error) {
	var student models.Student
	if err := utils.DB.Preload("AvailableTimes").First(&student, id).Error; err != nil {
		return student, err
	}
	return student, nil
}

func CreateOrUpdateStudent(userID uint, input StudentInput) (models.Student, error) {
	var student models.Student
	if err := utils.DB.Where("user_id = ?", userID).First(&student).Error; err != nil {
		if err.Error() == "record not found" {
			student.UserID = userID
		} else {
			return student, err
		}
	}

	student.SubjectsNeeded = strings.Join(input.SubjectsNeeded, ",")
	student.Grades = strings.Join(input.Grades, ",")

	if err := utils.DB.Save(&student).Error; err != nil {
		return student, err
	}
	utils.DB.Model(&student).Association("AvailableTimes").Clear()
	for _, at := range input.AvailableTime {
		availableTime := models.AvailableTime{
			Day:       at.Day,
			TimeRange: at.TimeRange,
			TutorID:   student.ID,
		}
		utils.DB.Create(&availableTime)
	}

	utils.DB.Preload("AvailableTimes").Find(&student)
	return student, nil
}
