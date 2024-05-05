package pasien

import (
	"errors"
	"seno-medika.com/config/db"
)

func DeletePasienById(id int) error {
	val, err := db.DB.Exec("DELETE FROM pasien WHERE pasien_id = $1", id)
	if err != nil {
		return err
	}

	if rows, _ := val.RowsAffected(); rows == 0 {
		return errors.New("pasien_id not found")
	}
	return nil
}

func DeletePasienByUuid(uid string) error {
	val, err := db.DB.Exec("DELETE FROM pasien WHERE pasien_uuid = $1", uid)
	if err != nil {
		return err
	}

	if rows, _ := val.RowsAffected(); rows == 0 {
		return errors.New("uuid not found")
	}
	return nil
}
