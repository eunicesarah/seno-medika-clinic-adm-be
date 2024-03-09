package person

import (
	"github.com/google/uuid"
)

type Pasien struct {
	PasienID           int       `json:"pasien_id"`
	NoERM              string    `json:"no_erm"`
	PasienUUID         uuid.UUID `json:"pasien_uuid"`
	NoRMLama           string    `json:"no_rm_lama"`
	NoDokRM            string    `json:"no_dok_rm"`
	Penjamin           string    `json:"penjamin"`
	NoPenjamin         string    `json:"no_penjamin"`
	NIK                string    `json:"nik"`
	NoKK               string    `json:"no_kk"`
	Nama               string    `json:"nama"`
	TempatLahir        string    `json:"tempat_lahir"`
	TanggalLahir       string    `json:"tanggal_lahir"`
	NoIHS              string    `json:"no_ihs"`
	JenisKelamin       string    `json:"jenis_kelamin"`
	GolonganDarah      string    `json:"golongan_darah"`
	NoTelpon           string    `json:"no_telpon"`
	Email              string    `json:"email"`
	Provinsi           string    `json:"provinsi"`
	KabupatenKota      string    `json:"kabupaten_kota"`
	Kecamatan          string    `json:"kecamatan"`
	Kelurahan          string    `json:"kelurahan"`
	Alamat             string    `json:"alamat"`
	NamaKontakDarurat  string    `json:"nama_kontak_darurat"`
	NomorKontakDarurat string    `json:"nomor_kontak_darurat"`
	Pekerjaan          string    `json:"pekerjaan"`
	Agama              string    `json:"agama"`
	WargaNegara        string    `json:"warga_negara"`
	Pendidikan         string    `json:"pendidikan"`
	StatusPerkawinan   string    `json:"status_perkawinan"`
	CreatedAt          string    `json:"created_at"`
	CreatedBy          string    `json:"created_by"`
	UpdatedAt          string    `json:"updated_at"`
	UpdatedBy          string    `json:"updated_by"`
}
