package services

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/ukique/taxi-service/internal/features/driver/repository"
	"github.com/ukique/taxi-service/internal/models"
)

func CreateOrder(ctx context.Context, pool *pgxpool.Pool) (models.Order, error) {
	driverID, err := repository.SearchAvailableDriver(ctx, pool)
	if err != nil {
		return models.Order{}, fmt.Errorf("failed to find driver: %w", err)
	}

	const orderStatus = models.OrderStatus("created")

	sqlQuery := `
	INSERT INTO orders(driver_id,status,created_at)
	VALUES ($1, $2, $3)
	RETURNING id, driver_id, status, created_at
`
	driverStatus := "driving"
	// unlock driver in db after too (check SearchAvailableDriver func)
	if err := repository.ChangeDriverStatus(ctx, pool, driverID, driverStatus); err != nil {
		return models.Order{}, fmt.Errorf("fail to change driver status: %w", err)
	}

	var order models.Order
	err = pool.QueryRow(ctx, sqlQuery, driverID, orderStatus, time.Now()).
		Scan(&order.ID, &order.DriverID, &order.Status, &order.CreatedAt)
	if err != nil {
		return models.Order{}, fmt.Errorf("fail to insert into orders: %w", err)
	}

	return order, nil
}
