package models

import "time"

type Coordinates struct {
	Lat           float64   `json:"lat"`
	Lon           float64   `json:"lon"`
	GeneratedTime time.Time `json:"generated_time"`
}
