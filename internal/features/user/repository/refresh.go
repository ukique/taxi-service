package repository

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/ukique/taxi-service/internal/models"
)

func SaveRefreshToken(ctx context.Context, pool *pgxpool.Pool, token models.RefreshToken) error {
	sqlQuery := `
	INSERT INTO refresh_tokens(username, refresh_token,created_at, expires_at)
	VALUES ($1,$2,$3,$4);
	`
	_, err := pool.Exec(ctx, sqlQuery, token.UserName, token.RefreshToken, token.CreatedAt, token.ExpiresAt)
	if err != nil {
		return err
	}

	return nil
}

func SearchRefreshToken(ctx context.Context, pool *pgxpool.Pool, clientToken string) (models.RefreshToken, error) {
	sqlQuery := `
	SELECT username, refresh_token, created_at, expires_at FROM refresh_tokens WHERE refresh_token = $1;
`
	var refreshToken models.RefreshToken
	row := pool.QueryRow(ctx, sqlQuery, clientToken)
	err := row.Scan(&refreshToken.UserName, &refreshToken.RefreshToken, &refreshToken.CreatedAt, &refreshToken.ExpiresAt)
	if errors.Is(err, pgx.ErrNoRows) {
		return refreshToken, err
	}
	if err != nil {
		return refreshToken, err
	}
	return refreshToken, nil
}
