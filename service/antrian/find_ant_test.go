package antrian

import (
	"github.com/stretchr/testify/require"
	"seno-medika.com/config/db"
	"seno-medika.com/model/antrian"
	"testing"
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

