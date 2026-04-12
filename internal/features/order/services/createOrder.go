package services

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/ukique/taxi-service/internal/features/driver/repository"
	"github.com/ukique/taxi-service/internal/models"
)

func CreateOrder(ctx context.Context, pool *pgxpool.Pool) error {
	driverID, err := repository.SearchAvailableDriver(ctx, pool)
	if err != nil {
		return fmt.Errorf("fail to find driver: %w", err)
	}

	const orderStatus = models.OrderStatus("created")

	sqlQuery := `
	INSERT INTO orders(driver_id,status)
	VALUES ($1, $2)
`
	driverStatus := "driving"
	// unlock driver in db after too (check SearchAvailableDriver func)
	if err := repository.ChangeDriverStatus(ctx, pool, driverID, driverStatus); err != nil {
		return fmt.Errorf("fail to change driver status: %w", err)
	}
	_, err = pool.Exec(ctx, sqlQuery, driverID, orderStatus)
	if err != nil {
		return fmt.Errorf("fail to insert into orders: %w", err)
	}

	return nil
}
