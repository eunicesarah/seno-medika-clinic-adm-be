package helper

import "seno-medika.com/config/db"

func IsEmailExists(email string, errChan chan error) {
	val, err := db.DB.Query("SELECT email FROM user WHERE email = %1", email)
	if err != nil {
		errChan <- err
		return
	}
	if val.Next() {
		errChan <- nil
		return
	}
	errChan <- nil
	return
}
