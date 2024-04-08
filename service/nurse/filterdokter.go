package nurse

import (
	"seno-medika.com/config/db"
	"seno-medika.com/model/antrian"
	"seno-medika.com/helper"
	"time"
)


func FindAntrianByDoctorName(name string) ([]antrian.Antrian, error) {
	var antrianVar []antrian.Antrian

	rows, err := db.DB.Query("SELECT * FROM antrian a JOIN pasien p ON p.pasien_id = a.pasien_id join anamnesis am on am.pasien_id = p.pasien_id join dokter d on am.dokter_id = d.dokter_id join users u on d.dokter_id = u.user_id WHERE u.nama = $1 and u.role = 'dokter" , name)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var antrianO antrian.Antrian
		err := rows.Scan(&antrianO.AntrianID, &antrianO.PasienID, &antrianO.NomorAntrian, &antrianO.Status, &antrianO.Poli, &antrianO.Instalasi, &antrianO.CreatedAt)
		if err != nil {
			return nil, err
		}
		antrianVar = append(antrianVar, antrianO)
	}

	return antrianVar, nil

}





func FindAntrianByDoctorPoli(poli string) ([]antrian.Antrian, error) {
	var antrianVar []antrian.Antrian

	rows, err := db.DB.Query("SELECT * FROM antrian a JOIN pasien p ON p.pasien_id = a.pasien_id join anamnesis am on am.pasien_id = p.pasien_id join dokter d on am.dokter_id = d.dokter_id join users u on d.dokter_id = u.user_id WHERE u.role = 'dokter' and d.jaga_poli_mana = $1 " ,poli)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var antrianO antrian.Antrian
		err := rows.Scan(&antrianO.AntrianID, &antrianO.PasienID, &antrianO.NomorAntrian, &antrianO.Status, &antrianO.Poli, &antrianO.Instalasi, &antrianO.CreatedAt)
		if err != nil {
			return nil, err
		}
		antrianVar = append(antrianVar, antrianO)
	}
	return antrianVar, nil

}

func FindAntrianToday () ([]antrian.Antrian, error) {
	day := helper.TranslateDay(time.Now().Weekday().String())
	var antrianVar []antrian.Antrian

	rows, err := db.DB.Query(`
		SELECT * FROM antrian a 
		JOIN pasien p ON p.pasien_id = a.pasien_id
		join anamnesis am on am.pasien_id = p.pasien_id 
		join dokter d on am.dokter_id = d.dokter_id 
		join users u on d.dokter_id = u.user_id 
		join list_jadwal_dokter l ON l.dokter_id = u.user_id  
		WHERE u.role = 'dokter' and d.jaga_poli_mana = $1 and l.hari = ?`,
	day)


	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var antrianO antrian.Antrian
		err := rows.Scan(&antrianO.AntrianID, &antrianO.PasienID, &antrianO.NomorAntrian, &antrianO.Status, &antrianO.Poli, &antrianO.Instalasi, &antrianO.CreatedAt)
		if err != nil {
			return nil, err
		}
		antrianVar = append(antrianVar, antrianO)
	}
	return antrianVar, nil
}

func FindAntrianByDoctorShift(shift int) ([]antrian.Antrian, error) {
	var antrianVar []antrian.Antrian

	rows, err := db.DB.Query(`
	SELECT * FROM antrian a 
	JOIN pasien p ON p.pasien_id = a.pasien_id
	join anamnesis am on am.pasien_id = p.pasien_id 
	join dokter d on am.dokter_id = d.dokter_id 
	join users u on d.dokter_id = u.user_id 
	join list_jadwal_dokter l ON l.dokter_id = u.user_id  
	WHERE u.role = 'dokter' and d.jaga_poli_mana = $1 and l.shift = ?`, shift)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var antrianO antrian.Antrian
		err := rows.Scan(&antrianO.AntrianID, &antrianO.PasienID, &antrianO.NomorAntrian, &antrianO.Status, &antrianO.Poli, &antrianO.Instalasi, &antrianO.CreatedAt)
		if err != nil {
			return nil, err
		}
		antrianVar = append(antrianVar, antrianO)
	}
	return antrianVar, nil

}

