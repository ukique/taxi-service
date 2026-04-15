package ws

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/ukique/taxi-service/internal/features/order/repository"
)

type WSHandler struct {
	pool *pgxpool.Pool
}

func NewWSHandler(pool *pgxpool.Pool) *WSHandler {
	return &WSHandler{pool: pool}
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024 * 4,
	WriteBufferSize: 1024 * 32,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func (h *WSHandler) WebSocketHandler(c *gin.Context) {
	channel := c.Param("channel")
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println("failed to upgrade ws connection:", err)
		return
	}
	defer conn.Close()
	
	switch channel {
	case "orders":
		orderID := c.Param("id")
		intOrderID, err := strconv.Atoi(orderID)
		if err != nil {
			log.Println("failed to get order_id:", err)
			break
		}
		orders, err := repository.GetOrdersData(c.Request.Context(), h.pool, intOrderID)
		if err != nil {
			log.Println("failed to get ordersData:", err)
			break
		}
		err = conn.WriteJSON(orders)
		if err != nil {
			log.Println("failed to send orders:", err)
			break
		}
	}
}
