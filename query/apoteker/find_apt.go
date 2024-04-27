package apoteker

import (
	"seno-medika.com/config/db"
	"seno-medika.com/model/person"
)

func FindAllApoteker() ([]person.Apoteker, error) {
	var apotekerVars []person.Apoteker

	val, err := db.DB.Query("SELECT * FROM users WHERE role = 'apoteker'")
	if err != nil {
		return nil, err
	}

	for val.Next() {
		var apotekerVar person.Apoteker
		if err := val.Scan(&apotekerVar.UserID, &apotekerVar.UserUUID, &apotekerVar.Nama, &apotekerVar.Password, &apotekerVar.Email, &apotekerVar.Role); err != nil {
			return nil, err
		}

		val, err := db.DB.Query("SELECT * FROM apoteker WHERE apoteker_id = $1", apotekerVar.UserID)
		if err != nil {
			return nil, err
		}

		if val.Next() {
			if err := val.Scan(&apotekerVar.ApotekerData.ApotekerID, &apotekerVar.ApotekerData.NomorLisensi); err != nil {
				return nil, err
			}
		}

		apotekerVars = append(apotekerVars, apotekerVar)
	}

	return apotekerVars, nil
}

func FindApotekerByID(id string) (person.Apoteker, error) {
	var apotekerVar person.Apoteker

	val, err := db.DB.Query("SELECT * FROM users WHERE user_id = $1 AND role = 'apoteker'", id)
	if err != nil {
		return person.Apoteker{}, err
	}

	if val.Next() {
		if err := val.Scan(&apotekerVar.UserID, &apotekerVar.UserUUID, &apotekerVar.Nama, &apotekerVar.Password, &apotekerVar.Email, &apotekerVar.Role); err != nil {
			return person.Apoteker{}, err
		}
	}

	val, err = db.DB.Query("SELECT * FROM apoteker WHERE apoteker_id = $1", apotekerVar.UserID)
	if err != nil {
		return person.Apoteker{}, err
	}

	if val.Next() {
		if err := val.Scan(&apotekerVar.ApotekerData.ApotekerID, &apotekerVar.ApotekerData.NomorLisensi); err != nil {
			return person.Apoteker{}, err
		}
	}

	return apotekerVar, nil
}
