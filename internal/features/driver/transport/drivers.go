package transport

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	orderTransport "github.com/ukique/taxi-service/internal/features/order/transport"
	"github.com/ukique/taxi-service/internal/models"
)

type DriverHandler struct {
	pool             *pgxpool.Pool
	secretKey        string
	hub              orderTransport.Broadcaster
	driverRepository DriversGetter
}

type DriversGetter interface {
	GetDriversData(ctx context.Context, pageID int) ([]models.Driver, error)
	GetDriversHistory(ctx context.Context, driverID int, pageID int) ([]models.OrderCoordinateEvent, error)
}

func NewDriverHandler(pool *pgxpool.Pool, secretKey string, hub orderTransport.Broadcaster, driveRepository DriversGetter) *DriverHandler {
	return &DriverHandler{pool: pool, secretKey: secretKey, hub: hub, driverRepository: driveRepository}
}
