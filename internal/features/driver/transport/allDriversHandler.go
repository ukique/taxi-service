package transport

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/ukique/taxi-service/internal/features/driver/repository"
)

func AllDriversHandler(pool *pgxpool.Pool) func(c *gin.Context) {
	return func(c *gin.Context) {
		drivers, err := repository.GetAllDrivers(c.Request.Context(), pool)
		if err != nil {
			log.Println("fail to get all drivers data:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"message": "fail to get all drivers data"})
			return
		}
		c.JSON(http.StatusOK, drivers)
	}
}
