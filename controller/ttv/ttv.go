package ttv

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"seno-medika.com/config/db"
	"seno-medika.com/model/common"
	"seno-medika.com/model/doctorstation"
	"seno-medika.com/model/nursestation"
	"seno-medika.com/service/ttv"
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
		if err := db.DB.QueryRow("INSERT INTO skrining_awal(disabilitas, ambulansi, hambatan_komunikasi, jalan_tidak_seimbang, jalan_alat_bantu, menopang_saat_duduk, hasil_cara_jalan, skala_nyeri, nyeri_berulang) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING skrin_awal_id",
			ttvVar.SkriningAwal.Disabilitas, ttvVar.SkriningAwal.Ambulansi, ttvVar.SkriningAwal.HambatanKomunikasi, ttvVar.SkriningAwal.JalanTidakSeimbang, ttvVar.SkriningAwal.JalanAlatBantu, ttvVar.SkriningAwal.MenopangSaatDuduk, ttvVar.SkriningAwal.HasilCaraJalan, ttvVar.SkriningAwal.SkalaNyeri, ttvVar.SkriningAwal.NyeriBerulang).Scan(&skrinAwalId); err != nil {
			errChan <- err
			return
		}
	}()

	go func() {
		defer wg.Done()
		if err := db.DB.QueryRow("INSERT INTO skrining_gizi(penurunan_bb, tdk_nafsu_makan, diagnosis_khusus, nama_penyakit, skala_nyeri, nyeri_berulang, sifat_nyeri) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING skrin_gizi_id",
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
			err := ttv.DeleteSkriningAwalById(target)
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
			err := ttv.DeleteSkriningGiziById(target)
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
			err := ttv.DeleteSkriningPenyakitById(target)
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
			err := ttv.DeleteTTVById(target)
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
			err := ttv.DeleteAnamnesisById(target)
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
			err := ttv.DeleteAlergiById(target)
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
			err := ttv.DeleteRiwayatPenyakitById(target)
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
	findType := c.Query("find_type")
	target := c.Query("target")

	switch findType {
	case "skrining_awal":
		switch findBy {
		case "id":
			skriningAwal, err := ttv.FindSkriningAwalById(target)
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
				Message:    "Successfully found skrining_awal",
				Status:     "ok",
				StatusCode: http.StatusOK,
				Data:       skriningAwal,
			})
			return

		default:
			skriningAwal, err := ttv.FindAllSkriningAwal()
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
				Message:    "Successfully found skrining_awal",
				Status:     "ok",
				StatusCode: http.StatusOK,
				Data:       skriningAwal,
			})
			return
		}

	case "skrining_gizi":
		switch findBy {
		case "id":
			skriningGizi, err := ttv.FindSkriningGiziById(target)
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
				Message:    "Successfully found skrining_gizi",
				Status:     "ok",
				StatusCode: http.StatusOK,
				Data:       skriningGizi,
			})
			return

		default:
			skriningGizi, err := ttv.FindAllSkriningGizi()
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
				Message:    "Successfully found skrining_gizi",
				Status:     "ok",
				StatusCode: http.StatusOK,
				Data:       skriningGizi,
			})
			return
		}

	case "ttv":
		switch findBy {
		case "id":
			ttv, err := ttv.FindTTVById(target)
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
				Message:    "Successfully found ttv",
				Status:     "ok",
				StatusCode: http.StatusOK,
				Data:       ttv,
			})
			return

		default:
			ttv, err := ttv.FindAllTTV()
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
				Message:    "Successfully found ttv",
				Status:     "ok",
				StatusCode: http.StatusOK,
				Data:       ttv,
			})
			return
		}

	case "alergi":
		switch findBy {
		case "id":
			alergi, err := ttv.FindAlergiById(target)
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
				Message:    "Successfully found alergi",
				Status:     "ok",
				StatusCode: http.StatusOK,
				Data:       alergi,
			})
			return

		default:
			alergi, err := ttv.FindAllAlergi()
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
				Message:    "Successfully found alergi",
				Status:     "ok",
				StatusCode: http.StatusOK,
				Data:       alergi,
			})
			return
		}

	case "riwayat_penyakit":
		switch findBy {
		case "id":
			riwayatPenyakit, err := ttv.FindRiwayatPenyakitById(target)
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
				Message:    "Successfully found riwayat_penyakit",
				Status:     "ok",
				StatusCode: http.StatusOK,
				Data:       riwayatPenyakit,
			})
			return

		default:
			riwayatPenyakit, err := ttv.FindAllRiwayatPenyakit()
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
				Message:    "Successfully found riwayat_penyakit",
				Status:     "ok",
				StatusCode: http.StatusOK,
				Data:       riwayatPenyakit,
			})
			return
		}

	case "anamnesis":
		switch findBy {
		case "id":
			anamnesis, err := ttv.FindAnamnesisById(target)
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
				Message:    "Successfully found anamnesis",
				Status:     "ok",
				StatusCode: http.StatusOK,
				Data:       anamnesis,
			})
			return

		default:
			anamnesis, err := ttv.FindAllAnamnesis()
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
				Message:    "Successfully found anamnesis",
				Status:     "ok",
				StatusCode: http.StatusOK,
				Data:       anamnesis,
			})

			return
		}

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

			err := ttv.ChangeSkriningAwalById(target, skriningAwal)
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

			err := ttv.ChangeSkriningGiziById(target, skriningGizi)
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

			err := ttv.ChangeTTVById(target, ttvVar)
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
			var anamnesis doctorstation.Anamnesis
			if err := c.ShouldBind(&anamnesis); err != nil {
				c.JSON(http.StatusBadRequest, common.Response{
					Message:    err.Error(),
					Status:     "Bad Request",
					StatusCode: http.StatusBadRequest,
					Data:       nil,
				})
				return
			}

			err := ttv.ChangeAnamnesisById(target, anamnesis)
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
			var alergi doctorstation.Alergi
			if err := c.ShouldBind(&alergi); err != nil {
				c.JSON(http.StatusBadRequest, common.Response{
					Message:    err.Error(),
					Status:     "Bad Request",
					StatusCode: http.StatusBadRequest,
					Data:       nil,
				})
				return
			}

			err := ttv.ChangeAlergiById(target, alergi)
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

			err := ttv.ChangeRiwayatPenyakitById(target, riwayatPenyakit)
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
