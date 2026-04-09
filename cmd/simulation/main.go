package main

import (
	"context"
	"log"
	"os"
	"sync"
	"time"

	"github.com/joho/godotenv"
	"github.com/ukique/taxi-service/internal/core/database"
	"github.com/ukique/taxi-service/internal/core/rabbitmq"
	orderFeatures "github.com/ukique/taxi-service/internal/features/order/repository"
	"github.com/ukique/taxi-service/internal/features/order/services"
	"github.com/ukique/taxi-service/internal/models"
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
	pool, err := database.CreateConnection(ctx, dataBaseURL)
	if err != nil {
		log.Println("fail connect to database:", err)
		os.Exit(1)
	}

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

	var processed sync.Map

	for {
		orders, err := orderFeatures.GetCreatedOrders(ctx, pool)
		if err != nil {
			log.Println("fail to get orders:", err)
			time.Sleep(1 * time.Second)
			continue
		}

		for _, order := range orders {
			if _, exists := processed.Load(order.ID); exists {
				continue
			}
			processed.Store(order.ID, true)

			go func(order models.Order) {
				defer processed.Delete(order.ID)

				for i := 0; i < 50; i++ {
					err := services.SendCoordinates(ctx, pool, brokerChannel, orderCoordinatesQueue, order)
					if err != nil {
						log.Println("fail to send coordinate:", err)
					}
					time.Sleep(5 * time.Second)
				}
			}(order)
		}
		time.Sleep(1 * time.Second)
	}
}
