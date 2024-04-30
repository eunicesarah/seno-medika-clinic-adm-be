package pasien

import (
	"seno-medika.com/config/db"
	"seno-medika.com/model/person"
)

func FindPasienById(id int) (person.Pasien, error) {
	var pasienVar person.Pasien

	err := db.DB.QueryRow("SELECT * FROM pasien WHERE pasien_id = $1", id).Scan(
		&pasienVar.PasienID,
		&pasienVar.NoERM,
		&pasienVar.PasienUUID,
		&pasienVar.NoRMLama,
		&pasienVar.NoDokRM,
		&pasienVar.Penjamin,
		&pasienVar.NoPenjamin,
		&pasienVar.NIK,
		&pasienVar.NoKK,
		&pasienVar.Nama,
		&pasienVar.TempatLahir,
		&pasienVar.TanggalLahir,
		&pasienVar.NoIHS,
		&pasienVar.JenisKelamin,
		&pasienVar.GolonganDarah,
		&pasienVar.NoTelpon,
		&pasienVar.Email,
		&pasienVar.Provinsi,
		&pasienVar.KabupatenKota,
		&pasienVar.Kecamatan,
		&pasienVar.Kelurahan,
		&pasienVar.Alamat,
		&pasienVar.NamaKontakDarurat,
		&pasienVar.NomorKontakDarurat,
		&pasienVar.Pekerjaan,
		&pasienVar.Agama,
		&pasienVar.WargaNegara,
		&pasienVar.Pendidikan,
		&pasienVar.StatusPerkawinan,
		&pasienVar.CreatedAt,
		&pasienVar.CreatedBy,
		&pasienVar.UpdatedAt,
		&pasienVar.UpdatedBy,
	)
	if err != nil {
		return person.Pasien{}, err
	}

	return pasienVar, nil
}

func FindPasienByUuid(uid string) (person.Pasien, error) {
	var pasienVar person.Pasien

	err := db.DB.QueryRow("SELECT * FROM pasien WHERE pasien_uuid = $1", uid).Scan(
		&pasienVar.PasienID,
		&pasienVar.NoERM,
		&pasienVar.PasienUUID,
		&pasienVar.NoRMLama,
		&pasienVar.NoDokRM,
		&pasienVar.Penjamin,
		&pasienVar.NoPenjamin,
		&pasienVar.NIK,
		&pasienVar.NoKK,
		&pasienVar.Nama,
		&pasienVar.TempatLahir,
		&pasienVar.TanggalLahir,
		&pasienVar.NoIHS,
		&pasienVar.JenisKelamin,
		&pasienVar.GolonganDarah,
		&pasienVar.NoTelpon,
		&pasienVar.Email,
		&pasienVar.Provinsi,
		&pasienVar.KabupatenKota,
		&pasienVar.Kecamatan,
		&pasienVar.Kelurahan,
		&pasienVar.Alamat,
		&pasienVar.NamaKontakDarurat,
		&pasienVar.NomorKontakDarurat,
		&pasienVar.Pekerjaan,
		&pasienVar.Agama,
		&pasienVar.WargaNegara,
		&pasienVar.Pendidikan,
		&pasienVar.StatusPerkawinan,
		&pasienVar.CreatedAt,
		&pasienVar.CreatedBy,
		&pasienVar.UpdatedAt,
		&pasienVar.UpdatedBy,
	)
	if err != nil {
		return person.Pasien{}, err
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
			&eachPasien.NoPenjamin,
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

func FindPasienByNIK(nik string) (person.Pasien, error) {
	var pasienVar person.Pasien

	err := db.DB.QueryRow("SELECT * FROM pasien WHERE nik = $1", nik).Scan(
		&pasienVar.PasienID,
		&pasienVar.NoERM,
		&pasienVar.PasienUUID,
		&pasienVar.NoRMLama,
		&pasienVar.NoDokRM,
		&pasienVar.Penjamin,
		&pasienVar.NoPenjamin,
		&pasienVar.NIK,
		&pasienVar.NoKK,
		&pasienVar.Nama,
		&pasienVar.TempatLahir,
		&pasienVar.TanggalLahir,
		&pasienVar.NoIHS,
		&pasienVar.JenisKelamin,
		&pasienVar.GolonganDarah,
		&pasienVar.NoTelpon,
		&pasienVar.Email,
		&pasienVar.Provinsi,
		&pasienVar.KabupatenKota,
		&pasienVar.Kecamatan,
		&pasienVar.Kelurahan,
		&pasienVar.Alamat,
		&pasienVar.NamaKontakDarurat,
		&pasienVar.NomorKontakDarurat,
		&pasienVar.Pekerjaan,
		&pasienVar.Agama,
		&pasienVar.WargaNegara,
		&pasienVar.Pendidikan,
		&pasienVar.StatusPerkawinan,
		&pasienVar.CreatedAt,
		&pasienVar.CreatedBy,
		&pasienVar.UpdatedAt,
		&pasienVar.UpdatedBy,
	)
	if err != nil {
		return person.Pasien{}, err
	}

	return pasienVar, nil
}

func FindPasienByName(name string) ([]person.Pasien, error) {
	var pasienVar []person.Pasien

	val, err := db.DB.Query("SELECT * FROM pasien WHERE nama LIKE $1", name)
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
			&eachPasien.NoPenjamin,
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
