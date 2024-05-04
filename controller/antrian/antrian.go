package antrian

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"seno-medika.com/config/db"
	"seno-medika.com/model/antrian"
	"seno-medika.com/model/common"
	antrian2 "seno-medika.com/query/antrian"
	antrian3 "seno-medika.com/query/nurse"
)

type filterResponse struct {
	Antrian []antrian.AntrianNurse `json:"antrian"`
	Size    int                    `json:"size"`
}

func AddAntrian(c *gin.Context) {
	var antr antrian.PendaftaranAntrian

	if err := c.ShouldBind(&antr); err != nil {
		c.JSON(http.StatusBadRequest, common.Response{
			Message:    err.Error(),
			Status:     "Bad Request",
			StatusCode: http.StatusBadRequest,
			Data:       nil,
		})
		return
	}

	var count int

	check := db.DB.QueryRow("SELECT pasien_id FROM pasien WHERE no_erm = $1 and nama = $2", antr.NoERM, antr.Nama).Scan(&antr.PasienID)
	if check != nil {
		c.JSON(http.StatusBadRequest, common.Response{
			Message:    "Pasien tidak ditemukan",
			Status:     "Bad Request",
			StatusCode: http.StatusBadRequest,
			Data:       nil,
		})
		return
	}

	err := db.DB.QueryRow("SELECT COUNT(*) FROM pasien WHERE pasien_id = $1", antr.PasienID).Scan(&count)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.Response{
			Message:    err.Error(),
			Status:     "Internal Server Error",
			StatusCode: http.StatusInternalServerError,
			Data:       nil,
		})
		return
	}

	if count == 0 {
		c.JSON(http.StatusBadRequest, common.Response{
			Message:    "Pasien tidak ditemukan",
			Status:     "Bad Request",
			StatusCode: http.StatusBadRequest,
			Data:       nil,
		})
		return
	}

	var jumlahAntrian int

	err = db.DB.QueryRow("SELECT COUNT(*) FROM antrian WHERE created_at = $1", time.Now().Local().Format("2006-01-02")).Scan(&jumlahAntrian)

	if err != nil {
		c.JSON(http.StatusInternalServerError, common.Response{
			Message:    err.Error(),
			Status:     "Internal Server Error",
			StatusCode: http.StatusInternalServerError,
			Data:       nil,
		})
		return
	}

	antr.NomorAntrian = jumlahAntrian + 1
	antr.CreatedAt = time.Now().Local().Format("2006-01-02")
	// antr.Status = "false"

	_, err = db.DB.Exec("INSERT INTO antrian (pasien_id, nomor_antrian, poli, instalasi, created_at) VALUES ($1, $2, $3, $4, $5)", antr.PasienID, antr.NomorAntrian, antr.Poli, antr.Instalasi, antr.CreatedAt)

	if err != nil {
		c.JSON(http.StatusInternalServerError, common.Response{
			Message:    err.Error(),
			Status:     "Internal Server Error",
			StatusCode: http.StatusInternalServerError,
			Data:       nil,
		})
		return
	}

	c.JSON(http.StatusCreated, common.Response{
		Message:    "Antrian berhasil ditambahkan",
		Status:     "Created",
		StatusCode: http.StatusOK,
		Data:       nil,
	})
	return
}

func DeleteAntrian(c *gin.Context) {
	id := c.Query("id")

	_, err := db.DB.Exec(
		`DELETE FROM antrian WHERE antrian_id = $1`,
		id)

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
		Message:    "Antrian berhasil dihapus",
		Status:     "ok",
		StatusCode: http.StatusOK,
		Data:       nil,
	})
}

