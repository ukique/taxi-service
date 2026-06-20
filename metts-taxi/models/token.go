package models

import "time"

type RefreshToken struct {
	Username     string    `json:"username"`
	RefreshToken string    `json:"refreshToken"`
	CreatedAt    time.Time `json:"created_at"`
	ExpiresAt    time.Time `json:"expires_at"`
}
