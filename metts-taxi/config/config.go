package config

import "time"

// TokenDuration is how long refreshToken will live.
var TokenDuration = time.Minute * 60 * 24 * 14 // 14 days
