package database

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func CreateConnection(ctx context.Context, databaseURL string) (*pgx.Conn, error) {
	conn, err := pgx.Connect(ctx, databaseURL)
	if err != nil {
		return nil, fmt.Errorf("fail connect to database: %w", err)
	}
	return conn, nil
}
