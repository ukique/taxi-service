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

type AuthHandler struct {
	pool      *pgxpool.Pool
	secretKey string
}

func NewAuthUserHandler(pool *pgxpool.Pool, secretKey string) *AuthHandler {
	return &AuthHandler{pool: pool, secretKey: secretKey}
}

func (h *AuthHandler) AuthenticationUserHandler(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		log.Println("fail to read JSON:", err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "fail to read data"})
		return
	}
	isValid := service.VerifyUserCredentials(c.Request.Context(), h.pool, user.Email, user.Username, user.Password)
	if isValid {
		tokenJWT, err := middleware.GenerateJWT(user, h.secretKey)
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
