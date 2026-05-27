package repository

import (
	"context"
)

func (d *DriversRepository) ChangeDriverName(ctx context.Context, id int, username string) error {
	sqlQuery := `
	UPDATE drivers
	SET username = $1
	WHERE id = $2;
`
	if _, err := d.pool.Exec(ctx, sqlQuery, username, id); err != nil {
		return err
	}
	return nil
}

func (d *DriversRepository) ChangeDriverStatus(ctx context.Context, id int, status string) error {
	sqlQuery := `
	UPDATE drivers
	SET status = $1
	WHERE id = $2
`
	if _, err := d.pool.Exec(ctx, sqlQuery, status, id); err != nil {
		return err
	}
	return nil
}
