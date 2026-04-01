package transport

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	driver "github.com/ukique/taxi-service/internal/features/driver/repository"
	order "github.com/ukique/taxi-service/internal/features/order/repository"
	"github.com/ukique/taxi-service/internal/features/order/services"

	"github.com/ukique/taxi-service/internal/models"
)

func CreateOrderHandler(conn *pgx.Conn) func(c *gin.Context) {
	return func(c *gin.Context) {
		var order models.Order
		if err := c.ShouldBindJSON(&order); err != nil {
			log.Println("fail to read JSON body:", err)
			c.JSON(http.StatusBadRequest, gin.H{"message": "incorrect data!"})
			return
		}
		if err := services.CreateOrder(c.Request.Context(), conn, order.UserID); err != nil {
			log.Println("fail to create Order:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"message": "fail to create Order"})
			return
		}
		c.JSON(http.StatusCreated, gin.H{"message": "order created!"})
	}
}

// CompleteOrderHandler didn't update!
func CompleteOrderHandler(conn *pgx.Conn) func(*gin.Context) {
	return func(c *gin.Context) {
		var body struct {
			OrderID int `json:"order_id"`
		}
		if err := c.ShouldBindJSON(&body); err != nil {
			log.Println("fail to read json body:", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": "fail to read json body"})
			return
		}
		//search driverID from DataBase
		driverID, err := order.GetDriverIDByOrder(c.Request.Context(), conn, body.OrderID)
		if err != nil {
			log.Println("fail to get driverID:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "fail to get driverID"})
			return
		}

		// unlock driver (because we use FOR UPDATE SKIP LOCKED in SearchAvailableDriver func)
		if err := driver.UnlockDriver(c.Request.Context(), conn, driverID); err != nil {
			log.Println("fail to unlock driver:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "fail to unlock driver"})
			return
		}

		//update order status to false (closed)
		if err := order.UpdateOrder(c.Request.Context(), conn, body.OrderID); err != nil {
			log.Println("fail to update order:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "fail to update order"})
			return
		}
		c.IndentedJSON(http.StatusOK, nil)
	}
}
