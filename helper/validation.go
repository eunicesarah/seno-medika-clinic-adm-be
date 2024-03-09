package helper

import (
	"errors"
	"github.com/asaskevich/govalidator"
)

func ValidationEmail(email string, err chan error) {
	if govalidator.IsNull(email) {
		err <- errors.New("email is null")
		return
	}

	if !govalidator.IsEmail(email) {
		err <- errors.New("Email isn't valid")
		return
	}

	return
}

func ValidationPassword(password string, err chan error) {
	if govalidator.IsNull(password) {
		err <- errors.New("password is null")
		return
	}

	if len(password) <= 6 {
		err <- errors.New("Password isn't valid")
		return
	}
	return
}
