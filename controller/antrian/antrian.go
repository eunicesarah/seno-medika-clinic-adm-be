package antrian

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"seno-medika.com/config/db"
	"seno-medika.com/model/common"
	"seno-medika.com/model/antrian"
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

