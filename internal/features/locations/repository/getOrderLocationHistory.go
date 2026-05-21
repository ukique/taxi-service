package repository

import (
	"context"

	"github.com/ukique/taxi-service/internal/models"
)

func (r *LocationRepository) GetOrderLocationHistory(ctx context.Context, orderID int) ([]models.OrderCoordinateEvent, error) {
	sqlQuery := `
	SELECT order_id, driver_id, lat, lon
	FROM driver_locations
	WHERE order_id = $1;
`
	var events []models.OrderCoordinateEvent
	rows, err := r.pool.Query(ctx, sqlQuery, orderID)
	if err != nil {
		return []models.OrderCoordinateEvent{}, err
	}

	for rows.Next() {
		var e models.OrderCoordinateEvent
		if err := rows.Scan(&e.Order.ID, &e.DriverID, &e.Coordinates.Lat, &e.Coordinates.Lon); err != nil {
			return []models.OrderCoordinateEvent{}, err
		}
		events = append(events, e)
	}
	return events, nil
}
