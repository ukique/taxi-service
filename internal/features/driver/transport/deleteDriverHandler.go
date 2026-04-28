package transport

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ukique/taxi-service/internal/features/driver/repository"
	"github.com/ukique/taxi-service/internal/middleware"
)

func (h *DriverHandler) DeleteDriverHandler(c *gin.Context) {
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
		log.Println("fail convert data", err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "id must be number"})
		return
	}

	if err := repository.DeleteDriverByID(c.Request.Context(), h.pool, idInt); err != nil {
		log.Println("fail to delete driver", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "fail to delete driver"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "driver deleted"})
}
