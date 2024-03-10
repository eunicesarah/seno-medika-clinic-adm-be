package antrian

import (
	"seno-medika.com/config/db"
	"seno-medika.com/model/antrian"
)

func UpdateStatusAntrianById(antrianID int, ant antrian.Antrian) error {
	_, err := db.DB.Exec(
		`UPDATE antrian
        SET status = true
        WHERE antrian_id = $1`,
		antrianID,
	)
	if err != nil {
		return err
	}
	return nil
}
