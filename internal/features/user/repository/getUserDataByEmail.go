package repository

import (
	"context"

	"github.com/ukique/taxi-service/internal/models"
)

func (u *UserRepository) GetDataByEmail(ctx context.Context, email string) (models.User, error) {
	sqlQuery := `
	SELECT username, password FROM users WHERE email = $1 
`
	var userData models.User
	err := u.pool.QueryRow(ctx, sqlQuery, email).Scan(&userData.Username, &userData.Password)
	if err != nil {
		return models.User{}, err
	}
	userData.Email = email
	return userData, nil
}
