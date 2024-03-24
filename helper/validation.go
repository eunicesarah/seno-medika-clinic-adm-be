package helper

import (
	"errors"
	"github.com/asaskevich/govalidator"
)

func ValidationEmail(email string, errorChan chan error) {
	if govalidator.IsNull(email) {
		errorChan <- errors.New("Email is null")
		return
	}

	if !govalidator.IsEmail(email) {
		errorChan <- errors.New("Email isn't valid")
		return
	}

	return
}

func ValidationPassword(password string, errorChan chan error) {
	if govalidator.IsNull(password) {
		errorChan <- errors.New("password is null")
	}

	if len(password) <= 6 {
		errorChan <- errors.New("Password isn't valid")
	}
	return
}
