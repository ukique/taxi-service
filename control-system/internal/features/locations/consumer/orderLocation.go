package consumer

import (
	"context"
	"encoding/json"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/ukique/taxi-service/internal/models"
)

func (c *Consumer) OrderLocationConsumer(delivery amqp.Delivery) {

	var eventBody models.OrderCoordinateEvent
	if err := json.Unmarshal(delivery.Body, &eventBody); err != nil {
		log.Println("failed to unmarshal delivery.Body:", err)
		err := delivery.Nack(false, false)
		if err != nil {
			log.Println("failed to Nack:", err)
			return
		}
		return
	}

	//Saving to driver_locations table
	if err := c.locationRepository.SaveLocation(context.Background(), eventBody); err != nil {
		log.Println("failed to SaveLocation: ", err)
		err := delivery.Nack(false, true)
		if err != nil {
			log.Println("failed to Nack:", err)
			return
		}
		return
	}

	// Sending to BroadCast where subscribe_orderDetails
	messageBody := models.OutgoingMessage[models.OrderCoordinateEvent]{
		Type: "coordinates",
		Page: eventBody.Order.ID,
		Data: eventBody,
	}

	message, err := json.Marshal(messageBody)
	if err != nil {
		log.Println("failed marshal coordinates OutGoingMessage:", err)
		err := delivery.Nack(false, true)
		if err != nil {
			log.Println("failed to Nack:", err)
			return
		}
		return
	}
	c.hub.SendToBroadcast(message)
	if eventBody.EventID == 1 {
		updOrdersData, err := c.orderRepository.GetOrdersData(context.Background(), 1)
		if err != nil {
			return
		}
		ordersBody := models.OutgoingMessage[[]models.Order]{
			Type: "orders",
			Data: updOrdersData,
		}
		orders, err := json.Marshal(ordersBody)
		if err != nil {
			return
		}
		c.hub.SendToBroadcast(orders)
	}

	if err := delivery.Ack(false); err != nil {
		log.Println("failed to send Ack message:", err)
	}
	if eventBody.Status == "done" {
		//search driverID from DataBase
		driverID, err := c.orderRepository.GetDriverIDByOrder(context.Background(), eventBody.Order.ID)
		if err != nil {
			log.Println("failed to get driverID:", err)
			return
		}

		// unlock driver (because we use FOR UPDATE SKIP LOCKED in SearchAvailableDriver func)
		if err := c.driverRepository.UnlockDriver(context.Background(), driverID); err != nil {
			log.Println("failed to unlock driver:", err)
			return
		}

		//update order status to false (closed)
		if err := c.orderRepository.UpdateOrder(context.Background(), eventBody.Order.ID); err != nil {
			log.Println("failed to update order:", err)
			return
		}

		ordersData, err := c.orderRepository.GetOrdersData(context.Background(), 1)
		if err != nil {
			return
		}
		ordersBody := models.OutgoingMessage[[]models.Order]{
			Type: "orders",
			Data: ordersData,
		}
		orders, err := json.Marshal(ordersBody)
		if err != nil {
			return
		}
		c.hub.SendToBroadcast(orders)
	}
}
