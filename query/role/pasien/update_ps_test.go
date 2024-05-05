package pasien

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"seno-medika.com/config/db"
	"seno-medika.com/model/person"
)

func TestUpdatePasienById_Success(t *testing.T) {
	_db := db.DB
	defer func() {
		_db.Exec("DELETE FROM pasien WHERE pasien_id = 8911")
		db.DB = _db
	}()
	_db.Query(`INSERT INTO pasien (
		pasien_id,
		no_erm,
		pasien_uuid,
		no_rm_lama,
		no_dok_rm,
		penjamin,
		no_penjamin,
		nik,
		no_kk,
		nama,
		tempat_lahir,
		tanggal_lahir,
		no_ihs,
		jenis_kelamin,
		golongan_darah,
		no_telpon,
		email,
		provinsi,
		kabupaten_kota,
		kecamatan,
		kelurahan,
		alamat,
		nama_kontak_darurat,
		nomor_kontak_darurat,
		pekerjaan,
		agama,
		warga_negara,
		pendidikan,
		status_perkawinan,
		created_at,
		created_by,
		updated_at,
		updated_by
	   ) VALUES (
		8911, '123', $1, '123', '123', 'BPJS', '123', '330102', '124212',
		'testNama123', 'sumedang', $2, '12345', 'perempuan', 'E', '086123',
		'123@test.go', 'jawa bali', 'jatinangor', 'sayang', 'cikeruh', 'jalan-jalan no 12', 'jokoTest', '0857123',
		'dokter','islam	', 'WNA', 'SMA', 'belum-kawin', $3, '123test', $4, '123test'
	   )
	   `, uuid.New(), time.Now().Local().Format("2006-01-02"), time.Now().Local().String(), time.Now().Local().String())

	err := UpdatePasienById(8911, person.Pasien{
		NoERM:              "31222",
		NoRMLama:           "312321",
		NoDokRM:            "315342",
		Penjamin:           "3143532",
		NoPenjamin:         "31652",
		NIK:                "3126456",
		NoKK:               "312645",
		Nama:               "Joko Doe",
		TempatLahir:        "Indonesia",
		TanggalLahir:       time.Now().Local().Format("2006-01-02"),
		NoIHS:              "31dd2",
		JenisKelamin:       "Perempuan",
		GolonganDarah:      "AB",
		NoTelpon:           "0384312",
		Email:              "alala@gmailc.om",
		Provinsi:           "Java",
		KabupatenKota:      "Bandung",
		Kecamatan:          "Banding",
		Kelurahan:          "Bandeng",
		Alamat:             "Jalan Bandong",
		NamaKontakDarurat:  "nineoneone",
		NomorKontakDarurat: "911",
		Pekerjaan:          "GUru",
		Agama:              "Ateis",
		WargaNegara:        "Indo",
		Pendidikan:         "S0",
		StatusPerkawinan:   "Kawing",
		CreatedAt:          time.Now().Local().String(),
		CreatedBy:          "312",
		UpdatedAt:          time.Now().Local().String(),
		UpdatedBy:          "312",
	})

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	_db.Exec("DELETE FROM pasien WHERE pasien_id = 8911")
	db.DB = _db
}

