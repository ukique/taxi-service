package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

func DeleteDriverByID(ctx context.Context, pool *pgxpool.Pool, id int) error {
	tx, err := pool.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	sqlQuery := `
	DELETE FROM orders WHERE driver_id=$1
`
	if _, err := tx.Exec(ctx, sqlQuery, id); err != nil {
		return err
	}
	sqlQuery = `
	DELETE FROM drivers WHERE id = $1;
`
	if _, err := tx.Exec(ctx, sqlQuery, id); err != nil {
		return err
	}
	return tx.Commit(ctx)
}
