package service

import (
	"context"
	"time"

	"github.com/ukique/taxi-service/internal/models"
)

func (s *OrderServices) CreateOrder(ctx context.Context) (models.Order, error) {
	driverID, err := s.driversRepository.SearchAvailableDriver(ctx)
	if err != nil {
		return models.Order{}, err
	}

	const orderStatus = models.OrderStatus("created")

	sqlQuery := `
	INSERT INTO orders(driver_id,status,created_at)
	VALUES ($1, $2, $3)
	RETURNING id, driver_id, status, created_at
`
	driverStatus := "driving"
	// unlock driver in db after too (check SearchAvailableDriver func)
	if err := s.driversRepository.ChangeDriverStatus(ctx, driverID, driverStatus); err != nil {
		return models.Order{}, err
	}

	var order models.Order
	err = s.pool.QueryRow(ctx, sqlQuery, driverID, orderStatus, time.Now()).
		Scan(&order.ID, &order.DriverID, &order.Status, &order.CreatedAt)
	if err != nil {
		return models.Order{}, err
	}

	return order, nil
}
