package ttv

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"seno-medika.com/config/db"
	"seno-medika.com/model/common"
	"seno-medika.com/model/nursestation"
	"sync"
)

func AddTTV(c *gin.Context) {
	var (
		ttvVar            nursestation.NurseStation
		wg                sync.WaitGroup
		ttvId             int
		alergiId          int
		riwayatPenyakitId int
		skrinGiziId       int
		skrinAwalId       int
	)

	if err := c.ShouldBind(&ttvVar); err != nil {
		c.JSON(http.StatusBadRequest, common.Response{
			Message:    err.Error(),
			Status:     "Bad Request",
			StatusCode: http.StatusBadRequest,
			Data:       nil,
		})
		return
	}

	errChan := make(chan error, 5)
	wg.Add(5)

	go func() {
		defer wg.Done()
		if err := db.DB.QueryRow("INSERT INTO skrining_awal(disabilitas, ambulalnsi, hambatan_komunikasi, jalan_tidak_seimbang, jalan_alat_bantu, menopang_saat_duduk, hasil_cara_jalan, skala_nyeri, nyeri_berulang) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING skrining_awal_id",
			ttvVar.SkriningAwal.Disabilitas, ttvVar.SkriningAwal.Ambulansi, ttvVar.SkriningAwal.HambatanKomunikasi, ttvVar.SkriningAwal.JalanTidakSeimbang, ttvVar.SkriningAwal.JalanAlatBantu, ttvVar.SkriningAwal.MenopangSaatDuduk, ttvVar.SkriningAwal.HasilCaraJalan, ttvVar.SkriningAwal.SkalaNyeri, ttvVar.SkriningAwal.NyeriBerulang).Scan(&skrinAwalId); err != nil {
			errChan <- err
			return
		}
	}()

	go func() {
		defer wg.Done()
		if err := db.DB.QueryRow("INSERT INTO skrining_gizi(penurunan_bb, tdk_nafsu_makan, diagnosis_khusus, nama_penyakit, skala_nyeri, nyeri_berulang, sifat_nyeri) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING skrining_gizi_id",
			ttvVar.SkriningGizi.PenurunanBB, ttvVar.SkriningGizi.TdkNafsuMakan, ttvVar.SkriningGizi.DiagnosisKhusus, ttvVar.SkriningGizi.NamaPenyakit, ttvVar.SkriningGizi.SkalaNyeri, ttvVar.SkriningGizi.NyeriBerulang, ttvVar.SkriningGizi.SifatNyeri).Scan(&skrinGiziId); err != nil {
			errChan <- err
			return
		}
	}()

	go func() {
		defer wg.Done()
		if err := db.DB.QueryRow("INSERT INTO riwayat_penyakit(rps, rpd, rpk) VALUES ($1, $2, $3) RETURNING riwayat_penyakit_id",
			ttvVar.RiwayatPenyakit.RPS, ttvVar.RiwayatPenyakit.RPD, ttvVar.RiwayatPenyakit.RPK).Scan(&riwayatPenyakitId); err != nil {
			errChan <- err
			return
		}
	}()

	go func() {
		defer wg.Done()
		if err := db.DB.QueryRow("INSERT INTO alergi(obat, makanan, lainnya) VALUES ($1, $2, $3) RETURNING alergi_id",
			ttvVar.Alergi.Obat, ttvVar.Alergi.Makanan, ttvVar.Alergi.Lainnya).Scan(&alergiId); err != nil {
			errChan <- err
			return
		}
	}()

	go func() {
		defer wg.Done()
		if err := db.DB.QueryRow("INSERT INTO ttv(kesadaran, sistole, diastole, tinggi_badan, cara_ukur_tb, berat_badan, lingkar_perut, detak_nadi, nafas, saturasi, suhu, detak_jantung, triage, psikolososial_spirit, keterangan) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15) RETURNING ttv_id",
			ttvVar.TTV.Kesadaran, ttvVar.TTV.Sistole, ttvVar.TTV.Diastole, ttvVar.TTV.TinggiBadan, ttvVar.TTV.CaraUkurTB, ttvVar.TTV.BeratBadan, ttvVar.TTV.LingkarPerut, ttvVar.TTV.DetakNadi, ttvVar.TTV.Nafas, ttvVar.TTV.Saturasi, ttvVar.TTV.Suhu, ttvVar.TTV.DetakJantung, ttvVar.TTV.Triage, ttvVar.TTV.PsikolososialSpirit, ttvVar.TTV.Keterangan).Scan(&ttvId); err != nil {
			errChan <- err
			return
		}
	}()

	wg.Wait()
	close(errChan)

	for val := range errChan {
		if val != nil {
			c.JSON(http.StatusInternalServerError, common.Response{
				Message:    val.Error(),
				Status:     "Internal Server Error",
				StatusCode: http.StatusInternalServerError,
				Data:       nil,
			})
			return
		}
	}

	if _, err := db.DB.Exec("INSERT INTO anamnesis(pasien_id, skrin_awal_id, skrin_gizi_id, ttv_id, riwayat_penyakit_id, alergi_id, dokter_id, perawat_id, keluhan_utama, keluhan_tambahan, lama_sakit) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11)",
		ttvVar.Anamnesis.PasienID, skrinAwalId, skrinGiziId, ttvId, riwayatPenyakitId, alergiId, ttvVar.Anamnesis.DokterID, ttvVar.Anamnesis.PerawatID, ttvVar.Anamnesis.KeluhanUtama, ttvVar.Anamnesis.KeluhanTambahan, ttvVar.Anamnesis.LamaSakit); err != nil {
		c.JSON(http.StatusInternalServerError, common.Response{
			Message:    err.Error(),
			Status:     "Internal Server Error",
			StatusCode: http.StatusInternalServerError,
			Data:       nil,
		})
		return
	}

	c.JSON(http.StatusOK, common.Response{
		Message:    "Successfully created ttv",
		Status:     "ok",
		StatusCode: http.StatusOK,
		Data:       nil,
	})
	return
}
