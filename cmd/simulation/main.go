package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/ukique/taxi-service/internal/core/database"
	"github.com/ukique/taxi-service/internal/core/rabbitmq"
	orderFeatures "github.com/ukique/taxi-service/internal/features/order/repository"
	"github.com/ukique/taxi-service/internal/features/order/services"
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

	//get RABBITMQ_URL from .env
	rabbitmqURL := os.Getenv("RABBITMQ_URL")

	// create *Conn for database features
	ctx := context.Background()
	conn, err := database.CreateConnection(ctx, dataBaseURL)
	if err != nil {
		log.Println("fail connect to database:", err)
		os.Exit(1)
	}
	defer func() {
		if err := conn.Close(ctx); err != nil {
			log.Println("fail to close database connection:", err)
		}
	}()
	// create *amqp.Connection for RabbitMQ features
	brokerConn, err := rabbitmq.CreateRabbitMQConnection(rabbitmqURL)
	if err != nil {
		log.Println("fail to connect to RabbitMQ:", err)
		os.Exit(1)
	}
	defer func() {
		if err := brokerConn.Close(); err != nil {
			log.Println("fail to close RabbitMQ connection:", err)
		}
	}()

	// create *amqp.Channel (broker channel)
	brokerChannel, err := rabbitmq.CreateChannel(brokerConn)
	if err != nil {
		log.Println("fail to create RabbitMQ channel:", err)
		os.Exit(1)
	}
	defer func() {
		if err := brokerChannel.Close(); err != nil {
			log.Println("fail to close RabbitMQ channel:", err)
			os.Exit(1)
		}
	}()

	// create Queue Declare for Orders Coordinates *amqp.Queue (broker queue)
	orderCoordinatesQueue, err := rabbitmq.QueueDeclareOrdersCoordinates(brokerChannel)
	if err != nil {
		log.Println("fail to create Queue Declare Orders Coordinates:", err)
		os.Exit(1)
	}

	for {
		orders, err := orderFeatures.GetCreatedOrders(ctx, conn)
		if err != nil {
			log.Println("fail to get orders:", err)
		}
		//Send coordinate to RabbitMQ
		for i := 0; i < 50; i++ {
			err := services.SendCoordinates(ctx, conn, brokerChannel, orderCoordinatesQueue, orders)
			if err != nil {
				log.Println("fail to send Coordinates:", err)
			}
			time.Sleep(5 * time.Second)
		}
	}
}
