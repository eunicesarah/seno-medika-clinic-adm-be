package role

import (
	"net/http"
	kasir2 "seno-medika.com/query/role/kasir"
	"strconv"

	"github.com/gin-gonic/gin"
	"seno-medika.com/config/db"
	"seno-medika.com/model/cashierstation"
	"seno-medika.com/model/common"
)

func GetNota(c *gin.Context) {
	find_by := c.Query("find_by")
	target := c.Query("target")

	switch find_by {
	case "id":
		val, _ := strconv.Atoi(target)
		nota, err := kasir2.FindNotaById(val)
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
		nota, err := kasir2.FindNotaByPasienID(val)
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
		nota, err := kasir2.FindNotaByResepId(val)
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
		nota, err := kasir2.FindNotaByMetodePembayaran(target)
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

	case "detail_resep":
		val, _ := strconv.Atoi(target)
		nota, err := kasir2.FindDetailByResepId(val)
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

	case "detail_tindakan":
		val, _ := strconv.Atoi(target)
		nota, err := kasir2.FindTindakanByNotaId(val)
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
			Message:    "Successfully get detail tindakan",
			Status:     "ok",
			StatusCode: http.StatusOK,
			Data:       nota,
		})
		return

	default:
		nota, err := kasir2.FindNotaAll()
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

	_, err := db.DB.Exec(`INSERT INTO nota (pasien_id, dokter_id, resep_id, list_tindakan_id, total_biaya, metode_pembayaran)  VALUES ($1, $2, $3, $4, $5, $6)`,
		notaVar.PasienID,
		notaVar.DokterID,
		notaVar.ResepID,
		notaVar.ListTindakanID,
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

func PatchNota(c *gin.Context) {
	changeType := c.Query("change_type")
	target := c.Query("target")
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

	switch changeType {
	case "metode_pembayaran":
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
		err = kasir2.UpdateMetodePembayaran(val, notaVar.MetodePembayaran)
		if err != nil {
			c.JSON(http.StatusInternalServerError, common.Response{
				Message:    err.Error(),
				Status:     "Internal Server Error",
				StatusCode: http.StatusInternalServerError,
				Data:       nil,
			})
			return
		}
	case "total_biaya":
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

		err = kasir2.UpdateTotalBiaya(val, notaVar.TotalBiaya)
		if err != nil {
			c.JSON(http.StatusInternalServerError, common.Response{
				Message:    err.Error(),
				Status:     "Internal Server Error",
				StatusCode: http.StatusInternalServerError,
				Data:       nil,
			})
			return
		}
	default:
		c.JSON(http.StatusBadRequest, common.Response{
			Message:    "Invalid update by",
			Status:     "Bad Request",
			StatusCode: http.StatusBadRequest,
			Data:       nil,
		})
		return
	}
	c.JSON(http.StatusOK, common.Response{
		Message:    "Successfully update nota",
		Status:     "ok",
		StatusCode: http.StatusOK,
		Data:       nil,
	})

}
