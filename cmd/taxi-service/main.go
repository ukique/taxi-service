package main

import (
	"context"
	"log"
	"os"

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

	// create *Conn for database features
	ctx := context.Background()
	conn, err := database.CreateConnection(ctx, dataBaseURL)
	if err != nil {
		log.Fatal("fail connect to database:", err)
	}

	//GIN setup
	router := gin.Default()
	router.POST("/users/register", userTransport.RegisterUserHandler(conn))
	router.POST("/drivers/register", driverTransport.RegisterDriverHandler(conn))
	router.POST("/orders", orderTransport.CreateOrderHandler(conn))
	router.POST("/orders/complete", orderTransport.CompleteOrderHandler(conn))
	if err := router.Run(":8080"); err != nil {
		log.Fatal("fail run server on port 8080:", err)
	}
}
