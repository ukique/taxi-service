package transport

import (
	"context"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ukique/taxi-service/internal/middleware"
)

func (h *Handler) OrderLocationHistoryHandler(c *gin.Context) {
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

	orderIDParam := c.Param("id")

	orderID, err := strconv.Atoi(orderIDParam)
	if err != nil {
		log.Println("failed to Atoi orderIDParam in OrderLocationHistoryHandler: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Server error!"})
		return
	}

	location, err := h.getter.GetOrderLocationHistory(context.Background(), orderID)
	if err != nil {
		log.Println("failed to get OrderLocationHistory:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "this order doesn't have coordinates!"})
		return
	}

	c.JSON(http.StatusOK, location)
}
