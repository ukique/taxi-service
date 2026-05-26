package repository

import (
	"context"

	"github.com/ukique/taxi-service/internal/models"
)

// GetLastCoordinatesEvent gets last coordinates from DataBase
func (r *LocationRepository) GetLastCoordinatesEvent(ctx context.Context, orderID int) (models.OrderCoordinateEvent, error) {
	sqlQuery := `
	SELECT l.id, l.order_id, l.driver_id, l.lat, l.lon, o.status
	FROM driver_locations l
	JOIN orders o ON o.id = l.order_id
	WHERE l.order_id = $1
	ORDER BY id DESC
	LIMIT 1;
`
	var event models.OrderCoordinateEvent
	err := r.pool.QueryRow(ctx, sqlQuery, orderID).Scan(&event.ID, &event.Order.ID, &event.DriverID, &event.Coordinates.Lat, &event.Coordinates.Lon, &event.Order.Status)
	if err != nil {
		return models.OrderCoordinateEvent{}, err
	}
	return event, nil
}
