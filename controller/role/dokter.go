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
	"seno-medika.com/model/station/cashierstation"
	dokter2 "seno-medika.com/query/role/dokter"
	"sync"
)

func GetDokter(c *gin.Context) {
	findBy := c.Query("find_by")
	target := c.Query("target")

	switch findBy {
	case "id":
		dokterVar, err := dokter2.FindDokterByID(target)
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
			Message:    "Successfully get dokter",
			Status:     "ok",
			StatusCode: http.StatusOK,
			Data:       dokterVar,
		})
		return
	default:
		dokterVars, err := dokter2.FindAllDokter()
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
			Message:    "Successfully get dokter",
			Status:     "ok",
			StatusCode: http.StatusOK,
			Data:       dokterVars,
		})
		return
	}
}

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
	dokterVar.Role = "dokter"

	errChan := make(chan error, 3)
	wg.Add(3)
	go func() {
		defer wg.Done()
		helper.ValidationEmail(dokterVar.Email, errChan)
	}()
	go func() {
		defer wg.Done()
		helper.IsEmailExists(dokterVar.Email, errChan)
	}()
	go func() {
		defer wg.Done()
		helper.ValidationPassword(dokterVar.Password, errChan)
	}()
	wg.Wait()

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

	var dokterId string
	if _, err := db.DB.Query(
		"INSERT INTO users(user_uuid, nama, password, email, role) VALUES($1,$2,$3,$4,$5)", dokterVar.UserUUID, dokterVar.Nama, dokterVar.Password,
		dokterVar.Email, dokterVar.Role); err != nil {
		c.JSON(http.StatusInternalServerError, common.Response{
			Message:    err.Error(),
			Status:     "Internal Server Error",
			StatusCode: http.StatusInternalServerError,
			Data:       nil,
		})
		return
	}

	if err := db.DB.QueryRow("SELECT user_id FROM users WHERE user_uuid = $1", dokterVar.UserUUID).Scan(&dokterId); err != nil {
		c.JSON(http.StatusInternalServerError, common.Response{
			Message:    err.Error(),
			Status:     "Internal Server Error",
			StatusCode: http.StatusInternalServerError,
			Data:       nil,
		})
		return
	}

	if _, err = db.DB.Query(
		"INSERT INTO dokter(dokter_id, jaga_poli_mana, jadwal_jaga, nomor_lisensi) VALUES ($1,$2,$3,$4)",
		dokterId, dokterVar.DokterData.JagaPoliMana, dokterVar.DokterData.JadwalJaga, dokterVar.DokterData.NomorLisensi); err != nil {
		c.JSON(http.StatusInternalServerError, common.Response{
			Message:    err.Error(),
			Status:     "Internal Server Error " + dokterId,
			StatusCode: http.StatusInternalServerError,
			Data:       nil,
		})
		return
	}

	for _, value := range dokterVar.DokterData.ListJadwalDokter {
		if _, err := db.DB.Query(
			"INSERT INTO list_jadwal_dokter(dokter_id, hari, shift) VALUES ($1,$2,$3)",
			dokterId, value.Hari, value.Shift); err != nil {
			c.JSON(http.StatusInternalServerError, common.Response{
				Message:    err.Error(),
				Status:     "Internal Server Error",
				StatusCode: http.StatusInternalServerError,
				Data:       nil,
			})
			return
		}
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
		if err := dokter2.DeleteDokterById(target); err != nil {
			c.JSON(http.StatusInternalServerError, common.Response{
				Message:    err.Error(),
				Status:     "Internal Server Error",
				StatusCode: http.StatusInternalServerError,
				Data:       nil,
			})
			return
		}
	case "jadwal":
		if err := dokter2.DeleteListJadwalById(target); err != nil {
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

		if err := dokter2.ChangeDokterById(target, dokterVar); err != nil {
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

		if err := dokter2.ChangeListJadwalById(target, dokterVar); err != nil {
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

func AddTindakan(c *gin.Context) {
	var tindVal cashierstation.Tindakan

	if err := c.ShouldBind(&tindVal); err != nil {
		c.JSON(http.StatusBadRequest, common.Response{
			Message:    err.Error(),
			Status:     "Bad Request",
			StatusCode: http.StatusBadRequest,
			Data:       nil,
		})
		return
	}

	if err := dokter2.AddTindakan(tindVal); err != nil {
		c.JSON(http.StatusInternalServerError, common.Response{
			Message:    err.Error(),
			Status:     "Internal Server Error",
			StatusCode: http.StatusInternalServerError,
			Data:       nil,
		})
		return
	}

	c.JSON(http.StatusCreated, common.Response{
		Message:    "Successfully add tindakan",
		Status:     "Status Created",
		StatusCode: http.StatusCreated,
		Data:       nil,
	})
	return
}

func GetTindakan(c *gin.Context) {
	findBy := c.Query("find_by")
	target := c.Query("target")

	switch findBy {
	case "id":
		tindVal, err := dokter2.FindTindakanById(target)
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
			Message:    "Successfully get tindakan",
			Status:     "ok",
			StatusCode: http.StatusOK,
			Data:       tindVal,
		})
		return
	default:
		tindVal, err := dokter2.FindAllTindakan()
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
			Message:    "Successfully get tindakan",
			Status:     "ok",
			StatusCode: http.StatusOK,
			Data:       tindVal,
		})
		return
	}
}

func DeleteTindakan(c *gin.Context) {
	deleteBy := c.Query("delete_by")
	target := c.Query("target")

	switch deleteBy {
	case "id":
		if err := dokter2.DeleteTindakanById(target); err != nil {
			c.JSON(http.StatusInternalServerError, common.Response{
				Message:    err.Error(),
				Status:     "Internal Server Error",
				StatusCode: http.StatusInternalServerError,
				Data:       nil,
			})
			return
		}

		c.JSON(http.StatusOK, common.Response{
			Message:    "Successfully delete tindakan",
			Status:     "ok",
			StatusCode: http.StatusOK,
			Data:       nil,
		})
		return

	default:
		c.JSON(http.StatusBadRequest, common.Response{
			Message:    "Invalid delete by",
			Status:     "Bad Request",
			StatusCode: http.StatusBadRequest,
			Data:       nil,
		})
		return
	}
}

func PutTindakan(c *gin.Context) {
	changeBy := c.Query("update_by")
	target := c.Query("target")

	switch changeBy {
	case "id":
		var tindVal cashierstation.Tindakan
		if err := c.ShouldBind(&tindVal); err != nil {
			c.JSON(http.StatusBadRequest, common.Response{
				Message:    err.Error(),
				Status:     "Bad Request",
				StatusCode: http.StatusBadRequest,
				Data:       nil,
			})
			return
		}

		if err := dokter2.PutTindakanById(target, tindVal); err != nil {
			c.JSON(http.StatusInternalServerError, common.Response{
				Message:    err.Error(),
				Status:     "Internal Server Error",
				StatusCode: http.StatusInternalServerError,
				Data:       nil,
			})
			return
		}

		c.JSON(http.StatusOK, common.Response{
			Message:    "Successfully update tindakan",
			Status:     "ok",
			StatusCode: http.StatusOK,
			Data:       nil,
		})
		return

	default:
		c.JSON(http.StatusBadRequest, common.Response{
			Message:    "Invalid change by",
			Status:     "Bad Request",
			StatusCode: http.StatusBadRequest,
			Data:       nil,
		})
		return
	}
}
