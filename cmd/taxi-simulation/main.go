package main

import (
	"log"
	"os"

	"github.com/goccy/go-yaml"
	"github.com/ukique/taxi-service/config"
	"github.com/ukique/taxi-service/internal/core/connections"
	"github.com/ukique/taxi-service/internal/core/rabbitmq"
	"github.com/ukique/taxi-service/internal/features/order"
	"github.com/ukique/taxi-service/internal/features/order/repository"
)

func main() {
	connection := connections.LoadConnections()
	defer connection.Pool.Close()

	defer func() {
		if err := connection.Broker.Close(); err != nil {
			log.Printf("failed to close broker: %v", err)
		}
	}()

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

	orderCoordinatesQueueConfig := rabbitmq.QueueConfig{
		Name:       "order.coordinates",
		Durable:    true,
		AutoDelete: false,
		Exclusive:  false,
		NoWait:     false,
		Args:       nil,
	}
	_, err = connection.Broker.DeclareQueue(orderCoordinatesQueueConfig)
	if err != nil {
		log.Println("fail to Declare Queue order.coordinates :", err)
		os.Exit(1)
	}
	orderCreatedQueueConfig := rabbitmq.QueueConfig{
		Name:       "order.created",
		Durable:    true,
		AutoDelete: false,
		Exclusive:  false,
		NoWait:     false,
		Args:       nil,
	}
	_, err = connection.Broker.DeclareQueue(orderCreatedQueueConfig)
	if err != nil {
		log.Println("fail to Declare Queue order.created :", err)
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
	orderRepository := repository.NewOrderRepository(connection.Pool)
	orderConsumer := order.NewOrderConsumer(simulationData, connection.Broker, orderRepository)
	connection.Broker.Consumer(orderCreatedConsumerConfig, orderConsumer.OrderCreatedConsumer)
}
