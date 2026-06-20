package repository

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/ukique/taxi-service/metts-taxi/models"
)

type Repository struct {
	pool *pgxpool.Pool
}

func NewAuthRepository(pool *pgxpool.Pool) *Repository {
	return &Repository{pool: pool}
}

func (r *Repository) RegisterUser(ctx context.Context, user models.User) error {
	sqlQuery := `
	INSERT INTO users(username, password, email, created_at)
	VALUES ($1, $2, $3, $4)
`
	_, err := r.pool.Exec(ctx, sqlQuery,
		user.Username, user.Password, user.Email, time.Now())
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) GetUserCredentials(ctx context.Context, username string) (models.User, error) {
	var user models.User

	sqlQuery := `
	SELECT password, email FROM users
	WHERE username = $1;
`
	err := r.pool.QueryRow(ctx, sqlQuery, username).Scan(&user.Password, &user.Email)
	if err != nil {
		return models.User{}, err
	}
	user.Username = username

	return user, nil
}

func (r *Repository) SaveRefreshToken(ctx context.Context, refreshToken models.RefreshToken) error {
	sqlQuery := `
	INSERT INTO refresh_tokens (username, refresh_token, created_at, expires_at)
	VALUES ($1,$2,$3,$4)
`
	_, err := r.pool.Exec(ctx, sqlQuery, refreshToken.Username, refreshToken.RefreshToken,
		refreshToken.CreatedAt, refreshToken.ExpiresAt)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) SearchRefreshToken(ctx context.Context, username string) (models.RefreshToken, error) {
	sqlQuery := `
	SELECT refresh_token, created_at, expires_at
	FROM refresh_tokens
	WHERE username = $1;
`
	var refreshToken models.RefreshToken
	err := r.pool.QueryRow(ctx, sqlQuery, username).Scan(
		&refreshToken.RefreshToken,
		&refreshToken.CreatedAt,
		&refreshToken.ExpiresAt)
	if err != nil {
		return models.RefreshToken{}, err
	}
	return refreshToken, nil
}
