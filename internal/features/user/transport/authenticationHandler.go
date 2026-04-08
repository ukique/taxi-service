package transport

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/ukique/taxi-service/internal/features/user/service"
	"github.com/ukique/taxi-service/internal/middleware"
	"github.com/ukique/taxi-service/internal/models"
)

func AuthenticationUserHandler(pool *pgxpool.Pool, secretKey string) func(c *gin.Context) {
	return func(c *gin.Context) {
		var user models.User
		if err := c.ShouldBindJSON(&user); err != nil {
			log.Println("fail to read JSON:", err)
			c.JSON(http.StatusBadRequest, gin.H{"message": "fail to read data"})
			return
		}
		isValid := service.VerifyUserCredentials(c.Request.Context(), pool, user.Email, user.Username, user.Password)
		if isValid {
			tokenJWT, err := middleware.GenerateJWT(user, secretKey)
			if err != nil {
				log.Println("fail to sign key:", err)
				c.JSON(http.StatusInternalServerError, gin.H{"message": "server error"})
				return
			}
			c.JSON(http.StatusOK, gin.H{"token": tokenJWT})
		} else {
			log.Println("user data isn't valid")
			c.JSON(http.StatusBadRequest, gin.H{"message": "data isn't valid"})
			return
		}

	}
}
