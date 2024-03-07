package antrian

import (
	"seno-medika.com/config/db"
	"seno-medika.com/model/antrian"
)

func FindAntrianById(id int) (antrian.Antrian, error) {
	var antrianO antrian.Antrian
	err := db.DB.QueryRow("SELECT * FROM antrian WHERE antrian_id = $1", id).
		Scan(&antrianO.AntrianID, &antrianO.PasienID, &antrianO.NomorAntrian, &antrianO.Status, &antrianO.Poli, &antrianO.Instalasi, &antrianO.CreatedAt)
	if err != nil {
		return antrian.Antrian{}, err
	}
	return antrianO, nil
}

func FindAntrianAll() ([]antrian.Antrian, error) {
	var antrianO []antrian.Antrian
	rows, err := db.DB.Query("SELECT * FROM antrian")
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var eachAntrian antrian.Antrian
		err := rows.Scan(&eachAntrian.AntrianID, &eachAntrian.PasienID, &eachAntrian.NomorAntrian, &eachAntrian.Status, &eachAntrian.Poli, &eachAntrian.Instalasi, &eachAntrian.CreatedAt)
		if err != nil {
			return nil, err
		}
		antrianO = append(antrianO, eachAntrian)
	}
	return antrianO, nil
}
