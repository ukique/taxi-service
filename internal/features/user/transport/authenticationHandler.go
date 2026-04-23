package transport

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/ukique/taxi-service/internal/features/user/repository"
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
		c.JSON(http.StatusBadRequest, gin.H{"message": "data isn't correct"})
		return
	}
	//Validate UserData
	isValid := service.VerifyUserCredentials(c.Request.Context(), h.pool, user.Email, user.Username, user.Password)
	if isValid {
		refreshToken, err := middleware.GenerateRefreshToken(16)
		if err != nil {
			log.Println("failed to create refreshToken", err)
			c.JSON(http.StatusInternalServerError, gin.H{"message": "server error"})
			return
		}
		createdAt := time.Now()
		expireAt := createdAt.Add(time.Minute * 60 * 24 * 7) // 7 days
		dataBaseRefreshToken := models.RefreshToken{
			UserName:     user.Username,
			RefreshToken: refreshToken,
			CreatedAt:    createdAt,
			ExpiresAt:    expireAt,
		}
		err = repository.SaveRefreshToken(c.Request.Context(), h.pool, dataBaseRefreshToken)
		if err != nil {
			return
		}
		accessToken, err := middleware.GenerateJWT(h.secretKey, user.Username)
		if err != nil {
			log.Println("fail to create JWT accessToken:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"message": "server error"})
			return
		}
		c.SetCookie(
			"refreshToken", // name
			refreshToken,   // data
			60*60*24*7,     // maxAge (here 7 days)
			"/",            // path
			"",             // domain
			false,          // secure WARNING: when add NGINX(https) need to do it 'true'
			true,           // httpOnly
		)
		c.SetCookie(
			"accessToken", //name
			accessToken,   //data
			60*5,          // maxAge (here 5minutes)
			"/",           // path
			"",            // domain
			false,         // secure WARNING: when add NGINX(https) need to do it 'true'
			false,         // httpOnly
		)
	} else {
		log.Println("user data isn't valid")
		c.JSON(http.StatusBadRequest, gin.H{"message": "data isn't valid"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "auth successful!"})
}
