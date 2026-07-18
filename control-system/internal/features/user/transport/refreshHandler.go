package transport

import (
	"errors"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ukique/taxi-service/config"
)

func (h *Handler) RefreshTokenHandler(c *gin.Context) {
	clientToken, err := c.Cookie("refreshToken")
	log.Println("refreshToken:", clientToken)
	if err != nil {
		log.Println("failed to get clientRefreshToken: ", err)
		c.JSON(http.StatusUnauthorized, gin.H{"message": "you aren't authorized!"})
		return
	}
	accessToken, err := h.userService.RefreshTokenService(c.Request.Context(), clientToken)
	if err != nil {
		if errors.Is(err, config.ErrInvalidRefreshToken) {
			log.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{"message": "you aren't authorized!"})
			return
		}
		log.Println("refreshToken service error:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "server error!"})
		return
	}
	c.SetCookie("accessToken", //name
		accessToken, //data
		60*5,        // maxAge (here 5minutes)
		"/",         // path
		"",          // domain
		true,        // secure WARNING: when add NGINX(https) need to do it 'true'
		false,       // httpOnly
	)
	c.JSON(http.StatusOK, gin.H{"message": "refresh successful!"})
}
