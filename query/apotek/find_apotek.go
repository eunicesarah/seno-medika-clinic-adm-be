package apotek

import (
	pharmacystation2 "seno-medika.com/model/station/pharmacystation"
	"time"

	"seno-medika.com/config/db"
)

func FindAllAntrianApotekToday() ([]pharmacystation2.DashboardApotek, error) {

	var apotekVars []pharmacystation2.DashboardApotek
	todayDate := time.Now().Format("2006-01-02")

	rows, err := db.DB.Query("SELECT a.nomor_antrian, a.poli, p.no_erm, a.created_at, p.nik, p.nama, p.jenis_kelamin, p.tempat_lahir, p.tanggal_lahir, n.metode_pembayaran, r.status_obat FROM antrian a, pasien p, nota n, resep r WHERE a.pasien_id = p.pasien_id AND a.pasien_id=n.pasien_id AND n.resep_id = r.resep_id  AND a.created_at = $1", todayDate)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var apotekVar pharmacystation2.DashboardApotek
		if err := rows.Scan(&apotekVar.NomorAntrian, &apotekVar.Poli, &apotekVar.NoERM, &apotekVar.CreatedAt, &apotekVar.NIK, &apotekVar.Nama, &apotekVar.JenisKelamin, &apotekVar.TempatLahir, &apotekVar.TanggalLahir, &apotekVar.MetodePembayaran, &apotekVar.Status); err != nil {
			return nil, err
		}
		apotekVars = append(apotekVars, apotekVar)
	}
	return apotekVars, nil
}

func FindAllAntrianApotek() ([]pharmacystation2.DashboardApotek, error) {

	var apotekVars []pharmacystation2.DashboardApotek

	rows, err := db.DB.Query("SELECT a.nomor_antrian, a.poli, p.no_erm, a.created_at, p.nik, p.nama, p.jenis_kelamin, p.tempat_lahir, p.tanggal_lahir, n.metode_pembayaran, r.status_obat FROM antrian a, pasien p, nota n, resep r WHERE a.pasien_id = p.pasien_id AND a.pasien_id=n.pasien_id AND n.resep_id = r.resep_id")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var apotekVar pharmacystation2.DashboardApotek
		if err := rows.Scan(&apotekVar.NomorAntrian, &apotekVar.Poli, &apotekVar.NoERM, &apotekVar.CreatedAt, &apotekVar.NIK, &apotekVar.Nama, &apotekVar.JenisKelamin, &apotekVar.TempatLahir, &apotekVar.TanggalLahir, &apotekVar.MetodePembayaran, &apotekVar.Status); err != nil {
			return nil, err
		}
		apotekVars = append(apotekVars, apotekVar)
	}
	return apotekVars, nil
}

func FindAllAntrianApotekByDate(date string) ([]pharmacystation2.DashboardApotek, error) {
	var apotekVars []pharmacystation2.DashboardApotek

	rows, err := db.DB.Query("SELECT a.nomor_antrian, a.poli, p.no_erm, a.created_at, p.nik, p.nama, p.jenis_kelamin, p.tempat_lahir, p.tanggal_lahir, n.metode_pembayaran, r.status_obat FROM antrian a, pasien p, nota n, resep r WHERE a.pasien_id = p.pasien_id AND a.pasien_id=n.pasien_id AND n.resep_id = r.resep_id  AND a.created_at = $1", date)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var apotekVar pharmacystation2.DashboardApotek
		if err := rows.Scan(&apotekVar.NomorAntrian, &apotekVar.Poli, &apotekVar.NoERM, &apotekVar.CreatedAt, &apotekVar.NIK, &apotekVar.Nama, &apotekVar.JenisKelamin, &apotekVar.TempatLahir, &apotekVar.TanggalLahir, &apotekVar.MetodePembayaran, &apotekVar.Status); err != nil {
			return nil, err
		}
		apotekVars = append(apotekVars, apotekVar)
	}
	return apotekVars, nil
}

func FindDetailResepByNoAntrian(no_antrian int) ([]pharmacystation2.DetailObat, error) {
	var details []pharmacystation2.DetailObat

	pasien_id := db.DB.QueryRow("SELECT pasien_id FROM antrian WHERE nomor_antrian = $1", no_antrian)

	var pasien_id_var int
	if err := pasien_id.Scan(&pasien_id_var); err != nil {
		return nil, err
	}

	rows, err := db.DB.Query("SELECT o.nama_obat,  o.satuan, lo.jumlah, lo.dosis, lo.keterangan, lo.aturan_pakai FROM nota n "+
		"INNER JOIN resep r ON n.resep_id = r.resep_id "+
		"INNER JOIN list_obat lo ON lo.resep_id = r.resep_id "+
		"INNER JOIN obat o ON lo.obat_id = o.obat_id "+
		"WHERE n.pasien_id = $1", pasien_id_var)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var detail pharmacystation2.DetailObat
		err := rows.Scan(&detail.Obat.NamaObat, &detail.Obat.Satuan, &detail.ListObat.Jumlah, &detail.ListObat.Dosis, &detail.ListObat.Keterangan, &detail.ListObat.AturanPakai)
		if err != nil {
			return nil, err
		}
		details = append(details, detail)
	}

	if len(details) == 0 {
		return nil, err
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return details, nil
}
