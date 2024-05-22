package resep

import (
	"errors"
	"seno-medika.com/config/db"
	"seno-medika.com/model/station/pharmacystation"
)

func PutResepById(id string, resep pharmacystation.Resep) error {
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
		return err
	}

	if rows, _ := val.RowsAffected(); rows == 0 {
		return errors.New("id obat not found")
	}

	return nil
}
