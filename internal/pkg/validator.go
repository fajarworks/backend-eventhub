package pkg

import (
	"errors"
	"net/mail"
)

func ValidateLengthPass(pass string) error {
	if len(pass) < 8 {
		return errors.New("password must be more than 8 characters")
	}
	return nil
}

func ValidateEmailFormat(email string) bool {
	_, err := mail.ParseAddress(email)

	return err == nil

}
