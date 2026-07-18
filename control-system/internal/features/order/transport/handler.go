package transport

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/ukique/taxi-service/internal/core/rabbitmq"
	"github.com/ukique/taxi-service/internal/models"
)

type Broadcaster interface {
	SendToBroadcast(payload []byte)
}
type Broker interface {
	PublisherWithContext(ctx context.Context, config rabbitmq.PublisherConfig) error
}
type OrderServices interface {
	CreateOrder(ctx context.Context) (models.Order, error)
}
type OrderRepository interface {
	GetOrdersData(ctx context.Context, pageID int) ([]models.Order, error)
}
type Handler struct {
	pool            *pgxpool.Pool
	secretKey       string
	hub             Broadcaster
	orderRepository OrderRepository
	orderServices   OrderServices
	broker          Broker
}

func NewOrderHandler(pool *pgxpool.Pool, secretKey string, hub Broadcaster,
	orderRepository OrderRepository, orderServices OrderServices, broker Broker) *Handler {
	return &Handler{
		pool:            pool,
		secretKey:       secretKey,
		hub:             hub,
		orderRepository: orderRepository,
		orderServices:   orderServices,
		broker:          broker}
}
