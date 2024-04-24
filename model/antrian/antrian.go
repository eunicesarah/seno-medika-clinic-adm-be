package antrian

type Antrian struct {
	AntrianID    int    `json:"antrian_id"`
	PasienID     int    `json:"pasien_id"`
	NomorAntrian int    `json:"nomor_antrian"`
	Status       bool   `json:"status"`
	Poli         string `json:"poli"`
	Instalasi    string `json:"instalasi"`
	CreatedAt    string `json:"created_at"`
}

type PendaftaranAntrian struct {
	NIK 		 string `json:"nik"`
	Nama		 string `json:"nama"`
	AntrianID    int    `json:"antrian_id"`
	PasienID     int    `json:"pasien_id"`
	NomorAntrian int    `json:"nomor_antrian"`
	Status       bool   `json:"status"`
	Poli         string `json:"poli"`
	Instalasi    string `json:"instalasi"`
	CreatedAt    string `json:"created_at"`
}

type AntrianNurse struct {
	NomorAntrian int    `json:"nomor_antrian"`
	Poli         string `json:"poli"`
	CreatedAt    string `json:"created_at"`
	PasienID     int    `json:"pasien_id"`
	NoERM        string `json:"no_erm"`
	NIK          string `json:"nik"`
	Nama         string `json:"nama"`
	JenisKelamin string `json:"jenis_kelamin"`
	TempatLahir  string `json:"tempat_lahir"`
	TanggalLahir string `json:"tanggal_lahir"`
	Penjamin     string `json:"penjamin"`
}
