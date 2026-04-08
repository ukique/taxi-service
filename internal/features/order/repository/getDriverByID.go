package repository

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

func GetDriverIDByOrder(ctx context.Context, pool *pgxpool.Pool, orderID int) (int, error) {
	sqlQuery := `
    SELECT driver_id FROM orders WHERE id = $1
`
	var driverID int
	err := pool.QueryRow(ctx, sqlQuery, orderID).Scan(&driverID)
	if err != nil {
		return 0, fmt.Errorf("fail to select driverID:%w", err)
	}
	return driverID, nil
}
