package service

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type DriversRepository interface {
	ChangeDriverStatus(ctx context.Context, id int, status string) error
	SearchAvailableDriver(ctx context.Context) (int, error)
}
type OrderServices struct {
	pool              *pgxpool.Pool
	driversRepository DriversRepository
}

func NewOrderServices(pool *pgxpool.Pool, driverRepository DriversRepository) *OrderServices {
	return &OrderServices{
		pool:              pool,
		driversRepository: driverRepository}
}
