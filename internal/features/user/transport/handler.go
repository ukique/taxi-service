package transport

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/ukique/taxi-service/internal/models"
)

type UserRepository interface {
	SaveRefreshToken(ctx context.Context, token models.RefreshToken) error
	SearchRefreshToken(ctx context.Context, clientToken string) (models.RefreshToken, error)
	RegisterUser(ctx context.Context, username, password, email string) error
}
type UserService interface {
	VerifyUserCredentials(ctx context.Context, email, userName, password string) bool
}
type Handler struct {
	pool           *pgxpool.Pool
	secretKey      string
	userRepository UserRepository
	userService    UserService
}

func NewUserHandler(pool *pgxpool.Pool, secretKey string, userRepository UserRepository, userService UserService) *Handler {
	return &Handler{
		pool:           pool,
		secretKey:      secretKey,
		userRepository: userRepository,
		userService:    userService,
	}
}
