package antrian

import (
	"log"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"seno-medika.com/config/db"
	"seno-medika.com/model/antrian"
)

// func TestFindAntrianById_Success(t *testing.T) {
// 	_db := db.DB
// 	defer func() {
// 		_db.Exec("DELETE FROM antrian WHERE nik = 123123")
// 		db.DB = _db
// 	}()

// 	_db.Exec("INSERT INTO antrian (nik, nama, status, poli, instalasi, created_at) VALUES ($2, $1, 'test', true, 'umum','akjfkjdfkskdf')", 1)

// 	val, err := FindAntrianById(1)
// 	require.NoError(t, err)
// 	require.NotEqual(t, antrian.Antrian{}, val)
// }

func TestFindAntrianById_Fail(t *testing.T) {
	_db := db.DB
	defer func() {
		_db = db.Conn()
		db.DB = _db
	}()

	// Test when ID doesn't exist
	val, err := FindAntrianById(111)
	require.Error(t, err)
	require.Equal(t, antrian.Antrian{}, val)

	// Test when database connection is closed
	_db.Close()
	val, err = FindAntrianById(111)
	require.Error(t, err)
	require.Equal(t, antrian.Antrian{}, val)
}

func TestFindAntrianAll_Success(t *testing.T) {
	_db := db.DB
	defer func() {
		_db.Exec("DELETE FROM antrian WHERE antrian_id = 1")
		db.DB = _db
	}()

	_db.Exec("INSERT INTO antrian (antrian_id, pasien_id, nomor_antrian, status, poli, instalasi, created_at) VALUES ($2, $1, 'test', true, 'test','akjfkjdfkskdf')", 1)

	val, err := FindAntrianAll()
	require.NoError(t, err)
	require.NotEqual(t, []antrian.Antrian{}, val)
}

func TestFindAntrianAll_Fail(t *testing.T) {
	_db := db.DB
	defer func() {
		_db = db.Conn()
		db.DB = _db
	}()

	// Test when database connection is closed
	_db.Close()
	val, err := FindAntrianAll()
	require.Error(t, err)
	require.Equal(t, []antrian.Antrian(nil), val)
}

func TestFindAntrianFilterPemeriksaan(t *testing.T) {
	var pasienId int
	_db := db.DB
	defer func() {
		_db = db.Conn()
		_db.Exec("DELETE FROM pasien WHERE nama ILIKE $1", "testNamadfb123")
		db.DB = _db
	}()

	if err := _db.QueryRow(`INSERT INTO pasien (
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
	   ) VALUES ( '12321312', $1, '123', '123', 'BPJS', '123', '3301023322', '124212',
		'testNamadfb123', 'sumedfbdfbdang', $2, '123dfbdfb45', 'perempuan', 'E', '0822136123',
		'122213213@tesst.go', 'jawa bali', 'jatinandfbdfbgor', 'saydfbdfbang', 'cikedfbdfbruh', 'jalan-jaldfbdfban no 12', 'jokoTesdfbdfbt', '0852345237123',
		'dokter','islam	', 'WNA', 'SMA', 'belum-dfbdfbkawin', $3, '123test', $4, '123test'
	   ) RETURNING pasien_id
	   `, uuid.New(), time.Now().Local().Format("2006-01-02"), time.Now().Local().String(), time.Now().Local().String()).Scan(&pasienId); err != nil {
		t.Error(err)
		return
	}

	_, err := _db.Exec(`INSERT INTO antrian (
		pasien_id,
		nomor_antrian,
		poli,
		instalasi,
		created_at
	) VALUES (
		$1,
		'1',
		'umum siang',
		'pratama',
		'2024-02-04'
	)`, pasienId)

	antrian0, size, err := FindAntrianFilterPemeriksaan("", "1", "10", "2024-02-04", "", "pemeriksaan_ttv")
	log.Println(antrian0)
	require.NoError(t, err)
	require.NotEqual(t, []antrian.Antrian{}, antrian0)
	require.NotEqual(t, 0, size)

}

