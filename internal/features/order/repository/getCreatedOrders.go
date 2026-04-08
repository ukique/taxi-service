package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/ukique/taxi-service/internal/models"
)

func GetCreatedOrders(ctx context.Context, pool *pgxpool.Pool) ([]models.Order, error) {
	sqlQuery := `
	SELECT id, driver_id, status FROM orders WHERE status = 'created'
`
	var orders []models.Order
	rows, err := pool.Query(ctx, sqlQuery)
	if err != nil {
		return []models.Order{}, err
	}
	defer rows.Close()

	for rows.Next() {
		var o models.Order
		if err := rows.Scan(&o.ID, &o.DriverID, &o.Status); err != nil {
			return []models.Order{}, err
		}
		orders = append(orders, o)
	}
	return orders, nil
}
