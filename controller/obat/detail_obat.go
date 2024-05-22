package obat

import (
	"net/http"
	"seno-medika.com/model/station/pharmacystation"
	"seno-medika.com/query/obat/detail"
	"strconv"

	"github.com/gin-gonic/gin"
	"seno-medika.com/config/db"
	"seno-medika.com/model/common"
)

func AddObat(c *gin.Context) {
	var obatVar pharmacystation.Obat

	if err := c.ShouldBind(&obatVar); err != nil {
		c.JSON(http.StatusBadRequest, common.Response{
			Message:    err.Error(),
			Status:     "Bad Request",
			StatusCode: http.StatusBadRequest,
			Data:       nil,
		})
		return
	}

	val, err := db.DB.Query(
		`SELECT *
		FROM obat WHERE nama_obat = $1`,
		obatVar.NamaObat)

	if err != nil {
		c.JSON(http.StatusInternalServerError, common.Response{
			Message:    err.Error(),
			Status:     "Internal Server Error",
			StatusCode: http.StatusInternalServerError,
			Data:       nil,
		})
		return
	}

	if val.Next() {
		c.JSON(http.StatusBadRequest, common.Response{
			Message:    "Obat already exist",
			Status:     "Bad Request",
			StatusCode: http.StatusBadRequest,
			Data:       nil,
		})
		return
	}

	_, err = db.DB.Exec(
		`INSERT INTO obat (
			nama_obat,
			jenis_asuransi,
			harga,
			stock,
			satuan
		) VALUES (
			$1, $2, $3, $4, $5
		)
		`,
		obatVar.NamaObat,
		obatVar.JenisAsuransi,
		obatVar.Harga,
		obatVar.Stock,
		obatVar.Satuan)

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
		Message:    "Successfully insert obat",
		Status:     "ok",
		StatusCode: http.StatusOK,
		Data:       nil,
	})
	return
}

func UpdateObat(c *gin.Context) {
	updateBy := c.Query("update_by")
	target := c.Query("target")
	var obatVar pharmacystation.Obat

	if err := c.ShouldBind(&obatVar); err != nil {
		c.JSON(http.StatusBadRequest, common.Response{
			Message:    err.Error(),
			Status:     "Bad Request",
			StatusCode: http.StatusBadRequest,
			Data:       nil,
		})
		return
	}

	switch updateBy {
	case "id":
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
		err = detail.UpdateObatById(val, obatVar)
		if err != nil {
			c.JSON(http.StatusInternalServerError, common.Response{
				Message:    err.Error(),
				Status:     "Internal Server Error",
				StatusCode: http.StatusInternalServerError,
				Data:       nil,
			})
			return
		}
	case "name":
		err := detail.UpdateObatByName(target, obatVar)
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
		})
		return
	}

	c.JSON(http.StatusOK, common.Response{
		Message:    "Successfully update obat",
		Status:     "ok",
		StatusCode: http.StatusOK,
		Data:       nil,
	})
	return
}

func DeleteObat(c *gin.Context) {
	deleteBy := c.Query("delete_by")
	target := c.Query("target")

	switch deleteBy {
	case "id":
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

		err = detail.DeleteObatById(val)
		if err != nil {
			c.JSON(http.StatusInternalServerError, common.Response{
				Message:    err.Error(),
				Status:     "Internal Server Error",
				StatusCode: http.StatusInternalServerError,
				Data:       nil,
			})
			return
		}
	case "name":
		err := detail.DeleteObatByName(target)
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
		})
		return
	}

	c.JSON(http.StatusOK, common.Response{
		Message:    "Successfully delete obat",
		Status:     "ok",
		StatusCode: http.StatusOK,
		Data:       nil,
	})
}

func GetObat(c *gin.Context) {

	find_by := c.Query("find_by")
	target := c.Query("target")

	switch find_by {
	case "id":
		val, _ := strconv.Atoi(target)
		obat, err := detail.FindObatById(val)
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
			Message:    "Successfully get obat",
			Status:     "ok",
			StatusCode: http.StatusOK,
			Data:       obat,
		})
		return

	case "name":
		obat, err := detail.FindObatByName(target)
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
			Message:    "Successfully get obat",
			Status:     "ok",
			StatusCode: http.StatusOK,
			Data:       obat,
		})
		return

	default:
		obat, err := detail.FindObatAll()
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
			Message:    "Successfully get obat",
			Status:     "ok",
			StatusCode: http.StatusOK,
			Data:       obat,
		})
		return
	}
}

func PatchObat(c *gin.Context) {

	updateBy := c.Query("update_by")
	target := c.Query("target")
	var obatVar pharmacystation.Obat

	if err := c.ShouldBind(&obatVar); err != nil {
		c.JSON(http.StatusBadRequest, common.Response{
			Message:    err.Error(),
			Status:     "Bad Request",
			StatusCode: http.StatusBadRequest,
			Data:       nil,
		})
		return
	}

	switch updateBy {
	case "id_harga":
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
		err = detail.PatchHargaObatById(val, obatVar.Harga)
		if err != nil {
			c.JSON(http.StatusInternalServerError, common.Response{
				Message:    err.Error(),
				Status:     "Internal Server Error",
				StatusCode: http.StatusInternalServerError,
				Data:       nil,
			})
			return
		}
	case "name_harga":
		err := detail.PatchHargaObatByName(target, obatVar.Harga)
		if err != nil {
			c.JSON(http.StatusInternalServerError, common.Response{
				Message:    err.Error(),
				Status:     "Internal Server Error",
				StatusCode: http.StatusInternalServerError,
				Data:       nil,
			})
			return
		}
	case "id_stock":
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
		err = detail.PatchStockObatById(val, obatVar.Stock)
		if err != nil {
			c.JSON(http.StatusInternalServerError, common.Response{
				Message:    err.Error(),
				Status:     "Internal Server Error",
				StatusCode: http.StatusInternalServerError,
				Data:       nil,
			})
			return
		}
	case "name_stock":
		err := detail.PatchStockObatByName(target, obatVar.Stock)
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
		})
		return
	}

	c.JSON(http.StatusOK, common.Response{
		Message:    "Successfully update obat",
		Status:     "ok",
		StatusCode: http.StatusOK,
		Data:       nil,
	})
	return

}
