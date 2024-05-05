package ttv

import (
	"seno-medika.com/config/db"
	"seno-medika.com/model/doctorstation"
	"seno-medika.com/model/nursestation"
)

func ChangeSkriningAwalById(id string, skriningAwal nursestation.SkriningAwal) error {
	var skrinAwalID string
	errID := db.DB.QueryRow("SELECT skrin_awal_id FROM anamnesis WHERE pasien_id = $1 ORDER BY skrin_awal_id DESC LIMIT 1", id).Scan(&skrinAwalID)
	if errID != nil {
		return errID
	}
	val, err := db.DB.Exec(
		`UPDATE skrining_awal SET
	           disabilitas = $1,
	           ambulansi = $2,
	           hambatan_komunikasi = $3,
	           jalan_tidak_seimbang = $4,
	           jalan_alat_bantu = $5,
	           menopang_saat_duduk = $6,
	           hasil_cara_jalan = $7,
	           skala_nyeri = $8,
	           nyeri_berulang = $9,
			   sifat_nyeri = $10
	           WHERE skrin_awal_id = $11`,
		skriningAwal.Disabilitas, skriningAwal.Ambulansi, skriningAwal.HambatanKomunikasi, skriningAwal.JalanTidakSeimbang, skriningAwal.JalanAlatBantu, skriningAwal.MenopangSaatDuduk, skriningAwal.HasilCaraJalan, skriningAwal.SkalaNyeri, skriningAwal.NyeriBerulang, skriningAwal.SifatNyeri, skrinAwalID)
	if err != nil {
		return err
	}

	if rows, _ := val.RowsAffected(); rows == 0 {
		return nil
	}

	return nil
}

func ChangeSkriningGiziById(id string, skriningGizi nursestation.SkriningGizi) error {
	var skrinGiziID string
	errID := db.DB.QueryRow("SELECT skrin_gizi_id FROM anamnesis WHERE pasien_id = $1 ORDER BY skrin_gizi_id DESC LIMIT 1", id).Scan(&skrinGiziID)
	if errID != nil {
		return errID
	}
	val, err := db.DB.Exec(
		`UPDATE skrining_gizi SET
	           penurunan_bb = $1,
	           tdk_nafsu_makan = $2,
	           diagnosis_khusus = $3,
	           nama_penyakit = $4,
	           WHERE skrin_gizi_id = $5`,
		skriningGizi.PenurunanBB, skriningGizi.TdkNafsuMakan, skriningGizi.DiagnosisKhusus, skriningGizi.NamaPenyakit, skrinGiziID)
	if err != nil {
		return err
	}

	if rows, _ := val.RowsAffected(); rows == 0 {
		return nil
	}

	return nil
}

func ChangeTTVById(id string, ttv nursestation.TTV) error {
	var ttvID string
	errID := db.DB.QueryRow("SELECT ttv_id FROM anamnesis WHERE pasien_id = $1 ORDER BY ttv_id DESC LIMIT 1", id).Scan(&ttvID)
	if errID != nil {
		return errID
	}
	val, err := db.DB.Exec(
		`UPDATE ttv SET
	           kesadaran = $1,
	           sistole = $2,
	           diastole = $3,
	           tinggi_badan = $4,
	           cara_ukur_tb = $5,
	           berat_badan = $6,
	           lingkar_perut = $7,
	           detak_nadi = $8,
	           nafas = $9,
	           saturasi = $10,
	           suhu = $11,
	           detak_jantung = $12,
	           triage = $13,
	           psikolososial_spirit = $14,
	           keterangan = $15
	           WHERE ttv_id = $16`,
		ttv.Kesadaran, ttv.Sistole, ttv.Diastole, ttv.TinggiBadan, ttv.CaraUkurTB, ttv.BeratBadan, ttv.LingkarPerut, ttv.DetakNadi, ttv.Nafas, ttv.Saturasi, ttv.Suhu, ttv.DetakJantung, ttv.Triage, ttv.PsikolososialSpirit, ttv.Keterangan, ttvID)
	if err != nil {
		return err
	}

	if rows, _ := val.RowsAffected(); rows == 0 {
		return nil
	}

	return nil
}

func ChangeRiwayatPenyakitById(id string, riwayatPenyakit nursestation.RiwayatPenyakit) error {
	var riwayatID string
	errID := db.DB.QueryRow("SELECT riwayat_penyakit_id FROM anamnesis WHERE pasien_id = $1 ORDER BY riwayat_penyakit_id DESC LIMIT 1", id).Scan(&riwayatID)
	if errID != nil {
		return errID
	}
	val, err := db.DB.Exec(
		`UPDATE riwayat_penyakit SET
			   rps = $1,
			   rpd = $2,
			   rpk = $3
			   WHERE riwayat_penyakit_id = $4`,
		riwayatPenyakit.RPS, riwayatPenyakit.RPD, riwayatPenyakit.RPK, riwayatID)
	if err != nil {
		return err
	}

	if rows, _ := val.RowsAffected(); rows == 0 {
		return nil
	}

	return nil
}

func ChangeAlergiById(id string, alergi doctorstation.Alergi) error {
	var alergiID string
	errID := db.DB.QueryRow("SELECT alergi_id FROM anamnesis WHERE pasien_id = $1 ORDER BY alergi_id DESC LIMIT 1", id).Scan(&alergiID)
	if errID != nil {
		return errID
	}
	val, err := db.DB.Exec(
		`UPDATE alergi SET
	           obat = $1,
	           makanan = $2,
	           lainnya = $3
	           WHERE alergi_id = $4`,
		alergi.Obat, alergi.Makanan, alergi.Lainnya, alergiID)
	if err != nil {
		return err
	}

	if rows, _ := val.RowsAffected(); rows == 0 {
		return nil
	}

	return nil
}

func ChangeAnamnesisById(id string, anamnesis doctorstation.Anamnesis) error {
	val, err := db.DB.Exec(
		`UPDATE anamnesis SET
	           pasien_id = $1,
	           skrin_awal_id = $2,
	           skrin_gizi_id = $3,
	           ttv_id = $4,
	           riwayat_penyakit_id = $5,
	           alergi_id = $6,
	           dokter_id = $7,
	           perawat_id = $8,
	           keluhan_utama = $9,
	           keluhan_tambahan = $10,
	           lama_sakit = $11
	           WHERE anamnesis_id = $12`,
		anamnesis.PasienID, anamnesis.SkrinAwalID, anamnesis.SkrinGiziID, anamnesis.TTVID, anamnesis.RiwayatPenyakitID, anamnesis.AlergiID, anamnesis.DokterID, anamnesis.PerawatID, anamnesis.KeluhanUtama, anamnesis.KeluhanTambahan, anamnesis.LamaSakit, id)
	if err != nil {
		return err
	}

	if rows, _ := val.RowsAffected(); rows == 0 {
		return nil
	}

	return nil
}
