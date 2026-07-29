package repository

import (
	"context"

	models2 "github.com/ukique/taxi-service/internal/models"
)

func (d *DriversRepository) GetDriversData(ctx context.Context, pageID int) ([]models2.Driver, error) {
	sqlQuery := ` 
  SELECT id,username, status FROM drivers
  ORDER BY id DESC 
  LIMIT $1 OFFSET $2;
`
	recordsLimit := 50
	offest := recordsLimit * (pageID - 1)
	var drivers []models2.Driver
	rows, err := d.pool.Query(ctx, sqlQuery, recordsLimit, offest)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var d models2.Driver
		if err := rows.Scan(&d.ID, &d.Username, &d.Status); err != nil {
			return nil, err
		}
		drivers = append(drivers, d)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return drivers, nil
}

func (d *DriversRepository) GetDriversHistory(ctx context.Context, driverID int, pageID int) ([]models2.OrderCoordinateEvent, error) {
	sqlQuery := `
	SELECT order_id, driver_id, lat, lon FROM driver_locations
	WHERE driver_id = $1                                     
	ORDER BY id DESC 
	LIMIT $2 OFFSET $3;
`
	recordsLimit := 50
	offest := recordsLimit * (pageID - 1)
	var history []models2.OrderCoordinateEvent
	rows, err := d.pool.Query(ctx, sqlQuery, driverID, recordsLimit, offest)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var h models2.OrderCoordinateEvent
		if err := rows.Scan(&h.Order.ID, &h.Order.DriverID, &h.Coordinates.Lat, &h.Coordinates.Lon); err != nil {
			return nil, err
		}
		history = append(history, h)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return history, nil
}
