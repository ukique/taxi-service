package ws

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) WebSocketHandler(c *gin.Context) {

	subscribeType := c.Query("type")
	if subscribeType != "subscribe_orders" && subscribeType != "subscribe_drivers" && subscribeType != "subscribe_locations" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid subscribe type"})
		return
	}

	wsConn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println("failed to upgrade ws connection:", err)
		return
	}
	client := &Client{
		conn:             wsConn,
		send:             make(chan []byte, 256),
		hub:              h.hub,
		subscribeType:    subscribeType, //subscribe_orders, subscribe_drivers, subscribe_locations etc.
		orderRepository:  h.orderRepository,
		driverRepository: h.driverRepository,
	}
	h.hub.register <- client

	go client.ReadPump()
	go client.WritePump()
}
