package kasir

import (
	"testing"

	"github.com/stretchr/testify/require"
	"seno-medika.com/config/db"
	"seno-medika.com/model/cashierstation"
	"seno-medika.com/model/pharmacystation"
)

func TestFindDetailByResepId_Fail(t *testing.T) {
	_db := db.DB
	defer func() {
		_db = db.Conn()
		db.DB = _db
	}()

	val, err := FindDetailByResepId(994)
	require.Error(t, err)
	require.Equal(t, []pharmacystation.DetailObat(nil), val)
	_db.Close()
	if err == nil {
		t.Errorf("expecting error, got nil")
	}
	if len(val) != 0 {
		t.Errorf("expecting empty, got %+v", val)
	}

}

// func TestFindDetailByResepId_Success(t *testing.T) {
// 	_db := db.DB
// 	defer func() {
// 		_db = db.Conn()
// 		db.DB = _db
// 	}()

// 	val, err := FindDetailByResepId(1)
// 	require.NoError(t, err)
// 	// require.NotEqual(t, pharmacystation.DetailObat{}, val)
// 	_db.Close()
// 	if err != nil {
// 		t.Errorf("expecting nil, got %+v", err)
// 	}
// 	if len(val) == 0 {
// 		t.Errorf("expecting not empty, got empty")
// 	}
// }

func TestFindTindakanByNotaId_Fail(t *testing.T) {
	_db := db.DB
	defer func() {
		_db = db.Conn()
		db.DB = _db
	}()

	val, err := FindTindakanByNotaId(994)
	require.Error(t, err)
	require.Equal(t, []cashierstation.Tindakan(nil), val)
	_db.Close()
	if err == nil {
		t.Errorf("expecting error, got nil")
	}
	if len(val) != 0 {
		t.Errorf("expecting empty, got %+v", val)
	}
}