package transport

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"github.com/ukique/taxi-service/internal/features/user/repository"
	"github.com/ukique/taxi-service/internal/models"
)

func RegisterUserHandler(conn *pgx.Conn) func(c *gin.Context) {
	return func(c *gin.Context) {
		var user models.User
		if err := c.ShouldBindJSON(&user); err != nil {
			log.Println("fail to read JSON body:", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": "fail to read JSON"})
			return
		}
		//validating data
		if user.Username == "" {
			log.Println("username can't be empty:")
			c.JSON(http.StatusBadRequest, gin.H{"error": "username can't be empty"})
			return
		}
		if user.Password == "" {
			log.Println("password can't be empty:")
			c.JSON(http.StatusBadRequest, gin.H{"error": "password can't be empty"})
			return
		}

		if err := repository.RegisterUser(c.Request.Context(), conn, user.Username, user.Password, user.Email); err != nil {
			log.Println("fail to register user:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "fail to register user"})
			return
		}
		c.JSON(http.StatusCreated, gin.H{"message:": "user created!"})
	}
}
