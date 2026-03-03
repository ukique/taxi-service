package repository

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

// SearchAvailableDriver finds and locks a random available driver.
// Must be paired with UnlockDriver func after order completion,
// otherwise the driver will remain locked in the DB permanently.
func SearchAvailableDriver(ctx context.Context, conn *pgx.Conn) (int, error) {
	sqlQuery := `
    UPDATE drivers
    SET status = false
    WHERE id = (  
    SELECT id FROM drivers
    WHERE status = true
	ORDER BY random()
	LIMIT 1 
	FOR UPDATE SKIP LOCKED -- Lock driver to avoid duplicate requests in goroutines
    )
    RETURNING id
`
	var id int
	err := conn.QueryRow(ctx, sqlQuery).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("fail query data: %w", err)
	}
	return id, nil
}
