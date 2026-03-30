package repository

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func DeleteDriverByID(ctx context.Context, conn *pgx.Conn, id int) error {
	sqlQuery := `
	DELETE FROM users WHERE id = $1;
`
	if _, err := conn.Exec(ctx, sqlQuery, id); err != nil {
		return err
	}
	return nil
}
