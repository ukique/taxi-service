package main

import (
	"context"
	"log"
	"os"

	"github.com/goccy/go-yaml"
	"github.com/joho/godotenv"
	"github.com/ukique/taxi-service/config"
	"github.com/ukique/taxi-service/internal/core/database"
	"github.com/ukique/taxi-service/internal/core/rabbitmq"
	"github.com/ukique/taxi-service/internal/features/order"
	"github.com/ukique/taxi-service/internal/features/order/repository"
)

func main() {
	// Load .env
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file", err)
		os.Exit(1)
	}

	// pool must be replaced (after refactor UpdateOrderStatus)
	dataBaseURL := os.Getenv("DATABASE_URL")
	pool, err := database.CreateConnection(context.Background(), dataBaseURL)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	//get RABBITMQ_URL from .env
	rabbitmqURL := os.Getenv("RABBITMQ_URL")

	//Get LoadTesting Simulation Data
	var simulationData config.Config

	file, err := os.ReadFile("config/simulation.yaml")
	if err != nil {
		log.Println("failed to load simulation.yaml:", err)
		os.Exit(1)
	}
	if err := yaml.Unmarshal(file, &simulationData); err != nil {
		log.Println("failed to unmarshal simulationData:", err)
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

	orderCoordinatesQueueConfig := rabbitmq.QueueConfig{
		Name:       "order.coordinates",
		Durable:    true,
		AutoDelete: false,
		Exclusive:  false,
		NoWait:     false,
		Args:       nil,
	}
	_, err = broker.DeclareQueue(orderCoordinatesQueueConfig)
	if err != nil {
		log.Println("fail to Declare Queue order.coordinates :", err)
		os.Exit(1)
	}
	orderCreatedConsumerConfig := rabbitmq.ConsumerConfig{
		QueueName:   "order.created",
		ConsumerTag: "",
		AutoAck:     false,
		Exclusive:   false,
		NoLocal:     false,
		NoWait:      false,
		Args:        nil,
	}
	log.Println("simulation is working!")
	orderRepository := repository.NewOrderRepository(pool)
	orderConsumer := order.NewOrderConsumer(simulationData, broker, orderRepository)
	err = broker.Consumer(orderCreatedConsumerConfig, orderConsumer.OrderCreatedConsumer)
	if err != nil {
		log.Println("failed to consume order.created:", err)
	}
}
