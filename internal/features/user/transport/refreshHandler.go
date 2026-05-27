package transport

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ukique/taxi-service/internal/middleware"
)

func (h *Handler) RefreshTokenHandler(c *gin.Context) {
	clientToken, err := c.Cookie("refreshToken")
	if err != nil {
		log.Println("failed to get clientRefreshToken: ", err)
		c.JSON(http.StatusUnauthorized, gin.H{"message": "you aren't authorized!"})
		return
	}
	refreshToken, err := h.userRepository.SearchRefreshToken(c.Request.Context(), clientToken)
	if err != nil {
		log.Println("refreshToken is invalid!:", err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "You need to Log In first!"})
		return
	}
	if time.Now().After(refreshToken.ExpiresAt) {
		log.Println("refresh token had been expire:", err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "Session finished yet. Log In again."})
		return
	}
	accessToken, err := middleware.GenerateJWT(h.secretKey, refreshToken.UserName)
	if err != nil {
		log.Println("fail to create JWT accessToken:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "server error"})
		return
	}
	c.SetCookie("accessToken", //name
		accessToken, //data
		60*5,        // maxAge (here 5minutes)
		"/",         // path
		"",          // domain
		false,       // secure WARNING: when add NGINX(https) need to do it 'true'
		false,       // httpOnly
	)
	c.JSON(http.StatusOK, gin.H{"message": "refresh successful!"})
}
