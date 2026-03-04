package repository

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"golang.org/x/crypto/bcrypt"
)

func RegisterDriver(ctx context.Context, conn *pgx.Conn, username, password, email string) error {
	// hash the password to bcrypt with 12 cost
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return fmt.Errorf("fail to hash password: %w", err)
	}

	sqlQuery := `
	INSERT INTO drivers(username, password, email)
	VALUES ($1, $2, $3);
`
	if _, err := conn.Exec(ctx, sqlQuery, username, string(hashedPassword), email); err != nil {
		return fmt.Errorf("fail to create driver: %w", err)
	}
	return nil
}
