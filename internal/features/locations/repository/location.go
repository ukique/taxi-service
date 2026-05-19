package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/ukique/taxi-service/internal/models"
)

type LocationRepository struct {
	pool *pgxpool.Pool
}

func NewLocationRepository(pool *pgxpool.Pool) *LocationRepository {
	return &LocationRepository{pool: pool}
}

func (r *LocationRepository) SaveLocation(ctx context.Context, orderBody models.OrderCoordinateEvent) error {
	sqlQuery := `
	INSERT INTO driver_locations (order_id, driver_id, lat,lon)
	VALUES ($1,$2,$3,$4);
`
	_, err := r.pool.Exec(ctx, sqlQuery, orderBody.Order.ID, orderBody.DriverID, orderBody.Coordinates.Lat, orderBody.Coordinates.Lon)
	if err != nil {
		return err
	}
	return nil
}