func GetAntrian(c *gin.Context) {
	var target = c.Query("target")
	var findBy = c.Query("find_by")

	if findBy == "dashboard" {
		poli := c.Query("poli")
		limit := c.Query("limit")
		page := c.Query("page")
		date := c.Query("date")
		search := c.Query("search")

		if limit == "" {
			limit = "10"
		}

		if page == "" {
			page = "0"
		} else {
			val, _ := strconv.Atoi(page)
			lim, _ := strconv.Atoi(limit)
			page = strconv.Itoa(val*lim - lim)
		}

		if date == "" {
			date = time.Now().Local().Format("2006-01-02")
		}

		data, size, err := antrian2.FindAntrianFilter(search, page, limit, date, poli)

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
			Message:    "Successfully get antrian",
			Status:     "ok",
			StatusCode: http.StatusOK,
			Data: filterResponse{
				Antrian: data,
				Size:    size,
			},
		})
		return
	}

	if findBy == "id" {
		val, err := strconv.Atoi(target)
		if err != nil {
			c.JSON(http.StatusBadRequest, common.Response{
				Message:    "Invalid target",
				Status:     "Bad Request",
				StatusCode: http.StatusBadRequest,
				Data:       nil,
			})
			return
		}
		data, err := antrian2.FindAntrianById(val)
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
			Message:    "Successfully get antrian",
			Status:     "ok",
			StatusCode: http.StatusOK,
			Data:       data,
		})
		return
	}
	if findBy == "pasienid" {
		val, err := strconv.Atoi(target)
		if err != nil {
			c.JSON(http.StatusBadRequest, common.Response{
				Message:    "Invalid target",
				Status:     "Bad Request",
				StatusCode: http.StatusBadRequest,
				Data:       nil,
			})
			return
		}
		data, err := antrian2.FindAntrianByPasienId(val)
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
			Message:    "Successfully get antrian",
			Status:     "ok",
			StatusCode: http.StatusOK,
			Data:       data,
		})
		return
	}

	if findBy == "doktername" {
		data, err := antrian3.FindAntrianByDoctorName(target)

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
			Message:    "Successfully get antrian",
			Status:     "ok",
			StatusCode: http.StatusOK,
			Data:       data,
		})
		return
	}

	if findBy == "dokterpoli" {
		data, err := antrian3.FindAntrianByDoctorPoli(target)
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
			Message:    "Successfully get antrian",
			Status:     "ok",
			StatusCode: http.StatusOK,
			Data:       data,
		})
		return
	}

	if findBy == "poli" {
		data, err := antrian3.FindAntrianByPoli(target)
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
			Message:    "Successfully get antrian",
			Status:     "ok",
			StatusCode: http.StatusOK,
			Data:       data,
		})
		return
	}

	if findBy == "shift" {
		val, err := strconv.Atoi(target)
		if err != nil {
			c.JSON(http.StatusBadRequest, common.Response{
				Message:    "Invalid target",
				Status:     "Bad Request",
				StatusCode: http.StatusBadRequest,
				Data:       nil,
			})
			return
		}

		data, err := antrian3.FindAntrianByDoctorShift(val)

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
			Message:    "Successfully get antrian",
			Status:     "ok",
			StatusCode: http.StatusOK,
			Data:       data,
		})
		return
	}

	// if (findBy == "day") {
	// 	data, err := antrian3.FindAntrianToday()

	// 	if err != nil {
	// 		c.JSON(http.StatusInternalServerError, common.Response{
	// 			Message:    err.Error(),
	// 			Status:     "Internal Server Error",
	// 			StatusCode: http.StatusInternalServerError,
	// 			Data:       nil,
	// 		})
	// 		return
	// 	}

	// 	c.JSON(http.StatusOK, common.Response{
	// 		Message:    "Successfully get antrian",
	// 		Status:     "ok",
	// 		StatusCode: http.StatusOK,
	// 		Data:       data,
	// 	})
	// 	return
	// }

	antrianList, err := antrian2.FindAntrianAll()

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
		Message:    "Successfully get antrian",
		Status:     "ok",
		StatusCode: http.StatusOK,
		Data:       antrianList,
	})

	return
}

func PatchAntrian(c *gin.Context) {
	var patchInput common.PatchInput

	changeBy := c.Query("change_by")
	changeType := c.Query("change_type")

	if err := c.ShouldBind(&patchInput); err != nil {
		c.JSON(http.StatusBadRequest, common.Response{
			Message:    err.Error(),
			Status:     "Bad Request",
			StatusCode: http.StatusBadRequest,
			Data:       nil,
		})
		return
	}

	switch changeType {
	case "status":

		switch changeBy {
		case "id":
			err := antrian2.ChangeStatusAntrianById(patchInput.Key.(int), patchInput.Value.(string))

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
				Message:    "Successfully update antrian",
				Status:     "ok",
				StatusCode: http.StatusOK,
				Data:       nil,
			})
			return

		case "poli":
			err := antrian2.ChangeStatusByPoli(patchInput.Key.(string), patchInput.Value.(string))

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
				Message:    "Successfully update antrian",
				Status:     "ok",
				StatusCode: http.StatusOK,
				Data:       nil,
			})
			return

		case "instalasi":
			err := antrian2.ChangeStatusByInstalasi(patchInput.Key.(string), patchInput.Value.(string))

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
				Message:    "Successfully update antrian",
				Status:     "ok",
				StatusCode: http.StatusOK,
				Data:       nil,
			})
			return

		default:
			c.JSON(http.StatusBadRequest, common.Response{
				Message:    "Invalid change_by",
				Status:     "Bad Request",
				StatusCode: http.StatusBadRequest,
				Data:       nil,
			})
			return
		}
	}
}
