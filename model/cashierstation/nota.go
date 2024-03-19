package cashierstation

type Nota struct {
	NotaID     int `json:"nota_id"`
	PasienID   int `json:"pasien_id"`
	DokterID   int `json:"dokter_id"`
	ResepID    int `json:"resep_id"`
	TotalBiaya int `json:"total_biaya"`
}
