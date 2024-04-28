package cashierstation

import "seno-medika.com/model/pharmacystation"

type Nota struct {
	NotaID           int    `json:"nota_id"`
	PasienID         int    `json:"pasien_id"`
	DokterID         int    `json:"dokter_id"`
	ResepID          int    `json:"resep_id"`
	ListTindakanID   int    `json:"list_tindakan_id"`
	TotalBiaya       int    `json:"total_biaya"`
	MetodePembayaran string `json:"metode_pembayaran"`
}

type ListTindakan struct {
	ListTindakanID int `json:"list_tindakan_id"`
	TindakanID     int `json:"tindakan_id"`
}

type Tindakan struct {
	TindakanID    int    `json:"tindakan_id"`
	NamaTindakan  string `json:"nama_tindakan"`
	Deskripsi     string `json:"deskripsi"`
	HargaTindakan int    `json:"harga_tindakan"`
}

type Penanganan struct {
	ListTindakanID int `json:"list_tindakan_id"`
}
type DetailNota struct {
	pharmacystation.Obat     `json:"-"`
	pharmacystation.ListObat `json:"-"`
	NamaObat                 string `json:"nama_obat"`
	Harga                    int    `json:"harga"`
	Jumlah                   int    `json:"jumlah"`
	Dosis                    string `json:"dosis"`
}

type DetailTindakan struct {
	NamaTindakan  string `json:"nama_tindakan"`
	Deskripsi     string `json:"deskripsi"`
	HargaTindakan int    `json:"harga_tindakan"`
}
