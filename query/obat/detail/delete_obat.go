package detail

import (
	"errors"
	"seno-medika.com/config/db"
)

func DeleteObatById(id int) error {
	val, err := db.DB.Exec("DELETE FROM obat WHERE obat_id = $1", id)
	if err != nil {
		return err
	}

	if rows, _ := val.RowsAffected(); rows == 0 {
		return errors.New("obat_id not found")
	}
	return nil
}

func DeleteObatByName(name string) error {
	val, err := db.DB.Exec("DELETE FROM obat WHERE nama_obat = $1", name)
	if err != nil {
		return err
	}

	if rows, _ := val.RowsAffected(); rows == 0 {
		return errors.New("uuid not found")
	}
	return nil
}
