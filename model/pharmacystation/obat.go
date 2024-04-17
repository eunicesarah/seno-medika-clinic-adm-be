package pharmacystation

type Obat struct {
	ObatID        int    `json:"obat_id"`
	NamaObat      string `json:"nama_obat"`
	JenisAsuransi string `json:"jenis_asuransi"`
	Harga         int    `json:"harga"`
}

type ListObat struct {
	ObatID  int    `json:"obat_id"`
	ResepID int    `json:"resep_id"`
	Jumlah  int    `json:"jumlah"`
	Dosis   string `json:"dosis"`
}

type DetailObat struct {
	Obat     Obat
	ListObat ListObat
}