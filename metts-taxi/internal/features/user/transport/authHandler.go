package transport

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ukique/taxi-service/metts-taxi/internal/core/validation"
	"github.com/ukique/taxi-service/metts-taxi/models"
)

func (h *Handler) LoginHandler(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body!"})
		return
	}

	tokenPair, err := h.service.Login(c.Request.Context(), user)
	if err != nil {
		if errors.Is(err, validation.ErrInvalidCredentials) {
			c.JSON(http.StatusConflict, gin.H{"error": "invalid credentials."})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Server error. Try later."})
		return
	}

	c.SetCookie(
		"accessToken",
		tokenPair.AccessToken,
		60*5, //5 minutes
		"/",
		"",
		false,
		false,
	)

	c.SetCookie(
		"refreshToken",
		tokenPair.RefreshToken,
		60*60*24*14, //14 days
		"/",
		"",
		false,
		true,
	)

	c.JSON(http.StatusOK, nil)
}
