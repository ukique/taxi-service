package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/ukique/taxi-service/internal/models"
)

func GetAllOrders(ctx context.Context, pool *pgxpool.Pool) ([]models.Order, error) {
	sqlQuery := `
  SELECT id,user_id, driver_id, pickup_lat, pickup_lon, dropout_lat, dropout_lon, status, created_at, updated_at FROM orders;
`
	var orders []models.Order
	rows, err := pool.Query(ctx, sqlQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var o models.Order
		if err := rows.Scan(&o.ID, &o.UserID, &o.DriverID, &o.PickUpLat, &o.PickUpLon, &o.DropOutLat, &o.DropOutLon, &o.Status, &o.CreatedAt, &o.UpdatedAt); err != nil {
			return nil, err
		}
		orders = append(orders, o)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return orders, nil
}
