package transport

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"github.com/ukique/taxi-service/internal/features/driver/repository"
	"github.com/ukique/taxi-service/internal/models"
)

func ChangeDriverNameHandler(conn *pgx.Conn) func(c *gin.Context) {
	return func(c *gin.Context) {
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
		err = repository.ChangeDriverName(c.Request.Context(), conn, idInt, driver.Username)
		if err != nil {
			log.Println("fail to change driver name:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"message": "fail to change driver name"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "driverName changed!"})
	}
}
