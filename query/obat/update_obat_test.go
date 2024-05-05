package obat

import (
	"seno-medika.com/config/db"
	"seno-medika.com/model/pharmacystation"
	"testing"
)

func TestUpdateObatById_Success(t *testing.T) {
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
	err := UpdateObatById(8911, pharmacystation.Obat{
		NamaObat:      "testNama123",
		JenisAsuransi: "BPJS",
		Harga:         12345,
		Stock: 	   10,
		Satuan:        "testSatuan",
	})
	if err != nil {
		t.Errorf("This should not be error, but have %v", err)
	}
	_db.Exec("DELETE FROM obat WHERE obat_id = 8911")
	db.DB = _db
}

func TestUpdateObatById_Failed(t *testing.T) {
	_db := db.DB
	defer func() {
		_db.Exec("DELETE FROM obat WHERE obat_id = 8911")
		db.DB = _db
	}()
	err := UpdateObatById(8911, pharmacystation.Obat{
		NamaObat:      "testNama123",
		JenisAsuransi: "BPJS",
		Harga:         12345,
		Stock: 	   10,
		Satuan:        "testSatuan",
	})
	if err == nil {
		t.Errorf("This should be error")
	}
}

func TestUpdateObatByName_Success(t *testing.T) {
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
	err := UpdateObatByName("testNama123", pharmacystation.Obat{
		ObatID:        8911,
		JenisAsuransi: "BPJS",
		Harga:         12345,
		Stock: 	   10,
		Satuan:        "testSatuan",
	})
	if err != nil {
		t.Errorf("This should not be error, but have %v", err)
	}
	_db.Exec("DELETE FROM obat WHERE obat_id = 8911")
	db.DB = _db
}

func TestUpdateObatByName_Failed(t *testing.T) {
	_db := db.DB
	defer func() {
		_db.Exec("DELETE FROM obat WHERE obat_id = 8911")
		db.DB = _db
	}()
	err := UpdateObatByName("testNama123", pharmacystation.Obat{
		ObatID:        8911,
		JenisAsuransi: "BPJS",
		Harga:         12345,
		Stock: 	   10,
		Satuan:        "testSatuan",
	})
	if err == nil {
		t.Errorf("This should be error")
	}
}
