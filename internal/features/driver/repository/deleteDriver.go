package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

func DeleteDriverByID(ctx context.Context, pool *pgxpool.Pool, id int) error {
	sqlQuery := `
	DELETE FROM users WHERE id = $1;
`
	if _, err := pool.Exec(ctx, sqlQuery, id); err != nil {
		return err
	}
	return nil
}
