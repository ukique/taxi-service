package transport

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ukique/taxi-service/internal/features/driver/repository"
	"github.com/ukique/taxi-service/internal/middleware"
	"github.com/ukique/taxi-service/internal/models"
)

func (h *DriverHandler) CreateDriverHandler(c *gin.Context) {
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
	var driver models.Driver
	if err := c.ShouldBindJSON(&driver); err != nil {
		log.Println("fail to read JSON body:", err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "fail to read JSON"})
		return
	}
	//validating data
	if driver.Username == "" {
		log.Println("username can't be empty:")
		c.JSON(http.StatusBadRequest, gin.H{"message": "username can't be empty"})
		return
	}

	if err := repository.CreateDriver(c.Request.Context(), h.pool, driver.Username); err != nil {
		log.Println("fail to register driver:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "fail to register driver"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "driver created!"})
}
