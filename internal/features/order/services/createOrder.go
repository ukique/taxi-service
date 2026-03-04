package services

import (
	"context"
	"fmt"
	"math/rand"

	"github.com/jackc/pgx/v5"
	"github.com/ukique/taxi-service/internal/features/driver/repository"
)

func CreateOrder(ctx context.Context, conn *pgx.Conn, userID int) error {
	driverID, err := repository.SearchAvailableDriver(ctx, conn)
	if err != nil {
		return fmt.Errorf("fail to find driver: %w", err)
	}

	pickupLat := rand.Float64()
	pickupLon := rand.Float64()
	dropOffLat := rand.Float64()
	dropOffLon := rand.Float64()

	orderStatus := true // status:in_progress

	sqlQuery := `
	INSERT INTO orders(user_id, driver_id, pickup_lat, pickup_lon,  dropOff_lat,  dropOff_lon,  status)
	VALUES ($1, $2, $3, $4, $5, $6, $7)
		
`
	_, err = conn.Exec(ctx, sqlQuery, userID, driverID, pickupLat, pickupLon, dropOffLat, dropOffLon, orderStatus)
	if err != nil {
		return fmt.Errorf("fail to insert into orders: %w", err)
	}

	return nil
}
