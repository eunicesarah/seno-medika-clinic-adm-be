package obat

import (
	"github.com/stretchr/testify/require"
	"seno-medika.com/config/db"
	"seno-medika.com/model/pharmacystation"
	"testing"
)

func TestFindObatById_Fail(t *testing.T) {
	_db := db.DB
	defer func() {
		_db = db.Conn()
		db.DB = _db
	}()

	val, err := FindObatById(998)
	require.Error(t, err)
	require.Equal(t, pharmacystation.Obat{}, val)

}

func TestFindObatById_Success(t *testing.T) {
	_db := db.DB
	defer func() {
		_db.Exec("DELETE FROM obat WHERE obat_id = 8911")
		db.DB = _db
	}()

	_db.Query(`INSERT INTO obat (
		obat_id,
		nama_obat,
		jenis_asuransi,
		harga,
		stock,
		satuan
	) VALUES (
		8911, 'testNama123', 'BPJS', 12345, 10, 'testSatuan'
	)
	`)
	val, err := FindObatById(8911)
	require.NoError(t, err)
	require.NotEqual(t, pharmacystation.Obat{}, val)
}

func TestFindObatByName_Fail(t *testing.T) {
	_db := db.DB
	defer func() {
		_db = db.Conn()
		db.DB = _db
	}()

	val, err := FindObatByName("testNama123")
	require.Error(t, err)
	require.Equal(t, pharmacystation.Obat{}, val)
}

func TestFindObatByName_Success(t *testing.T) {
	_db := db.DB
	defer func() {
		_db.Exec("DELETE FROM obat WHERE obat_id = 8911")
		db.DB = _db
	}()

	_db.Query(`INSERT INTO obat (
		obat_id,
		nama_obat,
		jenis_asuransi,
		harga,
		stock,
		satuan
	) VALUES (
		8911, 'testNama123', 'BPJS', 12345, 10, 'testSatuan'
	)
	`)
	val, err := FindObatByName("testNama123")
	require.NoError(t, err)
	require.NotEqual(t, pharmacystation.Obat{}, val)
}

func TestFindObatAll_Fail(t *testing.T) {
	_db := db.DB
	defer func() {
		_db = db.Conn()
		db.DB = _db
	}()

	_db.Close()

	val, err := FindObatAll()
	require.Error(t, err)
	require.Equal(t, []pharmacystation.Obat(nil), val)
}

func TestFindObatAll_Success(t *testing.T) {
	_db := db.DB
	defer func() {
		_db.Exec("DELETE FROM obat WHERE obat_id = 8911")
		db.DB = _db
	}()

	_db.Query(`INSERT INTO obat (
		obat_id,
		nama_obat,
		jenis_asuransi,
		harga,
		stock,
		satuan
	) VALUES (
		8911, 'testNama123', 'BPJS', 12345, 10, 'testSatuan'
	)
	`)
	val, err := FindObatAll()
	require.NoError(t, err)
	require.NotEqual(t, 0, len(val))
}
