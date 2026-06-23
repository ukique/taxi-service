package transport

import (
	"errors"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/ukique/taxi-service/metts-taxi/internal/core/validation"
	"github.com/ukique/taxi-service/metts-taxi/models"
	"golang.org/x/crypto/bcrypt"
)

func (h *Handler) RegisterUserHandler(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body!"})
		return
	}

	if err := validation.ParseCredentials(user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 12)
	if err != nil {
		log.Println("failed to hash the password upon registration:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Server error. Try later."})
		return
	}
	user.Password = string(hashedPassword)

	if err := h.repository.RegisterUser(c.Request.Context(), user); err != nil {
		if pgErr, ok := errors.AsType[*pgconn.PgError](err); ok && pgErr.Code == "23505" { //23505 is duplicated data pgx error
			switch pgErr.ConstraintName {
			case "users_username_key":
				c.JSON(http.StatusConflict, gin.H{"error": "username already taken"})
			case "users_email_key":
				c.JSON(http.StatusConflict, gin.H{"error": "email already taken"})
			default:
				c.JSON(http.StatusConflict, gin.H{"error": "username or email already taken"})
			}
			return
		}
		log.Println("failed to register user:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Server error. Try later."})
		return
	}

	c.JSON(http.StatusCreated, nil)
}
