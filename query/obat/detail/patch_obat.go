package detail

import (
	"errors"
	"seno-medika.com/config/db"
)

func PatchStockObatById(id int, stock int) error {
	val, err := db.DB.Exec(
		`UPDATE obat SET
			stock = $1
		WHERE obat_id = $2
		`,
		stock,
		id)

	if err != nil {
		return err
	}

	if rows, _ := val.RowsAffected(); rows == 0 {
		return errors.New("id obat not found")
	}

	return nil
}

func PatchStockObatByName(nama string, stock int) error {
	val, err := db.DB.Exec(
		`UPDATE obat SET
			stock = $1
		WHERE nama_obat = $2
		`,
		stock,
		nama)

	if err != nil {
		return err
	}

	if rows, _ := val.RowsAffected(); rows == 0 {
		return errors.New("nama obat not found")
	}

	return nil
}

func PatchHargaObatById(id int, harga int64) error {
	val, err := db.DB.Exec(
		`UPDATE obat SET
			harga = $1
		WHERE obat_id = $2
		`,
		harga,
		id)

	if err != nil {
		return err
	}

	if rows, _ := val.RowsAffected(); rows == 0 {
		return errors.New("id obat not found")
	}

	return nil
}

func PatchHargaObatByName(nama string, harga int64) error {
	val, err := db.DB.Exec(
		`UPDATE obat SET
			harga = $1
		WHERE nama_obat = $2
		`,
		harga,
		nama)

	if err != nil {
		return err
	}

	if rows, _ := val.RowsAffected(); rows == 0 {
		return errors.New("nama obat not found")
	}

	return nil
}
