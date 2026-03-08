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
			c.JSON(http.StatusBadRequest, gin.H{"error": "fail to read JSON"})
			return
		}
		//validating data
		if driver.Username == "" {
			log.Println("username can't be empty:")
			c.JSON(http.StatusBadRequest, gin.H{"error": "username can't be empty"})
			return
		}
		if driver.Password == "" {
			log.Println("password can't be empty:")
			c.JSON(http.StatusBadRequest, gin.H{"error": "password can't be empty"})
			return
		}

		if err := repository.RegisterDriver(c.Request.Context(), conn, driver.Username, driver.Password, driver.Email); err != nil {
			log.Println("fail to register driver:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "fail to register driver"})
			return
		}

		c.IndentedJSON(http.StatusCreated, nil)
	}
}
