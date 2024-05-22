package resep

import "seno-medika.com/config/db"

func DeleteResepById(id string) error {
	if _, err := db.DB.Exec("DELETE FROM resep WHERE resep_id = $1", id); err != nil {
		return err
	}
	return nil
}
