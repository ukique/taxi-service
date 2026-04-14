package models

import "time"

type Order struct {
	ID        int         `json:"id"`
	DriverID  int         `json:"driver_id"`
	Status    OrderStatus `json:"status"`
	CreatedAt time.Time   `json:"created_at"`
}

// OrderCoordinateEvent is a struct for send broker message
// Example structure you can check in docs/examples/message_broker_structure.md
type OrderCoordinateEvent struct {
	DriverID int
	Coordinates
	Order
}

// OrderStatus is a ENUM for orderStatus in Orders Table
type OrderStatus string
