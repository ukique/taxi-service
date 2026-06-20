package transport

import (
	"github.com/gin-gonic/gin"
)

func (h *Handler) RefreshTokenHandler(c *gin.Context) {
	/*
		refreshToken, err := c.Cookie("refreshToken")
		if err != nil {
			return
		}
		userName, err := c.Cookie("accessToken")
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": ""})
			return
		}
	*/
}
