package nurse

import (
	"seno-medika.com/config/db"
	"seno-medika.com/model/person"
	"seno-medika.com/helper"
	"time"
)


func FindDoctorsByName(name string) ([]person.Dokter, error) {
	var doctors []person.Dokter


	rows, err := db.DB.Query(`
		SELECT *
		FROM users u
		WHERE u.role = 'dokter' AND u.nama = $1`, name)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var dokter person.Dokter
		err := rows.Scan(&dokter.UserID, &dokter.UserUUID, &dokter.Nama, &dokter.Password, &dokter.Email, &dokter.Role)
		if err != nil {
			return nil, err
		}
		doctors = append(doctors, dokter)
	}
	return doctors, nil

}


func FindDoctorsByPoli(poli string) ([]person.Dokter, error) {
	var doctors []person.Dokter

	rows, err := db.DB.Query(`
		SELECT u.user_id, u.user_uuid, u.nama, u.password, u.email, u.role
		FROM users u
		JOIN dokter d ON u.user_id = d.dokter_id
		WHERE d.jaga_poli_mana = $1`, poli)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var dokter person.Dokter
		err := rows.Scan(&dokter.UserID, &dokter.UserUUID, &dokter.Nama, &dokter.Password, &dokter.Email, &dokter.Role)
		if err != nil {
			return nil, err
		}
		doctors = append(doctors, dokter)
	}
	return doctors, nil

}

func FindDoctorToday () ([]person.Dokter, error) {
	day := helper.TranslateDay(time.Now().Weekday().String())
	var doctors []person.Dokter

	rows, err := db.DB.Query(`
		SELECT u.user_id, u.user_uuid, u.nama, u.password, u.email, u.role
		FROM users u
		JOIN list_jadwal_dokter l
		ON l.dokter_id = u.user_id
		WHERE l.hari = ?`, 
	day)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var dokter person.Dokter
		err := rows.Scan(&dokter.UserID, &dokter.UserUUID, &dokter.Nama, &dokter.Password, &dokter.Email, &dokter.Role)
		if err != nil {
			return nil, err
		}
		doctors = append(doctors, dokter)
	}
	return doctors, nil
}

func FindDoctorsByShift(shift int) ([]person.Dokter, error) {
	var doctors []person.Dokter

	rows, err := db.DB.Query(`
		SELECT u.user_id, u.user_uuid, u.nama, u.password, u.email, u.role
		FROM users u
		JOIN list_jadwal_dokter l
		WHERE l.shift = $1`, shift)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var dokter person.Dokter
		err := rows.Scan(&dokter.UserID, &dokter.UserUUID, &dokter.Nama, &dokter.Password, &dokter.Email, &dokter.Role)
		if err != nil {
			return nil, err
		}
		doctors = append(doctors, dokter)
	}
	return doctors, nil

}

