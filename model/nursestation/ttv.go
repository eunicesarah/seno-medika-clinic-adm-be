package nursestation

import "seno-medika.com/model/doctorstation"

type NurseStation struct {
	NurseStationID  int                     `json:"nurse_station_id"`
	SkriningAwal    SkriningAwal            `json:"skrining_awal"`
	SkriningGizi    SkriningGizi            `json:"skrining_gizi"`
	TTV             TTV                     `json:"ttv"`
	RiwayatPenyakit RiwayatPenyakit         `json:"riwayat_penyakit"`
	Alergi          doctorstation.Alergi    `json:"alergi"`
	Anamnesis       doctorstation.Anamnesis `json:"anamnesis"`
}

type SkriningAwal struct {
	SkriningAwalID     int    `json:"skrining_awal_id"`
	Disabilitas        bool   `json:"disabilitas"`
	Ambulansi          bool   `json:"ambulansi"`
	HambatanKomunikasi bool   `json:"hambatan_komunikasi"`
	JalanTidakSeimbang bool   `json:"jalan_tidak_seimbang"`
	JalanAlatBantu     bool   `json:"jalan_alat_bantu"`
	MenopangSaatDuduk  bool   `json:"menopang_saat_duduk"`
	HasilCaraJalan     string `json:"hasil_cara_jalan"`
	SkalaNyeri         int    `json:"skala_nyeri"`
	NyeriBerulang      string `json:"nyeri_berulang"`
}

type SkriningGizi struct {
	SkriningGiziID  int    `json:"skrining_gizi_id"`
	PenurunanBB     int    `json:"penurunan_bb"`
	TdkNafsuMakan   bool   `json:"tdk_nafsu_makan"`
	DiagnosisKhusus bool   `json:"diagnosis_khusus"`
	NamaPenyakit    string `json:"nama_penyakit"`
	SkalaNyeri      int    `json:"skala_nyeri"`
	NyeriBerulang   string `json:"nyeri_berulang"`
	SifatNyeri      string `json:"sifat_nyeri"`
}

type TTV struct {
	TTVID               int    `json:"ttv_id"`
	Kesadaran           string `json:"kesadaran"`
	Sistole             int    `json:"sistole"`
	Diastole            int    `json:"diastole"`
	TinggiBadan         int    `json:"tinggi_badan"`
	CaraUkurTB          string `json:"cara_ukur_tb"`
	BeratBadan          int    `json:"berat_badan"`
	LingkarPerut        int    `json:"lingkar_perut"`
	DetakNadi           int    `json:"detak_nadi"`
	Nafas               int    `json:"nafas"`
	Saturasi            int    `json:"saturasi"`
	Suhu                int    `json:"suhu"`
	DetakJantung        bool   `json:"detak_jantung"`
	Triage              string `json:"triage"`
	PsikolososialSpirit string `json:"psikolososial_spirit"`
	Keterangan          string `json:"keterangan"`
}

type RiwayatPenyakit struct {
	RiwayatPenyakitID int    `json:"riwayat_penyakit_id"`
	RPS               string `json:"rps"`
	RPD               string `json:"rpd"`
	RPK               string `json:"rpk"`
}
