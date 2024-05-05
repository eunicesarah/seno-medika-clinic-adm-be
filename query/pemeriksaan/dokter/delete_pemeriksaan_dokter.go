package dokter

import "seno-medika.com/config/db"

func DeletePemeriksaanDokter(id string) error {
	_, err := db.DB.Exec("DELETE FROM pemeriksaan_dokter WHERE pemeriksaan_dokter_id = $1", id)
	if err != nil {
		return err
	}

	return nil
}
