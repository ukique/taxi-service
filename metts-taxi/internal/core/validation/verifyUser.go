package validation

import (
	"context"
	"errors"

	"github.com/ukique/taxi-service/metts-taxi/models"
	"golang.org/x/crypto/bcrypt"
)

type UserRepository interface {
	GetUserCredentials(ctx context.Context, username string) (models.User, error)
}

var ErrInvalidCredentials = errors.New("invalid credentials")

func VerifyUser(ctx context.Context, userRepository UserRepository, user models.User) error {
	dbUser, err := userRepository.GetUserCredentials(ctx, user.Username)
	if err != nil {
		return ErrInvalidCredentials
	}

	if dbUser.Email != user.Email {
		return ErrInvalidCredentials
	}

	err = bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(user.Password))
	if err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			return ErrInvalidCredentials
		}
		return err
	}

	return nil
}
