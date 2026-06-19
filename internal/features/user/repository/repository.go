package repository

import (
	"context"
	"errors"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/ukique/taxi-service/internal/models"
	"golang.org/x/crypto/bcrypt"
)

type UserRepository struct {
	pool *pgxpool.Pool
}

func NewUserRepository(pool *pgxpool.Pool) *UserRepository {
	return &UserRepository{
		pool: pool,
	}
}

func (u *UserRepository) RegisterUser(ctx context.Context, username, password, email string) error {
	// hash the password to bcrypt with 12 cost
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return err
	}

	sqlQuery := `
	INSERT INTO user(username, password, email, created_at)
	VALUES ($1, $2, $3, $4);
`
	if _, err := u.pool.Exec(ctx, sqlQuery, username, string(hashedPassword), email, time.Now()); err != nil {
		return err
	}
	return nil
}

func (u *UserRepository) GetDataByUsername(ctx context.Context, username string) (models.User, error) {
	sqlQuery := `
	SELECT username, password, email FROM users WHERE username = $1;
`
	var userData models.User
	err := u.pool.QueryRow(ctx, sqlQuery, username).Scan(&userData.Username, &userData.Password, &userData.Email)
	if err != nil {
		return models.User{}, err
	}
	return userData, nil
}

func (u *UserRepository) SaveRefreshToken(ctx context.Context, token models.RefreshToken) error {
	sqlQuery := `
	INSERT INTO refresh_tokens(username, refresh_token,created_at, expires_at)
	VALUES ($1,$2,$3,$4);
	`
	_, err := u.pool.Exec(ctx, sqlQuery, token.UserName, token.RefreshToken, token.CreatedAt, token.ExpiresAt)
	if err != nil {
		return err
	}

	return nil
}

func (u *UserRepository) SearchRefreshToken(ctx context.Context, clientToken string) (models.RefreshToken, error) {
	sqlQuery := `
	SELECT username, refresh_token, created_at, expires_at 
    FROM refresh_tokens 
    WHERE refresh_token = $1;
`
	var refreshToken models.RefreshToken
	row := u.pool.QueryRow(ctx, sqlQuery, clientToken)
	err := row.Scan(&refreshToken.UserName, &refreshToken.RefreshToken, &refreshToken.CreatedAt, &refreshToken.ExpiresAt)
	if errors.Is(err, pgx.ErrNoRows) {
		return refreshToken, err
	}

	if err != nil {
		return refreshToken, err
	}
	return refreshToken, nil
}
