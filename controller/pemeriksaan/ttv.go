package pemeriksaan

import (
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"seno-medika.com/config/db"
	"seno-medika.com/model/common"
	doctorstation2 "seno-medika.com/model/station/doctorstation"
	"seno-medika.com/model/station/nursestation"
	ttv2 "seno-medika.com/query/pemeriksaan/ttv"
	"seno-medika.com/model/doctorstation"
	"seno-medika.com/model/nursestation"
	"seno-medika.com/query/ttv"
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
		if err := db.DB.QueryRow("INSERT INTO skrining_awal(disabilitas, ambulansi, hambatan_komunikasi, jalan_tidak_seimbang, jalan_alat_bantu, menopang_saat_duduk, hasil_cara_jalan, skala_nyeri, nyeri_berulang, sifat_nyeri) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) RETURNING skrin_awal_id",
			ttvVar.SkriningAwal.Disabilitas, ttvVar.SkriningAwal.Ambulansi, ttvVar.SkriningAwal.HambatanKomunikasi, ttvVar.SkriningAwal.JalanTidakSeimbang, ttvVar.SkriningAwal.JalanAlatBantu, ttvVar.SkriningAwal.MenopangSaatDuduk, ttvVar.SkriningAwal.HasilCaraJalan, ttvVar.SkriningAwal.SkalaNyeri, ttvVar.SkriningAwal.NyeriBerulang, ttvVar.SkriningAwal.SifatNyeri).Scan(&skrinAwalId); err != nil {
			errChan <- err
			return
		}
	}()

	go func() {
		defer wg.Done()
		if err := db.DB.QueryRow("INSERT INTO skrining_gizi(penurunan_bb, tdk_nafsu_makan, diagnosis_khusus, nama_penyakit) VALUES ($1, $2, $3, $4) RETURNING skrin_gizi_id",
			ttvVar.SkriningGizi.PenurunanBB, ttvVar.SkriningGizi.TdkNafsuMakan, ttvVar.SkriningGizi.DiagnosisKhusus, ttvVar.SkriningGizi.NamaPenyakit).Scan(&skrinGiziId); err != nil {
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

func DeleteTTV(c *gin.Context) {
	changeBy := c.Query("delete_by")
	changeType := c.Query("delete_type")
	target := c.Query("target")

	if changeBy == "" || changeType == "" || target == "" {
		c.JSON(http.StatusBadRequest, common.Response{
			Message:    "change_by, change_type, and target are required",
			Status:     "Bad Request",
			StatusCode: http.StatusBadRequest,
			Data:       nil,
		})
		return
	}

	switch changeType {
	case "skrining_awal":
		switch changeBy {
		case "id":
			err := ttv2.DeleteSkriningAwalById(target)
			if err != nil {
				c.JSON(http.StatusInternalServerError, common.Response{
					Message:    err.Error(),
					Status:     "Internal Server Error",
					StatusCode: http.StatusInternalServerError,
					Data:       nil,
				})
				return
			}
		}
	case "skrining_gizi":
		switch changeBy {
		case "id":
			err := ttv2.DeleteSkriningGiziById(target)
			if err != nil {
				c.JSON(http.StatusInternalServerError, common.Response{
					Message:    err.Error(),
					Status:     "Internal Server Error",
					StatusCode: http.StatusInternalServerError,
					Data:       nil,
				})
				return
			}
		}
	case "skrining_penyakit":
		switch changeBy {
		case "id":
			err := ttv2.DeleteSkriningPenyakitById(target)
			if err != nil {
				c.JSON(http.StatusInternalServerError, common.Response{
					Message:    err.Error(),
					Status:     "Internal Server Error",
					StatusCode: http.StatusInternalServerError,
					Data:       nil,
				})
				return
			}
		}
	case "ttv":
		switch changeBy {
		case "id":
			err := ttv2.DeleteTTVById(target)
			if err != nil {
				c.JSON(http.StatusInternalServerError, common.Response{
					Message:    err.Error(),
					Status:     "Internal Server Error",
					StatusCode: http.StatusInternalServerError,
					Data:       nil,
				})
				return
			}
		}

	case "anamnesis":
		switch changeBy {
		case "id":
			err := ttv2.DeleteAnamnesisById(target)
			if err != nil {
				c.JSON(http.StatusInternalServerError, common.Response{
					Message:    err.Error(),
					Status:     "Internal Server Error",
					StatusCode: http.StatusInternalServerError,
					Data:       nil,
				})
				return
			}
		}
	case "alergi":
		switch changeBy {
		case "id":
			err := ttv2.DeleteAlergiById(target)
			if err != nil {
				c.JSON(http.StatusInternalServerError, common.Response{
					Message:    err.Error(),
					Status:     "Internal Server Error",
					StatusCode: http.StatusInternalServerError,
					Data:       nil,
				})
				return
			}
		}

	case "riwayat_penyakit":
		switch changeBy {
		case "id":
			err := ttv2.DeleteRiwayatPenyakitById(target)
			if err != nil {
				c.JSON(http.StatusInternalServerError, common.Response{
					Message:    err.Error(),
					Status:     "Internal Server Error",
					StatusCode: http.StatusInternalServerError,
					Data:       nil,
				})
				return
			}
		}
	default:
		c.JSON(http.StatusBadRequest, common.Response{
			Message:    "change_type not found",
			Status:     "Bad Request",
			StatusCode: http.StatusBadRequest,
			Data:       nil,
		})
		return
	}

	c.JSON(http.StatusOK, common.Response{
		Message:    "Successfully deleted ttv",
		Status:     "ok",
		StatusCode: http.StatusOK,
		Data:       nil,
	})
	return
}

func FindTTV(c *gin.Context) {
	findBy := c.Query("find_by")
	target := c.Query("target")

	switch findBy {
	case "id":
		ttvVar, err := ttv2.FindNurseStationById(target)
		if err != nil {
			c.JSON(http.StatusInternalServerError, common.Response{
				Message:    err.Error(),
				Status:     "Internal Server Error",
				StatusCode: http.StatusInternalServerError,
				Data:       nil,
			})
			return
		}

		c.JSON(http.StatusOK, common.Response{
			Message:    "Successfully get ttv",
			Status:     "ok",
			StatusCode: http.StatusOK,
			Data:       ttvVar,
		})
		return
	case "pasien_id":
		ttvVar, err := ttv.FindNurseStationByPasienId(target)
		if err != nil {
			c.JSON(http.StatusInternalServerError, common.Response{
				Message:    err.Error(),
				Status:     "Internal Server Error",
				StatusCode: http.StatusInternalServerError,
				Data:       nil,
			})
			return
		}

		c.JSON(http.StatusOK, common.Response{
			Message:    "Successfully get ttv",
			Status:     "ok",
			StatusCode: http.StatusOK,
			Data:       ttvVar,
		})
		return

	default:
		ttvVars, err := ttv2.FindAllNurseStation()
		if err != nil {
			c.JSON(http.StatusInternalServerError, common.Response{
				Message:    err.Error(),
				Status:     "Internal Server Error",
				StatusCode: http.StatusInternalServerError,
				Data:       nil,
			})
			return
		}

		c.JSON(http.StatusOK, common.Response{
			Message:    "Successfully get ttv",
			Status:     "ok",
			StatusCode: http.StatusOK,
			Data:       ttvVars,
		})
		return
	}
}

func PatchTTV(c *gin.Context) {
	changeType := c.Query("change_type")
	changeBy := c.Query("change_by")
	target := c.Query("target")

	if changeType == "" || changeBy == "" || target == "" {
		c.JSON(http.StatusBadRequest, common.Response{
			Message:    "change_type, change_by, and target are required",
			Status:     "Bad Request",
			StatusCode: http.StatusBadRequest,
			Data:       nil,
		})
		return
	}

	switch changeType {
	case "skrining_awal":
		switch changeBy {
		case "id":
			var skriningAwal nursestation.SkriningAwal
			if err := c.ShouldBind(&skriningAwal); err != nil {
				c.JSON(http.StatusBadRequest, common.Response{
					Message:    err.Error(),
					Status:     "Bad Request",
					StatusCode: http.StatusBadRequest,
					Data:       nil,
				})
				return
			}

			err := ttv2.ChangeSkriningAwalById(target, skriningAwal)
			if err != nil {
				c.JSON(http.StatusInternalServerError, common.Response{
					Message:    err.Error(),
					Status:     "Internal Server Error",
					StatusCode: http.StatusInternalServerError,
					Data:       nil,
				})
				return
			}
		}

	case "skrining_gizi":
		switch changeBy {
		case "id":
			var skriningGizi nursestation.SkriningGizi
			if err := c.ShouldBind(&skriningGizi); err != nil {
				c.JSON(http.StatusBadRequest, common.Response{
					Message:    err.Error(),
					Status:     "Bad Request",
					StatusCode: http.StatusBadRequest,
					Data:       nil,
				})
				return
			}

			err := ttv2.ChangeSkriningGiziById(target, skriningGizi)
			if err != nil {
				c.JSON(http.StatusInternalServerError, common.Response{
					Message:    err.Error(),
					Status:     "Internal Server Error",
					StatusCode: http.StatusInternalServerError,
					Data:       nil,
				})
				return
			}
		}

	case "ttv":
		switch changeBy {
		case "id":
			var ttvVar nursestation.TTV
			if err := c.ShouldBind(&ttvVar); err != nil {
				c.JSON(http.StatusBadRequest, common.Response{
					Message:    err.Error(),
					Status:     "Bad Request",
					StatusCode: http.StatusBadRequest,
					Data:       nil,
				})
				return
			}

			err := ttv2.ChangeTTVById(target, ttvVar)
			if err != nil {
				c.JSON(http.StatusInternalServerError, common.Response{
					Message:    err.Error(),
					Status:     "Internal Server Error",
					StatusCode: http.StatusInternalServerError,
					Data:       nil,
				})
				return
			}
		}

	case "anamnesis":
		switch changeBy {
		case "id":
			var anamnesis doctorstation2.Anamnesis
			if err := c.ShouldBind(&anamnesis); err != nil {
				c.JSON(http.StatusBadRequest, common.Response{
					Message:    err.Error(),
					Status:     "Bad Request",
					StatusCode: http.StatusBadRequest,
					Data:       nil,
				})
				return
			}

			err := ttv2.ChangeAnamnesisById(target, anamnesis)
			if err != nil {
				c.JSON(http.StatusInternalServerError, common.Response{
					Message:    err.Error(),
					Status:     "Internal Server Error",
					StatusCode: http.StatusInternalServerError,
					Data:       nil,
				})
				return
			}
		}

	case "alergi":
		switch changeBy {
		case "id":
			var alergi doctorstation2.Alergi
			if err := c.ShouldBind(&alergi); err != nil {
				c.JSON(http.StatusBadRequest, common.Response{
					Message:    err.Error(),
					Status:     "Bad Request",
					StatusCode: http.StatusBadRequest,
					Data:       nil,
				})
				return
			}

			err := ttv2.ChangeAlergiById(target, alergi)
			if err != nil {
				c.JSON(http.StatusInternalServerError, common.Response{
					Message:    err.Error(),
					Status:     "Internal Server Error",
					StatusCode: http.StatusInternalServerError,
					Data:       nil,
				})
				return
			}
		}

	case "riwayat_penyakit":
		switch changeBy {
		case "id":
			var riwayatPenyakit nursestation.RiwayatPenyakit
			if err := c.ShouldBind(&riwayatPenyakit); err != nil {
				c.JSON(http.StatusBadRequest, common.Response{
					Message:    err.Error(),
					Status:     "Bad Request",
					StatusCode: http.StatusBadRequest,
					Data:       nil,
				})
				return
			}

			err := ttv2.ChangeRiwayatPenyakitById(target, riwayatPenyakit)
			if err != nil {
				c.JSON(http.StatusInternalServerError, common.Response{
					Message:    err.Error(),
					Status:     "Internal Server Error",
					StatusCode: http.StatusInternalServerError,
					Data:       nil,
				})
				return
			}
		}

	default:
		c.JSON(http.StatusBadRequest, common.Response{
			Message:    "change_type not found",
			Status:     "Bad Request",
			StatusCode: http.StatusBadRequest,
			Data:       nil,
		})
		return

	}

	c.JSON(http.StatusOK, common.Response{
		Message:    "Successfully patched ttv",
		Status:     "ok",
		StatusCode: http.StatusOK,
		Data:       nil,
	})

	return
}
