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
