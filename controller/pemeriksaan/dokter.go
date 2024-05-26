package pemeriksaan

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"seno-medika.com/model/common"
	"seno-medika.com/model/station/doctorstation"
	"seno-medika.com/query/pemeriksaan/dokter"
)

func AddPemeriksaanDokter(c *gin.Context) {
	var (
		pemeriksaan doctorstation.PemeriksaanDokter
	)

	if err := c.ShouldBindJSON(&pemeriksaan); err != nil {
		c.JSON(http.StatusBadRequest, common.Response{
			Message:    err.Error(),
			Status:     "Bad Request",
			StatusCode: http.StatusBadRequest,
			Data:       nil,
		})
		return
	}

	if err := dokter.AddPemeriksaanDokterDefault(pemeriksaan); err != nil {
		c.JSON(http.StatusInternalServerError, common.Response{
			Message:    err.Error(),
			Status:     "Internal Server Error",
			StatusCode: http.StatusInternalServerError,
			Data:       nil,
		})
		return
	}

	c.JSON(http.StatusCreated, common.Response{
		Message:    "Successfully add pemeriksaan dokter",
		Status:     "ok",
		StatusCode: http.StatusCreated,
		Data:       nil,
	})
	return
}

func AddListAnatomiCtr(c *gin.Context) {
	var (
		res []doctorstation.Anatomi
	)

	if err := c.ShouldBind(&res); err != nil {
		c.JSON(http.StatusBadRequest, common.Response{
			Message:    err.Error(),
			Status:     "Bad Request",
			StatusCode: http.StatusBadRequest,
			Data:       nil,
		})
		return
	}

	if err := dokter.AddListAnatomi(res); err != nil {
		c.JSON(http.StatusInternalServerError, common.Response{
			Message:    err.Error(),
			Status:     "Internal Service Error",
			StatusCode: http.StatusInternalServerError,
			Data:       nil,
		})
		return
	}

	c.JSON(http.StatusCreated, common.Response{
		Message:    "Successfully add list",
		Status:     "ok",
		StatusCode: http.StatusCreated,
		Data:       nil,
	})
	return
}

func DeletePemeriksaanDokter(c *gin.Context) {
	deleteBy := c.Query("delete_by")
	target := c.Query("target")

	switch deleteBy {
	case "id":
		if err := dokter.DeletePemeriksaanDokter(target); err != nil {
			c.JSON(http.StatusInternalServerError, common.Response{
				Message:    err.Error(),
				Status:     "Internal Server Error",
				StatusCode: http.StatusInternalServerError,
				Data:       nil,
			})
			return
		}

		c.JSON(http.StatusOK, common.Response{
			Message:    "Successfully delete pemeriksaan dokter",
			Status:     "ok",
			StatusCode: http.StatusOK,
			Data:       nil,
		})
		return
	}
}

