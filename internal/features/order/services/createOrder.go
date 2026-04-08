package services

import (
	"context"
	"fmt"
	"math/rand"

	"github.com/jackc/pgx/v5"
	"github.com/ukique/taxi-service/internal/features/driver/repository"
	"github.com/ukique/taxi-service/internal/models"
)

func CreateOrder(ctx context.Context, conn *pgx.Conn, userID int) error {
	driverID, err := repository.SearchAvailableDriver(ctx, conn)
	if err != nil {
		return fmt.Errorf("fail to find driver: %w", err)
	}

	const orderStatus = models.OrderStatus("created")
	pickupLat := rand.Float64()*180 - 90
	pickupLon := rand.Float64()*360 - 180
	dropoutLat := rand.Float64()*180 - 90
	dropoutLon := rand.Float64()*360 - 180

	sqlQuery := `
	INSERT INTO orders(user_id, driver_id, pickup_lat, pickup_lon ,dropout_lat,dropout_lon,status)
	VALUES ($1, $2, $3, $4, $5, $6, $7)
		
`
	driverStatus := "driving"
	// unlock driver in db after too (check SearchAvailableDriver func)
	if err := repository.ChangeDriverStatus(ctx, conn, driverID, driverStatus); err != nil {
		return fmt.Errorf("fail to change driver status: %w", err)
	}
	_, err = conn.Exec(ctx, sqlQuery, userID, driverID, pickupLat, pickupLon, dropoutLat, dropoutLon, orderStatus)
	if err != nil {
		return fmt.Errorf("fail to insert into orders: %w", err)
	}

	return nil
}
