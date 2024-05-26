package obat

import (
	"seno-medika.com/config/db"
	"seno-medika.com/model/station/pharmacystation"
)

func AddListObat(list pharmacystation.ListObat) error {
	if _, err := db.DB.Exec("INSERT INTO list_obat(resep_id, jumlah, dosis, aturan_pakai, keterangan, obat_id) VALUES ($1, $2, $3, $4, $5, $6)", list.ResepID, list.Jumlah, list.Dosis, list.AturanPakai, list.Keterangan, list.ObatID); err != nil {
		return err
	}

	return nil
}
