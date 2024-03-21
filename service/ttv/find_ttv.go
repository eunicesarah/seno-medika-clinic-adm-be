package ttv

import (
	"seno-medika.com/config/db"
	"seno-medika.com/model/doctorstation"
	"seno-medika.com/model/nursestation"
)

func FindAllSkriningAwal() ([]nursestation.SkriningAwal, error) {
	var skriningAwal []nursestation.SkriningAwal
	rows, err := db.DB.Query("SELECT * FROM skrining_awal")
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var skriningAwalVar nursestation.SkriningAwal
		if err := rows.Scan(&skriningAwalVar.SkriningAwalID, &skriningAwalVar.Disabilitas, &skriningAwalVar.Ambulansi, &skriningAwalVar.HambatanKomunikasi, &skriningAwalVar.JalanTidakSeimbang, &skriningAwalVar.JalanAlatBantu, &skriningAwalVar.MenopangSaatDuduk, &skriningAwalVar.HasilCaraJalan, &skriningAwalVar.SkalaNyeri, &skriningAwalVar.NyeriBerulang); err != nil {
			return nil, err
		}
		skriningAwal = append(skriningAwal, skriningAwalVar)
	}

	return skriningAwal, nil
}

func FindSkriningAwalById(id string) (nursestation.SkriningAwal, error) {
	var skriningAwalVar nursestation.SkriningAwal
	if err := db.DB.QueryRow("SELECT * FROM skrining_awal WHERE skrining_awal_id = $1", id).Scan(&skriningAwalVar.SkriningAwalID, &skriningAwalVar.Disabilitas, &skriningAwalVar.Ambulansi, &skriningAwalVar.HambatanKomunikasi, &skriningAwalVar.JalanTidakSeimbang, &skriningAwalVar.JalanAlatBantu, &skriningAwalVar.MenopangSaatDuduk, &skriningAwalVar.HasilCaraJalan, &skriningAwalVar.SkalaNyeri, &skriningAwalVar.NyeriBerulang); err != nil {
		return nursestation.SkriningAwal{}, err
	}

	return skriningAwalVar, nil
}

func FindAllSkriningGizi() ([]nursestation.SkriningGizi, error) {
	var skriningGizi []nursestation.SkriningGizi
	rows, err := db.DB.Query("SELECT * FROM skrining_gizi")
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var skriningGiziVar nursestation.SkriningGizi
		if err := rows.Scan(&skriningGiziVar.SkriningGiziID, &skriningGiziVar.PenurunanBB, &skriningGiziVar.TdkNafsuMakan, &skriningGiziVar.DiagnosisKhusus, &skriningGiziVar.NamaPenyakit, &skriningGiziVar.SkalaNyeri, &skriningGiziVar.NyeriBerulang, &skriningGiziVar.SifatNyeri); err != nil {
			return nil, err
		}
		skriningGizi = append(skriningGizi, skriningGiziVar)
	}

	return skriningGizi, nil
}

func FindSkriningGiziById(id string) (nursestation.SkriningGizi, error) {
	var skriningGiziVar nursestation.SkriningGizi
	if err := db.DB.QueryRow("SELECT * FROM skrining_gizi WHERE skrining_gizi_id = $1", id).Scan(&skriningGiziVar.SkriningGiziID, &skriningGiziVar.PenurunanBB, &skriningGiziVar.TdkNafsuMakan, &skriningGiziVar.DiagnosisKhusus, &skriningGiziVar.NamaPenyakit, &skriningGiziVar.SkalaNyeri, &skriningGiziVar.NyeriBerulang, &skriningGiziVar.SifatNyeri); err != nil {
		return nursestation.SkriningGizi{}, err
	}

	return skriningGiziVar, nil
}

func FindAllTTV() ([]nursestation.TTV, error) {
	var ttv []nursestation.TTV
	rows, err := db.DB.Query("SELECT * FROM ttv")
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var ttvVar nursestation.TTV
		if err := rows.Scan(&ttvVar.TTVID, &ttvVar.Kesadaran, &ttvVar.Sistole, &ttvVar.Diastole, &ttvVar.TinggiBadan, &ttvVar.CaraUkurTB, &ttvVar.BeratBadan, &ttvVar.LingkarPerut, &ttvVar.DetakNadi, &ttvVar.Nafas, &ttvVar.Saturasi, &ttvVar.Suhu, &ttvVar.DetakJantung, &ttvVar.Triage, &ttvVar.PsikolososialSpirit, &ttvVar.Keterangan); err != nil {
			return nil, err
		}
		ttv = append(ttv, ttvVar)
	}

	return ttv, nil
}

