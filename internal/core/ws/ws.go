package ws

import (
	"log"

	"github.com/gin-gonic/gin"
)

func (h *Handler) WebSocketHandler(c *gin.Context) {
	wsConn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println("failed to upgrade ws connection:", err)
		return
	}
	client := &Client{
		conn:               wsConn,
		send:               make(chan []byte, 256),
		hub:                h.hub,
		orderRepository:    h.orderRepository,
		driverRepository:   h.driverRepository,
		locationRepository: h.locationRepository,
	}
	h.hub.register <- client

	go client.ReadPump()
	go client.WritePump()
}
