package ws

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/ukique/taxi-service/internal/models"
)

type OrderRepository interface {
	GetOrdersData(ctx context.Context, pageID int) ([]models.Order, error)
}
type DriverRepository interface {
	GetDriversData(ctx context.Context, pageID int) ([]models.Driver, error)
}
type LocationRepository interface {
	GetLastCoordinatesEvent(ctx context.Context, orderID int) (models.OrderCoordinateEvent, error)
}
type Client struct {
	conn               *websocket.Conn
	send               chan []byte
	hub                *Hub
	orderRepository    OrderRepository
	driverRepository   DriverRepository
	locationRepository LocationRepository
	subscribeType      string
	subscribedPage     int
}
type Handler struct {
	pool               *pgxpool.Pool
	hub                *Hub
	orderRepository    OrderRepository
	driverRepository   DriverRepository
	locationRepository LocationRepository
	secretKey          string
}

func NewWSHandler(pool *pgxpool.Pool, hub *Hub, orderRepository OrderRepository, driverRepository DriverRepository, locationRepository LocationRepository,
	secretKey string) *Handler {
	return &Handler{pool: pool, hub: hub, orderRepository: orderRepository, driverRepository: driverRepository, locationRepository: locationRepository, secretKey: secretKey}
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024 * 4,
	WriteBufferSize: 1024 * 32,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func (c *Client) ReadPump() {
	defer func() {
		c.hub.unregister <- c
		c.conn.Close()
	}()
	for {
		var message models.IncomingMessage
		err := c.conn.ReadJSON(&message)
		if err != nil {
			log.Println("failed to readMessage", err)
			break
		}
		switch message.Type {
		case "subscribe_orders":
			c.subscribeType = "orders"
			ordersData, err := c.orderRepository.GetOrdersData(context.Background(), message.Page)
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
			c.send <- orders
		case "subscribe_drivers":
			c.subscribeType = "drivers"
			driverData, err := c.driverRepository.GetDriversData(context.Background(), message.Page)
			if err != nil {
				return
			}
			driversBody := models.OutgoingMessage[[]models.Driver]{
				Type: "drivers",
				Data: driverData,
			}
			drivers, err := json.Marshal(driversBody)
			if err != nil {
				return
			}
			c.send <- drivers
		case "subscribe_orderDetails":
			c.subscribeType = "coordinates"
			c.subscribedPage = message.Page //order_ID
			lastEvent, err := c.locationRepository.GetLastCoordinatesEvent(context.Background(), c.subscribedPage)
			if err != nil {
				log.Println("failed to GetLastCoordinatesEvent", err)
				return
			}

			eventBody := models.OutgoingMessage[models.OrderCoordinateEvent]{
				Type: "coordinates",
				Page: c.subscribedPage,
				Data: lastEvent,
			}

			event, err := json.Marshal(eventBody)
			if err != nil {
				log.Println("failed to Marshal eventBody:", err)
				return
			}
			c.send <- event
		}
	}
}

func (c *Client) WritePump() {
	defer func() {
		c.conn.Close()
	}()
	for {
		select {
		case message, ok := <-c.send:
			if !ok {
				if err := c.conn.WriteMessage(websocket.CloseMessage, []byte{}); err != nil {
					log.Println("failed to close send connection", err)
					return
				}
			}
			if err := c.conn.WriteMessage(websocket.TextMessage, message); err != nil {
				return
			}
		}
	}
}
