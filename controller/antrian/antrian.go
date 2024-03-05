package antrian

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"seno-medika.com/config/db"
	"seno-medika.com/model/common"
	"seno-medika.com/model/antrian"
	"seno-medika.com/service/antrianpasien"
)

// TODO: get pasien from list
func AddAntrianOffline(c *gin.Context) {
	var antrian antrian.Antrian
	// var pasien antrian.Pasien

	if err := c.ShouldBind(&antrian); err != nil {
		c.JSON(http.StatusBadRequest, common.Response{
			Message:    err.Error(),
			Status:     "Bad Request",
			StatusCode: http.StatusBadRequest,
			Data:       nil,
		})
		return
	}

	antrian.CreatedAt = time.Now().Local().String()
	// antrian.PasienID = 

	_, err := db.DB.Exec(
		`INSERT INTO antrian (
			pasien_id,
			nomor_antrian,
			status,
			poli,
			instalasi,
			created_at
		) VALUES ($1, $2, $3, $4, $5, $6)`,
		antrian.PasienID,
		antrian.NomorAntrian,
		antrian.Status,
		antrian.Poli,
		antrian.Instalasi,
		antrian.CreatedAt)

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
}

func AddAntrianOnline(c *gin.Context) {
	var antrianOnline antrian.AntrianOnline
	var antrian antrian.Antrian
	
	if err := c.ShouldBind(&antrianOnline); err != nil {
		c.JSON(http.StatusBadRequest, common.Response{
			Message:    err.Error(),
			Status:     "Bad Request",
			StatusCode: http.StatusBadRequest,
			Data:       nil,
		})
		return
	}

	pasienID, err := antrianpasien.ConvertNIKtoPasienID(antrianOnline.NIK)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.Response{
			Message:    err.Error(),
			Status:     "Internal Server Error",
			StatusCode: http.StatusInternalServerError,
			Data:       nil,
		})
		return
	}

	antrian.PasienID = pasienID
	antrian.NomorAntrian = antrianOnline.NomorAntrian

	_, err = db.DB.Exec(
		`INSERT INTO antrian (
			pasien_id,
			nomor_antrian,
			status,
			poli,
			instalasi,
			created_at
		) VALUES ($1, $2, $3, $4, $5, $6)`,
		antrian.PasienID,
		antrian.NomorAntrian,
		antrian.Status,
		antrian.Poli,
		antrian.Instalasi,
		antrian.CreatedAt)

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

func GetAllAntrian(c *gin.Context) {
	var antrians []antrian.Antrian
	rows, err := db.DB.Query(
	`SELECT antrian_id, pasien_id, nomor_antrian, status, poli, instalasi, created_at FROM antrian`)

	if err != nil {
		c.JSON(http.StatusInternalServerError, common.Response{
			Message:    err.Error(),
			Status:     "Internal Server Error",
			StatusCode: http.StatusInternalServerError,
			Data:       nil,
		})
		return
	}
	
	for rows.Next() {
		var antrianItem antrian.Antrian
		err := rows.Scan(
			&antrianItem.AntrianID,
			&antrianItem.PasienID,
			&antrianItem.NomorAntrian,
			&antrianItem.Status,
			&antrianItem.Poli,
			&antrianItem.Instalasi,
			&antrianItem.CreatedAt)

		if err != nil {
			c.JSON(http.StatusInternalServerError, common.Response{
				Message:    err.Error(),
				Status:     "Internal Server Error",
				StatusCode: http.StatusInternalServerError,
				Data:       nil,
			})
			return
		}

		antrians = append(antrians, antrianItem)
	}

	c.JSON(http.StatusOK, common.Response{
		Message:    "Berhasil mendapatkan data antrian",
		Status:     "ok",
		StatusCode: http.StatusOK,
		Data:       antrians,
	})
}