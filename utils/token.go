package utils

//
//import (
//	"fmt"
//	"github.com/dgrijalva/jwt-go"
//	"github.com/gin-gonic/gin"
//	"os"
//	"strconv"
//	"time"
//	"tutoringsystem/models"
//)
//
//func GenerateToken(user models.User) (string, error) {
//	tokenLifespan, err := strconv.Atoi(os.Getenv("TOKEN_HOUR_LIFESPAN"))
//	if err != nil {
//		return "", err
//	}
//
//	claims := jwt.MapClaims{}
//	claims["authorized"] = true
//	claims["user_id"] = user.ID
//	claims["exp"] = time.Now().Add(time.Hour * time.Duration(tokenLifespan)).Unix()
//
//	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
//	return token.SignedString([]byte(os.Getenv("API_SECRET")))
//}
//
//func ValidateToken(signedToken string) (*jwt.Token, error) {
//	token, err := jwt.Parse(signedToken, func(token *jwt.Token) (interface{}, error) {
//		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
//			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
//		}
//		return []byte(os.Getenv("API_SECRET")), nil
//	})
//
//	if err != nil {
//		return nil, err
//	}
//
//	return token, nil
//}
//
//func ExtractTokenUserID(c *gin.Context) uint {
//	user, exists := c.Get("user")
//	if !exists {
//		return 0
//	}
//	userId := user.(*jwt.Token).Claims.(jwt.MapClaims)["user_id"].(uint)
//	return userId
//}
