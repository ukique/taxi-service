package validation

import (
	"errors"
	"net/mail"
	"regexp"

	"github.com/ukique/taxi-service/metts-taxi/models"
)

func IsValidEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}
func IsOnlyEnglishLetters(username string) bool {
	re := regexp.MustCompile(`^[a-zA-Z]+$`)
	return re.MatchString(username)
}
func ParseCredentials(user models.User) error {
	if user.Username == "" {
		return errors.New("username can't be empty")
	}
	if !IsOnlyEnglishLetters(user.Username) {
		return errors.New("username must contain only English letters")
	}
	if user.Email == "" {
		return errors.New("email can't be empty")
	}
	if user.Password == "" {
		return errors.New("password can't be empty")
	}

	if len(user.Username) > 16 {
		return errors.New("username can't be more than 16 characters")
	}
	if len(user.Email) > 254 {
		return errors.New("incorrect email address")
	}
	if !IsValidEmail(user.Email) {
		return errors.New("this isn't email address")
	}
	if len(user.Password) > 72 {
		return errors.New("password is too long")
	}
	if len(user.Password) < 8 {
		return errors.New("password must be at least 8 characters")
	}
	return nil
}
