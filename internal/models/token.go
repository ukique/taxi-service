package models

import "time"

type RefreshToken struct {
	ID           int       `json:"id"`
	UserName     string    `json:"username"`
	RefreshToken string    `json:"refresh_token"`
	CreatedAt    time.Time `json:"created_at"`
	ExpiresAt    time.Time `json:"expires_at"`
}
