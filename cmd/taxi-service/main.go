package main

import (
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/ukique/taxi-service/internal/core/connections"
	"github.com/ukique/taxi-service/internal/core/rabbitmq"
	"github.com/ukique/taxi-service/internal/core/ws"
	driversRepository "github.com/ukique/taxi-service/internal/features/driver/repository"
	driverTransport "github.com/ukique/taxi-service/internal/features/driver/transport"
	"github.com/ukique/taxi-service/internal/features/locations"
	locationrepository "github.com/ukique/taxi-service/internal/features/locations/repository"
	locationtransport "github.com/ukique/taxi-service/internal/features/locations/transport"
	"github.com/ukique/taxi-service/internal/features/order/repository"
	userTransport "github.com/ukique/taxi-service/internal/features/user/transport"

	orderTransport "github.com/ukique/taxi-service/internal/features/order/transport"
)

func main() {
	connection := connections.LoadConnections()
	defer connection.Pool.Close()

	defer func() {
		if err := connection.Broker.Close(); err != nil {
			log.Printf("failed to close broker: %v", err)
		}
	}()

	//Run Hub for ws connections
	hub := ws.NewHub()
	go hub.Run()

	orderRepository := repository.NewOrderRepository(connection.Pool)
	driverRepository := driversRepository.NewDriversRepository(connection.Pool)
	userHandler := userTransport.NewUserRegisterHandler(connection.Pool)
	authUserHandler := userTransport.NewAuthUserHandler(connection.Pool, connection.SecretKey)
	driverHandler := driverTransport.NewDriverHandler(connection.Pool, connection.SecretKey, hub, driverRepository)
	orderHandler := orderTransport.NewOrderHandler(connection.Pool, connection.SecretKey, hub, orderRepository, connection.Broker)
	refreshTokenHandler := userTransport.NewRefreshHandler(connection.Pool, connection.SecretKey)
	locationRepository := locationrepository.NewLocationRepository(connection.Pool)
	locationHandler := locationtransport.NewLocationHandler(locationRepository)
	websocket := ws.NewWSHandler(connection.Pool, hub, orderRepository, driverRepository, locationRepository)

	locationConsumer := locations.NewLocationConsumer(locationRepository, hub)
	orderCreatedConfig := rabbitmq.QueueConfig{
		Name:       "order.created",
		Durable:    true,
		AutoDelete: false,
		Exclusive:  false,
		NoWait:     false,
		Args:       nil,
	}
	_, err := connection.Broker.DeclareQueue(orderCreatedConfig)
	if err != nil {
		log.Println("fail to Declare Queue order.created :", err)
		os.Exit(1)
	}
	orderCoordinatesConsumerConfig := rabbitmq.ConsumerConfig{
		QueueName:   "order.coordinates",
		ConsumerTag: "",
		AutoAck:     false,
		Exclusive:   false,
		NoLocal:     false,
		NoWait:      false,
		Args:        nil,
	}
	go connection.Broker.Consumer(orderCoordinatesConsumerConfig, locationConsumer.OrderLocationConsumer)
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
	//router.GET("/orders/details/:id")
	//coordinates
	router.GET("/location/:id", locationHandler.OrderLocationHistoryHandler)
	if err := router.Run(":8080"); err != nil {
		log.Println("fail run server on port 8080:", err)
		os.Exit(1)
	}
}
