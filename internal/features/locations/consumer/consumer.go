package consumer

import (
	"context"

	"github.com/ukique/taxi-service/internal/features/order/transport"
	"github.com/ukique/taxi-service/internal/models"
)

type Consumer struct {
	hub                transport.Broadcaster
	locationRepository LocationRepository
	orderRepository    OrderRepository
	driverRepository   DriverRepository
}

func NewLocationConsumer(locationRepository LocationRepository, orderRepository OrderRepository, driverRepository DriverRepository, hub transport.Broadcaster) *Consumer {
	return &Consumer{
		locationRepository: locationRepository,
		orderRepository:    orderRepository,
		driverRepository:   driverRepository,
		hub:                hub}
}

type LocationRepository interface {
	SaveLocation(ctx context.Context, orderBody models.OrderCoordinateEvent) error
}
type OrderRepository interface {
	UpdateOrder(ctx context.Context, orderID int) error
	GetDriverIDByOrder(ctx context.Context, orderID int) (int, error)
}
type DriverRepository interface {
	UnlockDriver(ctx context.Context, driverID int) error
}
