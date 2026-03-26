package service

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5"
	"github.com/ukique/taxi-service/internal/features/user/repository"
	"golang.org/x/crypto/bcrypt"
)

func VerifyUserCredentials(ctx context.Context, conn *pgx.Conn, email, userName, password string) bool {
	userData, err := repository.GetDataByEmail(ctx, conn, email)
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
