package repository

import (
	"context"

	"github.com/ukique/taxi-service/internal/models"
)

// GetLastCoordinatesEvent gets last coordinates from DataBase
func (r *LocationRepository) GetLastCoordinatesEvent(ctx context.Context, orderID int) (models.OrderCoordinateEvent, error) {
	sqlQuery := `
	SELECT id, order_id, driver_id, lat, lon
	FROM driver_locations
	WHERE order_id = $1
	ORDER BY id DESC
	LIMIT 1;
`
	var event models.OrderCoordinateEvent
	err := r.pool.QueryRow(ctx, sqlQuery, orderID).Scan(&event.ID, &event.Order.ID, &event.DriverID, &event.Coordinates.Lat, &event.Coordinates.Lon)
	if err != nil {
		return models.OrderCoordinateEvent{}, err
	}
	return event, nil
}
