package doctorstation

type ListCppt struct {
	PasienId int `json:"pasien_id"`
	NoErm    int `json:"no_erm"`
	CpptId   int `json:"cppt_id"`
}

type Cppt struct {
	CpptId          int    `json:"cppt_id"`
	Unit            string `json:"unit"`
	Tanggal         string `json:"tanggal"`
	Objektif        string `json:"objektif"`
	Assessment      string `json:"assessment"`
	Penatalaksanaan string `json:"penatalaksanaan"`
}

type PemeriksaanDokter struct {
	PemeriksaanDokterId int `json:"pemeriksaan_dokter_id"`
	PasienId            int `json:"pasien_id"`
	DokterId            int `json:"dokter_id"`
	PerawatId           int `json:"perawat_id"`
}

type PemeriksaanFisik struct {
	PemeriksaanFisikId   int    `json:"pemeriksaan_fisik_id"`
	PemeriksaanDokterId  int    `json:"pemeriksaan_dokter_id"`
	TerapiYgSdhDilakukan string `json:"terapi_yg_sdh_dilakukan"`
	RencanaTindakan      string `json:"rencana_tindakan"`
	TindakanKeperawatan  string `json:"tindakan_keperawatan"`
	Observasi            string `json:"observasi"`
	Merokok              bool   `json:"merokok"`
	KonsumsiAlkohol      bool   `json:"konsumsi_alkohol"`
	KurangSayur          bool   `json:"kurang_sayur"`
}

type RiwayatPemeriksaan struct {
	RiwayatPemeriksaanId int    `json:"riwayat_pemeriksaan_id"`
	PemeriksaanDokterId  int    `json:"pemeriksaan_dokter_id"`
	PasienId             int    `json:"pasien_id"`
	Tanggal              string `json:"tanggal"`
	Pemeriksaan          string `json:"pemeriksaan"`
	Keterangan           string `json:"keterangan"`
}

type KeadaanFisik struct {
	KeadaanFisikId              int  `json:"keadaan_fisik_id"`
	PemeriksaanDokterId         int  `json:"pemeriksaan_dokter_id"`
	PemeriksaanKulit            bool `json:"pemeriksaan_kulit"`
	PemeriksaanKuku             bool `json:"pemeriksaan_kuku"`
	PemeriksaanKepala           bool `json:"pemeriksaan_kepala"`
	PemeriksaanMata             bool `json:"pemeriksaan_mata"`
	PemeriksaanTelinga          bool `json:"pemeriksaan_telinga"`
	PemeriksaanHidungSinus      bool `json:"pemeriksaan_hidung_sinus"`
	PemeriksaanMulutBibir       bool `json:"pemeriksaan_mulut_bibir"`
	PemeriksaanLeher            bool `json:"pemeriksaan_leher"`
	PemeriksaanDadaPunggung     bool `json:"pemeriksaan_dada_punggung"`
	PemeriksaanKardiovaskuler   bool `json:"pemeriksaan_kardiovaskuler"`
	PemeriksaanAbdomenPerut     bool `json:"pemeriksaan_abdomen_perut"`
	PemeriksaanEkstremitasAtas  bool `json:"pemeriksaan_ekstremitas_atas"`
	PemeriksaanEkstremitasBawah bool `json:"pemeriksaan_ekstremitas_bawah"`
	PemeriksaanGenitaliaPria    bool `json:"pemeriksaan_genitalia_pria"`
}

type Diagnosa struct {
	DiagnosaId          int    `json:"diagnosa_id"`
	PemeriksaanDokterId int    `json:"pemeriksaan_dokter_id"`
	Diagnosa            string `json:"diagnosa"`
	Jenis               string `json:"jenis"`
	Kasus               string `json:"kasus"`
	StatusDiagnosis     string `json:"status_diagnosis"`
}

type Anatomi struct {
	AnatomiId           int    `json:"anatomi_id"`
	PasienId            int    `json:"pasien_id"`
	PemeriksaanDokterId int    `json:"pemeriksaan_dokter_id"`
	BagianTubuh         string `json:"bagian_tubuh"`
	Keterangan          string `json:"keterangan"`
}
