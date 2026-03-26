package transport

import (
	"errors"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/ukique/taxi-service/internal/core"
	"github.com/ukique/taxi-service/internal/features/user/repository"
	"github.com/ukique/taxi-service/internal/models"
)

func RegisterUserHandler(conn *pgx.Conn) func(c *gin.Context) {
	return func(c *gin.Context) {
		var user models.User
		if err := c.ShouldBindJSON(&user); err != nil {
			log.Println("fail to read JSON body:", err)
			c.JSON(http.StatusBadRequest, gin.H{"message": "fail to read JSON"})
			return
		}
		//validating data
		if user.Username == "" {
			log.Println("username can't be empty:")
			c.JSON(http.StatusBadRequest, gin.H{"message": "username can't be empty"})
			return
		}
		if user.Email == "" {
			log.Println("email can't be empty:")
			c.JSON(http.StatusBadRequest, gin.H{"message": "email can't be empty"})
			return
		}
		if user.Password == "" {
			log.Println("password can't be empty:")
			c.JSON(http.StatusBadRequest, gin.H{"message": "password can't be empty"})
			return
		}

		if len(user.Username) > 16 {
			log.Println("username can't be more than 16 characters")
			c.JSON(http.StatusBadRequest, gin.H{"message": "username can't be more than 16 characters"})
			return
		}
		if len(user.Email) > 254 {
			log.Println("incorrect email address")
			c.JSON(http.StatusBadRequest, gin.H{"message": "incorrect email address"})
			return
		}
		if !core.EmailValid(user.Email) {
			log.Println("this isn't email address")
			c.JSON(http.StatusBadRequest, gin.H{"message": "this isn't email address"})
			return
		}
		if len(user.Password) > 72 {
			log.Println("password is too long")
			c.JSON(http.StatusBadRequest, gin.H{"message": "password is too long"})
			return
		}
		if len(user.Password) < 8 {
			log.Println("password must be at least 8 characters")
			c.JSON(http.StatusBadRequest, gin.H{"message": "password must be at least 8 characters"})
			return
		}
		//duplicated data
		var pgErr *pgconn.PgError
		err := repository.RegisterUser(c.Request.Context(), conn, user.Username, user.Password, user.Email)
		if err != nil {
			if errors.As(err, &pgErr) && pgErr.Code == "23505" { //23505 is duplicated data pgx error
				switch pgErr.ConstraintName {
				case "users_username_key":
					c.JSON(http.StatusConflict, gin.H{"message": "username already taken"})
				case "users_email_key":
					c.JSON(http.StatusConflict, gin.H{"message": "email already taken"})
				default:
					c.JSON(http.StatusConflict, gin.H{"message": "username or email already taken"})
				}
				return
			}
			log.Println("fail to register user:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"message": "server error"})
			return
		}

		c.JSON(http.StatusCreated, gin.H{"message": "user created!"})
	}
}
