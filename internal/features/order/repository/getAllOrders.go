package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/ukique/taxi-service/internal/models"
)

func GetAllOrders(ctx context.Context, pool *pgxpool.Pool) ([]models.Order, error) {
	sqlQuery := `
  SELECT id, driver_id, status, created_at FROM orders;
`
	var orders []models.Order
	rows, err := pool.Query(ctx, sqlQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var o models.Order
		if err := rows.Scan(&o.ID, &o.DriverID, &o.Status, &o.CreatedAt); err != nil {
			return nil, err
		}
		orders = append(orders, o)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return orders, nil
}
