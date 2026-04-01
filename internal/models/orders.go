package models

type Order struct {
	UserID int `json:"user_id"`
}

// OrderStatus is a ENUM for orderStatus in Orders Table
type OrderStatus string
