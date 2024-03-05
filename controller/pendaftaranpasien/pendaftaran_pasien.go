package pendaftaranpasien


import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"seno-medika.com/config/db"
	"seno-medika.com/model/common"
	"seno-medika.com/model/person"
	
)

// TODO: get user that created and updated the data
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
	

	pasien.PasienUUID = uuid.New()
	pasien.CreatedAt = time.Now().Local().String()
	pasien.UpdatedAt = time.Now().Local().String()
	// pasien.CreatedBy 
	// pasien.UpdatedBy

	_, err := db.DB.Exec(
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
		Message:    "Successfully insert user",
		Status:     "ok",
		StatusCode: http.StatusOK,
		Data:       nil,
	})

}

// TODO: get user that updated the data
