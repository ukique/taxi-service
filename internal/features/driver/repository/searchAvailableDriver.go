package repository

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

// SearchAvailableDriver finds and locks a random available driver.
// Must be paired with UnlockDriver func after order completion,
// otherwise the driver will remain locked in the DB permanently.
func SearchAvailableDriver(ctx context.Context, pool *pgxpool.Pool) (int, error) {
	sqlQuery := `
    UPDATE drivers
    SET status = 'driving'
    WHERE id = (  
    SELECT id FROM drivers
    WHERE status = 'searching'
	ORDER BY random()
	LIMIT 1 
	FOR UPDATE SKIP LOCKED -- Lock driver to avoid duplicate requests in goroutines
    )
    RETURNING id
`
	var id int
	err := pool.QueryRow(ctx, sqlQuery).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("fail query data: %w", err)
	}
	return id, nil
}

// UnlockDriver sets driver status to available after order completion.
func UnlockDriver(ctx context.Context, pool *pgxpool.Pool, driverID int) error {
	sqlQuery := `
    UPDATE drivers
	SET status = 'offline'
    WHERE id = $1
`
	_, err := pool.Exec(ctx, sqlQuery, driverID)
	if err != nil {
		return fmt.Errorf("fail unlock driver: %w", err)
	}
	return nil
}
