package transport

import (
	"context"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ukique/taxi-service/internal/middleware"
)

func (h *DriverHandler) GetDriversHistoryHandler(c *gin.Context) {
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
