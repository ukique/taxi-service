package repository

import (
	"context"
	"time"

	"github.com/ukique/taxi-service/internal/models"
)

func (r *LocationRepository) SaveLocationBatch(ctx context.Context, events []models.OrderCoordinateEvent) error {
	orderIDs := make([]int64, len(events))
	driverIDs := make([]int64, len(events))
	lats := make([]float64, len(events))
	lons := make([]float64, len(events))
	createdAts := make([]time.Time, len(events))

	for i, e := range events {
		orderIDs[i] = int64(e.Order.ID)
		driverIDs[i] = int64(e.Order.DriverID)
		lats[i] = e.Coordinates.Lat
		lons[i] = e.Coordinates.Lon
		createdAts[i] = e.Coordinates.CreatedAt
	}

	sqlQuery := `
    INSERT INTO driver_locations (order_id, driver_id, lat, lon, created_at)
    SELECT * FROM unnest($1::bigint[], $2::bigint[], $3::float8[], $4::float8[], $5::timestamp[])
`
	_, err := r.pool.Exec(ctx, sqlQuery, orderIDs, driverIDs, lats, lons, createdAts)
	if err != nil {
		return err
	}
	return nil
}
