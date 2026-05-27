package service

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/ukique/taxi-service/internal/models"
)

type UserRepository interface {
	GetDataByEmail(ctx context.Context, email string) (models.User, error)
}

type UserService struct {
	pool           *pgxpool.Pool
	userRepository UserRepository
}

func NewUserService(pool *pgxpool.Pool, userRepository UserRepository) *UserService {
	return &UserService{pool: pool, userRepository: userRepository}
}
