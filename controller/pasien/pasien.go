package pasien

import (
	"net/http"
	"seno-medika.com/service/pasien"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"seno-medika.com/config/db"
	"seno-medika.com/model/common"
	"seno-medika.com/model/person"
)

func AddPasien(c *gin.Context) {
	var pasien person.Pasien

	if err := c.ShouldBind(&pasien); err != nil {
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
		FROM pasien WHERE nik = $1`,
		pasien.NIK)

	if err == nil {
		c.JSON(http.StatusBadRequest, common.Response{
			Message:    "Pasien already exist",
			Status:     "Bad Request",
			StatusCode: http.StatusBadRequest,
			Data:       nil,
		})
		return
	}

	if val.Next() {
		c.JSON(http.StatusBadRequest, common.Response{
			Message:    "Pasien already exist",
			Status:     "Bad Request",
			StatusCode: http.StatusBadRequest,
		})
		return
	}

	pasien.PasienUUID = uuid.New()
	pasien.CreatedAt = time.Now().Local().String()
	pasien.UpdatedAt = time.Now().Local().String()

	_, err = db.DB.Exec(
		`INSERT INTO pasien (
			no_erm,
			pasien_uuid,
			no_rm_lama,
			no_dok_rm,
			penjamin,
			no_penjamin,
			nik,
			no_kk,
			nama,
			tempat_lahir,
			tanggal_lahir,
			no_ihs,
			jenis_kelamin,
			golongan_darah,
			no_telpon,
			email,
			provinsi,
			kabupaten_kota,
			kecamatan,
			kelurahan,
			alamat,
			nama_kontak_darurat,
			nomor_kontak_darurat,
			pekerjaan,
			agama,
			warga_negara,
			pendidikan,
			status_perkawinan,
			created_at,
			created_by,
			updated_at,
			updated_by
		) VALUES (
			$1, $2, $3, $4, $5, $6, $7, $8,
			$9, $10, $11, $12, $13, $14, $15, $16,
			$17, $18, $19, $20, $21, $22, $23, $24,
			$25, $26, $27, $28, $29, $30, $31, $32
		)
		`,
		pasien.NoERM,
		pasien.PasienUUID,
		pasien.NoRMLama,
		pasien.NoDokRM,
		pasien.Penjamin,
		pasien.NoPenjamin,
		pasien.NIK,
		pasien.NoKK,
		pasien.Nama,
		pasien.TempatLahir,
		pasien.TanggalLahir,
		pasien.NoIHS,
		pasien.JenisKelamin,
		pasien.GolonganDarah,
		pasien.NoTelpon,
		pasien.Email,
		pasien.Provinsi,
		pasien.KabupatenKota,
		pasien.Kecamatan,
		pasien.Kelurahan,
		pasien.Alamat,
		pasien.NamaKontakDarurat,
		pasien.NomorKontakDarurat,
		pasien.Pekerjaan,
		pasien.Agama,
		pasien.WargaNegara,
		pasien.Pendidikan,
		pasien.StatusPerkawinan,
		pasien.CreatedAt,
		pasien.CreatedBy,
		pasien.UpdatedAt,
		pasien.UpdatedBy)

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
		Message:    "Successfully insert pasien",
		Status:     "ok",
		StatusCode: http.StatusOK,
		Data:       nil,
	})
	return
}

func UpdatePasien(c *gin.Context) {
	updateBy := c.Query("update_by")
	target := c.Query("target")
	var pasienVar person.Pasien

	if err := c.ShouldBind(&pasienVar); err != nil {
		c.JSON(http.StatusBadRequest, common.Response{
			Message:    err.Error(),
			Status:     "Bad Request",
			StatusCode: http.StatusBadRequest,
			Data:       nil,
		})
		return
	}
	pasienVar.UpdatedAt = time.Now().Local().String()

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
		err = pasien.UpdatePasienById(val, pasienVar)
		if err != nil {
			c.JSON(http.StatusInternalServerError, common.Response{
				Message:    err.Error(),
				Status:     "Internal Server Error",
				StatusCode: http.StatusInternalServerError,
				Data:       nil,
			})
			return
		}
	case "uuid":
		err := pasien.UpdatePasienByUuid(target, pasienVar)
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
		Message:    "Successfully update pasien",
		Status:     "ok",
		StatusCode: http.StatusOK,
		Data:       nil,
	})
	return
}

func DeletePasien(c *gin.Context) {
	updateBy := c.Query("update_by")
	target := c.Query("target")

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

		err = pasien.DeletePasienById(val)
		if err != nil {
			c.JSON(http.StatusInternalServerError, common.Response{
				Message:    err.Error(),
				Status:     "Internal Server Error",
				StatusCode: http.StatusInternalServerError,
				Data:       nil,
			})
			return
		}
	case "uuid":
		err := pasien.DeletePasienByUuid(target)
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
		Message:    "Successfully delete pasien",
		Status:     "ok",
		StatusCode: http.StatusOK,
		Data:       nil,
	})
}

func GetPasien(c *gin.Context) {
	var pasienList []person.Pasien

	updateBy := c.Query("update_by")
	target := c.Query("target")

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

		pasienList, err = pasien.FindPasienById(val)
		if err != nil {
			c.JSON(http.StatusInternalServerError, common.Response{
				Message:    err.Error(),
				Status:     "Internal Server Error",
				StatusCode: http.StatusInternalServerError,
				Data:       nil,
			})
			return
		}

	case "uuid":
		var err error
		pasienList, err = pasien.FindPasienByUuid(target)
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
		var err error
		pasienList, err = pasien.FindPasienAll()
		if err != nil {
			c.JSON(http.StatusInternalServerError, common.Response{
				Message:    err.Error(),
				Status:     "Internal Server Error",
				StatusCode: http.StatusInternalServerError,
				Data:       nil,
			})
			return
		}
		return
	}

	c.JSON(http.StatusOK, common.Response{
		Message:    "Successfully get pasien",
		Status:     "ok",
		StatusCode: http.StatusOK,
		Data:       pasienList,
	})
	return
}