func TestUpdatePasienById_Fail(t *testing.T) {
	err := UpdatePasienById(221, person.Pasien{
		NoERM:              "31222",
		NoRMLama:           "312321",
		NoDokRM:            "315342",
		Penjamin:           "3143532",
		NoPenjamin:         "31652",
		NIK:                "3126456",
		NoKK:               "312645",
		Nama:               "Joko Doe",
		TempatLahir:        "Indonesia",
		TanggalLahir:       time.Now().Local().Format("2006-01-02"),
		NoIHS:              "31dd2",
		JenisKelamin:       "Perempuan",
		GolonganDarah:      "AB",
		NoTelpon:           "0384312",
		Email:              "alala@gmailc.om",
		Provinsi:           "Java",
		KabupatenKota:      "Bandung",
		Kecamatan:          "Banding",
		Kelurahan:          "Bandeng",
		Alamat:             "Jalan Bandong",
		NamaKontakDarurat:  "nineoneone",
		NomorKontakDarurat: "911",
		Pekerjaan:          "GUru",
		Agama:              "Ateis",
		WargaNegara:        "Indo",
		Pendidikan:         "S0",
		StatusPerkawinan:   "Kawing",
		CreatedAt:          time.Now().Local().String(),
		CreatedBy:          "312",
		UpdatedAt:          time.Now().Local().String(),
		UpdatedBy:          "312",
	})

	if err == nil {
		t.Errorf("Expected error, got nil 5")
	}

	_db := db.DB
	defer func() {
		_db = db.Conn()
		db.DB = _db
	}()

	_db.Close()
	err = UpdatePasienById(221, person.Pasien{
		NoERM:              "31222",
		NoRMLama:           "312321",
		NoDokRM:            "315342",
		Penjamin:           "3143532",
		NoPenjamin:         "31652",
		NIK:                "3126456",
		NoKK:               "312645",
		Nama:               "Joko Doe",
		TempatLahir:        "Indonesia",
		TanggalLahir:       time.Now().Local().Format("2006-01-02"),
		NoIHS:              "31dd2",
		JenisKelamin:       "Perempuan",
		GolonganDarah:      "AB",
		NoTelpon:           "0384312",
		Email:              "alala@gmailc.om",
		Provinsi:           "Java",
		KabupatenKota:      "Bandung",
		Kecamatan:          "Banding",
		Kelurahan:          "Bandeng",
		Alamat:             "Jalan Bandong",
		NamaKontakDarurat:  "nineoneone",
		NomorKontakDarurat: "911",
		Pekerjaan:          "GUru",
		Agama:              "Ateis",
		WargaNegara:        "Indo",
		Pendidikan:         "S0",
		StatusPerkawinan:   "Kawing",
		CreatedAt:          time.Now().Local().String(),
		CreatedBy:          "312",
		UpdatedAt:          time.Now().Local().String(),
		UpdatedBy:          "312",
	})
	if err == nil {
		t.Errorf("Expected error, got nil 1")
	}
}

func TestUpdatePasienByUuid_Success(t *testing.T) {
	_db := db.DB
	defer func() {
		_db.Exec("DELETE FROM pasien WHERE pasien_id = 1")
		db.DB = _db
	}()

	uid := uuid.New()
	_db.Query(`INSERT INTO pasien (
		no_erm,
		pasien_uuid,
		no_rm_lama,
		no_dok_rm,
		penjamin,
		no_penjamin,
		nik,
		no_kk,
		nama,
		tempat_lahir,
		tanggal_lahir,
		no_ihs,
		jenis_kelamin,
		golongan_darah,
		no_telpon,
		email,
		provinsi,
		kabupaten_kota,
		kecamatan,
		kelurahan,
		alamat,
		nama_kontak_darurat,
		nomor_kontak_darurat,
		pekerjaan,
		agama,
		warga_negara,
		pendidikan,
		status_perkawinan,
		created_at,
		created_by,
		updated_at,
		updated_by
	   ) VALUES (
		'123', $1, '123', '123', 'BPJS', '123', '330102', '124212',
		'testNama123', 'sumedang', $2, '12345', 'perempuan', 'E', '086123',
		'123@test.go', 'jawa bali', 'jatinangor', 'sayang', 'cikeruh', 'jalan-jalan no 12', 'jokoTest', '0857123',
		'dokter','islam	', 'WNA', 'SMA', 'belum-kawin', $3, '123test', $4, '123test'
	   )
	   `, uid.String(), time.Now().Local().Format("2006-01-02"), time.Now().Local().String(), time.Now().Local().String())

	err := UpdatePasienByUuid(uid.String(), person.Pasien{
		NoERM:              "31222",
		NoRMLama:           "312321",
		NoDokRM:            "315342",
		Penjamin:           "3143532",
		NoPenjamin:         "31652",
		NIK:                "3126456",
		NoKK:               "312645",
		Nama:               "Joko Doe",
		TempatLahir:        "Indonesia",
		TanggalLahir:       time.Now().Local().Format("2006-01-02"),
		NoIHS:              "31dd2",
		JenisKelamin:       "Perempuan",
		GolonganDarah:      "AB",
		NoTelpon:           "0384312",
		Email:              "alala@gmailc.om",
		Provinsi:           "Java",
		KabupatenKota:      "Bandung",
		Kecamatan:          "Banding",
		Kelurahan:          "Bandeng",
		Alamat:             "Jalan Bandong",
		NamaKontakDarurat:  "nineoneone",
		NomorKontakDarurat: "911",
		Pekerjaan:          "GUru",
		Agama:              "Ateis",
		WargaNegara:        "Indo",
		Pendidikan:         "S0",
		StatusPerkawinan:   "Kawing",
		CreatedAt:          time.Now().Local().String(),
		CreatedBy:          "312",
		UpdatedAt:          time.Now().Local().String(),
		UpdatedBy:          "312",
	})

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	_db.Exec("DELETE FROM pasien WHERE pasien_id = 1")
	db.DB = _db
}

