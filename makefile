include .env
export
service-run:
	go run cmd/taxi-service/main.go
migrate-up:
	goose -dir migrations postgres 	"$(DATABASE_URL)" up