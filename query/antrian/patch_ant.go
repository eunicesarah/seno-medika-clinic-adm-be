package antrian

import (
	"seno-medika.com/config/db"
)

func ChangeStatusAntrianById(antrianID string, status string) error {
	_, err := db.DB.Exec(
		`UPDATE antrian
        SET status = $1
        WHERE antrian_id = $2`,
		status, antrianID,
	)
	if err != nil {
		return err
	}
	return nil
}

func ChangeStatusByPoli(poli string, status string) error {
	_, err := db.DB.Exec(
		`UPDATE antrian
		SET status = $1
		WHERE poli = $2`,
		status, poli,
	)
	if err != nil {
		return err
	}
	return nil
}

func ChangeStatusByInstalasi(instalasi string, status string) error {
	_, err := db.DB.Exec(
		`UPDATE antrian
		SET status = $1
		WHERE instalasi = $2`,
		status, instalasi,
	)
	if err != nil {
		return err
	}
	return nil
}