func PatchPemeriksaanDokter(c *gin.Context) {
	updateBy := c.Query("update_by")
	updateType := c.Query("update_type")
	target := c.Query("target")

	switch updateBy {
	case "antrian_id":
		switch updateType {
		case "riwayat_pemeriksaan":
			var (
				riwayatPemeriksaan doctorstation.RiwayatPemeriksaan
			)

			if err := c.ShouldBindJSON(&riwayatPemeriksaan); err != nil {
				c.JSON(http.StatusBadRequest, common.Response{
					Message:    err.Error(),
					Status:     "Bad Request",
					StatusCode: http.StatusBadRequest,
					Data:       nil,
				})
				return
			}

			if err := dokter.PatchRiwayatPemeriksaan(target, riwayatPemeriksaan); err != nil {
				c.JSON(http.StatusInternalServerError, common.Response{
					Message:    err.Error(),
					Status:     "Internal Server Error",
					StatusCode: http.StatusInternalServerError,
					Data:       nil,
				})
				return
			}

			c.JSON(http.StatusOK, common.Response{
				Message:    "Successfully update riwayat pemeriksaan",
				Status:     "ok",
				StatusCode: http.StatusOK,
				Data:       nil,
			})
			return

		case "keadaan_fisik":
			var (
				keadaanFisik doctorstation.KeadaanFisik
			)

			if err := c.ShouldBindJSON(&keadaanFisik); err != nil {
				c.JSON(http.StatusBadRequest, common.Response{
					Message:    err.Error(),
					Status:     "Bad Request",
					StatusCode: http.StatusBadRequest,
					Data:       nil,
				})
				return
			}

			if err := dokter.PatchKeadaanFisik(target, keadaanFisik); err != nil {
				c.JSON(http.StatusInternalServerError, common.Response{
					Message:    err.Error(),
					Status:     "Internal Server Error",
					StatusCode: http.StatusInternalServerError,
					Data:       nil,
				})
				return
			}

			c.JSON(http.StatusOK, common.Response{
				Message:    "Successfully update keadaan fisik",
				Status:     "ok",
				StatusCode: http.StatusOK,
				Data:       nil,
			})
			return

		case "diagnosa":
			var (
				diagnosa doctorstation.Diagnosa
			)

			if err := c.ShouldBindJSON(&diagnosa); err != nil {
				c.JSON(http.StatusBadRequest, common.Response{
					Message:    err.Error(),
					Status:     "Bad Request",
					StatusCode: http.StatusBadRequest,
					Data:       nil,
				})
				return
			}

			if err := dokter.PatchDiagnosa(target, diagnosa); err != nil {
				c.JSON(http.StatusInternalServerError, common.Response{
					Message:    err.Error(),
					Status:     "Internal Server Error",
					StatusCode: http.StatusInternalServerError,
					Data:       nil,
				})
				return
			}

			c.JSON(http.StatusOK, common.Response{
				Message:    "Successfully update diagnosa",
				Status:     "ok",
				StatusCode: http.StatusOK,
				Data:       nil,
			})
			return

		case "anatomi":
			var (
				anatomi doctorstation.Anatomi
			)

			if err := c.ShouldBindJSON(&anatomi); err != nil {
				c.JSON(http.StatusBadRequest, common.Response{
					Message:    err.Error(),
					Status:     "Bad Request",
					StatusCode: http.StatusBadRequest,
					Data:       nil,
				})
				return
			}

			if err := dokter.PatchAnatomi(target, anatomi); err != nil {
				c.JSON(http.StatusInternalServerError, common.Response{
					Message:    err.Error(),
					Status:     "Internal Server Error",
					StatusCode: http.StatusInternalServerError,
					Data:       nil,
				})
				return
			}

			c.JSON(http.StatusOK, common.Response{
				Message:    "Successfully update anatomi",
				Status:     "ok",
				StatusCode: http.StatusOK,
				Data:       nil,
			})
			return

		case "pemeriksaan_fisik":
			var (
				pemeriksaanFisik doctorstation.PemeriksaanFisik
			)

			if err := c.ShouldBindJSON(&pemeriksaanFisik); err != nil {
				c.JSON(http.StatusBadRequest, common.Response{
					Message:    err.Error(),
					Status:     "Bad Request",
					StatusCode: http.StatusBadRequest,
					Data:       nil,
				})
				return
			}

			if err := dokter.PatchPemeriksaanFisik(target, pemeriksaanFisik); err != nil {
				c.JSON(http.StatusInternalServerError, common.Response{
					Message:    err.Error(),
					Status:     "Internal Server Error",
					StatusCode: http.StatusInternalServerError,
					Data:       nil,
				})
				return
			}

			c.JSON(http.StatusOK, common.Response{
				Message:    "Successfully update pemeriksaan fisik",
				Status:     "ok",
				StatusCode: http.StatusOK,
				Data:       nil,
			})
			return

		}
	}
}

func GetPemeriksaanDokter(c *gin.Context) {
	findBy := c.Query("find_by")
	target := c.Query("target")

	switch findBy {
	case "id":
		data, err := dokter.FindPemeriksaanDokterById(target)

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
			Message:    "Successfully get pemeriksaan dokter",
			Status:     "ok",
			StatusCode: http.StatusOK,
			Data:       data,
		})
		return

	case "antrian_id":
		data, err := dokter.FindPemeriksaanDokterByAntrianId(target)

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
			Message:    "Successfully get pemeriksaan dokter",
			Status:     "ok",
			StatusCode: http.StatusOK,
			Data:       data,
		})

		return
	}
}