func FindTTVById(id string) (nursestation.TTV, error) {
	var ttvVar nursestation.TTV
	if err := db.DB.QueryRow("SELECT * FROM ttv WHERE ttv_id = $1", id).Scan(&ttvVar.TTVID, &ttvVar.Kesadaran, &ttvVar.Sistole, &ttvVar.Diastole, &ttvVar.TinggiBadan, &ttvVar.CaraUkurTB, &ttvVar.BeratBadan, &ttvVar.LingkarPerut, &ttvVar.DetakNadi, &ttvVar.Nafas, &ttvVar.Saturasi, &ttvVar.Suhu, &ttvVar.DetakJantung, &ttvVar.Triage, &ttvVar.PsikolososialSpirit, &ttvVar.Keterangan); err != nil {
		return nursestation.TTV{}, err
	}

	return ttvVar, nil
}

func FindAllRiwayatPenyakit() ([]nursestation.RiwayatPenyakit, error) {
	var riwayatPenyakit []nursestation.RiwayatPenyakit
	rows, err := db.DB.Query("SELECT * FROM riwayat_penyakit")
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var riwayatPenyakitVar nursestation.RiwayatPenyakit
		if err := rows.Scan(&riwayatPenyakitVar.RiwayatPenyakitID, &riwayatPenyakitVar.RPS, &riwayatPenyakitVar.RPD, &riwayatPenyakitVar.RPK); err != nil {
			return nil, err
		}
		riwayatPenyakit = append(riwayatPenyakit, riwayatPenyakitVar)
	}

	return riwayatPenyakit, nil
}

func FindRiwayatPenyakitById(id string) (nursestation.RiwayatPenyakit, error) {
	var riwayatPenyakitVar nursestation.RiwayatPenyakit
	if err := db.DB.QueryRow("SELECT * FROM riwayat_penyakit WHERE riwayat_penyakit_id = $1", id).Scan(&riwayatPenyakitVar.RiwayatPenyakitID, &riwayatPenyakitVar.RPS, &riwayatPenyakitVar.RPD, &riwayatPenyakitVar.RPK); err != nil {
		return nursestation.RiwayatPenyakit{}, err
	}

	return riwayatPenyakitVar, nil
}

func FindAllAlergi() ([]doctorstation.Alergi, error) {
	var alergi []doctorstation.Alergi
	rows, err := db.DB.Query("SELECT * FROM alergi")
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var alergiVar doctorstation.Alergi
		if err := rows.Scan(&alergiVar.AlergiID, &alergiVar.Obat, &alergiVar.Makanan, &alergiVar.Lainnya); err != nil {
			return nil, err
		}
		alergi = append(alergi, alergiVar)
	}

	return alergi, nil
}

func FindAlergiById(id string) (doctorstation.Alergi, error) {
	var alergiVar doctorstation.Alergi
	if err := db.DB.QueryRow("SELECT * FROM alergi WHERE alergi_id = $1", id).Scan(&alergiVar.AlergiID, &alergiVar.Obat, &alergiVar.Makanan, &alergiVar.Lainnya); err != nil {
		return doctorstation.Alergi{}, err
	}

	return alergiVar, nil
}

func FindAllAnamnesis() ([]doctorstation.Anamnesis, error) {
	var anamnesis []doctorstation.Anamnesis
	rows, err := db.DB.Query("SELECT * FROM anamnesis")
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var anamnesisVar doctorstation.Anamnesis
		if err := rows.Scan(&anamnesisVar.AnamnesisID, &anamnesisVar.PasienID, &anamnesisVar.SkrinAwalID, &anamnesisVar.SkrinGiziID, &anamnesisVar.TTVID, &anamnesisVar.RiwayatPenyakitID, &anamnesisVar.AlergiID, &anamnesisVar.DokterID, &anamnesisVar.PerawatID); err != nil {
			return nil, err
		}
		anamnesis = append(anamnesis, anamnesisVar)
	}

	return anamnesis, nil
}

func FindAnamnesisById(id string) (doctorstation.Anamnesis, error) {
	var anamnesisVar doctorstation.Anamnesis
	if err := db.DB.QueryRow("SELECT * FROM anamnesis WHERE anamnesis_id = $1", id).Scan(&anamnesisVar.AnamnesisID, &anamnesisVar.PasienID, &anamnesisVar.SkrinAwalID, &anamnesisVar.SkrinGiziID, &anamnesisVar.TTVID, &anamnesisVar.RiwayatPenyakitID, &anamnesisVar.AlergiID, &anamnesisVar.DokterID, &anamnesisVar.PerawatID); err != nil {
		return doctorstation.Anamnesis{}, err
	}

	return anamnesisVar, nil
}
