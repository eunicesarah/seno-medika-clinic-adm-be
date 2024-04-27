package dokter

import "seno-medika.com/config/db"

func DeleteListJadwalById(id string) error {
	if _, err := db.DB.Exec("DELETE FROM jadwal WHERE dokter_id = $1", id); err != nil {
		return err
	}
	return nil
}

func DeleteDokterById(id string) error {
	if _, err := db.DB.Exec("DELETE FROM dokter WHERE dokter_id = $1", id); err != nil {
		return err
	}
	return nil
}
