package obat

import (
	"seno-medika.com/config/db"
	"seno-medika.com/model/pharmacystation"
)

func FindObatById(id int) (pharmacystation.Obat, error) {
	var obatVar pharmacystation.Obat

	err := db.DB.QueryRow("SELECT * FROM obat WHERE obat_id = $1", id).Scan(
		&obatVar.ObatID,
		&obatVar.NamaObat,
		&obatVar.JenisAsuransi,
		&obatVar.Harga,
	)
	if err != nil {
		return pharmacystation.Obat{}, err
	}

	return obatVar, nil
}
func FindObatByName(name string) (pharmacystation.Obat, error) {
	var obatVar pharmacystation.Obat

	err := db.DB.QueryRow("SELECT * FROM obat WHERE nama_obat = $1", name).Scan(
		&obatVar.ObatID,
		&obatVar.NamaObat,
		&obatVar.JenisAsuransi,
		&obatVar.Harga,
	)
	if err != nil {
		return pharmacystation.Obat{}, err
	}

	return obatVar, nil
}
func FindObatAll() ([]pharmacystation.Obat, error) {
	var obatVar []pharmacystation.Obat

	val, err := db.DB.Query("SELECT * FROM obat")
	if err != nil {
		return nil, err
	}
	for val.Next() {
		var eachObat pharmacystation.Obat
		err := val.Scan(
			&eachObat.ObatID,
			&eachObat.NamaObat,
			&eachObat.JenisAsuransi,
			&eachObat.Harga,
		)

		if err != nil {
			return nil, err
		}
		obatVar = append(obatVar, eachObat)
	}

	return obatVar, nil
}