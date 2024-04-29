package kasir

import (
	"testing"
	"github.com/stretchr/testify/require"
	"seno-medika.com/config/db"
	// "seno-medika.com/model/cashierstation"
)

// func TestUpdateMetodePembayaran_Success(t *testing.T) {

// }
func TestUpdateMetodePembayaran_Fail(t *testing.T) {
	_db := db.DB
	defer func() {
		_db = db.Conn()
		db.DB = _db
	}()

	err := UpdateMetodePembayaran(998, "ovo")
	require.Error(t, err)

	_db.Close()
	err = UpdateMetodePembayaran(998, "ovo")
	require.Error(t, err)
}
// func TestUpdateTotalBiaya_Success(t *testing.T) {

// }
func TestUpdateTotalBiaya_Fail(t *testing.T) {
	_db := db.DB
	defer func() {
		_db = db.Conn()
		db.DB = _db
	}()

	err := UpdateTotalBiaya(998, 1000)
	require.Error(t, err)

	_db.Close()
	err = UpdateTotalBiaya(998, 1000)
	require.Error(t, err)
}