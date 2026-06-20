package jwt

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/ukique/taxi-service/metts-taxi/models"
	"golang.org/x/crypto/bcrypt"
)

type Repository interface {
	SaveRefreshToken(ctx context.Context, refreshToken models.RefreshToken) error
}

// TokenMaker handles JWT token generation, validating using HS256
type TokenMaker struct {
	secretKey  string // secretKey, loaded from AUTH_SECRET_KEY
	repository Repository
}

type TokenPair struct {
	AccessToken  string
	RefreshToken string
}

func NewTokenMaker(secretKey string, repository Repository) *TokenMaker {
	return &TokenMaker{secretKey: secretKey, repository: repository}
}

// ErrInvalidToken is returned by VerifyToken when token method, is invalid, or token expired.
var ErrInvalidToken = errors.New("invalid token")

// GenerateJWT create a signed JWT token with HS256 that valid for 5 minutes.
// after expires you should refresh it.
func (t *TokenMaker) GenerateJWT(username string) (string, error) {
	claims := jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Minute * 5).Unix(),
	}
	// We use HS256, but after adding AWS we should use RS256 instead of HS256
	// by reason of using microservice architecture
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(t.secretKey))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func (t *TokenMaker) GenerateRefreshToken() (string, error) {
	bytes := make([]byte, 32) // refreshToken is a rand 32 bytes with hex
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	refreshToken := hex.EncodeToString(bytes)
	return refreshToken, nil
}

// VerifyJWT validate userToken
// return ErrInvalidToken is method, token is invalid or token expired
func (t *TokenMaker) VerifyJWT(userToken string) (*jwt.Token, error) {
	token, err := jwt.Parse(userToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, ErrInvalidToken
		}
		return []byte(t.secretKey), nil
	})
	if err != nil {
		return token, err
	}
	if !token.Valid {
		return token, ErrInvalidToken
	}
	return token, nil
}

// CreateTokenPair creates and return refresh and accessToken
func (t *TokenMaker) CreateTokenPair(ctx context.Context, username string, tokenDuration time.Duration) (TokenPair, error) {

	accessToken, err := t.GenerateJWT(username)
	if err != nil {
		return TokenPair{}, err
	}

	refreshToken, err := t.GenerateRefreshToken()
	if err != nil {
		return TokenPair{}, err
	}

	hashedRefreshToken, err := bcrypt.GenerateFromPassword([]byte(refreshToken), 12)
	createdAt := time.Now()
	expiresAt := createdAt.Add(tokenDuration)

	refreshTokenData := models.RefreshToken{
		Username:     username,
		RefreshToken: string(hashedRefreshToken),
		CreatedAt:    createdAt,
		ExpiresAt:    expiresAt,
	}

	if err := t.repository.SaveRefreshToken(ctx, refreshTokenData); err != nil {
		return TokenPair{}, err
	}

	var tokenPair TokenPair
	tokenPair.AccessToken = accessToken
	tokenPair.RefreshToken = refreshToken

	return tokenPair, nil
}
