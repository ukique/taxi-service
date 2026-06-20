package main

import (
	"context"
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
	"github.com/ukique/taxi-service/metts-taxi/internal/core/jwt"
	"github.com/ukique/taxi-service/metts-taxi/internal/features/user/repository"
	"github.com/ukique/taxi-service/metts-taxi/internal/features/user/service"
	"github.com/ukique/taxi-service/metts-taxi/internal/features/user/transport"
)

func main() {
	err := godotenv.Load(".env.metts")
	if err != nil {
		log.Fatal("No .env file, using environment variables:", err)
	}
	port := os.Getenv("METTS_PORT")
	dbURL := os.Getenv("METTS_DATABASE_URL")
	secretKey := os.Getenv("METTS_SECRET_KEY")

	pool, err := pgxpool.New(context.Background(), dbURL)
	if err != nil {
		log.Fatal("failed to connect to AuthDB: ", err)
	}

	authRepository := repository.NewAuthRepository(pool)
	jwtMaker := jwt.NewTokenMaker(secretKey, authRepository)
	authService := service.NewAuthService(authRepository, jwtMaker)
	authHandler := transport.NewAuthHandler(authRepository, authService)
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PATCH", "PUT", "DELETE"},
		AllowHeaders:     []string{"Content-Type"},
		AllowCredentials: true,
	}))
	r.POST("/register", authHandler.RegisterUserHandler)
	r.POST("/login", authHandler.LoginHandler)
	r.PUT("/refresh/token", authHandler.RefreshTokenHandler)
	if err := r.Run(":" + port); err != nil {
		log.Fatal("failed to run auth server: ", err)
	}
}
