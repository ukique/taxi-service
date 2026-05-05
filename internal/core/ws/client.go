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
type Client struct {
	conn            *websocket.Conn
	send            chan []byte
	hub             *Hub
	orderRepository OrderRepository
}
type Handler struct {
	pool            *pgxpool.Pool
	hub             *Hub
	orderRepository OrderRepository
}

func NewWSHandler(pool *pgxpool.Pool, hub *Hub, service OrderRepository) *Handler {
	return &Handler{pool: pool, hub: hub, orderRepository: service}
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
		var message models.Message
		err := c.conn.ReadJSON(&message)
		if err != nil {
			log.Println("failed to readMessage", err)
			break
		}
		switch message.Type {
		case "orders":
			order, err := json.Marshal(message)
			if err != nil {
				break
			}
			c.hub.broadcast <- order
		case "drivers":
			drivers, err := json.Marshal(message)
			if err != nil {
				break
			}
			c.hub.broadcast <- drivers
		}
	}
}

func (c *Client) WritePump() {
	defer func() {
		c.conn.Close()
	}()
	for {
		select {
		case orders, ok := <-c.send:
			if !ok {
				if err := c.conn.WriteMessage(websocket.CloseMessage, []byte{}); err != nil {
					log.Println("failed to close send connection", err)
					return
				}
			}
			var ordersMessage models.Message
			err := json.Unmarshal(orders, &ordersMessage)
			if err != nil {
				return
			}
			ordersData, err := c.orderRepository.GetOrdersData(context.Background(), ordersMessage.Page)
			if err != nil {
				return
			}
			if err := c.conn.WriteJSON(ordersData); err != nil {
				return
			}
			log.Println(ordersMessage.Type)
			if err != nil {
				return
			}
		}
	}
}
