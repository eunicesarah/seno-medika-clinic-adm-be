package kasir

import (
	"errors"
	"seno-medika.com/config/db"
)

func UpdateMetodePembayaran(nota_id int, metode_pembayaran string) error {
	var id int
	err := db.DB.QueryRow("SELECT nota_id FROM nota WHERE nota_id = $1", nota_id).Scan(&id)
	if err != nil {
		return errors.New("nota_id not found")
	}

	_, err = db.DB.Exec("UPDATE nota SET metode_pembayaran = $1 WHERE nota_id = $2", metode_pembayaran, nota_id)
	if err != nil {
		return err
	}
	return nil
}

func UpdateTotalBiaya(nota_id int, total_biaya int64) error {
	var id int
	err := db.DB.QueryRow("SELECT nota_id FROM nota WHERE nota_id = $1", nota_id).Scan(&id)
	if err != nil {
		return errors.New("nota_id not found")
	}

	_, err = db.DB.Exec("UPDATE nota SET total_biaya = $1 WHERE nota_id = $2", total_biaya, nota_id)
	if err != nil {
		return err
	}
	return nil
}
