package repository

import (
	"context"
)

func (o *OrderRepository) GetDriverIDByOrder(ctx context.Context, orderID int) (int, error) {
	sqlQuery := `
    SELECT driver_id FROM orders WHERE id = $1
`
	var driverID int
	err := o.pool.QueryRow(ctx, sqlQuery, orderID).Scan(&driverID)
	if err != nil {
		return 0, err
	}
	return driverID, nil
}
