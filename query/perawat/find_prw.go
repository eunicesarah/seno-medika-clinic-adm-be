package perawat

import (
	"seno-medika.com/config/db"
	"seno-medika.com/model/person"
)

func FindAllPerawat() ([]person.Perawat, error) {
	var perawatVars []person.Perawat

	val, err := db.DB.Query("SELECT * FROM users WHERE role = 'perawat'")
	if err != nil {
		return nil, err
	}

	for val.Next() {
		var perawatVar person.Perawat
		if err := val.Scan(&perawatVar.UserID, &perawatVar.UserUUID, &perawatVar.Nama, &perawatVar.Password, &perawatVar.Email, &perawatVar.Role); err != nil {
			return nil, err
		}

		val, err := db.DB.Query("SELECT * FROM perawat WHERE perawat_id = $1", perawatVar.UserID)
		if err != nil {
			return nil, err
		}

		if val.Next() {
			if err := val.Scan(&perawatVar.PerawatData.PerawatID, &perawatVar.PerawatData.NomorLisensi); err != nil {
				return nil, err
			}
		}

		perawatVars = append(perawatVars, perawatVar)
	}

	return perawatVars, nil
}

func FindPerawatByID(id string) (person.Perawat, error) {
	var perawatVar person.Perawat

	val, err := db.DB.Query("SELECT * FROM users WHERE user_id = $1 AND role = 'perawat'", id)
	if err != nil {
		return person.Perawat{}, err
	}

	if val.Next() {
		if err := val.Scan(&perawatVar.UserID, &perawatVar.UserUUID, &perawatVar.Nama, &perawatVar.Password, &perawatVar.Email, &perawatVar.Role); err != nil {
			return person.Perawat{}, err
		}
	}

	val, err = db.DB.Query("SELECT * FROM perawat WHERE perawat_id = $1", perawatVar.UserID)
	if err != nil {
		return person.Perawat{}, err
	}

	if val.Next() {
		if err := val.Scan(&perawatVar.PerawatData.PerawatID, &perawatVar.PerawatData.NomorLisensi); err != nil {
			return person.Perawat{}, err
		}
	}

	return perawatVar, nil
}
