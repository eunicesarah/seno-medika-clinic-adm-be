package pasien

import (
	"seno-medika.com/config/db"
	"seno-medika.com/model/person"
)

func FindPasienById(id int) ([]person.Pasien, error) {
	var pasienVar []person.Pasien

	val, err := db.DB.Query("SELECT * FROM pasien WHERE pasien_id = $1", id)
	if err != nil {
		return nil, err
	}

	for val.Next() {
		var eachPasien person.Pasien
		err := val.Scan(
			&eachPasien.PasienID,
			&eachPasien.NoERM,
			&eachPasien.PasienUUID,
			&eachPasien.NoRMLama,
			&eachPasien.NoDokRM,
			&eachPasien.Penjamin,
			&eachPasien.NIK,
			&eachPasien.NoKK,
			&eachPasien.Nama,
			&eachPasien.TempatLahir,
			&eachPasien.TanggalLahir,
			&eachPasien.NoIHS,
			&eachPasien.JenisKelamin,
			&eachPasien.GolonganDarah,
			&eachPasien.NoTelpon,
			&eachPasien.Email,
			&eachPasien.Provinsi,
			&eachPasien.KabupatenKota,
			&eachPasien.Kecamatan,
			&eachPasien.Kelurahan,
			&eachPasien.Alamat,
			&eachPasien.NamaKontakDarurat,
			&eachPasien.NomorKontakDarurat,
			&eachPasien.Pekerjaan,
			&eachPasien.Agama,
			&eachPasien.WargaNegara,
			&eachPasien.Pendidikan,
			&eachPasien.StatusPerkawinan,
			&eachPasien.CreatedAt,
			&eachPasien.CreatedBy,
			&eachPasien.UpdatedAt,
			&eachPasien.UpdatedBy,
		)

		if err != nil {
			return nil, err
		}
		pasienVar = append(pasienVar, eachPasien)
	}

	return pasienVar, nil
}

func FindPasienByUuid(uid string) ([]person.Pasien, error) {
	var pasienVar []person.Pasien

	val, err := db.DB.Query("SELECT * FROM pasien WHERE pasien_uuid = $1", uid)
	if err != nil {
		return nil, err
	}

	for val.Next() {
		var eachPasien person.Pasien
		err := val.Scan(
			&eachPasien.PasienID,
			&eachPasien.NoERM,
			&eachPasien.PasienUUID,
			&eachPasien.NoRMLama,
			&eachPasien.NoDokRM,
			&eachPasien.Penjamin,
			&eachPasien.NIK,
			&eachPasien.NoKK,
			&eachPasien.Nama,
			&eachPasien.TempatLahir,
			&eachPasien.TanggalLahir,
			&eachPasien.NoIHS,
			&eachPasien.JenisKelamin,
			&eachPasien.GolonganDarah,
			&eachPasien.NoTelpon,
			&eachPasien.Email,
			&eachPasien.Provinsi,
			&eachPasien.KabupatenKota,
			&eachPasien.Kecamatan,
			&eachPasien.Kelurahan,
			&eachPasien.Alamat,
			&eachPasien.NamaKontakDarurat,
			&eachPasien.NomorKontakDarurat,
			&eachPasien.Pekerjaan,
			&eachPasien.Agama,
			&eachPasien.WargaNegara,
			&eachPasien.Pendidikan,
			&eachPasien.StatusPerkawinan,
			&eachPasien.CreatedAt,
			&eachPasien.CreatedBy,
			&eachPasien.UpdatedAt,
			&eachPasien.UpdatedBy,
		)

		if err != nil {
			return nil, err
		}
		pasienVar = append(pasienVar, eachPasien)
	}

	return pasienVar, nil
}

func FindPasienAll() ([]person.Pasien, error) {
	var pasienVar []person.Pasien

	val, err := db.DB.Query("SELECT * FROM pasien")
	if err != nil {
		return nil, err
	}

	for val.Next() {
		var eachPasien person.Pasien
		err := val.Scan(
			&eachPasien.PasienID,
			&eachPasien.NoERM,
			&eachPasien.PasienUUID,
			&eachPasien.NoRMLama,
			&eachPasien.NoDokRM,
			&eachPasien.Penjamin,
			&eachPasien.NIK,
			&eachPasien.NoKK,
			&eachPasien.Nama,
			&eachPasien.TempatLahir,
			&eachPasien.TanggalLahir,
			&eachPasien.NoIHS,
			&eachPasien.JenisKelamin,
			&eachPasien.GolonganDarah,
			&eachPasien.NoTelpon,
			&eachPasien.Email,
			&eachPasien.Provinsi,
			&eachPasien.KabupatenKota,
			&eachPasien.Kecamatan,
			&eachPasien.Kelurahan,
			&eachPasien.Alamat,
			&eachPasien.NamaKontakDarurat,
			&eachPasien.NomorKontakDarurat,
			&eachPasien.Pekerjaan,
			&eachPasien.Agama,
			&eachPasien.WargaNegara,
			&eachPasien.Pendidikan,
			&eachPasien.StatusPerkawinan,
			&eachPasien.CreatedAt,
			&eachPasien.CreatedBy,
			&eachPasien.UpdatedAt,
			&eachPasien.UpdatedBy,
		)

		if err != nil {
			return nil, err
		}
		pasienVar = append(pasienVar, eachPasien)
	}

	return pasienVar, nil
}

func FindPasienByNIK(nik string) ([]person.Pasien, error) {
	var pasienVar []person.Pasien

	val, err := db.DB.Query("SELECT * FROM pasien WHERE nik = $1", nik)
	if err != nil {
		return nil, err
	}

	for val.Next() {
		var eachPasien person.Pasien
		err := val.Scan(
			&eachPasien.PasienID,
			&eachPasien.NoERM,
			&eachPasien.PasienUUID,
			&eachPasien.NoRMLama,
			&eachPasien.NoDokRM,
			&eachPasien.Penjamin,
			&eachPasien.NIK,
			&eachPasien.NoKK,
			&eachPasien.Nama,
			&eachPasien.TempatLahir,
			&eachPasien.TanggalLahir,
			&eachPasien.NoIHS,
			&eachPasien.JenisKelamin,
			&eachPasien.GolonganDarah,
			&eachPasien.NoTelpon,
			&eachPasien.Email,
			&eachPasien.Provinsi,
			&eachPasien.KabupatenKota,
			&eachPasien.Kecamatan,
			&eachPasien.Kelurahan,
			&eachPasien.Alamat,
			&eachPasien.NamaKontakDarurat,
			&eachPasien.NomorKontakDarurat,
			&eachPasien.Pekerjaan,
			&eachPasien.Agama,
			&eachPasien.WargaNegara,
			&eachPasien.Pendidikan,
			&eachPasien.StatusPerkawinan,
			&eachPasien.CreatedAt,
			&eachPasien.CreatedBy,
			&eachPasien.UpdatedAt,
			&eachPasien.UpdatedBy,
		)

		if err != nil {
			return nil, err
		}
		pasienVar = append(pasienVar, eachPasien)
	}

	return pasienVar, nil
}
