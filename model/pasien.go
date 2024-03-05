package model

type Pasien struct {
	Pasien_id           int    `json:"pasien_id"`
	No_erm              int    `json:"no_erm"`
	Pasien_uuid         string `json:"pasien_uuid"`
	No_rm_lama          string `json:"no_rm_lama"`
	No_dok_rm           string `json:"no_dok_rm"`
	Penjamin            string `json:"penjamin"`
	No_penjamin         string `json:"no_penjamin"`
	Nik                 string `json:"nik"`
	No_kk               string `json:"no_kk"`
	Nama                string `json:"nama"`
	Tempat_lahir        string `json:"tempat_lahir"`
	Tanggal_lahir       string `json:"tanggal_lahir"`
	No_ihs              string `json:"no_ihs"`
	Jenis_kelamin       string `json:"jenis_kelamin"`
	Golongan_darah      string `json:"golongan_darah"`
	No_telpon           string `json:"no_telpon"`
	Email               string `json:"email"`
	Provinsi            string `json:"provinsi"`
	Kabupaten_kota      string `json:"kabupaten_kota"`
	Kecamatan           string `json:"kecamatan"`
	Kelurahan           string `json:"kelurahan"`
	Alamat              string `json:"alamat"`
	Nama_kontak_darurat string `json:"nama_kontak_darurat"`
	Nomor_kontak_darurat string `json:"nomor_kontak_darurat"`
	Pekerjaan           string `json:"pekerjaan"`
	Agama               string `json:"agama"`
	Warga_negara        string `json:"warga_negara"`
	Pendidikan          string `json:"pendidikan"`
	Status_perkawinan  string `json:"status_perkawinan"`
	Created_at          string `json:"created_at"`
	Created_by          string `json:"created_by"`
	Updated_at          string `json:"updated_at"`
	Updated_by          string `json:"updated_by"`
}
