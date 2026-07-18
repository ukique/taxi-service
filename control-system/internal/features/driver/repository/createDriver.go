package repository

import (
	"context"
	"time"
)

func (d *DriversRepository) CreateDriver(ctx context.Context, username string) error {
	sqlQuery := `
	INSERT INTO drivers(username,created_at)
	VALUES ($1,$2);
`
	if _, err := d.pool.Exec(ctx, sqlQuery, username, time.Now()); err != nil {
		return err
	}
	return nil
}
