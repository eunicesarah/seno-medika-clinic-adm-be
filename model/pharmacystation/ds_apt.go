package pharmacystation

type DashboardApotek struct{
	NomorAntrian int `json:"nomor_antrian"`
	Poli string `json:"poli"`
	NoERM string `json:"no_erm"`
	NIK string `json:"nik"`
	Nama string `json:"nama"`
	JenisKelamin string `json:"jenis_kelamin"`
	TempatLahir string `json:"tempat_lahir"`
	TanggalLahir string `json:"tanggal_lahir"`
	Penjamin string `json:"penjamin"`
	Status string `json:"status_obat"`
}