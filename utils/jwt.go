package utils

import (
	"errors"
	"fmt"
	"os"
	"time"
	"tutoringsystem/models"

	"github.com/golang-jwt/jwt/v4"
)

var secretKey = []byte(os.Getenv("API_SECRET"))

type CustomClaims struct {
	UserID uint `json:"user_id"`
	jwt.StandardClaims
}

// GenerateToken 生成JWT Token
func GenerateToken(user models.User) (string, error) {
	claims := CustomClaims{
		UserID: user.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secretKey)
}

// ValidateToken 验证JWT Token
func ValidateToken(tokenString string) (models.User, error) {
	var user models.User

	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
	fmt.Println(token)

	if err != nil {
		return user, err
	}

	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		// 通过用户ID从数据库获取用户信息
		if err := DB.First(&user, claims.UserID).Error; err != nil {
			return user, err
		}
		return user, nil
	} else {
		return user, errors.New("invalid token")
	}
}
