package superadmin

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"seno-medika.com/config/db"
)

func ChangePasswordById(id int, password string) error {
	hashPass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	val, err := db.DB.Exec("UPDATE users SET password = $1 WHERE user_id = $2", string(hashPass), id)
	if err != nil {
		return err
	}

	if rows, _ := val.RowsAffected(); rows == 0 {
		return errors.New("id not found")
	}
	return nil
}

func ChangePasswordByUuid(uid string, password string) error {
	hashPass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	val, err := db.DB.Exec("UPDATE users SET password = $1 WHERE user_uuid = $2", string(hashPass), uid)
	if err != nil {
		return err
	}

	if rows, _ := val.RowsAffected(); rows == 0 {
		return errors.New("uuid not found")
	}
	return nil
}

func ChangeNameById(id int, name string) error {
	val, err := db.DB.Exec("UPDATE users SET nama = $1 WHERE user_id = $2", name, id)
	if err != nil {
		return err
	}

	if rows, _ := val.RowsAffected(); rows == 0 {
		return errors.New("id not found")
	}
	return nil
}

func ChangeNameByUuid(uid string, name string) error {
	val, err := db.DB.Exec("UPDATE users SET nama = $1 WHERE user_uuid = $2", name, uid)
	if err != nil {
		return err
	}

	if rows, _ := val.RowsAffected(); rows == 0 {
		return errors.New("uuid not found")
	}
	return nil
}

func ChangeEmailById(id string, email string) error {
	val, err := db.DB.Exec("UPDATE users SET email = $1 WHERE user_id = $2", email, id)
	if err != nil {
		return err
	}

	if rows, _ := val.RowsAffected(); rows == 0 {
		return errors.New("id not found")
	}
	return nil
}

func ChangeEmailByUuid(uid string, email string) error {
	val, err := db.DB.Exec("UPDATE users SET email = $1 WHERE user_uuid = $2", email, uid)
	if err != nil {
		return err
	}

	if rows, _ := val.RowsAffected(); rows == 0 {
		return errors.New("uuid not found")
	}
	return nil
}

func ChangeRoleById(id int, role string) error {
	val, err := db.DB.Exec("UPDATE users SET role = $1 WHERE user_id = $2", role, id)
	if err != nil {
		return err
	}

	if rows, _ := val.RowsAffected(); rows == 0 {
		return errors.New("id not found")
	}
	return nil
}

func ChangeRoleByUuid(uid string, role string) error {
	val, err := db.DB.Exec("UPDATE users SET role = $1 WHERE user_uuid = $2", role, uid)
	if err != nil {
		return err
	}

	if rows, _ := val.RowsAffected(); rows == 0 {
		return errors.New("uuid not found")
	}
	return nil
}
