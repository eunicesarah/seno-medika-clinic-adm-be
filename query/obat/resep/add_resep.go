package resep

import (
	"seno-medika.com/config/db"
	"seno-medika.com/model/station/pharmacystation"
)

func AddResep(resep pharmacystation.Resep) error {
	if _, err := db.DB.Exec("INSERT INTO resep(pemeriksaan_dokter_id, deskripsi, ruang_tujuan, status_obat) VALUES ($1, $2, $3, $4)", resep.PemeriksaanDokterID, resep.Deskripsi, resep.RuangTujuan, resep.StatusObat); err != nil {
		return err
	}
	return nil
}
