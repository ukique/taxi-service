package transport

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ukique/taxi-service/internal/middleware"
	"github.com/ukique/taxi-service/internal/models"
)

func (h *DriverHandler) ChangeDriverNameHandler(c *gin.Context) {
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

	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		log.Println("fail convert data:", err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "id must be number"})
		return
	}

	var driver models.Driver
	if err := c.ShouldBindJSON(&driver); err != nil {
		log.Println("invalid data:", err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid data"})
		return
	}
	err = h.driverRepository.ChangeDriverName(c.Request.Context(), idInt, driver.Username)
	if err != nil {
		log.Println("fail to change driver name:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "fail to change driver name"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "driverName changed!"})
}

func (h *DriverHandler) ChangeDriverStatusHandler(c *gin.Context) {
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

	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		log.Println("fail convert data:", err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "id must be number"})
		return
	}
	var driver models.Driver
	if err := c.ShouldBindJSON(&driver); err != nil {
		log.Println("invalid data:", err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid data"})
		return
	}
	if driver.Status != "offline" && driver.Status != "searching" && driver.Status != "driving" {
		log.Println("invalid status")
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid status"})
		return
	}
	if err := h.driverRepository.ChangeDriverStatus(c.Request.Context(), idInt, driver.Status); err != nil {
		log.Println("fail to change driver status:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "fail to change driver status"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "driver status changed!"})
}
