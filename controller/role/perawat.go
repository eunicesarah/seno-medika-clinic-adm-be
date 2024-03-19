package role

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"seno-medika.com/config/db"
	"seno-medika.com/model/common"
	"seno-medika.com/model/person"
	"seno-medika.com/service/perawat"
)

func GetPerawat(c *gin.Context) {
	findBy := c.Query("find_by")
	target := c.Query("target")

	switch findBy {
	case "id":
		perawatVar, err := perawat.FindPerawatByID(target)
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
			Message:    "Successfully get perawat",
			Status:     "ok",
			StatusCode: http.StatusOK,
			Data:       perawatVar,
		})
		return
	default:
		perawatVars, err := perawat.FindAllPerawat()
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
			Message:    "Successfully get perawat",
			Status:     "ok",
			StatusCode: http.StatusOK,
			Data:       perawatVars,
		})
		return
	}
}

func AddPerawat(c *gin.Context) {
	var perawatVar person.Perawat

	if err := c.ShouldBind(&perawatVar); err != nil {
		c.JSON(http.StatusBadRequest, common.Response{
			Message:    err.Error(),
			Status:     "Bad Request",
			StatusCode: http.StatusBadRequest,
			Data:       nil,
		})
		return
	}

	perawatVar.UserUUID = uuid.New()
	pass, err := bcrypt.GenerateFromPassword([]byte(perawatVar.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.Response{
			Message:    err.Error(),
			Status:     "Internal Server Error",
			StatusCode: http.StatusInternalServerError,
			Data:       nil,
		})
		return
	}
	perawatVar.Password = string(pass)

	if _, err := db.DB.Exec("INSERT INTO users (user_uuid, nama, email, password, role) VALUES ($1, $2, $3, $4, $5)",
		perawatVar.UserUUID, perawatVar.Nama, perawatVar.Email, perawatVar.Password, perawatVar.Role); err != nil {
		c.JSON(http.StatusInternalServerError, common.Response{
			Message:    err.Error(),
			Status:     "Internal Server Error",
			StatusCode: http.StatusInternalServerError,
			Data:       nil,
		})
		return
	}

	if _, err := db.DB.Exec("INSERT INTO perawat(perwat_id, nomor_lisensi) VALUES ($1, $2)",
		perawatVar.UserID, perawatVar.PerawatData.NomorLisensi); err != nil {
		c.JSON(http.StatusInternalServerError, common.Response{
			Message:    err.Error(),
			Status:     "Internal Server Error",
			StatusCode: http.StatusInternalServerError,
			Data:       nil,
		})
		return
	}

	c.JSON(http.StatusCreated, common.Response{
		Message:    "Successfully created perawat",
		Status:     "ok",
		StatusCode: http.StatusCreated,
		Data:       nil,
	})
	return
}

func PatchPerawat(c *gin.Context) {
	var perawatVar person.PerawatData
	target := c.Query("target")

	if err := c.ShouldBind(&perawatVar); err != nil {
		c.JSON(http.StatusBadRequest, common.Response{
			Message:    err.Error(),
			Status:     "Bad Request",
			StatusCode: http.StatusBadRequest,
			Data:       nil,
		})
		return
	}

	val, err := db.DB.Exec("UPDATE perawat SET nomor_lisensi = $1 WHERE perawat_id = $2",
		perawatVar.NomorLisensi, target)

	if err != nil {
		c.JSON(http.StatusInternalServerError, common.Response{
			Message:    err.Error(),
			Status:     "Internal Server Error",
			StatusCode: http.StatusInternalServerError,
			Data:       nil,
		})
		return
	}

	if rows, _ := val.RowsAffected(); rows == 0 {
		c.JSON(http.StatusBadRequest, common.Response{
			Message:    err.Error(),
			Status:     "Bad Request",
			StatusCode: http.StatusBadRequest,
			Data:       nil,
		})
		return
	}

	c.JSON(http.StatusCreated, common.Response{
		Message:    "Successfully updated perawat",
		Status:     "ok",
		StatusCode: http.StatusCreated,
		Data:       nil,
	})
	return
}

func DeletePerawat(c *gin.Context) {
	target := c.Query("target")

	val, err := db.DB.Exec("DELETE FROM perawat WHERE perawat_id = $1", target)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.Response{
			Message:    err.Error(),
			Status:     "Internal Server Error",
			StatusCode: http.StatusInternalServerError,
			Data:       nil,
		})
		return
	}

	if rows, _ := val.RowsAffected(); rows == 0 {
		c.JSON(http.StatusBadRequest, common.Response{
			Message:    err.Error(),
			Status:     "Bad Request",
			StatusCode: http.StatusBadRequest,
			Data:       nil,
		})
		return
	}

	c.JSON(http.StatusCreated, common.Response{
		Message:    "Successfully deleted perawat",
		Status:     "ok",
		StatusCode: http.StatusCreated,
		Data:       nil,
	})
	return
}
