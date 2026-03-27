package repository

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/ukique/taxi-service/internal/models"
)

func GetDataByEmail(ctx context.Context, conn *pgx.Conn, email string) (models.User, error) {
	sqlQuery := `
	SELECT username, password FROM users WHERE email = $1 
`
	var userData models.User
	err := conn.QueryRow(ctx, sqlQuery, email).Scan(&userData.Username, &userData.Password)
	if err != nil {
		return models.User{}, err
	}
	userData.Email = email
	return userData, nil
}
