package transport

import (
	"context"

	orderTransport "github.com/ukique/taxi-service/internal/features/order/transport"
	"github.com/ukique/taxi-service/internal/models"
)

type DriversRepository interface {
	GetDriversData(ctx context.Context, pageID int) ([]models.Driver, error)
	GetDriversHistory(ctx context.Context, driverID int, pageID int) ([]models.OrderCoordinateEvent, error)
	ChangeDriverName(ctx context.Context, id int, username string) error
	ChangeDriverStatus(ctx context.Context, id int, status string) error
	CreateDriver(ctx context.Context, username string) error
	DeleteDriverByID(ctx context.Context, id int) error
}

type DriverHandler struct {
	secretKey        string
	hub              orderTransport.Broadcaster
	driverRepository DriversRepository
}

func NewDriverHandler(secretKey string, hub orderTransport.Broadcaster, driveRepository DriversRepository) *DriverHandler {
	return &DriverHandler{secretKey: secretKey, hub: hub, driverRepository: driveRepository}
}
