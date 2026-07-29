package main

import (
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	batch2 "github.com/ukique/taxi-service/internal/core/batch"
	"github.com/ukique/taxi-service/internal/core/connections"
	rabbitmq2 "github.com/ukique/taxi-service/internal/core/rabbitmq"
	ws2 "github.com/ukique/taxi-service/internal/core/ws"
	driversRepository "github.com/ukique/taxi-service/internal/features/driver/repository"
	driverTransport "github.com/ukique/taxi-service/internal/features/driver/transport"
	"github.com/ukique/taxi-service/internal/features/locations/consumer"
	locationrepository "github.com/ukique/taxi-service/internal/features/locations/repository"
	locationtransport "github.com/ukique/taxi-service/internal/features/locations/transport"
	"github.com/ukique/taxi-service/internal/features/order/repository"
	orderService "github.com/ukique/taxi-service/internal/features/order/service"
	orderTransport "github.com/ukique/taxi-service/internal/features/order/transport"
	userRepository "github.com/ukique/taxi-service/internal/features/user/repository"
	userService "github.com/ukique/taxi-service/internal/features/user/service"
	userTransport "github.com/ukique/taxi-service/internal/features/user/transport"
	"github.com/ukique/taxi-service/internal/models"
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
	hub := ws2.NewHub()
	go hub.Run()

	//Create the coordinate channel for CoordinatesBatch
	coordinatesChan := make(chan models.OrderCoordinateEvent, connection.Config.Batch.CoordinatesChanSize)

	//drivers
	driverRepository := driversRepository.NewDriversRepository(connection.Pool)
	driverHandler := driverTransport.NewDriverHandler(connection.SecretKey, hub, driverRepository)
	//orders
	orderRepository := repository.NewOrderRepository(connection.Pool)
	orderServices := orderService.NewOrderServices(connection.Pool, driverRepository)
	orderHandler := orderTransport.NewOrderHandler(connection.Pool, connection.SecretKey, hub, orderRepository, orderServices, connection.Broker)
	//user
	usersRepository := userRepository.NewUserRepository(connection.Pool)
	usersService := userService.NewUserService(usersRepository, connection.SecretKey)
	usersHandler := userTransport.NewUserHandler(connection.Pool, connection.SecretKey, usersRepository, usersService)
	//locations
	locationRepository := locationrepository.NewLocationRepository(connection.Pool)
	locationHandler := locationtransport.NewLocationHandler(locationRepository, connection.SecretKey)
	locationConsumer := consumer.NewLocationConsumer(orderRepository, driverRepository, hub, coordinatesChan)
	//batch
	batch := batch2.NewBatch(connection.Config.Batch.BatchSize, connection.Config.Batch.WaitTimeout, coordinatesChan, locationRepository)
	//ws
	websocket := ws2.NewWSHandler(connection.Pool, hub, orderRepository, driverRepository, locationRepository, connection.SecretKey)

	orderCreatedConfig := rabbitmq2.QueueConfig{
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
	orderCoordinatesConfig := rabbitmq2.QueueConfig{
		Name:       "order.coordinates",
		Durable:    true,
		AutoDelete: false,
		Exclusive:  false,
		NoWait:     false,
		Args:       nil,
	}
	_, err = connection.Broker.DeclareQueue(orderCoordinatesConfig)
	if err != nil {
		log.Println("fail to Declare Queue order.coordinates :", err)
		os.Exit(1)
	}
	coordinatesBatchConfig := rabbitmq2.QueueConfig{
		Name:       "coordinates.batch",
		Durable:    true,
		AutoDelete: false,
		Exclusive:  false,
		NoWait:     false,
		Args:       nil,
	}
	_, err = connection.Broker.DeclareQueue(coordinatesBatchConfig)
	if err != nil {
		log.Println("failed to Declare Queue coordinates.batch :", err)
		os.Exit(1)
	}
	orderCoordinatesConsumerConfig := rabbitmq2.ConsumerConfig{
		QueueName:   "order.coordinates",
		ConsumerTag: "",
		AutoAck:     false,
		Exclusive:   false,
		NoLocal:     false,
		NoWait:      false,
		Args:        nil,
	}
	coordinatesBatchConsumerConfig := rabbitmq2.ConsumerConfig{
		QueueName:   "coordinates.batch",
		ConsumerTag: "",
		AutoAck:     false,
		Exclusive:   false,
		NoLocal:     false, //always false
		NoWait:      false,
		Args:        nil,
	}
	//Batch
	go batch.CoordinatesBatch()

	go connection.Broker.Consumer(orderCoordinatesConsumerConfig, locationConsumer.OrderLocationConsumer)
	go connection.Broker.Consumer(coordinatesBatchConsumerConfig, locationConsumer.CoordinatesBatchConsumer)

	//GIN setup
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost"},
		AllowMethods:     []string{"GET", "POST", "PATCH", "PUT", "DELETE"},
		AllowHeaders:     []string{"Content-Type"},
		AllowCredentials: true,
	}))
	//websocket
	router.GET("/ws", websocket.WebSocketHandler)
	//user
	router.POST("/users/register", usersHandler.RegisterUserHandler)
	router.POST("/users/authentication", usersHandler.AuthenticationUserHandler)
	router.POST("/refreshToken", usersHandler.RefreshTokenHandler)
	//drivers
	router.GET("/drivers/:id/page/:pageID", driverHandler.GetDriversHistoryHandler)
	router.POST("/drivers/create", driverHandler.CreateDriverHandler)
	router.DELETE("/drivers/:id", driverHandler.DeleteDriverHandler)
	router.PATCH("/drivers/:id/username", driverHandler.ChangeDriverNameHandler)
	router.PATCH("/drivers/:id/status", driverHandler.ChangeDriverStatusHandler)
	//order
	router.POST("/orders", orderHandler.CreateOrderHandler)
	router.GET("/orders/:id", orderHandler.GetOrdersDataHandler)
	//coordinates
	router.GET("/location/:id", locationHandler.OrderLocationHistoryHandler)
	if err := router.Run(":" + connection.AppPort); err != nil {
		log.Println("fail run server on port 8080:", err)
		os.Exit(1)
	}
}
