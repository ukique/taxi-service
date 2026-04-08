package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/ukique/taxi-service/internal/models"
)

func SaveLocation(ctx context.Context, pool *pgxpool.Pool, orderBody models.OrderCoordinateEvent) error {
	sqlQuery := `
	INSERT INTO driver_locations (order_id, driver_id, lat,lon, generated_time)
	VALUES ($1,$2,$3,$4,$5);
`
	_, err := pool.Exec(ctx, sqlQuery, orderBody.Order.ID, orderBody.DriverID, orderBody.Coordinates.Lat, orderBody.Coordinates.Lon, orderBody.Coordinates.GeneratedTime)
	if err != nil {
		return err
	}
	return nil
}
