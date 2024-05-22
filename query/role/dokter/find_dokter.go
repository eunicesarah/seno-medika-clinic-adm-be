package dokter

import (
	"seno-medika.com/config/db"
	"seno-medika.com/model/person"
	"seno-medika.com/model/station/cashierstation"
	"sync"
)

func FindAllDokter() ([]person.Dokter, error) {
	var dokterVar []person.Dokter

	val, err := db.DB.Query("SELECT * FROM users WHERE role = 'dokter'")
	if err != nil {
		return nil, err
	}

	for val.Next() {
		var dokter person.Dokter
		if err := val.Scan(&dokter.UserID, &dokter.UserUUID, &dokter.Nama, &dokter.Password, &dokter.Email, &dokter.Role); err != nil {
			return nil, err
		}
		var wg sync.WaitGroup
		wg.Add(2)

		errChan := make(chan error, 3)
		go func() {
			defer wg.Done()
			val, err := db.DB.Query("SELECT * FROM dokter WHERE dokter_id = $1", dokter.UserID)
			if err != nil {
				return
			}
			if val.Next() {
				if err := val.Scan(&dokter.DokterData.DokterID, &dokter.DokterData.JagaPoliMana, &dokter.DokterData.JadwalJaga, &dokter.DokterData.NomorLisensi); err != nil {
					errChan <- err
					return
				}
			}
		}()

		go func() {
			defer wg.Done()
			val, err := db.DB.Query("SELECT * FROM list_jadwal_dokter WHERE dokter_id = $1", dokter.UserID)
			if err != nil {
				errChan <- err
				return
			}
			for val.Next() {
				var listJadwalDokter person.ListJadwalDokter
				if err := val.Scan(&listJadwalDokter.DokterID, &listJadwalDokter.Hari, &listJadwalDokter.Shift); err != nil {
					errChan <- err
					return
				}
				dokter.DokterData.ListJadwalDokter = append(dokter.DokterData.ListJadwalDokter, listJadwalDokter)
			}
		}()

		wg.Wait()
		close(errChan)
		for err := range errChan {
			if err != nil {
				return nil, err
			}
		}

		dokterVar = append(dokterVar, dokter)
	}

	return dokterVar, nil
}

func FindDokterByID(id string) (person.Dokter, error) {
	var dokter person.Dokter

	val, err := db.DB.Query("SELECT * FROM users WHERE user_id = $1 AND role = 'dokter'", id)
	if err != nil {
		return dokter, err
	}

	if val.Next() {
		if err := val.Scan(&dokter.UserID, &dokter.UserUUID, &dokter.Nama, &dokter.Password, &dokter.Email, &dokter.Role); err != nil {
			return dokter, err
		}
		var wg sync.WaitGroup
		wg.Add(2)

		errChan := make(chan error, 3)
		go func() {
			defer wg.Done()
			val, err := db.DB.Query("SELECT * FROM dokter WHERE dokter_id = $1", dokter.UserID)
			if err != nil {
				return
			}
			if val.Next() {
				if err := val.Scan(&dokter.DokterData.DokterID, &dokter.DokterData.JagaPoliMana, &dokter.DokterData.JadwalJaga, &dokter.DokterData.NomorLisensi); err != nil {
					errChan <- err
					return
				}
			}
		}()

		go func() {
			defer wg.Done()
			val, err := db.DB.Query("SELECT * FROM list_jadwal_dokter WHERE dokter_id = $1", dokter.UserID)
			if err != nil {
				errChan <- err
				return
			}
			for val.Next() {
				var listJadwalDokter person.ListJadwalDokter
				if err := val.Scan(&listJadwalDokter.DokterID, &listJadwalDokter.Hari, &listJadwalDokter.Shift); err != nil {
					errChan <- err
					return
				}
				dokter.DokterData.ListJadwalDokter = append(dokter.DokterData.ListJadwalDokter, listJadwalDokter)
			}
		}()

		wg.Wait()
		close(errChan)
		for err := range errChan {
			if err != nil {
				return dokter, err
			}
		}
	}

	return dokter, nil
}

func FindAllTindakan() ([]cashierstation.Tindakan, error) {
	var tindVal []cashierstation.Tindakan

	val, err := db.DB.Query("SELECT * FROM tindakan")
	if err != nil {
		return nil, err
	}

	for val.Next() {
		var temp cashierstation.Tindakan
		if err := val.Scan(
			&temp.TindakanID,
			&temp.ProsedurTindakan,
			&temp.Jumlah,
			&temp.Keterangan,
			&temp.TanggalRencana,
			&temp.HargaTindakan,
			&temp.IndikasiTindakan,
			&temp.Tujuan,
			&temp.Risiko,
			&temp.Komplikasi,
			&temp.AlternatifRisiko,
		); err != nil {
			return nil, err
		}

		tindVal = append(tindVal, temp)
	}

	return tindVal, nil
}
func FindTindakanById(id string) (cashierstation.Tindakan, error) {
	var tindVal cashierstation.Tindakan
	if err := db.DB.QueryRow("SELECT * FROM tindakan WHERE tindakan_id = $1", id).Scan(
		&tindVal.TindakanID,
		&tindVal.ProsedurTindakan,
		&tindVal.Jumlah,
		&tindVal.Keterangan,
		&tindVal.TanggalRencana,
		&tindVal.HargaTindakan,
		&tindVal.IndikasiTindakan,
		&tindVal.Tujuan,
		&tindVal.Risiko,
		&tindVal.Komplikasi,
		&tindVal.AlternatifRisiko,
	); err != nil {
		return tindVal, err
	}

	return tindVal, nil
}
