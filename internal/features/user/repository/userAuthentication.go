package repository

import (
	"context"

	"golang.org/x/crypto/bcrypt"
)

func (u *UserRepository) RegisterUser(ctx context.Context, username, password, email string) error {
	// hash the password to bcrypt with 12 cost
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return err
	}

	sqlQuery := `
	INSERT INTO users(username, password, email)
	VALUES ($1, $2, $3);
`
	if _, err := u.pool.Exec(ctx, sqlQuery, username, string(hashedPassword), email); err != nil {
		return err
	}
	return nil
}
