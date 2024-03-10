package superadmin

import (
	"errors"
	"seno-medika.com/config/db"
	"seno-medika.com/model/person"
)

func UpdateUserByUuid(uid string, user person.User) error {
	val, err := db.DB.Exec(
		"UPDATE users SET nama = $1, email = $2, role = $3 WHERE user_uuid = $4",
		user.Nama, user.Email, user.Role, uid)

	if err != nil {
		return err
	}

	if rows, _ := val.RowsAffected(); rows == 0 {
		return errors.New("uuid not found")
	}
	return nil
}

func UpdateUserById(id int, user person.User) error {
	val, err := db.DB.Exec(
		"UPDATE users SET nama = $1, email = $2, role = $3 WHERE user_id = $4",
		user.Nama, user.Email, user.Role, id)

	if err != nil {
		return err
	}

	if rows, _ := val.RowsAffected(); rows == 0 {
		return errors.New("id not found")
	}
	return nil
}
