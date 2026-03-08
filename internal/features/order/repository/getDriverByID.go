package repository

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func GetDriverIDByOrder(ctx context.Context, conn *pgx.Conn, orderID int) (int, error) {
	sqlQuery := `
    SELECT driver_id FROM orders WHERE id = $1
`
	var driverID int
	err := conn.QueryRow(ctx, sqlQuery, orderID).Scan(&driverID)
	if err != nil {
		return 0, fmt.Errorf("fail to select driverID:%w", err)
	}
	return driverID, nil
}
