package models

type Order struct {
	UserID     int     `json:"user_id"`
	PickupLat  float64 `json:"pickup_lat"`
	PickupLon  float64 `json:"pickup_lon"`
	DropoutLat float64 `json:"dropout_lat"`
	DropoutLon float64 `json:"dropout_lon"`
}
