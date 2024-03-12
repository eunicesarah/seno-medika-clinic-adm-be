package antrian

// import (
// 	"github.com/stretchr/testify/require"
// 	"seno-medika.com/config/db"
// 	"seno-medika.com/model/antrian"
// 	"testing"
// )

// func TestFindAntrianById_Success(t *testing.T) {
// 	_db := db.DB
// 	val, _ := FindAntrianById(1)
// 	require.Equal(t, antrian.Antrian{}, val)
// 	defer func() {
// 		_db.Exec("DELETE FROM antrian WHERE antrian_id = 1")
// 		db.DB = _db
// 	}()

// 	_db.Exec("INSERT INTO antrian (antrian_id, pasien_id, nomor_antrian, status, poli, instalasi, created_at) VALUES ($2, $1, 'test', true, 'test','akjfkjdfkskdf')", 1)

// 	val, err := FindAntrianById(1)
// 	t.Log(val)
// 	require.NoError(t, err)

// }
// func TestFindAntrianById_Fail(t *testing.T) {
// 	_db := db.DB
// 	defer func() {
// 		_db = db.Conn()
// 		db.DB = _db
// 	}()

// 	val, err := FindAntrianById(111)
// 	require.Error(t, err)
// 	require.Equal(t, antrian.Antrian{}, val)

// 	_db.Close()
// 	val, err = FindAntrianById(111)
// 	require.Error(t, err)
// 	require.Equal(t, antrian.Antrian{}, val)
// }

// func TestFindAntrianAll_Success(t *testing.T) {
// 	_db := db.DB
// 	val, _ := FindAntrianAll()
// 	require.Equal(t, []antrian.Antrian{}, val)
// 	defer func() {
// 		_db.Exec("DELETE FROM antrian WHERE antrian_id = 1")
// 		db.DB = _db
// 	}()

// 	_db.Exec("INSERT INTO antrian (antrian_id, pasien_id, nomor_antrian, status, poli, instalasi, created_at) VALUES ($2, $1, 'test', true, 'test','akjfkjdfkskdf')", 1)

// 	val, err := FindAntrianAll()
// 	t.Log(val)
// 	require.NoError(t, err)

// }

// func TestFindAntrianAll_Fail(t *testing.T) {
// 	_db := db.DB
// 	defer func() {
// 		_db = db.Conn()
// 		db.DB = _db
// 	}()

// 	_db.Close()

// 	val, err := FindAntrianAll()
// 	require.Error(t, err)
// 	require.Equal(t, []antrian.Antrian{}, val)
// }
