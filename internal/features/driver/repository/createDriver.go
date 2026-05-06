package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

func CreateDriver(ctx context.Context, pool *pgxpool.Pool, username string) error {
	sqlQuery := `
	INSERT INTO drivers(username,created_at)
	VALUES ($1,$2);
`
	if _, err := pool.Exec(ctx, sqlQuery, username, time.Now()); err != nil {
		return fmt.Errorf("fail to create driver: %w", err)
	}
	return nil
}
