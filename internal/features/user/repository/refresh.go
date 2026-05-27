package repository

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/ukique/taxi-service/internal/models"
)

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
	SELECT username, refresh_token, created_at, expires_at FROM refresh_tokens WHERE refresh_token = $1;
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
