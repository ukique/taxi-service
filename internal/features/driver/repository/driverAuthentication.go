package repository

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

func CreateDriver(ctx context.Context, pool *pgxpool.Pool, username string) error {

	sqlQuery := `
	INSERT INTO drivers(username)
	VALUES ($1);
`
	if _, err := pool.Exec(ctx, sqlQuery, username); err != nil {
		return fmt.Errorf("fail to create driver: %w", err)
	}
	return nil
}
