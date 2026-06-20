package transport

import (
	"context"

	"github.com/ukique/taxi-service/metts-taxi/internal/core/jwt"
	"github.com/ukique/taxi-service/metts-taxi/models"
)

type AuthRepository interface {
	RegisterUser(ctx context.Context, user models.User) error
	GetUserCredentials(ctx context.Context, username string) (models.User, error)
}
type AuthService interface {
	Login(ctx context.Context, user models.User) (jwt.TokenPair, error)
}
type Handler struct {
	repository AuthRepository
	service    AuthService
}

func NewAuthHandler(repository AuthRepository, service AuthService) *Handler {
	return &Handler{
		repository: repository,
		service:    service,
	}
}