func TestFindAntrianFilterPemeriksaan_Poli(t *testing.T) {
	var pasienId int
	_db := db.DB
	defer func() {
		_db = db.Conn()
		_db.Exec("DELETE FROM pasien WHERE nama ILIKE $1", "tes8tNamadfb123")
		db.DB = _db
	}()

	if err := _db.QueryRow(`INSERT INTO pasien (
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
	   ) VALUES ( '12320812', $1, '123', '123', 'BPJS', '123', '9001023322', '124212',
		'tes8tNamadfb123', 'sumedfbdfbdang', $2, '123dfbdfb45', 'perempuan', 'E', '0822136123',
		'122213213@tesst.go', 'jawa bali', 'jatinandfbdfbgor', 'saydfbdfbang', 'cikedfbdfbruh', 'jalan-jaldfbdfban no 12', 'jokoTesdfbdfbt', '0852345237123',
		'dokter','islam	', 'WNA', 'SMA', 'belum-dfbdfbkawin', $3, '123test', $4, '123test'
	   ) RETURNING pasien_id
	   `, uuid.New(), time.Now().Local().Format("2006-01-02"), time.Now().Local().String(), time.Now().Local().String()).Scan(&pasienId); err != nil {
		t.Error(err)
		return
	}

	_, err := _db.Exec(`INSERT INTO antrian (
		pasien_id,
		nomor_antrian,
		poli,
		instalasi,
		created_at
	) VALUES (
		$1,
		'2',
		'umum siang',
		'pratama',
		'2024-02-04'
	)`, pasienId)

	antrian0, size, err := FindAntrianFilterPemeriksaan("", "1", "10", "2024-02-04", "umum siang", "pemeriksaan_ttv")
	log.Println(antrian0)
	require.NoError(t, err)
	require.NotEqual(t, []antrian.Antrian{}, antrian0)
	require.NotEqual(t, 0, size)

}

func TestFindAntrianFilterPemeriksaan_Search(t *testing.T) {
	var pasienId int
	_db := db.DB
	defer func() {
		_db = db.Conn()
		_db.Exec("DELETE FROM pasien WHERE nama ILIKE $1", "tes9tNamadfb123")
		db.DB = _db
	}()

	if err := _db.QueryRow(`INSERT INTO pasien (
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
	   ) VALUES ( '11321812', $1, '123', '123', 'BPJS', '123', '9301093322', '124212',
		'tes9tNamadfb123', 'sumedfbdfbdang', $2, '123dfbdfb45', 'perempuan', 'E', '0829136123',
		'122213213@tesst.go', 'jawa bali', 'jatinandfbdfbgor', 'saydfbdfbang', 'cikedfbdfbruh', 'jalan-jaldfbdfban no 12', 'jokoTesdfbdfbt', '0852345237123',
		'dokter','islam	', 'WNA', 'SMA', 'belum-dfbdfbkawin', $3, '123test', $4, '123test'
	   ) RETURNING pasien_id
	   `, uuid.New(), time.Now().Local().Format("2006-01-02"), time.Now().Local().String(), time.Now().Local().String()).Scan(&pasienId); err != nil {
		t.Error(err)
		return
	}

	_, err := _db.Exec(`INSERT INTO antrian (
		pasien_id,
		nomor_antrian,
		poli,
		instalasi,
		created_at
	) VALUES (
		$1,
		'3',
		'umum siang',
		'pratama',
		'2024-02-04'
	)`, pasienId)

	antrian0, size, err := FindAntrianFilterPemeriksaan("tes9t", "1", "10", "2024-02-04", "", "pemeriksaan_ttv")
	log.Println(antrian0)
	require.NoError(t, err)
	require.NotEqual(t, []antrian.Antrian{}, antrian0)
	require.NotEqual(t, 0, size)

}


func TestFindAntrianFilterPemeriksaan_SearchPoli(t *testing.T) {
	var pasienId int
	_db := db.DB
	defer func() {
		_db = db.Conn()
		_db.Exec("DELETE FROM pasien WHERE nama ILIKE $1", "tes0tNamadfb123")
		db.DB = _db
	}()

	if err := _db.QueryRow(`INSERT INTO pasien (
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
	   ) VALUES ( '12391812', $1, '123', '123', 'BPJS', '123', '9301723322', '124212',
		'tes0tNamadfb123', 'sumedfbdfbdang', $2, '123dfbdfb45', 'perempuan', 'E', '0812136123',
		'122213213@tesst.go', 'jawa bali', 'jatinandfbdfbgor', 'saydfbdfbang', 'cikedfbdfbruh', 'jalan-jaldfbdfban no 12', 'jokoTesdfbdfbt', '0852345237123',
		'dokter','islam	', 'WNA', 'SMA', 'belum-dfbdfbkawin', $3, '123test', $4, '123test'
	   ) RETURNING pasien_id
	   `, uuid.New(), time.Now().Local().Format("2006-01-02"), time.Now().Local().String(), time.Now().Local().String()).Scan(&pasienId); err != nil {
		t.Error(err)
		return
	}

	_, err := _db.Exec(`INSERT INTO antrian (
		pasien_id,
		nomor_antrian,
		poli,
		instalasi,
		created_at
	) VALUES (
		$1,
		'4',
		'umum siang',
		'pratama',
		'2024-02-04'
	)`, pasienId)

	antrian0, size, err := FindAntrianFilterPemeriksaan("tes0t", "1", "10", "2024-02-04", "umum siang", "pemeriksaan_ttv")
	log.Println(antrian0)
	require.NoError(t, err)
	require.NotEqual(t, []antrian.Antrian{}, antrian0)
	require.NotEqual(t, 0, size)

}
