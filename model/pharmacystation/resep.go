package pharmacystation

type Resep struct {
	ResepID   int    `json:"resep_id"`
	PemeriksaanDokterID int `json:"pemeriksaan_dokter_id"`
	Deskripsi string `json:"deskripsi"`
	RuangTujuan string `json:"ruang_tujuan"`
	StatusObat string `json:"status_obat"`
}