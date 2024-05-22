package resep

import (
	"seno-medika.com/config/db"
	"seno-medika.com/model/station/pharmacystation"
)

func FindAllResep() ([]pharmacystation.Resep, error) {
	var (
		resVal []pharmacystation.Resep
	)

	rows, err := db.DB.Query("SELECT * FROM resep")
	if err != nil {
		return resVal, err
	}

	for rows.Next() {
		var temp pharmacystation.Resep
		if err := rows.Scan(temp.ResepID, temp.PemeriksaanDokterID, temp.Deskripsi, temp.RuangTujuan, temp.StatusObat); err != nil {
			return nil, err
		}
		resVal = append(resVal, temp)
	}

	return resVal, nil
}

func FindResepById(id string) (pharmacystation.Resep, error) {
	var resVal pharmacystation.Resep
	if err := db.DB.QueryRow("SELECT * FROM resep WHERE resep_id = $1", id).Scan(resVal.ResepID, resVal.PemeriksaanDokterID, resVal.Deskripsi, resVal.RuangTujuan, resVal.StatusObat); err != nil {
		return resVal, err
	}

	return resVal, nil
}
