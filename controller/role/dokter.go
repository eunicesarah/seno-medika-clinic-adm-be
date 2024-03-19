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
	"seno-medika.com/service/dokter"
	"sync"
)

func AddDokter(c *gin.Context) {
	var dokterVar person.Dokter
	var wg sync.WaitGroup

	if err := c.ShouldBind(&dokterVar); err != nil {
		c.JSON(http.StatusBadRequest, common.Response{
			Message:    err.Error(),
			Status:     "Bad Request",
			StatusCode: http.StatusBadRequest,
			Data:       nil,
		})
		return
	}
	dokterVar.UserUUID = uuid.New()

	errChan := make(chan error)

	go helper.ValidationEmail(dokterVar.Email, errChan)
	go helper.IsEmailExists(dokterVar.Email, errChan)
	go helper.ValidationEmail(dokterVar.Email, errChan)

	if err := <-errChan; err != nil {
		c.JSON(http.StatusBadRequest, common.Response{
			Message:    err.Error(),
			Status:     "Bad Request",
			StatusCode: http.StatusBadRequest,
			Data:       nil,
		})
		return
	}

	pass, err := bcrypt.GenerateFromPassword([]byte(dokterVar.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.Response{
			Message:    err.Error(),
			Status:     "Internal Server Error",
			StatusCode: http.StatusInternalServerError,
			Data:       nil,
		})
		return
	}

	dokterVar.Password = string(pass)

	if _, err = db.DB.Query(
		"INSERT INTO users(user_uuid, nama, password, email, role)"+
			" VALUES($1,$2,$3,$4,$5)", dokterVar.UserUUID, dokterVar.Nama, dokterVar.Password,
		dokterVar.Email, dokterVar.Role); err != nil {
		c.JSON(http.StatusInternalServerError, common.Response{
			Message:    err.Error(),
			Status:     "Internal Server Error",
			StatusCode: http.StatusInternalServerError,
			Data:       nil,
		})
		return
	}

	errChan = make(chan error)

	wg.Add(2)
	go func() {
		defer wg.Done()
		if _, err = db.DB.Query(
			"INSERT INTO dokter(dokter_id, jaga_poli_mana, jadwal_jaga, nomor_lisensi) VALUES ($1,$2,$3,$4)",
			dokterVar.UserID, dokterVar.DokterData.JagaPoliMana, dokterVar.DokterData.JadwalJaga, dokterVar.DokterData.NomorLisensi); err != nil {
			errChan <- err
			return
		}
	}()

	go func() {
		defer wg.Done()
		for _, value := range dokterVar.DokterData.ListJadwalDokter {
			if _, err := db.DB.Query(
				"INSERT INTO list_jadwal_dokter(dokter_id, hari, shift) VALUES ($1,$2,$3)",
				dokterVar.UserID, value.Hari, value.Shift); err != nil {
				errChan <- err
				return
			}
		}
	}()

	wg.Wait()

	if err = <-errChan; err != nil {
		c.JSON(http.StatusInternalServerError, common.Response{
			Message:    err.Error(),
			Status:     "Internal Server Error",
			StatusCode: http.StatusInternalServerError,
			Data:       nil,
		})
		return
	}

	c.JSON(http.StatusCreated, common.Response{
		Message:    "Dokter created",
		Status:     "ok",
		StatusCode: http.StatusCreated,
		Data:       nil,
	})
	return
}

func DeleteDokter(c *gin.Context) {
	changeType := c.Query("change_type")
	target := c.Query("target")

	switch changeType {
	case "dokter":
		if err := dokter.DeleteDokterById(target); err != nil {
			c.JSON(http.StatusInternalServerError, common.Response{
				Message:    err.Error(),
				Status:     "Internal Server Error",
				StatusCode: http.StatusInternalServerError,
				Data:       nil,
			})
			return
		}
	case "jadwal":
		if err := dokter.DeleteListJadwalById(target); err != nil {
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
			Message:    "Invalid change type",
			Status:     "Bad Request",
			StatusCode: http.StatusBadRequest,
			Data:       nil,
		})
		return
	}

	c.JSON(http.StatusOK, common.Response{
		Message:    "Successfully delete dokter",
		Status:     "ok",
		StatusCode: http.StatusOK,
		Data:       nil,
	})
	return
}

func PatchDokter(c *gin.Context) {
	changeType := c.Query("change_type")
	target := c.Query("target")

	switch changeType {
	case "dokter":
		var dokterVar person.DokterData
		if err := c.ShouldBind(&dokterVar); err != nil {
			c.JSON(http.StatusBadRequest, common.Response{
				Message:    err.Error(),
				Status:     "Bad Request",
				StatusCode: http.StatusBadRequest,
				Data:       nil,
			})
			return
		}

		if err := dokter.ChangeDokterById(target, dokterVar); err != nil {
			c.JSON(http.StatusInternalServerError, common.Response{
				Message:    err.Error(),
				Status:     "Internal Server Error",
				StatusCode: http.StatusInternalServerError,
				Data:       nil,
			})
			return
		}
	case "jadwal":
		var dokterVar person.ListJadwalDokter
		if err := c.ShouldBind(&dokterVar); err != nil {
			c.JSON(http.StatusBadRequest, common.Response{
				Message:    err.Error(),
				Status:     "Bad Request",
				StatusCode: http.StatusBadRequest,
				Data:       nil,
			})
			return
		}

		if err := dokter.ChangeListJadwalById(target, dokterVar); err != nil {
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
			Message:    "Invalid change type",
			Status:     "Bad Request",
			StatusCode: http.StatusBadRequest,
			Data:       nil,
		})
		return
	}

	c.JSON(http.StatusOK, common.Response{
		Message:    "Successfully delete dokter",
		Status:     "ok",
		StatusCode: http.StatusOK,
		Data:       nil,
	})
	return
}
