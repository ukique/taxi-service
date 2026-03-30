package repository

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func ChangeDriverName(ctx context.Context, conn *pgx.Conn, id int, username string) error {
	sqlQuery := `
	UPDATE drivers
	SET username = $1
	WHERE id = $2;
`
	if _, err := conn.Exec(ctx, sqlQuery, username, id); err != nil {
		return err
	}
	return nil
}

func ChangeDriverStatus(ctx context.Context, conn *pgx.Conn, id int, status string) error {
	sqlQuery := `
	UPDATE drivers
	SET status = $1
	WHERE id = $2
`
	if _, err := conn.Exec(ctx, sqlQuery, status, id); err != nil {
		return err
	}
	return nil
}
