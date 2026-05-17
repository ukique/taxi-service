package main

import (
	"context"
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/ukique/taxi-service/internal/core/database"
	"github.com/ukique/taxi-service/internal/core/rabbitmq"
	"github.com/ukique/taxi-service/internal/core/ws"
	driversRepository "github.com/ukique/taxi-service/internal/features/driver/repository"
	driverTransport "github.com/ukique/taxi-service/internal/features/driver/transport"
	"github.com/ukique/taxi-service/internal/features/order/repository"
	userTransport "github.com/ukique/taxi-service/internal/features/user/transport"

	orderTransport "github.com/ukique/taxi-service/internal/features/order/transport"
)

func main() {
	// Load .env
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file", err)
		os.Exit(1)
	}
	// get DATABASE_URL from .env
	dataBaseURL := os.Getenv("DATABASE_URL")

	//get SECRET_KEY for JWT
	secretKey := os.Getenv("SECRET_KEY")

	//get RABBITMQ_URL from .env
	rabbitmqURL := os.Getenv("RABBITMQ_URL")

	// create *Conn for database features
	ctx := context.Background()
	pool, err := database.CreateConnection(ctx, dataBaseURL)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	broker, err := rabbitmq.NewBroker(rabbitmqURL)
	if err != nil {
		log.Println("failed to create NewBroker:", err)
	}
	defer func() {
		if err := broker.Close(); err != nil {
			log.Printf("failed to close broker: %v", err)
		}
	}()

	orderCreatedConfig := rabbitmq.QueueConfig{
		Name:       "order.created",
		Durable:    true,
		AutoDelete: false,
		Exclusive:  false,
		NoWait:     false,
		Args:       nil,
	}
	_, err = broker.DeclareQueue(orderCreatedConfig)
	if err != nil {
		log.Println("fail to Declare Queue order.created :", err)
		os.Exit(1)
	}

	//Run Hub for ws connections
	hub := ws.NewHub()
	go hub.Run()

	orderRepository := repository.NewOrderRepository(pool)
	driverRepository := driversRepository.NewDriversRepository(pool)
	userHandler := userTransport.NewUserRegisterHandler(pool)
	authUserHandler := userTransport.NewAuthUserHandler(pool, secretKey)
	driverHandler := driverTransport.NewDriverHandler(pool, secretKey, hub, driverRepository)
	orderHandler := orderTransport.NewOrderHandler(pool, secretKey, hub, orderRepository, broker)
	refreshTokenHandler := userTransport.NewRefreshHandler(pool, secretKey)

	websocket := ws.NewWSHandler(pool, hub, orderRepository, driverRepository)
	//GIN setup
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PATCH", "PUT", "DELETE"},
		AllowHeaders:     []string{"Content-Type"},
		AllowCredentials: true,
	}))
	//websocket
	router.GET("/ws", websocket.WebSocketHandler)
	//users
	router.POST("/users/register", userHandler.RegisterUserHandler)
	router.POST("/users/authentication", authUserHandler.AuthenticationUserHandler)
	router.POST("/refreshToken", refreshTokenHandler.RefreshTokenHandler)
	//drivers
	router.POST("/drivers/create", driverHandler.CreateDriverHandler)
	router.DELETE("/drivers/:id", driverHandler.DeleteDriverHandler)
	router.PATCH("/drivers/:id/username", driverHandler.ChangeDriverNameHandler)
	router.PATCH("/drivers/:id/status", driverHandler.ChangeDriverStatusHandler)
	//order
	router.POST("/orders", orderHandler.CreateOrderHandler)
	//router.GET("/orders/complete", orderHandler.CompleteOrderHandler)
	router.GET("/orders/details/:id")
	if err := router.Run(":8080"); err != nil {
		log.Println("fail run server on port 8080:", err)
		os.Exit(1)
	}
}
