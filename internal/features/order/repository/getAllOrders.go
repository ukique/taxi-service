package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/ukique/taxi-service/internal/models"
)

func GetOrdersData(ctx context.Context, pool *pgxpool.Pool, pageID int) ([]models.Order, error) {
	sqlQuery := `
  SELECT id, driver_id, status, created_at FROM orders
  ORDER BY id DESC 
  LIMIT $1 OFFSET $2;
`
	recordsLimit := 50
	offest := recordsLimit * (pageID - 1)
	var orders []models.Order
	rows, err := pool.Query(ctx, sqlQuery, recordsLimit, offest)
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
