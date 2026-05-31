package repository

import (
	"context"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func (u *UserRepository) RegisterUser(ctx context.Context, username, password, email string) error {
	// hash the password to bcrypt with 12 cost
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return err
	}

	sqlQuery := `
	INSERT INTO users(username, password, email, created_at)
	VALUES ($1, $2, $3, $4);
`
	if _, err := u.pool.Exec(ctx, sqlQuery, username, string(hashedPassword), email, time.Now()); err != nil {
		return err
	}
	return nil
}
