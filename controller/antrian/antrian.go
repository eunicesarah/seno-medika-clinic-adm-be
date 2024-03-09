package antrian

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"seno-medika.com/config/db"
	"seno-medika.com/model/antrian"
	"seno-medika.com/model/common"
	antrian2 "seno-medika.com/service/antrian"
	"strconv"
	"time"
)

func AddAntrian(c *gin.Context) {
	var antr antrian.Antrian

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

	_, err = db.DB.Exec("INSERT INTO antrian (pasien_id, nomor_antrian, status, poli, instalasi, created_at) VALUES ($1, $2, $3, $4, $5, $6)", antr.PasienID, antr.NomorAntrian, antr.Status, antr.Poli, antr.Instalasi, antr.CreatedAt)

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
