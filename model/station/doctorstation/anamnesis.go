package doctorstation

type Anamnesis struct {
	AnamnesisID       int    `json:"anamnesis_id"`
	PasienID          int    `json:"pasien_id"`
	SkrinAwalID       int    `json:"skrin_awal_id"`
	SkrinGiziID       int    `json:"skrin_gizi_id"`
	TTVID             int    `json:"ttv_id"`
	RiwayatPenyakitID int    `json:"riwayat_penyakit_id"`
	AlergiID          int    `json:"alergi_id"`
	DokterID          int    `json:"dokter_id"`
	PerawatID         int    `json:"perawat_id"`
	KeluhanUtama      string `json:"keluhan_utama"`
	KeluhanTambahan   string `json:"keluhan_tambahan"`
	LamaSakit         int    `json:"lama_sakit"`
}
