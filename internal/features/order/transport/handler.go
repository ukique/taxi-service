package transport

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"github.com/ukique/taxi-service/internal/features/order/services"
	"github.com/ukique/taxi-service/internal/models"
)

func CreateOrderHandler(conn *pgx.Conn) func(c *gin.Context) {
	return func(c *gin.Context) {
		var order models.Order
		if err := c.ShouldBindJSON(&order); err != nil {
			log.Println("fail to read JSON body:", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": "fail to read JSON"})
			return
		}
		if err := services.CreateOrder(c.Request.Context(), conn, order.UserID); err != nil {
			log.Println("fail to create Order:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "fail to create Order"})
			return
		}
		c.IndentedJSON(http.StatusCreated, nil)
	}
}
