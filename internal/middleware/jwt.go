package middleware

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateJWT(secretKey string, userName string) (string, error) {
	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		jwt.MapClaims{
			"username": userName,
			"exp":      time.Now().Add(time.Minute * 5).Unix(),
		})
	jwtToken, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", fmt.Errorf("fail to sign key:%w", err)
	}
	return jwtToken, err
}

func GenerateRefreshToken(length int) (string, error) {
	bytes := make([]byte, length)
	if _, err := rand.Read(bytes); err != nil {
		return "", nil
	}
	return hex.EncodeToString(bytes), nil
}

func VerifyJWT(secretKey string, userToken string) (*jwt.Token, error) {
	token, err := jwt.Parse(userToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("invalid signing method: %v", token.Header["alg"])
		}
		return []byte(secretKey), nil
	})
	if err != nil || !token.Valid {
		return token, err
	}
	return token, nil
}
