package services

import (
	"math/rand"
	"time"
)

// GenerateCoordinates generate random float64 coordinates
// for example lat = 47.842658, lon = 34.811989
func GenerateCoordinates() (float64, float64, time.Time) {
	driverLat := rand.Float64()*180 - 90
	driverLon := rand.Float64()*360 - 180
	generatedTime := time.Now()
	return driverLat, driverLon, generatedTime
}
