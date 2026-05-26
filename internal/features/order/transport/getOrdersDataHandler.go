package transport

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetOrdersDataHandler(c *gin.Context) {
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
