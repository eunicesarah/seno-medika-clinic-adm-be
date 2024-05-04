package antrian

import (
	"testing"

	"seno-medika.com/config/db"
)

func TestChangeStatusAntrianById_Success(t *testing.T) {
	_db := db.DB
	defer func() {
		_db.Exec("DELETE FROM antrian WHERE antrian_id = 2")
		db.DB = _db
	}()

	_db.Exec("INSERT INTO antrian (antrian_id, pasien_id, nomor_antrian, status, poli, instalasi, created_at) VALUES ($2, $1, 'test', true, 'test','akjfkjdfkskdf')", 2)

	err := ChangeStatusAntrianById(2, "true")
	if err != nil {
		t.Errorf("Error ChangeStatusAntrianById: %s", err)
	}

}

// func TestChangeStatusAntrianById_Fail(t *testing.T) {
// 	err := ChangeStatusAntrianById(2, true)
// 	if err == nil {
// 		t.Errorf("Expected error, got nil")
// 	}

// 	_db := db.DB
// 	defer func() {
// 		_db = db.Conn()
// 		db.DB = _db
// 	}()

// 	_db.Close()
// 	err = ChangeStatusAntrianById(2, true)
// 	if err == nil {
// 		t.Errorf("Expected error, got nil")
// 	}
// }

func TestChangeStatusByPoli_Success(t *testing.T) {
	_db := db.DB
	defer func() {
		_db.Exec("DELETE FROM antrian WHERE antrian_id = 2")
		db.DB = _db
	}()

	_db.Exec("INSERT INTO antrian (antrian_id, pasien_id, nomor_antrian, status, poli, instalasi, created_at) VALUES ($2, $1, 'test', true, 'test','akjfkjdfkskdf')", 2)

	err := ChangeStatusByPoli("test", "true")
	if err != nil {
		t.Errorf("Error ChangeStatusByPoli: %s", err)
	}

}

// func TestChangeStatusByPoli_Fail(t *testing.T) {
// 	err := ChangeStatusByPoli("test", true)
// 	if err == nil {
// 		t.Errorf("Expected error, got nil")
// 	}

// 	_db := db.DB
// 	defer func() {
// 		_db = db.Conn()
// 		db.DB = _db
// 	}()

// 	_db.Close()
// 	err = ChangeStatusByPoli("test", true)
// 	if err == nil {
// 		t.Errorf("Expected error, got nil")
// 	}
// }

func TestChangeStatusByInstalasi_Success(t *testing.T) {
	_db := db.DB
	defer func() {
		_db.Exec("DELETE FROM antrian WHERE antrian_id = 2")
		db.DB = _db
	}()

	_db.Exec("INSERT INTO antrian (antrian_id, pasien_id, nomor_antrian, status, poli, instalasi, created_at) VALUES ($2, $1, 'test', true, 'test','akjfkjdfkskdf')", 2)

	err := ChangeStatusByInstalasi("test", "true")
	if err != nil {
		t.Errorf("Error ChangeStatusByInstalasi: %s", err)
	}

}

// func TestChangeStatusByInstalasi_Fail(t *testing.T) {
// 	err := ChangeStatusByInstalasi("test", true)
// 	if err == nil {
// 		t.Errorf("Expected error, got nil")
// 	}

// 	_db := db.DB
// 	defer func() {
// 		_db = db.Conn()
// 		db.DB = _db
// 	}()

// 	_db.Close()
// 	err = ChangeStatusByInstalasi("test", true)
// 	if err == nil {
// 		t.Errorf("Expected error, got nil")
// 	}
// }

func TestChangeStatusAntrianById_Fail(t *testing.T) {
	// Simulate DB error by closing the connection
	db.DB.Close()

	err := ChangeStatusAntrianById(2, "true")
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
}

func TestChangeStatusByPoli_Fail(t *testing.T) {
	// Simulate DB error by closing the connection
	db.DB.Close()

	err := ChangeStatusByPoli("test", "true")
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
}

func TestChangeStatusByInstalasi_Fail(t *testing.T) {
	// Simulate DB error by closing the connection
	db.DB.Close()

	err := ChangeStatusByInstalasi("test", "true")
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
}
