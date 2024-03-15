package person

type Dokter struct {
	DokterID     int    `json:"dokter_id"`
	JagaPoliMana string `json:"jaga_poli_mana"`
	JadwalJaga   string `json:"jadwal_jaga"`
	NomorLisensi string `json:"nomor_lisensi"`
}

type ListJadwalDokter struct {
	DokterID int    `json:"dokter_id"`
	Hari     string `json:"hari"`
	Shift    int    `json:"shift"`
}
