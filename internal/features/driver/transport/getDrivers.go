package transport

import (
	"context"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *DriverHandler) GetDriversHistoryHandler(c *gin.Context) {
	driversPage := c.Param("id")
	pageID := c.Param("pageID")

	driversPageInt, err := strconv.Atoi(driversPage)
	if err != nil {
		log.Println("failed to Atoi driversPage:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Server error!"})
		return
	}

	pageIDInt, err := strconv.Atoi(pageID)
	if err != nil {
		log.Println("failed to Atoi driversPage:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Server error!"})
		return
	}
	log.Println(driversPageInt, " ", pageIDInt)
	driversHistory, err := h.driverRepository.GetDriversHistory(context.Background(), driversPageInt, pageIDInt)
	if err != nil {
		log.Println("failed to GetDriversHistory:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Server error!"})
		return
	}

	c.JSON(http.StatusOK, driversHistory)
}
