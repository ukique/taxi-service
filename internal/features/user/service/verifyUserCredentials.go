package service

import (
	"context"
	"log"

	"golang.org/x/crypto/bcrypt"
)

func (u *UserService) VerifyUserCredentials(ctx context.Context, email, userName, password string) bool {
	userData, err := u.userRepository.GetDataByEmail(ctx, email)
	if err != nil {
		log.Println("fail to get data by email:", err)
		return false
	}
	if userData.Username != userName {
		log.Println("incorrect username:", userData.Username)
		return false
	}
	err = bcrypt.CompareHashAndPassword([]byte(userData.Password), []byte(password))
	if err != nil {
		log.Println("incorrect password:", err)
		return false
	}
	return true
}
