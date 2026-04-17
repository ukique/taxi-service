package ws

import (
	"context"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
	driverRepository "github.com/ukique/taxi-service/internal/features/driver/repository"
	orderRepository "github.com/ukique/taxi-service/internal/features/order/repository"
)

func (h *Handler) WebSocketHandler(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println("failed to upgrade ws connection:", err)
		return
	}
	defer conn.Close()

	//cancel when client disconnected
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	safe := &safeConn{conn: conn}
	channel := c.Param("channel")

	//detects that client disconnected
	go func() {
		defer cancel()
		for {
			_, _, err := conn.ReadMessage()
			if err != nil {
				return
			}
		}
	}()

	switch channel {
	case "orders":
		orderPageID := c.Param("id")
		intOrderPageID, err := strconv.Atoi(orderPageID)
		if err != nil {
			log.Println("failed to get order_id:", err)
			return
		}
		orders, err := orderRepository.GetOrdersData(ctx, h.pool, intOrderPageID)
		if err != nil {
			return
		}
		if err := safe.WriteJSON(orders); err != nil {
			return
		}
		go h.OrdersHandler(ctx, safe, intOrderPageID)
	case "drivers":
		driverPageID := c.Param("id")
		intDriverPageID, err := strconv.Atoi(driverPageID)
		if err != nil {
			log.Println("failed to get driver_id:", err)
			return
		}
		drivers, err := driverRepository.GetDriversData(ctx, h.pool, intDriverPageID)
		if err != nil {
			return
		}
		if err := safe.WriteJSON(drivers); err != nil {
			return
		}
		go h.DriversHandler(ctx, safe, intDriverPageID)
	}
	<-ctx.Done()
}
