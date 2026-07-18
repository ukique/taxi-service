package repository

import (
	"context"

	"github.com/ukique/taxi-service/internal/models"
)

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
