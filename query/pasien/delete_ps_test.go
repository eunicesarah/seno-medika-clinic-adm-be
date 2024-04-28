package pasien

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"seno-medika.com/config/db"
	"testing"
	"time"
)

func TestDeletePasienById_Success(t *testing.T) {
	_db := db.DB
	var pasienId int

	defer func() {
		db.DB = _db
	}()

	_db.QueryRow(`INSERT INTO pasien (
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
	   ) VALUES ( '123', $1, '123', '123', 'BPJS', '123', '330102', '124212',
		'testNama123', 'sumedang', $2, '12345', 'perempuan', 'E', '086123',
		'123@test.go', 'jawa bali', 'jatinangor', 'sayang', 'cikeruh', 'jalan-jalan no 12', 'jokoTest', '0857123',
		'dokter','islam	', 'WNA', 'SMA', 'belum-kawin', $3, '123test', $4, '123test'
	   ) RETURNING pasien_id
	   `, uuid.New(), time.Now().Local().Format("2006-01-02"), time.Now().Local().String(), time.Now().Local().String()).Scan(&pasienId)
	err := DeletePasienById(pasienId)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}

func TestDeletePasienById_Fail(t *testing.T) {
	_db := db.DB

	defer func() {
		_db = db.Conn()
		db.DB = _db
	}()

	err := DeletePasienById(7986)
	require.Error(t, err)

	_db.Close()
	err = DeletePasienById(7986)
	require.Error(t, err)
}

func TestDeletePasienByUuid_Success(t *testing.T) {
	_db := db.DB

	defer func() {
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
	   `, uid, time.Now().Local().Format("2006-01-02"), time.Now().Local().String(), time.Now().Local().String())
	err := DeletePasienByUuid(uid.String())
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}

func TestDeletePasienByUuid_Fail(t *testing.T) {
	_db := db.DB

	defer func() {
		_db = db.Conn()
		db.DB = _db
	}()

	err := DeletePasienByUuid(uuid.New().String())
	require.Error(t, err)

	_db.Close()
	err = DeletePasienByUuid(uuid.New().String())
	require.Error(t, err)
}
