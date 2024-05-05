package obat

import (
	"testing"

	"github.com/stretchr/testify/require"
	"seno-medika.com/config/db"
)

func TestDeleteObatById_Success(t *testing.T) {
	_db := db.DB

	defer func() {
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
	err := DeleteObatById(8911)
	if err != nil {
		t.Errorf("This should not be error, but have %v", err)
	}

}

func TestDeleteObatById_Failed(t *testing.T) {
	_db := db.DB

	defer func() {
		_db = db.Conn()
		db.DB = _db
	}()

	err := DeleteObatById(8911)
	require.Error(t, err)

	_db.Close()
	err = DeleteObatById(8911)
	require.Error(t, err)
}

func TestDeleteObatByName_Success(t *testing.T) {
	_db := db.DB

	defer func() {
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
	err := DeleteObatByName("testNama123")
	if err != nil {
		t.Errorf("This should not be error, but have %v", err)
	}

}

func TestDeleteObatByName_Failed(t *testing.T) {
	_db := db.DB

	defer func() {
		_db = db.Conn()
		db.DB = _db
	}()

	err := DeleteObatByName("testNama123")
	require.Error(t, err)

	_db.Close()
	err = DeleteObatByName("testNama123")
	require.Error(t, err)
}
