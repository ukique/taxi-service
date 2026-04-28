package transport

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	driver "github.com/ukique/taxi-service/internal/features/driver/repository"
	order "github.com/ukique/taxi-service/internal/features/order/repository"
	"github.com/ukique/taxi-service/internal/features/order/services"
	"github.com/ukique/taxi-service/internal/middleware"
)

type Handler struct {
	pool      *pgxpool.Pool
	secretKey string
}

func NewOrderHandler(pool *pgxpool.Pool, secretKey string) *Handler {
	return &Handler{pool: pool, secretKey: secretKey}
}

func (h *Handler) CreateOrderHandler(c *gin.Context) {
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
	if err := services.CreateOrder(c.Request.Context(), h.pool); err != nil {
		log.Println("fail to create Order:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "fail to create Order"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "order created!"})
}

func (h *Handler) CompleteOrderHandler(c *gin.Context) {
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
	var body struct {
		OrderID int `json:"order_id"`
	}
	if err := c.ShouldBindJSON(&body); err != nil {
		log.Println("fail to read json body:", err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "fail to read json body"})
		return
	}

	//search driverID from DataBase
	driverID, err := order.GetDriverIDByOrder(c.Request.Context(), h.pool, body.OrderID)
	if err != nil {
		log.Println("fail to get driverID:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "fail to get driverID"})
		return
	}

	// unlock driver (because we use FOR UPDATE SKIP LOCKED in SearchAvailableDriver func)
	if err := driver.UnlockDriver(c.Request.Context(), h.pool, driverID); err != nil {
		log.Println("fail to unlock driver:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "fail to unlock driver"})
		return
	}

	//update order status to false (closed)
	if err := order.UpdateOrder(c.Request.Context(), h.pool, body.OrderID); err != nil {
		log.Println("fail to update order:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "fail to update order"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "order completed!"})
}