func TestUpdatePasienByUuid_Fail(t *testing.T) {
	err := UpdatePasienByUuid(uuid.New().String(), person.Pasien{
		NoERM:              "31222",
		NoRMLama:           "312321",
		NoDokRM:            "315342",
		Penjamin:           "3143532",
		NoPenjamin:         "31652",
		NIK:                "3126456",
		NoKK:               "312645",
		Nama:               "Joko Doe",
		TempatLahir:        "Indonesia",
		TanggalLahir:       time.Now().Local().Format("2006-01-02"),
		NoIHS:              "31dd2",
		JenisKelamin:       "Perempuan",
		GolonganDarah:      "AB",
		NoTelpon:           "0384312",
		Email:              "alala@gmailc.om",
		Provinsi:           "Java",
		KabupatenKota:      "Bandung",
		Kecamatan:          "Banding",
		Kelurahan:          "Bandeng",
		Alamat:             "Jalan Bandong",
		NamaKontakDarurat:  "nineoneone",
		NomorKontakDarurat: "911",
		Pekerjaan:          "GUru",
		Agama:              "Ateis",
		WargaNegara:        "Indo",
		Pendidikan:         "S0",
		StatusPerkawinan:   "Kawing",
		CreatedAt:          time.Now().Local().String(),
		CreatedBy:          "312",
		UpdatedAt:          time.Now().Local().String(),
		UpdatedBy:          "312",
	})

	if err == nil {
		t.Errorf("Expected error, got nil")
	}

	_db := db.DB
	defer func() {
		_db = db.Conn()
		db.DB = _db
	}()

	_db.Close()
	err = UpdatePasienByUuid(uuid.New().String(), person.Pasien{
		NoERM:              "31222",
		NoRMLama:           "312321",
		NoDokRM:            "315342",
		Penjamin:           "3143532",
		NoPenjamin:         "31652",
		NIK:                "3126456",
		NoKK:               "312645",
		Nama:               "Joko Doe",
		TempatLahir:        "Indonesia",
		TanggalLahir:       time.Now().Local().Format("2006-01-02"),
		NoIHS:              "31dd2",
		JenisKelamin:       "Perempuan",
		GolonganDarah:      "AB",
		NoTelpon:           "0384312",
		Email:              "alala@gmailc.om",
		Provinsi:           "Java",
		KabupatenKota:      "Bandung",
		Kecamatan:          "Banding",
		Kelurahan:          "Bandeng",
		Alamat:             "Jalan Bandong",
		NamaKontakDarurat:  "nineoneone",
		NomorKontakDarurat: "911",
		Pekerjaan:          "GUru",
		Agama:              "Ateis",
		WargaNegara:        "Indo",
		Pendidikan:         "S0",
		StatusPerkawinan:   "Kawing",
		CreatedAt:          time.Now().Local().String(),
		CreatedBy:          "312",
		UpdatedAt:          time.Now().Local().String(),
		UpdatedBy:          "312",
	})

	if err == nil {
		t.Errorf("Expected error, got nil")
	}
}
