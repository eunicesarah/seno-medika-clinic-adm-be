package pharmacystation

type DashboardApotek struct {
	NomorAntrian     int    `json:"nomor_antrian"`
	Poli             string `json:"poli"`
	NoERM            string `json:"no_erm"`
	CreatedAt        string `json:"created_at"`
	NIK              string `json:"nik"`
	Nama             string `json:"nama"`
	JenisKelamin     string `json:"jenis_kelamin"`
	TempatLahir      string `json:"tempat_lahir"`
	TanggalLahir     string `json:"tanggal_lahir"`
	MetodePembayaran string `json:"metode_pembayaran"`
	Status           string `json:"status_obat"`
}
