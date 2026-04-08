package transport

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/ukique/taxi-service/internal/features/order/repository"
)

func GetAllOrdersHandler(pool *pgxpool.Pool) func(c *gin.Context) {
	return func(c *gin.Context) {
		orders, err := repository.GetAllOrders(c.Request.Context(), pool)
		if err != nil {
			log.Println("fail to get ALl Orders:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Server Error"})
			return
		}
		c.JSON(http.StatusOK, orders)
	}
}
