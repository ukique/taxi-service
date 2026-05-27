package repository

import (
	"context"
)

// SearchAvailableDriver finds and locks a random available driver.
// Must be paired with UnlockDriver func after order completion,
// otherwise the driver will remain locked in the DB permanently.
func (d *DriversRepository) SearchAvailableDriver(ctx context.Context) (int, error) {
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
	err := d.pool.QueryRow(ctx, sqlQuery).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

// UnlockDriver sets driver status to available after order completion.
func (d *DriversRepository) UnlockDriver(ctx context.Context, driverID int) error {
	sqlQuery := `
    UPDATE drivers
	SET status = 'offline'
    WHERE id = $1
`
	_, err := d.pool.Exec(ctx, sqlQuery, driverID)
	if err != nil {
		return err
	}
	return nil
}
