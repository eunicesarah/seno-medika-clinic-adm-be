package pharmacystation

type Obat struct {
	ObatID        int    `json:"obat_id"`
	NamaObat      string `json:"nama_obat"`
	JenisAsuransi string `json:"jenis_asuransi"`
	Harga         int64  `json:"harga"`
	Stock         int    `json:"stock"`
	Satuan        string `json:"satuan"`
}

type ListObat struct {
	ObatID      int    `json:"obat_id"`
	ResepID     int    `json:"resep_id"`
	Jumlah      int    `json:"jumlah"`
	Dosis       string `json:"dosis"`
	AturanPakai string `json:"aturan_pakai"`
	Keterangan  string `json:"keterangan"`
}

type DetailObat struct {
	Obat     Obat
	ListObat ListObat
}
