package config

import "errors"

// ErrInvalidRefreshToken means that client refreshToken expired or invalid
var ErrInvalidRefreshToken = errors.New("invalid refreshToken")
