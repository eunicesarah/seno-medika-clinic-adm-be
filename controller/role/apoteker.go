package role

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"seno-medika.com/config/db"
	"seno-medika.com/helper"
	"seno-medika.com/model/common"
	"seno-medika.com/model/person"
	"seno-medika.com/service/apoteker"
)

func GetApoteker(c *gin.Context) {
	findBy := c.Query("find_by")
	target := c.Query("target")

	switch findBy {
	case "id":
		apotekerVar, err := apoteker.FindApotekerByID(target)
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
			Message:    "Successfully get apoteker",
			Status:     "ok",
			StatusCode: http.StatusOK,
			Data:       apotekerVar,
		})
		return
	default:
		apotekerVars, err := apoteker.FindAllApoteker()
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
			Message:    "Successfully get apoteker",
			Status:     "ok",
			StatusCode: http.StatusOK,
			Data:       apotekerVars,
		})
		return
	}
}

func AddApoteker(c *gin.Context) {
	var apotekerVar person.Apoteker
	if err := c.ShouldBind(&apotekerVar); err != nil {
		c.JSON(http.StatusBadRequest, common.Response{
			Message:    err.Error(),
			Status:     "Bad Request",
			StatusCode: http.StatusBadRequest,
			Data:       nil,
		})
		return
	}

	apotekerVar.UserUUID = uuid.New()
	pass, err := bcrypt.GenerateFromPassword([]byte(apotekerVar.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusBadRequest, common.Response{
			Message:    err.Error(),
			Status:     "Bad Request",
			StatusCode: http.StatusBadRequest,
			Data:       nil,
		})
		return
	}
	apotekerVar.Password = string(pass)

	errChan := make(chan error)

	go helper.ValidationEmail(apotekerVar.Email, errChan)
	go helper.IsEmailExists(apotekerVar.Email, errChan)
	go helper.ValidationEmail(apotekerVar.Email, errChan)

	close(errChan)

	if err := <-errChan; err != nil {
		c.JSON(http.StatusBadRequest, common.Response{
			Message:    err.Error(),
			Status:     "Bad Request",
			StatusCode: http.StatusBadRequest,
			Data:       nil,
		})
		return
	}

	row := db.DB.QueryRow("INSERT INTO users (user_uuid, nama, email, password, role) VALUES ($1, $2, $3, $4, $5)",
		apotekerVar.UserUUID, apotekerVar.Nama, apotekerVar.Email, apotekerVar.Password, apotekerVar.Role)

	var apotekerId string

	if err := row.Scan(&apotekerId); err != nil {
		c.JSON(http.StatusInternalServerError, common.Response{
			Message:    err.Error(),
			Status:     "Internal Server Error",
			StatusCode: http.StatusInternalServerError,
			Data:       nil,
		})
		return
	}

	if _, err := db.DB.Exec("INSERT INTO apoteker (apoteker_id, nomor_lisensi) VALUES ($1,$2)",
		apotekerId, apotekerVar.ApotekerData.NomorLisensi); err != nil {
		c.JSON(http.StatusInternalServerError, common.Response{
			Message:    err.Error(),
			Status:     "Internal Server Error",
			StatusCode: http.StatusInternalServerError,
			Data:       nil,
		})
		return
	}

	c.JSON(http.StatusCreated, common.Response{
		Message:    "Successfully created apoteker",
		Status:     "ok",
		StatusCode: http.StatusCreated,
		Data:       nil,
	})
	return
}

func PatchApoteker(c *gin.Context) {
	var apotekerVar person.ApotekerData
	target := c.Query("target")

	val, err := db.DB.Exec("UPDATE FROM apoteker SET nomor_lisensi = $1 WHERE apoteker_id = $2", apotekerVar.NomorLisensi, target)
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
		Message:    "Successfully updated apoteker",
		Status:     "ok",
		StatusCode: http.StatusCreated,
		Data:       nil,
	})
	return
}

func DeleteApoteker(c *gin.Context) {
	target := c.Query("target")

	val, err := db.DB.Exec("DELETE FROM apoteker WHERE perawat_id = $1", target)
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
		Message:    "Successfully deleted apoteker",
		Status:     "ok",
		StatusCode: http.StatusCreated,
		Data:       nil,
	})
	return
}
