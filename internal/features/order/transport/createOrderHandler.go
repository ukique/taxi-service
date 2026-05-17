package transport

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rabbitmq/amqp091-go"
	"github.com/ukique/taxi-service/internal/core/rabbitmq"
	"github.com/ukique/taxi-service/internal/core/ws"
	"github.com/ukique/taxi-service/internal/features/order/services"
	"github.com/ukique/taxi-service/internal/middleware"
	"github.com/ukique/taxi-service/internal/models"
)

type Handler struct {
	pool            *pgxpool.Pool
	secretKey       string
	hub             Broadcaster
	orderRepository ws.OrderRepository
	broker          Broker
}

type Broadcaster interface {
	SendToBroadcast(payload []byte)
}

type Broker interface {
	PublisherWithContext(ctx context.Context, config rabbitmq.PublisherConfig) error
}

func NewOrderHandler(pool *pgxpool.Pool, secretKey string, hub Broadcaster, orderRepository ws.OrderRepository, broker Broker) *Handler {
	return &Handler{pool: pool, secretKey: secretKey, hub: hub, orderRepository: orderRepository, broker: broker}
}

func (h *Handler) CreateOrderHandler(c *gin.Context) {
	clientToken, err := c.Cookie("accessToken")
	if err != nil {
		log.Println("failed to get clientAccessToken: ", err)
		c.JSON(http.StatusUnauthorized, gin.H{"message": "you aren't authorized!"})
		return
	}
	_, err = middleware.VerifyJWT(h.secretKey, clientToken)
	if err != nil {
		log.Println("Client token is fake:", err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "your token isn't correct, try authorize again."})
		return
	}
	orderData, err := services.CreateOrder(c.Request.Context(), h.pool)
	if err != nil {
		log.Println("failed to create Order:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "fail to create Order"})
		return
	}
	orderJson, err := json.Marshal(&orderData)
	if err != nil {
		log.Println("failed to marshal orderData:", err)
		return
	}
	message := amqp091.Publishing{
		Body: orderJson,
	}
	orderCreatedPublisherConfig := rabbitmq.PublisherConfig{
		Exchange:  "",
		Key:       "order.created",
		Mandatory: false,
		Immediate: false, // (always false)
		Message:   message,
	}
	if err := h.broker.PublisherWithContext(context.Background(), orderCreatedPublisherConfig); err != nil {
		log.Println("failed to publish order.created: ", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"ok":      true,
		"message": "order created!",
	})

	ordersData, err := h.orderRepository.GetOrdersData(c.Request.Context(), 1)
	if err != nil {
		return
	}
	ordersBody := models.OutgoingMessage[[]models.Order]{
		Type: "orders",
		Data: ordersData,
	}
	orders, err := json.Marshal(ordersBody)
	if err != nil {
		return
	}
	h.hub.SendToBroadcast(orders)
}
