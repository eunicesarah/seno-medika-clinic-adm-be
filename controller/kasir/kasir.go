package kasir

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"seno-medika.com/config/db"
	"seno-medika.com/model/cashierstation"
	"seno-medika.com/model/common"
	"seno-medika.com/query/kasir"
)

func GetNota(c *gin.Context) {
	find_by := c.Query("find_by")
	target := c.Query("target")

	switch find_by {
	case "id":
		val, _ := strconv.Atoi(target)
		nota, err := kasir.FindNotaById(val)
		if err != nil {
			c.JSON(http.StatusBadRequest, common.Response{
				Message:    "Invalid target",
				Status:     "Bad Request",
				StatusCode: http.StatusBadRequest,
				Data:       nil,
			})
			return
		}
		c.JSON(http.StatusOK, common.Response{
			Message:    "Successfully get nota",
			Status:     "ok",
			StatusCode: http.StatusOK,
			Data:       nota,
		})
		return

	case "pasien_id":
		val, _ := strconv.Atoi(target)
		nota, err := kasir.FindNotaByPasienID(val)
		if err != nil {
			c.JSON(http.StatusBadRequest, common.Response{
				Message:    "Invalid target",
				Status:     "Bad Request",
				StatusCode: http.StatusBadRequest,
				Data:       nil,
			})
			return
		}
		c.JSON(http.StatusOK, common.Response{
			Message:    "Successfully get nota",
			Status:     "ok",
			StatusCode: http.StatusOK,
			Data:       nota,
		})
		return

	case "resep_id":
		val, _ := strconv.Atoi(target)
		nota, err := kasir.FindNotaByResepId(val)
		if err != nil {
			c.JSON(http.StatusBadRequest, common.Response{
				Message:    "Invalid target",
				Status:     "Bad Request",
				StatusCode: http.StatusBadRequest,
				Data:       nil,
			})
			return
		}
		c.JSON(http.StatusOK, common.Response{
			Message:    "Successfully get nota",
			Status:     "ok",
			StatusCode: http.StatusOK,
			Data:       nota,
		})
		return

	case "metode_pembayaran":
		nota, err := kasir.FindNotaByMetodePembayaran(target)
		if err != nil {
			c.JSON(http.StatusBadRequest, common.Response{
				Message:    "Invalid target",
				Status:     "Bad Request",
				StatusCode: http.StatusBadRequest,
				Data:       nil,
			})
			return
		}
		c.JSON(http.StatusOK, common.Response{
			Message:    "Successfully get nota",
			Status:     "ok",
			StatusCode: http.StatusOK,
			Data:       nota,
		})
		return

	case "detail":
		val, _ := strconv.Atoi(target)
		nota, err := kasir.FindDetailByResepId(val)
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
			Message:    "Successfully get detail nota",
			Status:     "ok",
			StatusCode: http.StatusOK,
			Data:       nota,
		})
		return

	default:
		nota, err := kasir.FindNotaAll()
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
			Message:    "Successfully get nota",
			Status:     "ok",
			StatusCode: http.StatusOK,
			Data:       nota,
		})
		return
	}
}

func AddNota(c *gin.Context) {
	var notaVar cashierstation.Nota

	if err := c.ShouldBindJSON(&notaVar); err != nil {
		c.JSON(http.StatusBadRequest, common.Response{
			Message:    err.Error(),
			Status:     "Bad Request",
			StatusCode: http.StatusBadRequest,
			Data:       nil,
		})
		return
	}

	_, err := db.DB.Exec(`INSERT INTO nota (pasien_id, dokter_id, resep_id, total_biaya, metode_pembayaran) VALUES ($1, $2, $3, $4, $5)`,
		notaVar.PasienID,
		notaVar.DokterID,
		notaVar.ResepID,
		notaVar.TotalBiaya,
		notaVar.MetodePembayaran,
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
		Message:    "Successfully insert nota",
		Status:     "ok",
		StatusCode: http.StatusOK,
		Data:       nil,
	})
}
