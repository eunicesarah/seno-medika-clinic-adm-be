package kasir

import (
	"testing"
	"github.com/stretchr/testify/require"
	"seno-medika.com/config/db"
	"seno-medika.com/model/cashierstation"
)


// func TestFindNotaById_Success(t *testing.T) {

// }
func TestFindNotaById_Fail(t *testing.T) {
	_db := db.DB
	defer func() {
		_db = db.Conn()
		db.DB = _db
	}()

	val, err := FindNotaById(998)
	require.Error(t, err)
	require.Equal(t, cashierstation.Nota{}, val)

	_db.Close()
	val, err = FindNotaById(998)
	require.Error(t, err)
	require.Equal(t, cashierstation.Nota{}, val)
}
// func TestFindNotaByMetodePembayaran_Success(t *testing.T) {

// }
func TestFindNotaByMetodePembayaran_Fail(t *testing.T) {
	_db := db.DB
	defer func() {
		_db = db.Conn()
		db.DB = _db
	}()

	val, err := FindNotaByMetodePembayaran("ovo")
	require.Error(t, err)
	require.Equal(t, []cashierstation.Nota(nil), val)

	_db.Close()
	val, err = FindNotaByMetodePembayaran("ovo")
	require.Error(t, err)
	require.Equal(t, []cashierstation.Nota(nil), val)
}
// func TestFindNotaByResepId_Success(t *testing.T) {

// }
func TestFindNotaByResepId_Fail(t *testing.T) {
	_db := db.DB
	defer func() {
		_db = db.Conn()
		db.DB = _db
	}()

	val, err := FindNotaByResepId(998)
	require.Error(t, err)
	require.Equal(t, []cashierstation.Nota(nil), val)

	_db.Close()
	val, err = FindNotaByResepId(998)
	require.Error(t, err)
	require.Equal(t, []cashierstation.Nota(nil), val)
}
// func TestFindNotaByPasienID_Success(t *testing.T) {

// }
func TestFindNotaByPasienID_Fail(t *testing.T) {
	_db := db.DB
	defer func() {
		_db = db.Conn()
		db.DB = _db
	}()

	val, err := FindNotaByPasienID(998)
	require.Error(t, err)
	require.Equal(t, []cashierstation.Nota(nil), val)

	_db.Close()
	val, err = FindNotaByPasienID(998)
	require.Error(t, err)
	require.Equal(t, []cashierstation.Nota(nil), val)
}
// func TestFindNotaAll_Success(t *testing.T) {

// }
// func TestFindNotaAll_Fail(t *testing.T) {

// }