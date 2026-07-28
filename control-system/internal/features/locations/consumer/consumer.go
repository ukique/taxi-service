package consumer

import (
	"context"

	"github.com/ukique/taxi-service/internal/features/order/transport"
	"github.com/ukique/taxi-service/internal/models"
)

type Consumer struct {
	hub              transport.Broadcaster
	orderRepository  OrderRepository
	driverRepository DriverRepository
	coordinatesChan  chan models.OrderCoordinateEvent
}

func NewLocationConsumer(orderRepository OrderRepository, driverRepository DriverRepository, hub transport.Broadcaster, coordinatesChan chan models.OrderCoordinateEvent) *Consumer {
	return &Consumer{
		orderRepository:  orderRepository,
		driverRepository: driverRepository,
		hub:              hub,
		coordinatesChan:  coordinatesChan,
	}
}

type OrderRepository interface {
	UpdateOrder(ctx context.Context, orderID int) error
	GetDriverIDByOrder(ctx context.Context, orderID int) (int, error)
	GetOrdersData(ctx context.Context, pageID int) ([]models.Order, error)
}
type DriverRepository interface {
	UnlockDriver(ctx context.Context, driverID int) error
}
