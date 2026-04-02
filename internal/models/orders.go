package models

import "time"

type Order struct {
	ID         int         `json:"id"`
	UserID     int         `json:"user_id"`
	DriverID   int         `json:"driver_id"`
	Status     OrderStatus `json:"status"`
	PickUpLat  float64     `json:"pick_up_lat"`
	PickUpLon  float64     `json:"pick_up_lon"`
	DropOutLat float64     `json:"drop_out_lat"`
	DropOutLon float64     `json:"drop_out_lon"`
	CreatedAt  time.Time   `json:"created_at"`
	UpdatedAt  time.Time   `json:"updated_at"`
}

// OrderStatus is a ENUM for orderStatus in Orders Table
type OrderStatus string
