package service

import (
	"context"
	"time"

	"github.com/ukique/taxi-service/config"
	"github.com/ukique/taxi-service/internal/middleware"
	"github.com/ukique/taxi-service/internal/models"
)

type UserRepository interface {
	GetDataByEmail(ctx context.Context, email string) (models.User, error)
	SearchRefreshToken(ctx context.Context, clientToken string) (models.RefreshToken, error)
}

type UserService struct {
	userRepository UserRepository
	secretKey      string
}

func NewUserService(userRepository UserRepository, secretKey string) *UserService {
	return &UserService{
		userRepository: userRepository,
		secretKey:      secretKey}
}

func (u *UserService) RefreshTokenService(ctx context.Context, clientToken string) (string, error) {
	refreshToken, err := u.userRepository.SearchRefreshToken(ctx, clientToken)
	if err != nil {
		return "", config.ErrInvalidRefreshToken
	}
	// checking refresh token expiring
	if time.Now().After(refreshToken.ExpiresAt) {
		return "", config.ErrInvalidRefreshToken
	}
	accessToken, err := middleware.GenerateJWT(u.secretKey, refreshToken.UserName)
	if err != nil {
		return "", err
	}
	return accessToken, nil
}
