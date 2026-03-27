package middleware

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/ukique/taxi-service/internal/models"
)

func GenerateJWT(user models.User, secretKey string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"username": user.Username,
			"exp":      time.Now().Add(time.Hour * 48).Unix(),
		})

	signedKey, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", fmt.Errorf("fail to sign key:%w", err)
	}
	return signedKey, nil
}
