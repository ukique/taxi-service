package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

func ChangeDriverName(ctx context.Context, pool *pgxpool.Pool, id int, username string) error {
	sqlQuery := `
	UPDATE drivers
	SET username = $1
	WHERE id = $2;
`
	if _, err := pool.Exec(ctx, sqlQuery, username, id); err != nil {
		return err
	}
	return nil
}

func ChangeDriverStatus(ctx context.Context, pool *pgxpool.Pool, id int, status string) error {
	sqlQuery := `
	UPDATE drivers
	SET status = $1
	WHERE id = $2
`
	if _, err := pool.Exec(ctx, sqlQuery, status, id); err != nil {
		return err
	}
	return nil
}
