package pendaftaranpasien

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"seno-medika.com/config/db"
	"seno-medika.com/model"
)

func AddPasien(c *gin.Context) {
	var pasien model.Pasien
	_db := db.DB

	if err := c.ShouldBindJSON(&pasien); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if _, err := _db.Exec(
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
		pasien.No_erm,
		pasien.Pasien_uuid,
		pasien.No_rm_lama,
		pasien.No_dok_rm,
		pasien.Penjamin,
		pasien.No_penjamin,
		pasien.Nik,
		pasien.No_kk,
		pasien.Nama,
		pasien.Tempat_lahir,
		pasien.Tanggal_lahir,
		pasien.No_ihs,
		pasien.Jenis_kelamin,
		pasien.Golongan_darah,
		pasien.No_telpon,
		pasien.Email,
		pasien.Provinsi,
		pasien.Kabupaten_kota,
		pasien.Kecamatan,
		pasien.Kelurahan,
		pasien.Alamat,
		pasien.Nama_kontak_darurat,
		pasien.Nomor_kontak_darurat,
		pasien.Pekerjaan,
		pasien.Agama,
		pasien.Warga_negara,
		pasien.Pendidikan,
		pasien.Status_perkawinan,
		pasien.Created_at,
		pasien.Created_by,
		pasien.Updated_at,
		pasien.Updated_by,
	); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Pasien berhasil didaftarkan"})
}
