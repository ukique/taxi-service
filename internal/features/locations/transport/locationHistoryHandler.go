package transport

import (
	"context"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) OrderLocationHistoryHandler(c *gin.Context) {
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
