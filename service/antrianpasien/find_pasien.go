package antrianpasien

import (
	"seno-medika.com/config/db"
)

func FindIDbyNIK(nik string) (int, error) {
	var pasienID int
	err := db.DB.QueryRow("SELECT pasien_id FROM pasien WHERE nik = $1", nik).Scan(&pasienID)
	if err != nil {
		return 0, err
	}

	return pasienID, nil
}
