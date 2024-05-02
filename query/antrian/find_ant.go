package antrian

import (
	"seno-medika.com/config/db"
	"seno-medika.com/model/antrian"
)

func FindAntrianById(id int) (antrian.Antrian, error) {
	var antrianO antrian.Antrian
	err := db.DB.QueryRow("SELECT * FROM antrian WHERE antrian_id = $1", id).
		Scan(&antrianO.AntrianID, &antrianO.PasienID, &antrianO.NomorAntrian, &antrianO.Status, &antrianO.Poli, &antrianO.Instalasi, &antrianO.CreatedAt)
	if err != nil {
		return antrian.Antrian{}, err
	}
	return antrianO, nil
}
func FindAntrianByPasienId(id int) (antrian.Antrian, error) {
	var antrianO antrian.Antrian
	var nik string
	err := db.DB.QueryRow("SELECT nik FROM pasien WHERE id = $1", id).
		Scan(nik)
	if err != nil {
		return antrian.Antrian{}, err
	}
	err = db.DB.QueryRow("SELECT * FROM antrian WHERE nik = $1", nik).
		Scan(&antrianO.AntrianID, &antrianO.PasienID, &antrianO.NomorAntrian, &antrianO.Status, &antrianO.Poli, &antrianO.Instalasi, &antrianO.CreatedAt)
	if err != nil {
		return antrian.Antrian{}, err
	}
	return antrianO, nil
}

func FindAntrianAll() ([]antrian.Antrian, error) {
	var antrianO []antrian.Antrian
	rows, err := db.DB.Query("SELECT * FROM antrian")
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var eachAntrian antrian.Antrian
		err := rows.Scan(&eachAntrian.AntrianID, &eachAntrian.PasienID, &eachAntrian.NomorAntrian, &eachAntrian.Status, &eachAntrian.Poli, &eachAntrian.Instalasi, &eachAntrian.CreatedAt)
		if err != nil {
			return nil, err
		}
		antrianO = append(antrianO, eachAntrian)
	}
	return antrianO, nil
}

