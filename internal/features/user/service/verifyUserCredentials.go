package service

import (
	"context"
	"log"

	"github.com/ukique/taxi-service/internal/models"
	"golang.org/x/crypto/bcrypt"
)

func (u *UserService) VerifyUserCredentials(ctx context.Context, user models.User) bool {
	userData, err := u.userRepository.GetDataByUsername(ctx, user.Username)
	if err != nil {
		log.Println("fail to get data by email:", err)
		return false
	}
	if userData.Username != user.Username {
		log.Println("incorrect username:", userData.Username)
		return false
	}
	err = bcrypt.CompareHashAndPassword([]byte(userData.Password), []byte(user.Password))
	if err != nil {
		log.Println("incorrect password:", err)
		return false
	}
	return true
}
