package pasien

import (
	"errors"
	"seno-medika.com/config/db"
	"seno-medika.com/model/person"
)

func UpdatePasienByUuid(uid string, pasien person.Pasien) error {
	val, err := db.DB.Exec(
		`UPDATE pasien SET
			no_erm = $1,
			no_rm_lama = $2,
			no_dok_rm = $3,
			penjamin = $4,
			no_penjamin = $5,
			nik = $6,
			no_kk = $7,
			nama = $8,
			tempat_lahir = $9,
			tanggal_lahir = $10,
			no_ihs = $11,
			jenis_kelamin = $12,
			golongan_darah = $13,
			no_telpon = $14,
			email = $15,
			provinsi = $16,
			kabupaten_kota = $17,
			kecamatan = $18,
			kelurahan = $19,
			alamat = $20,
			nama_kontak_darurat = $21,
			nomor_kontak_darurat = $22,
			pekerjaan = $23,
			agama = $24,
			warga_negara = $25,
			pendidikan = $26,
			status_perkawinan = $27,
			updated_at = $28,
			updated_by = $29
		WHERE pasien_uuid = $30
		`,
		pasien.NoERM,
		pasien.NoRMLama,
		pasien.NoDokRM,
		pasien.Penjamin,
		pasien.NoPenjamin,
		pasien.NIK,
		pasien.NoKK,
		pasien.Nama,
		pasien.TempatLahir,
		pasien.TanggalLahir,
		pasien.NoIHS,
		pasien.JenisKelamin,
		pasien.GolonganDarah,
		pasien.NoTelpon,
		pasien.Email,
		pasien.Provinsi,
		pasien.KabupatenKota,
		pasien.Kecamatan,
		pasien.Kelurahan,
		pasien.Alamat,
		pasien.NamaKontakDarurat,
		pasien.NomorKontakDarurat,
		pasien.Pekerjaan,
		pasien.Agama,
		pasien.WargaNegara,
		pasien.Pendidikan,
		pasien.StatusPerkawinan,
		pasien.UpdatedAt,
		pasien.UpdatedBy,
		uid)

	if err != nil {
		return err
	}

	if rows, _ := val.RowsAffected(); rows == 0 {
		return errors.New("uuid not found")
	}

	return nil
}

func UpdatePasienById(id int, pasien person.Pasien) error {
	val, err := db.DB.Exec(
		`UPDATE pasien SET
			no_erm = $1,
			no_rm_lama = $2,
			no_dok_rm = $3,
			penjamin = $4,
			no_penjamin = $5,
			nik = $6,
			no_kk = $7,
			nama = $8,
			tempat_lahir = $9,
			tanggal_lahir = $10,
			no_ihs = $11,
			jenis_kelamin = $12,
			golongan_darah = $13,
			no_telpon = $14,
			email = $15,
			provinsi = $16,
			kabupaten_kota = $17,
			kecamatan = $18,
			kelurahan = $19,
			alamat = $20,
			nama_kontak_darurat = $21,
			nomor_kontak_darurat = $22,
			pekerjaan = $23,
			agama = $24,
			warga_negara = $25,
			pendidikan = $26,
			status_perkawinan = $27,
			updated_at = $28,
			updated_by = $29
		WHERE pasien_id = $30
		`,
		pasien.NoERM,
		pasien.NoRMLama,
		pasien.NoDokRM,
		pasien.Penjamin,
		pasien.NoPenjamin,
		pasien.NIK,
		pasien.NoKK,
		pasien.Nama,
		pasien.TempatLahir,
		pasien.TanggalLahir,
		pasien.NoIHS,
		pasien.JenisKelamin,
		pasien.GolonganDarah,
		pasien.NoTelpon,
		pasien.Email,
		pasien.Provinsi,
		pasien.KabupatenKota,
		pasien.Kecamatan,
		pasien.Kelurahan,
		pasien.Alamat,
		pasien.NamaKontakDarurat,
		pasien.NomorKontakDarurat,
		pasien.Pekerjaan,
		pasien.Agama,
		pasien.WargaNegara,
		pasien.Pendidikan,
		pasien.StatusPerkawinan,
		pasien.UpdatedAt,
		pasien.UpdatedBy,
		id)

	if err != nil {
		return err
	}

	if rows, _ := val.RowsAffected(); rows == 0 {
		return errors.New("id not found")
	}

	return nil
}
