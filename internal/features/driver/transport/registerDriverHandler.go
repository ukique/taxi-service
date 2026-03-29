package transport

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"github.com/ukique/taxi-service/internal/features/driver/repository"
	"github.com/ukique/taxi-service/internal/models"
)

func RegisterDriverHandler(conn *pgx.Conn) func(c *gin.Context) {
	return func(c *gin.Context) {
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

		if err := repository.CreateDriver(c.Request.Context(), conn, driver.Username); err != nil {
			log.Println("fail to register driver:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"message": "fail to register driver"})
			return
		}

		c.JSON(http.StatusCreated, gin.H{"message": "driver created!"})
	}
}