func FindAntrianFilter(search string, page string, limit string, date string, poli string) ([]antrian.AntrianNurse, int, error) {
	var antrianO []antrian.AntrianNurse
	var size int

	if search != "" {
		if poli != "" {
			rows, err := db.DB.Query("SELECT a.antrian_id, a.nomor_antrian, a.poli, a.created_at, p.pasien_id, p.no_erm, p.nik, p.nama, p.jenis_kelamin, p.tempat_lahir, p.tanggal_lahir, p.penjamin FROM antrian a JOIN pasien p ON p.pasien_id = a.pasien_id"+
				" WHERE p.nama ILIKE $1 AND a.poli = $2 AND a.created_at = $3 ORDER BY a.nomor_antrian ASC LIMIT $4 OFFSET $5", "%"+search+"%", poli, date, limit, page)

			_ = db.DB.QueryRow("SELECT COUNT(*) FROM antrian a JOIN pasien p ON p.pasien_id = a.pasien_id"+
				" WHERE p.nama ILIKE $1 AND a.poli = $2 AND a.created_at = $3", "%"+search+"%", poli, date).Scan(&size)

			if err != nil {
				return nil, 0, err
			}

			for rows.Next() {
				var eachAntrian antrian.AntrianNurse
				err := rows.Scan(&eachAntrian.AntrianID, &eachAntrian.NomorAntrian, &eachAntrian.Poli, &eachAntrian.CreatedAt, &eachAntrian.PasienID, &eachAntrian.NoERM, &eachAntrian.NIK, &eachAntrian.Nama, &eachAntrian.JenisKelamin, &eachAntrian.TempatLahir, &eachAntrian.TanggalLahir, &eachAntrian.Penjamin)
				if err != nil {
					return nil, 0, err
				}
				antrianO = append(antrianO, eachAntrian)
			}

			return antrianO, size, nil
		}

		rows, err := db.DB.Query("SELECT a.antrian_id, a.nomor_antrian, a.poli, a.created_at, p.pasien_id, p.no_erm, p.nik, p.nama, p.jenis_kelamin, p.tempat_lahir, p.tanggal_lahir, p.penjamin FROM antrian a JOIN pasien p ON p.pasien_id = a.pasien_id"+
			" WHERE p.nama ILIKE $1 AND a.created_at = $2 ORDER BY a.nomor_antrian ASC LIMIT $3 OFFSET $4", "%"+search+"%", date, limit, page)

		_ = db.DB.QueryRow("SELECT COUNT(*) FROM antrian a JOIN pasien p ON p.pasien_id = a.pasien_id"+
			" WHERE p.nama ILIKE $1 AND a.created_at = $2", "%"+search+"%", date).Scan(&size)

		if err != nil {
			return nil, 0, err
		}

		for rows.Next() {
			var eachAntrian antrian.AntrianNurse
			err := rows.Scan(&eachAntrian.AntrianID, &eachAntrian.NomorAntrian, &eachAntrian.Poli, &eachAntrian.CreatedAt, &eachAntrian.PasienID, &eachAntrian.NoERM, &eachAntrian.NIK, &eachAntrian.Nama, &eachAntrian.JenisKelamin, &eachAntrian.TempatLahir, &eachAntrian.TanggalLahir, &eachAntrian.Penjamin)
			if err != nil {
				return nil, 0, err
			}
			antrianO = append(antrianO, eachAntrian)
		}

		return antrianO, size, nil

	}

	if poli != "" {
		rows, err := db.DB.Query("SELECT a.antrian_id, a.nomor_antrian, a.poli, a.created_at, p.pasien_id, p.no_erm, p.nik, p.nama, p.jenis_kelamin, p.tempat_lahir, p.tanggal_lahir, p.penjamin FROM antrian a JOIN pasien p ON p.pasien_id = a.pasien_id"+
			" WHERE a.poli LIKE $1 AND a.created_at = $2 ORDER BY a.nomor_antrian ASC LIMIT $3 OFFSET $4", poli, date, limit, page)

		_ = db.DB.QueryRow("SELECT COUNT(*) FROM antrian a JOIN pasien p ON p.pasien_id = a.pasien_id"+
			" WHERE a.poli LIKE $1 AND a.created_at = $2", poli, date).Scan(&size)

		if err != nil {
			return nil, 0, err
		}

		for rows.Next() {
			var eachAntrian antrian.AntrianNurse
			err := rows.Scan(&eachAntrian.AntrianID, &eachAntrian.NomorAntrian, &eachAntrian.Poli, &eachAntrian.CreatedAt, &eachAntrian.PasienID, &eachAntrian.NoERM, &eachAntrian.NIK, &eachAntrian.Nama, &eachAntrian.JenisKelamin, &eachAntrian.TempatLahir, &eachAntrian.TanggalLahir, &eachAntrian.Penjamin)
			if err != nil {
				return nil, 0, err
			}
			antrianO = append(antrianO, eachAntrian)
		}

		return antrianO, size, nil
	}

	rows, err := db.DB.Query("SELECT a.antrian_id, a.nomor_antrian, a.poli, a.created_at, p.pasien_id, p.no_erm, p.nik, p.nama, p.jenis_kelamin, p.tempat_lahir, p.tanggal_lahir, p.penjamin FROM antrian a JOIN pasien p ON p.pasien_id = a.pasien_id"+
		" WHERE a.created_at = $1 ORDER BY a.nomor_antrian ASC LIMIT $2 OFFSET $3", date, limit, page)

	_ = db.DB.QueryRow("SELECT COUNT(*) FROM antrian a JOIN pasien p ON p.pasien_id = a.pasien_id"+
		" WHERE a.created_at = $1", date).Scan(&size)

	if err != nil {
		return nil, 0, err
	}

	for rows.Next() {
		var eachAntrian antrian.AntrianNurse
		err := rows.Scan(&eachAntrian.AntrianID, &eachAntrian.NomorAntrian, &eachAntrian.Poli, &eachAntrian.CreatedAt, &eachAntrian.PasienID, &eachAntrian.NoERM, &eachAntrian.NIK, &eachAntrian.Nama, &eachAntrian.JenisKelamin, &eachAntrian.TempatLahir, &eachAntrian.TanggalLahir, &eachAntrian.Penjamin)
		if err != nil {
			return nil, 0, err
		}
		antrianO = append(antrianO, eachAntrian)
	}

	return antrianO, size, nil
}
