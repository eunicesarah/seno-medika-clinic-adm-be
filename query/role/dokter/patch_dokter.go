package dokter

import (
	"errors"
	"seno-medika.com/config/db"
	"seno-medika.com/model/person"
)

func ChangeDokterById(id string, dokterData person.DokterData) error {
	val, err := db.DB.Exec("UPDATE dokter SET jaga_poli_mana = $1, jadwal_jaga = $2, nomor_lisensi = $3 WHERE dokter_id = $4",
		dokterData.JagaPoliMana, dokterData.JadwalJaga, dokterData.NomorLisensi, id)

	if err != nil {
		return err
	}

	if rows, _ := val.RowsAffected(); rows == 0 {
		return errors.New("Not changes anything")
	}

	return nil
}

func ChangeListJadwalById(id string, dokter person.ListJadwalDokter) error {
	val, err := db.DB.Exec("UPDATE list_jadwal_dokter SET hari = $1, shift = $2 WHERE dokter_id = $3",
		dokter.Hari, dokter.Shift, id)

	if err != nil {
		return err
	}

	if rows, _ := val.RowsAffected(); rows == 0 {
		return errors.New("Not change anything")
	}

	return nil
}
