package cashierstation

type Nota struct {
	NotaID     int `json:"nota_id"`
	PasienID   int `json:"pasien_id"`
	DokterID   int `json:"dokter_id"`
	ResepID    int `json:"resep_id"`
	ListTindakanID int `json:"list_tindakan_id"`
	TotalBiaya int `json:"total_biaya"`
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
	TindakanID     int `json:"tindakan_id"`
	ListTindakanID int `json:"list_tindakan_id"`
}