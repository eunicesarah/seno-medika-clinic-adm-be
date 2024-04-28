package obat

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"seno-medika.com/config/db"
	"seno-medika.com/model/common"
	"seno-medika.com/model/pharmacystation"
	"seno-medika.com/query/obat"
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
			harga
		) VALUES (
			$1, $2, $3
		)
		`,
		obatVar.NamaObat,
		obatVar.JenisAsuransi,
		obatVar.Harga)

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
		err = obat.UpdateObatById(val, obatVar)
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
		err := obat.UpdateObatByName(target, obatVar)
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

		err = obat.DeleteObatById(val)
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
		err := obat.DeleteObatByName(target)
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
		obat, err := obat.FindObatById(val)
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
		obat, err := obat.FindObatByName(target)
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
		obat, err := obat.FindObatAll()
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
