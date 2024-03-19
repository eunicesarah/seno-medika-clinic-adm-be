package helper

import (
	"errors"
	"seno-medika.com/config/db"
)

func IsEmailExists(email string, errorChan chan error) {
	val, err := db.DB.Query("SELECT email FROM users WHERE email = $1", email)
	if err != nil {
		errorChan <- err
		return

	}
	if val.Next() {
		errorChan <- errors.New("Email already exists")
		return
	}
	return
}
