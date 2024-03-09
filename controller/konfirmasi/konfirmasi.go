package konfirmasi

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"seno-medika.com/config/db"
	"seno-medika.com/model/antrian"
	"seno-medika.com/model/common"
)

func ConfirmPasien(c *gin.Context) {
	var antrianPasien antrian.Antrian
	if err := c.ShouldBind(&antrianPasien); err != nil {
		c.JSON(http.StatusBadRequest, common.Response{
			Message:    err.Error(),
			Status:     "Bad Request",
			StatusCode: http.StatusBadRequest,
			Data:       nil,
		})
		return
	}

	_, err := db.DB.Exec(
		`UPDATE antrian
        SET status = true
        WHERE antrian_id = $1`,
		antrianPasien.AntrianID,
	)

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
		Message:    "Successfully confirm pasien",
		Status:     "ok",
		StatusCode: http.StatusOK,
		Data:       nil,
	})
}
