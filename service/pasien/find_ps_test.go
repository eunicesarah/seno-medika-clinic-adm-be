package pasien

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"seno-medika.com/config/db"
	"seno-medika.com/model/person"
	"testing"
	"time"
)

func TestFindPasienById_Fail(t *testing.T) {
	_db := db.DB
	defer func() {
		_db = db.Conn()
		db.DB = _db
	}()

	val, err := FindPasienById(998)
	require.Error(t, err)
	require.Equal(t, person.Pasien{}, val)

	_db.Close()
	val, err = FindPasienById(998)
	require.Error(t, err)
	require.Equal(t, person.Pasien{}, val)
}

func TestFindPasienById_Success(t *testing.T) {
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
		8911, '123333', $1, '123', '123', 'BPJS', '123', '330102', '124212',
		'testNama123', 'sumedang', $2, '12345', 'perempuan', 'E', '086123',
		'123@test.go', 'jawa bali', 'jatinangor', 'sayang', 'cikeruh', 'jalan-jalan no 12', 'jokoTest', '0857123',
		'dokter','islam	', 'WNA', 'SMA', 'belum-kawin', $3, '123test', $4, '123test'
	   )
	   `, uuid.New(), time.Now().Local().Format("2006-02-01"), time.Now().Local().String(), time.Now().Local().String())

	val, err := FindPasienById(8911)
	t.Log(val)
	require.NoError(t, err)
	_db.Exec("DELETE FROM pasien WHERE pasien_id = 8911")
	db.DB = _db
}

func TestFindPasienByUuid_Fail(t *testing.T) {
	_db := db.DB
	defer func() {
		_db = db.Conn()
		db.DB = _db
	}()

	val, err := FindPasienByUuid(uuid.New().String())
	require.Error(t, err)
	require.Equal(t, person.Pasien{}, val)

	_db.Close()
	val, err = FindPasienByUuid(uuid.New().String())
	require.Error(t, err)
	require.Equal(t, person.Pasien{}, val)
}

func TestFindPasienByUuid_Success(t *testing.T) {
	_db := db.DB
	defer func() {
		_db.Exec("DELETE FROM pasien WHERE pasien_id = 123")
		db.DB = _db
	}()

	uid := uuid.New()
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
		123, '12333', $1, '123', '123', 'BPJS', '123', '330102', '124212',
		'testNama123', 'sumedang', $2, '12345', 'perempuan', 'E', '086123',
		'123@test.go', 'jawa bali', 'jatinangor', 'sayang', 'cikeruh', 'jalan-jalan no 12', 'jokoTest', '0857123',
		'dokter','islam	', 'WNA', 'SMA', 'belum-kawin', $3, '123test', $4, '123test'
	   )
	   `, uid, time.Now().Local().Format("2006-02-01"), time.Now().Local().String(), time.Now().Local().String())

	val, err := FindPasienByUuid(uid.String())
	require.NoError(t, err)
	require.Equal(t, uid, val.PasienUUID)
	_db.Exec("DELETE FROM pasien WHERE pasien_id = 123")
	db.DB = _db
}

func TestFindPasienAll_Success(t *testing.T) {
	_db := db.DB

	defer func() {
		_db.Exec("DELETE FROM pasien WHERE pasien_id = 991")
		db.DB = _db
	}()

	uid := uuid.New()
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
		8911, '1223', $1, '123', '123', 'BPJS', '123', '330102', '124212',
		'testNama123', 'sumedang', $2, '12345', 'perempuan', 'E', '086123',
		'123@test.go', 'jawa bali', 'jatinangor', 'sayang', 'cikeruh', 'jalan-jalan no 12', 'jokoTest', '0857123',
		'dokter','islam	', 'WNA', 'SMA', 'belum-kawin', $3, '123test', $4, '123test'
	   )
	   `, uid, time.Now().Local().Format("2006-02-01"), time.Now().Local().String(), time.Now().Local().String())

	vals, _ := _db.Exec("SELECT * FROM pasien")

	val, err := FindPasienAll()
	sum, er := vals.RowsAffected()
	require.NoError(t, err)
	require.NoError(t, er)
	require.Equal(t, sum, int64(len(val)))
}

func TestFindPasienAll_Fail(t *testing.T) {
	_db := db.DB
	defer func() {
		_db = db.Conn()
		db.DB = _db
	}()

	_db.Close()

	val, err := FindPasienAll()
	require.Error(t, err)
	require.Equal(t, []person.Pasien(nil), val)
}

func TestFindPasienByNIK_Sucsess(t *testing.T) {
	_db := db.DB
	defer func() {
		_db.Exec("DELETE FROM pasien WHERE nik = 330102")
		db.DB = _db
	}()

	uid := uuid.New()
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
	123, '12333', $1, '123', '123', 'BPJS', '123', '330102', '124212',
	'testNama123', 'sumedang', $2, '12345', 'perempuan', 'E', '086123',
	'123@test.go', 'jawa bali', 'jatinangor', 'sayang', 'cikeruh', 'jalan-jalan no 12', 'jokoTest', '0857123',
	'dokter','islam ', 'WNA', 'SMA', 'belum-kawin', $3, '123test', $4, '123test'
	)
	`, uid.String(), time.Now().Local().Format("2006-02-01"), time.Now().Local().String(), time.Now().Local().String())

	val, err := FindPasienByNIK("330102")
	t.Log(val)
	require.NoError(t, err)
	_db.Exec("DELETE FROM pasien WHERE pasien_id = 123")
	db.DB = _db

}
   
func TestFindPasienByNIK_Fail(t *testing.T) {
	_db := db.DB
	defer func() {
		_db = db.Conn()
		db.DB = _db
	}()

	val, err := FindPasienByNIK("123")
	require.Error(t, err)
	require.Equal(t, person.Pasien{}, val)

	_db.Close()
	val, err = FindPasienByNIK("123")
	require.Error(t, err)
	require.Equal(t, person.Pasien{}, val)
}