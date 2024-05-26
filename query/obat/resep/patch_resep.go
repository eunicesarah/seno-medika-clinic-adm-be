package resep

import (
	"errors"
	"seno-medika.com/config/db"
	"seno-medika.com/model/station/pharmacystation"
)

func PutResepById(id string, resep pharmacystation.Resep) (pharmacystation.Resep, error) {
	val, err := db.DB.Exec(
		`UPDATE resep SET 
                 pemeriksaan_dokter_id = $1,
                 deskripsi = $2,
                 ruang_tujuan = $3,
                 status_obat = $4
                 WHERE resep_id = $5`,
		resep.PemeriksaanDokterID,
		resep.Deskripsi,
		resep.RuangTujuan,
		resep.StatusObat,
		id)

	if err != nil {
		return pharmacystation.Resep{}, err
	}

	if rows, _ := val.RowsAffected(); rows == 0 {
		return pharmacystation.Resep{}, errors.New("id obat not found")
	}

	// Fetch the updated record from the database
	var updatedResep pharmacystation.Resep
	err = db.DB.QueryRow(
		`SELECT resep_id, pemeriksaan_dokter_id, deskripsi, ruang_tujuan, status_obat 
                 FROM resep WHERE resep_id = $1`, id).Scan(
		&updatedResep.ResepID,
		&updatedResep.PemeriksaanDokterID,
		&updatedResep.Deskripsi,
		&updatedResep.RuangTujuan,
		&updatedResep.StatusObat)

	if err != nil {
		return pharmacystation.Resep{}, err
	}

	return updatedResep, nil
}
