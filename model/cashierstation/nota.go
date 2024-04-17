package cashierstation

import "seno-medika.com/model/pharmacystation"

type Nota struct {
	NotaID           int    `json:"nota_id"`
	PasienID         int    `json:"pasien_id"`
	DokterID         int    `json:"dokter_id"`
	ResepID          int    `json:"resep_id"`
	TotalBiaya       int    `json:"total_biaya"`
	MetodePembayaran string `json:"metode_pembayaran"`
}

type DetailNota struct {
	pharmacystation.Obat `json:"-"`
	pharmacystation.ListObat `json:"-"`
	NamaObat             string `json:"nama_obat"`
	Harga                int    `json:"harga"`
	Jumlah               int    `json:"jumlah"`
	Dosis                string `json:"dosis"`
}