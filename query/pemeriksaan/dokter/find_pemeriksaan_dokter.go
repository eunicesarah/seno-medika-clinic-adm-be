package dokter

import (
	"seno-medika.com/config/db"
	"seno-medika.com/model/station/doctorstation"
	"sync"
)

func FindPemeriksaanDokterById(pemeriksaanId string) (doctorstation.PemeriksaanDokterMerge, error) {
	var (
		pemeriksaan doctorstation.PemeriksaanDokterMerge
		wg          sync.WaitGroup
	)

	wg.Add(6)
	errChan := make(chan error, 6)

	go func() {
		defer wg.Done()
		if err := db.DB.QueryRow("SELECT pemeriksaan_dokter_id, antrian_id, dokter_id, perawat_id FROM pemeriksaan_dokter WHERE pemeriksaan_dokter_id = $1", pemeriksaanId).Scan(&pemeriksaan.PemeriksaanDokter.PemeriksaanDokterId, &pemeriksaan.PemeriksaanDokter.AntrianId, &pemeriksaan.PemeriksaanDokter.DokterId, &pemeriksaan.PemeriksaanDokter.PerawatId); err != nil {
			errChan <- err
			return
		}
	}()

	go func() {
		defer wg.Done()
		if err := db.DB.QueryRow("SELECT riwayat_pemeriksaan_id, pasien_id, tanggal, pemeriksaan, keterangan FROM riwayat_pemeriksaan WHERE pemeriksaan_dokter_id = $1", pemeriksaanId).Scan(&pemeriksaan.RiwayatPemeriksaan.RiwayatPemeriksaanId, &pemeriksaan.RiwayatPemeriksaan.PasienId, &pemeriksaan.RiwayatPemeriksaan.Tanggal, &pemeriksaan.RiwayatPemeriksaan.Pemeriksaan, &pemeriksaan.RiwayatPemeriksaan.Keterangan); err != nil {
			errChan <- err
			return
		}
	}()

	go func() {
		defer wg.Done()
		if err := db.DB.QueryRow("SELECT pemeriksaan_fisik_id, terapi_yg_sdh_dilakukan, rencana_tindakan, tindakan_keperawatan, observasi, merokok, konsumsi_alkohol, kurang_sayur FROM pemeriksaan_fisik WHERE pemeriksaan_dokter_id = $1", pemeriksaanId).Scan(&pemeriksaan.PemeriksaanFisik.PemeriksaanFisikId, &pemeriksaan.PemeriksaanFisik.TerapiYgSdhDilakukan, &pemeriksaan.PemeriksaanFisik.RencanaTindakan, &pemeriksaan.PemeriksaanFisik.TindakanKeperawatan, &pemeriksaan.PemeriksaanFisik.Observasi, &pemeriksaan.PemeriksaanFisik.Merokok, &pemeriksaan.PemeriksaanFisik.KonsumsiAlkohol, &pemeriksaan.PemeriksaanFisik.KurangSayur); err != nil {
			errChan <- err
			return
		}
	}()

	go func() {
		defer wg.Done()
		if err := db.DB.QueryRow("SELECT keadaan_fisik_id, pemeriksaan_kulit, pemeriksaan_kuku, pemeriksaan_kepala, pemeriksaan_mata, pemeriksaan_telinga, pemeriksaan_hidung_sinus, pemeriksaan_mulut_bibir, pemeriksaan_leher, pemeriksaan_dada_punggung, pemeriksaan_kardiovaskuler, pemeriksaan_abdomen_perut, pemeriksaan_ekstremitas_atas, pemeriksaan_ekstremitas_bawah, pemeriksaan_genitalia_pria FROM keadaan_fisik WHERE pemeriksaan_dokter_id = $1", pemeriksaanId).Scan(&pemeriksaan.KeadaanFisik.KeadaanFisikId, &pemeriksaan.KeadaanFisik.PemeriksaanKulit, &pemeriksaan.KeadaanFisik.PemeriksaanKuku, &pemeriksaan.KeadaanFisik.PemeriksaanKepala, &pemeriksaan.KeadaanFisik.PemeriksaanMata, &pemeriksaan.KeadaanFisik.PemeriksaanTelinga, &pemeriksaan.KeadaanFisik.PemeriksaanHidungSinus, &pemeriksaan.KeadaanFisik.PemeriksaanMulutBibir, &pemeriksaan.KeadaanFisik.PemeriksaanLeher, &pemeriksaan.KeadaanFisik.PemeriksaanDadaPunggung, &pemeriksaan.KeadaanFisik.PemeriksaanKardiovaskuler, &pemeriksaan.KeadaanFisik.PemeriksaanAbdomenPerut, &pemeriksaan.KeadaanFisik.PemeriksaanEkstremitasAtas, &pemeriksaan.KeadaanFisik.PemeriksaanEkstremitasBawah, &pemeriksaan.KeadaanFisik.PemeriksaanGenitaliaPria); err != nil {
			errChan <- err
			return
		}
	}()

	go func() {
		defer wg.Done()
		if err := db.DB.QueryRow("SELECT diagnosa_id, diagnosa, jenis, kasus, status_diagnosis FROM diagnosa WHERE pemeriksaan_dokter_id = $1", pemeriksaanId).Scan(&pemeriksaan.Diagnosa.DiagnosaId, &pemeriksaan.Diagnosa.Diagnosa, &pemeriksaan.Diagnosa.Jenis, &pemeriksaan.Diagnosa.Kasus, &pemeriksaan.Diagnosa.StatusDiagnosis); err != nil {
			errChan <- err
			return
		}
	}()

	go func() {
		defer wg.Done()
		if err := db.DB.QueryRow("SELECT anatomi_id, pasien_id, bagian_tubuh, keterangan FROM anatomi WHERE pemeriksaan_dokter_id = $1", pemeriksaanId).Scan(&pemeriksaan.Anatomi.AnatomiId, &pemeriksaan.Anatomi.PasienId, &pemeriksaan.Anatomi.BagianTubuh, &pemeriksaan.Anatomi.Keterangan); err != nil {
			errChan <- err
			return
		}
	}()

	wg.Wait()
	close(errChan)

	for val := range errChan {
		if val != nil {
			return pemeriksaan, val
		}
	}

	return pemeriksaan, nil
}

func FindPemeriksaanDokterByAntrianId(antrianId int) (doctorstation.PemeriksaanDokterMerge, error) {
	var (
		pemeriksaanId string
		pemeriksaan   doctorstation.PemeriksaanDokterMerge
	)

	if err := db.DB.QueryRow("SELECT pemeriksaan_dokter_id FROM pemeriksaan_dokter WHERE antrian_id = $1", antrianId).Scan(&pemeriksaanId); err != nil {
		return pemeriksaan, err
	}

	return FindPemeriksaanDokterById(pemeriksaanId)
}
