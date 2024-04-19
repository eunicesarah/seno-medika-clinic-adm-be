package nurse

import (
	"seno-medika.com/config/db"
	"seno-medika.com/model/antrian"
)

func ListAntrianNurse() ([]antrian.AntrianNurse, error) {
    var antrianList []antrian.AntrianNurse

    rows, err := db.DB.Query("SELECT a.nomor_antrian, a.poli, a.created_at, p.no_erm, p.nik, p.nama, p.jenis_kelamin, p.tempat_lahir, p.tanggal_lahir, p.penjamin FROM antrian a JOIN pasien p ON p.pasien_id = a.pasien_id")
    if err != nil {
        return nil, err
    }
    for rows.Next() {
        var eachAntrian antrian.AntrianNurse
        err := rows.Scan(&eachAntrian.NomorAntrian, &eachAntrian.Poli, &eachAntrian.CreatedAt, &eachAntrian.NoERM, &eachAntrian.NIK, &eachAntrian.Nama, &eachAntrian.JenisKelamin, &eachAntrian.TempatLahir, &eachAntrian.TanggalLahir, &eachAntrian.Penjamin)
        if err != nil {
            return nil, err
        }
        antrianList = append(antrianList, eachAntrian)
    }
    return antrianList, nil
}
