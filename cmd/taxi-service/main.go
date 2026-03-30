package main

import (
	"context"
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/ukique/taxi-service/internal/core/database"

	driverTransport "github.com/ukique/taxi-service/internal/features/driver/transport"

	userTransport "github.com/ukique/taxi-service/internal/features/user/transport"

	orderTransport "github.com/ukique/taxi-service/internal/features/order/transport"
)

func main() {
	// get DATABASE_URL from .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}
	dataBaseURL := os.Getenv("DATABASE_URL")

	//get SECRET_KEY for JWT
	secretKey := os.Getenv("SECRET_KEY")

	// create *Conn for database features
	ctx := context.Background()
	conn, err := database.CreateConnection(ctx, dataBaseURL)
	if err != nil {
		log.Fatal("fail connect to database:", err)
	}
	defer func() {
		if err := conn.Close(ctx); err != nil {
			log.Println("fail to close connection", err)
		}
	}()

	//GIN setup
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:5173"},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders: []string{"Content-Type"},
	}))
	//users
	router.POST("/users/register", userTransport.RegisterUserHandler(conn))
	router.POST("/users/authentication", userTransport.AuthenticationUserHandler(conn, secretKey))

	//drivers
	router.POST("/drivers/register", driverTransport.RegisterDriverHandler(conn))
	router.GET("/drivers", driverTransport.AllDriversHandler(conn))
	router.DELETE("/drivers/:id", driverTransport.DeleteDriverHandler(conn))
	//orders
	router.POST("/orders", orderTransport.CreateOrderHandler(conn))
	router.POST("/orders/complete", orderTransport.CompleteOrderHandler(conn))

	if err := router.Run(":8080"); err != nil {
		log.Fatal("fail run server on port 8080:", err)
	}
}
