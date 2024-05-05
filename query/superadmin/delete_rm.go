package superadmin

import (
	"errors"
	"seno-medika.com/config/db"
)

func DeleteUserById(id int) error {
	val, err := db.DB.Exec("DELETE FROM users WHERE user_id = $1", id)
	if err != nil {
		return err
	}

	if rows, _ := val.RowsAffected(); rows == 0 {
		return errors.New("user_id not found")
	}
	return nil
}

func DeleteUserByUuid(uid string) error {
	val, err := db.DB.Exec("DELETE FROM users WHERE user_uuid = $1", uid)
	if err != nil {
		return err
	}

	if rows, _ := val.RowsAffected(); rows == 0 {
		return errors.New("uuid not found")
	}
	return nil
}

func DeleteUserByName(name string) error {
	val, err := db.DB.Exec("DELETE FROM users WHERE nama = $1", name)
	if err != nil {
		return err
	}

	if rows, _ := val.RowsAffected(); rows == 0 {
		return errors.New("name not found")
	}
	return nil
}

func DeleteUserByEmail(email string) error {
	val, err := db.DB.Exec("DELETE FROM users WHERE email = $1", email)
	if err != nil {
		return err
	}

	if rows, _ := val.RowsAffected(); rows == 0 {
		return errors.New("email not found")
	}
	return nil
}

func DeleteUserByRole(role string) error {
	val, err := db.DB.Exec("DELETE FROM users WHERE role = $1", role)
	if err != nil {
		return err
	}

	if rows, _ := val.RowsAffected(); rows == 0 {
		return errors.New("role not found")
	}
	return nil
}
