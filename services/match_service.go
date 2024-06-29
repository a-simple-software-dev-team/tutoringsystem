package services

import (
	"errors"
	"math/rand"
	"time"
	"tutoringsystem/models"
	"tutoringsystem/utils"
)

func GetMatchesByUser(userID uint) ([]models.Match, error) {
	println("jdiajsdkjfwkejdksjf")
	var matches []models.Match
	if err := utils.DB.Where("user_id = ?", userID).Or("match_id = ?", userID).Find(&matches).Error; err != nil {
		return matches, err
	}
	return matches, nil
}

// MatchStudentsAndTutors is a simplified example. You'd want a more sophisticated matching algorithm in production.
func MatchStudentsAndTutors(userID uint, role string) error {
	var users []models.User
	if role == "tutor" {
		if err := utils.DB.Where("role = ?", "student").Find(&users).Error; err != nil {
			return err
		}
	} else if role == "student" {
		if err := utils.DB.Where("role = ?", "tutor").Find(&users).Error; err != nil {
			return err
		}
	} else {
		return errors.New("invalid role")
	}

	var matches []models.Match
	rand.Seed(time.Now().Unix())
	for _, user := range users {
		match := models.Match{
			UserID:             userID,
			MatchID:            user.ID,
			CompatibilityScore: rand.Float64() * 100, // For demonstration purposes
			Role:               role,
		}
		matches = append(matches, match)
	}

	if err := utils.DB.Create(&matches).Error; err != nil {
		return err
	}

	return nil
}
