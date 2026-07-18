package transport

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	models2 "github.com/ukique/taxi-service/internal/models"
)

type UserRepository interface {
	SaveRefreshToken(ctx context.Context, token models2.RefreshToken) error
	SearchRefreshToken(ctx context.Context, clientToken string) (models2.RefreshToken, error)
	RegisterUser(ctx context.Context, username, password, email string) error
}
type UserService interface {
	VerifyUserCredentials(ctx context.Context, user models2.User) bool
	RefreshTokenService(ctx context.Context, clientToken string) (string, error)
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
