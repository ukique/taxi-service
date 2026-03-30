package repository

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func CreateDriver(ctx context.Context, conn *pgx.Conn, username string) error {

	sqlQuery := `
	INSERT INTO drivers(username)
	VALUES ($1);
`
	if _, err := conn.Exec(ctx, sqlQuery, username); err != nil {
		return fmt.Errorf("fail to create driver: %w", err)
	}
	return nil
}
