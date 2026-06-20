package service

import (
	"context"

	"github.com/ukique/taxi-service/metts-taxi/config"
	"github.com/ukique/taxi-service/metts-taxi/internal/core/jwt"
	"github.com/ukique/taxi-service/metts-taxi/internal/core/validation"
	"github.com/ukique/taxi-service/metts-taxi/models"
	"golang.org/x/crypto/bcrypt"
)

type AuthRepository interface {
	RegisterUser(ctx context.Context, user models.User) error
	GetUserCredentials(ctx context.Context, username string) (models.User, error)
	SearchRefreshToken(ctx context.Context, username string) (models.RefreshToken, error)
}
type Service struct {
	repository AuthRepository
	jwtMaker   *jwt.TokenMaker
}

func NewAuthService(repository AuthRepository, jwtMaker *jwt.TokenMaker) *Service {
	return &Service{
		repository: repository,
		jwtMaker:   jwtMaker,
	}
}

func (s *Service) Login(ctx context.Context, user models.User) (jwt.TokenPair, error) {
	err := validation.VerifyUser(ctx, s.repository, user)
	if err != nil {
		return jwt.TokenPair{}, err
	}

	tokenPair, err := s.jwtMaker.CreateTokenPair(ctx, user.Username, config.TokenDuration)
	if err != nil {
		return jwt.TokenPair{}, err
	}

	return tokenPair, nil
}

func (s *Service) RefreshToken(ctx context.Context, userData models.RefreshToken) (string, error) {
	dbData, err := s.repository.SearchRefreshToken(ctx, userData.Username)
	if err != nil {
		return "", err
	}
	err = bcrypt.CompareHashAndPassword([]byte(dbData.RefreshToken), []byte(userData.RefreshToken))
	if err != nil {
		return "", err
	}

	accessToken, err := s.jwtMaker.GenerateJWT(userData.Username)
	if err != nil {
		return "", err
	}
	return accessToken, nil
}
