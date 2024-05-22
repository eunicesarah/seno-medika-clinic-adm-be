package detail

import (
	"errors"
	"seno-medika.com/config/db"
	"seno-medika.com/model/station/pharmacystation"
)

func UpdateObatById(id int, obat pharmacystation.Obat) error {
	val, err := db.DB.Exec(
		`UPDATE obat SET
			nama_obat = $1,
			jenis_asuransi = $2,
			harga = $3,
			stock = $4,
			satuan = $5
		WHERE obat_id = $6
		`,
		obat.NamaObat,
		obat.JenisAsuransi,
		obat.Harga,
		obat.Stock,
		obat.Satuan,
		id)

	if err != nil {
		return err
	}

	if rows, _ := val.RowsAffected(); rows == 0 {
		return errors.New("id obat not found")
	}

	return nil
}

func UpdateObatByName(nama string, obat pharmacystation.Obat) error {
	val, err := db.DB.Exec(
		`UPDATE obat SET
			obat_id = $1,
			jenis_asuransi = $2,
			harga = $3,
			stock = $4,
			satuan = $5
		WHERE nama_obat = $6
		`,
		obat.ObatID,
		obat.JenisAsuransi,
		obat.Harga,
		obat.Stock,
		obat.Satuan,
		nama)

	if err != nil {
		return err
	}

	if rows, _ := val.RowsAffected(); rows == 0 {
		return errors.New("nama obat not found")
	}

	return nil
}
