package ws

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ukique/taxi-service/internal/middleware"
)

func (h *Handler) WebSocketHandler(c *gin.Context) {
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
