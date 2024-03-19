package doctorstation

type Alergi struct {
	AlergiID int    `json:"alergi_id"`
	Obat     string `json:"obat"`
	Makanan  string `json:"makanan"`
	Lainnya  string `json:"lainnya"`
}
