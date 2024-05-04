package dokter

import (
	"seno-medika.com/config/db"
	"seno-medika.com/model/station/doctorstation"
	"sync"
	"time"
)

func AddPemeriksaanDokterDefault(pemeriksaan doctorstation.PemeriksaanDokter) error {
	var (
		pasienId string
		wg       sync.WaitGroup
	)

	errChan := make(chan error, 2)

	go func() {
		defer wg.Done()
		if err := db.DB.QueryRow("INSERT INTO pemeriksaan_dokter (antrian_id, dokter_id, perawat_id) VALUES ($1, $2, $3) RETURNING pemeriksaan_dokter_id", pemeriksaan.AntrianId, pemeriksaan.DokterId, pemeriksaan.PerawatId).Scan(&pemeriksaan.PemeriksaanDokterId); err != nil {
			errChan <- err
			return
		}
	}()

	go func() {
		defer wg.Done()
		if err := db.DB.QueryRow("SELECT pasien_id FROM antrian WHERE antrian_id = $1", pemeriksaan.AntrianId).Scan(&pasienId); err != nil {
			errChan <- err
			return
		}
	}()

	wg.Wait()
	close(errChan)

	for val := range errChan {
		if val != nil {
			return val
		}
	}

	wg.Add(6)
	errChan = make(chan error, 6)

	go func() {
		defer wg.Done()
		if _, err := db.DB.Exec("INSERT INTO riwayat_pemeriksaan (pemeriksaan_dokter_id, pasien_id, tanggal) VALUES ($1, $2, $3)", pemeriksaan.PemeriksaanDokterId, pasienId, time.Now().Format("2006-01-02")); err != nil {
			errChan <- err
			return
		}
	}()

	go func() {
		defer wg.Done()
		if _, err := db.DB.Exec("INSERT INTO keadaan_fisik (pemeriksaan_dokter_id) VALUES ($1)", pemeriksaan.PemeriksaanDokterId); err != nil {
			errChan <- err
			return
		}
	}()

	go func() {
		defer wg.Done()
		if _, err := db.DB.Exec("INSERT INTO pemeriksaan_fisik (pemeriksaan_dokter_id) VALUES ($1)", pemeriksaan.PemeriksaanDokterId); err != nil {
			errChan <- err
			return
		}
	}()

	go func() {
		defer wg.Done()
		if _, err := db.DB.Exec("INSERT INTO diagnosa(pemeriksaan_dokter_id) VALUES ($1)", pemeriksaan.PemeriksaanDokterId); err != nil {
			errChan <- err
			return
		}
	}()

	go func() {
		defer wg.Done()
		if _, err := db.DB.Exec("INSERT INTO resep(pemeriksaan_dokter_id) VALUES ($1)", pemeriksaan.PemeriksaanDokterId); err != nil {
			errChan <- err
			return
		}
	}()

	go func() {
		defer wg.Done()
		if _, err := db.DB.Exec("INSERT INTO anatomi(pemeriksaan_dokter_id, pasien_id) VALUES ($1, $2)", pemeriksaan.PemeriksaanDokterId, pasienId); err != nil {
			errChan <- err
			return
		}
	}()

	wg.Wait()
	close(errChan)

	for val := range errChan {
		if val != nil {
			return val
		}
	}

	return nil
}
