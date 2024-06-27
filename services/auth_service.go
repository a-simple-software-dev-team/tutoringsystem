package services

import (
	"errors"
	"regexp"
	"tutoringsystem/models"
	"tutoringsystem/utils"

	"github.com/jinzhu/gorm"
)

type RegisterInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Role     string `json:"role" binding:"required"`
}

type LoginInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func IsValidEmail(email string) bool {
	// 定义Email的正则表达式
	const emailRegexPattern = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`

	// 编译正则表达式
	regex := regexp.MustCompile(emailRegexPattern)

	// 使用正则表达式匹配Email地址
	return regex.MatchString(email)
}

func RegisterUser(input RegisterInput) (models.User, error) {
	var user models.User
	hashedPassword, err := utils.HashPassword(input.Password)
	if err != nil {
		return user, err
	}

	if IsValidEmail(input.Email) == false {
		return user, errors.New("Email is invalid")
	}

	user = models.User{
		Username: input.Username,
		Password: hashedPassword,
		Email:    input.Email,
		Role:     input.Role,
	}

	if err := utils.DB.Create(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}

func LoginUser(input LoginInput) (string, error) {
	var user models.User
	if err := utils.DB.Where("username = ?", input.Username).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "", errors.New("user not found")
		}
		return "", err
	}

	if !utils.CheckPassword(user.Password, input.Password) {
		return "", errors.New("incorrect password")
	}

	token, err := utils.GenerateToken(user)
	if err != nil {
		return "", err
	}

	return token, nil
}
