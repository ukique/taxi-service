package transport

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/ukique/taxi-service/internal/features/driver/repository"
	"github.com/ukique/taxi-service/internal/models"
)

type DriverHandler struct {
	pool *pgxpool.Pool
}

func NewDriverHandler(pool *pgxpool.Pool) *DriverHandler {
	return &DriverHandler{pool: pool}
}

func (h *DriverHandler) ChangeDriverNameHandler(c *gin.Context) {
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
	err = repository.ChangeDriverName(c.Request.Context(), h.pool, idInt, driver.Username)
	if err != nil {
		log.Println("fail to change driver name:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "fail to change driver name"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "driverName changed!"})
}

func (h *DriverHandler) ChangeDriverStatusHandler(c *gin.Context) {
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
	if err := repository.ChangeDriverStatus(c.Request.Context(), h.pool, idInt, driver.Status); err != nil {
		log.Println("fail to change driver status:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "fail to change driver status"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "driver status changed!"})
}
