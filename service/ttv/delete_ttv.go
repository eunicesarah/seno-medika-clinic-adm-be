package ttv

import "seno-medika.com/config/db"

func DeleteSkriningAwalById(id string) error {
	if _, err := db.DB.Exec("DELETE FROM skrining_awal WHERE skrining_awal_id = $1", id); err != nil {
		return err
	}

	return nil
}

func DeleteSkriningGiziById(id string) error {
	if _, err := db.DB.Exec("DELETE FROM skrining_gizi WHERE skrining_gizi_id = $1", id); err != nil {
		return err
	}

	return nil
}

func DeleteSkriningPenyakitById(id string) error {
	if _, err := db.DB.Exec("DELETE FROM skrining_penyakit WHERE skrining_penyakit_id = $1", id); err != nil {
		return err
	}

	return nil
}

func DeleteTTVById(id string) error {
	if _, err := db.DB.Exec("DELETE FROM ttv WHERE ttv_id = $1", id); err != nil {
		return err
	}

	return nil
}

func DeleteAnamnesisById(id string) error {
	if _, err := db.DB.Exec("DELETE FROM anamnesis WHERE anamnesis_id = $1", id); err != nil {
		return err
	}

	return nil
}
