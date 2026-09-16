package benchmarks

import (
	"context"
	"os"
	"sync"
	"testing"

	"github.com/joho/godotenv"
	"github.com/rabbitmq/amqp091-go"
	"github.com/ukique/taxi-service/internal/core/rabbitmq"
)

func BenchmarkProduce_OrderCreated(b *testing.B) {
	err := godotenv.Load("../.env")
	if err != nil {
		b.Fatalf("No .env file, using environment variables")
	}
	rabbitURL := os.Getenv("RABBITMQ_URL_TEST")
	conn, err := amqp091.Dial(rabbitURL)
	if err != nil {
		b.Fatalf("dial error: %v", err)
	}
	defer conn.Close()
	ch, err := conn.Channel()
	if err != nil {
		b.Fatalf("channel error: %v", err)
	}
	defer ch.Close()

	_, err = ch.QueueDeclare("order.created", true, false, false, false, nil)
	if err != nil {
		b.Fatalf("queue declare failed: %v", err)
	}

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		message := amqp091.Publishing{
			Body: []byte("test order"),
		}
		config := rabbitmq.PublisherConfig{
			Exchange:  "",
			Key:       "order.created",
			Mandatory: false,
			Immediate: false, // (always false)
			Message:   message,
		}
		err := ch.PublishWithContext(
			context.Background(),
			config.Exchange,
			config.Key,
			config.Mandatory,
			config.Immediate,
			config.Message,
		)
		if err != nil {
			b.Fatalf("publish failed: %v", err)
		}

	}
}

func BenchmarkProduce_OrderCoordinates(b *testing.B) {
	err := godotenv.Load("../.env")
	if err != nil {
		b.Fatalf("No .env file, using environment variables")
	}
	rabbitURL := os.Getenv("RABBITMQ_URL_TEST")
	conn, err := amqp091.Dial(rabbitURL)
	if err != nil {
		b.Fatalf("dial error: %v", err)
	}
	defer conn.Close()
	ch, err := conn.Channel()
	if err != nil {
		b.Fatalf("channel error: %v", err)
	}
	defer ch.Close()

	_, err = ch.QueueDeclare("order.coordinates", true, false, false, false, nil)
	if err != nil {
		b.Fatalf("queue declare failed: %v", err)
	}
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		message := amqp091.Publishing{
			Body: []byte("test coordinates"),
		}
		config := rabbitmq.PublisherConfig{
			Exchange:  "",
			Key:       "order.coordinates",
			Mandatory: false,
			Immediate: false, // (always false)
			Message:   message,
		}
		err := ch.PublishWithContext(
			context.Background(),
			config.Exchange,
			config.Key,
			config.Mandatory,
			config.Immediate,
			config.Message,
		)
		if err != nil {
			b.Fatalf("publish failed: %v", err)
		}
	}
}

// BenchmarkProduce_Combined simulates real production traffic: order.created
// and order.coordinates being published AT THE SAME TIME, sharing the same
// connection/CPU/network — unlike running the two benchmarks separately,
// which never overlap.
//
// Ratio: for every 1 order.created message, 50 order.coordinates messages
// are published, matching the real-world ratio (each order sends ~50
// location updates over its lifetime).
func BenchmarkProduce_Combined(b *testing.B) {
	err := godotenv.Load("../.env")
	if err != nil {
		b.Fatalf("No .env file, using environment variables")
	}

	rabbitURL := os.Getenv("RABBITMQ_URL_TEST")
	conn, err := amqp091.Dial(rabbitURL)
	if err != nil {
		b.Fatalf("dial error: %v", err)
	}
	defer conn.Close()


	setupCh, err := conn.Channel()
	if err != nil {
		b.Fatalf("channel error: %v", err)
	}
	if _, err := setupCh.QueueDeclare("order.created", true, false, false, false, nil); err != nil {
		b.Fatalf("queue declare failed (order.created): %v", err)
	}
	if _, err := setupCh.QueueDeclare("order.coordinates", true, false, false, false, nil); err != nil {
		b.Fatalf("queue declare failed (order.coordinates): %v", err)
	}
	setupCh.Close()

	const coordinatesPerOrder = 50

	b.ResetTimer()
	b.ReportAllocs()

	var wg sync.WaitGroup

	// goroutine 1: publishes to order.created, b.N times
	wg.Add(1)
	go func() {
		defer wg.Done()

		ch, err := conn.Channel()
		if err != nil {
			panic(err)
		}
		defer ch.Close()

		msg := amqp091.Publishing{Body: []byte("test order")}

		for i := 0; i < b.N; i++ {
			err := ch.PublishWithContext(
				context.Background(),
				"", "order.created", false, false, msg,
			)
			if err != nil {
				panic(err)
			}
		}
	}()

	// goroutine 2: publishes to order.coordinates, b.N * 50 times,
	// matching the real ratio of coordinate updates per order
	wg.Add(1)
	go func() {
		defer wg.Done()

		ch, err := conn.Channel()
		if err != nil {
			panic(err)
		}
		defer ch.Close()

		msg := amqp091.Publishing{Body: []byte("test coordinates")}

		for i := 0; i < b.N*coordinatesPerOrder; i++ {
			err := ch.PublishWithContext(
				context.Background(),
				"", "order.coordinates", false, false, msg,
			)
			if err != nil {
				panic(err)
			}
		}
	}()

	wg.Wait()
}
