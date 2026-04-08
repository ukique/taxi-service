package database

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

func CreateConnection(ctx context.Context, dataBaseURL string) (*pgxpool.Pool, error) {
	pool, err := pgxpool.New(ctx, dataBaseURL)
	if err != nil {
		return nil, fmt.Errorf("failed connect to database %w", err)
	}
	return pool, err
}
