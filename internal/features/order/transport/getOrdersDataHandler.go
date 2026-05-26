package transport

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ukique/taxi-service/internal/middleware"
)

func (h *Handler) GetOrdersDataHandler(c *gin.Context) {
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

	pageIDParam := c.Param("id")

	pageID, err := strconv.Atoi(pageIDParam)
	if err != nil {
		log.Println("failed to Atoi orders PageID:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Server Error!"})
		return
	}

	orders, err := h.orderRepository.GetOrdersData(c.Request.Context(), pageID)
	if err != nil {
		log.Println("failed to GetOrdersData:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "No orders found!"})
		return
	}
	c.JSON(http.StatusOK, orders)
}
