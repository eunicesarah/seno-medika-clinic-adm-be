package superadmin

import (
	"github.com/google/uuid"
	"seno-medika.com/config/db"
	"testing"
)

func TestChangeEmailById_Success(t *testing.T) {
	_db := db.DB
	uid := uuid.New()
	defer func() {
		_db.Exec("DELETE FROM users WHERE user_id = 1")
		db.DB = _db
	}()

	_db.Exec("INSERT INTO users (user_id, user_uuid, nama, email, role, password) VALUES (1, $1, 'test', 'test', 'test','akjfkjdfkskdf')", uid)

	err := ChangeEmailById("1", "test")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}

func TestChangeEmailById_Fail(t *testing.T) {
	err := ChangeEmailById("1", "test")
	if err == nil {
		t.Errorf("Expected error, got nil")
	}

	_db := db.DB
	defer func() {
		_db = db.Conn()
		db.DB = _db
	}()

	_db.Close()
	err = ChangeEmailById("1", "test")
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
}

func TestChangeEmailByUuid_Success(t *testing.T) {
	_db := db.DB
	uid := uuid.New()
	defer func() {
		_db.Exec("DELETE FROM users WHERE user_id = 1")
		db.DB = _db
	}()

	_db.Exec("INSERT INTO users (user_id, user_uuid, nama, email, role, password) VALUES (1, $1, 'test', 'test', 'test','akjfkjdfkskdf')", uid)

	err := ChangeEmailByUuid(uid.String(), "test")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}

func TestChangeEmailByUuid_Fail(t *testing.T) {
	err := ChangeEmailByUuid(uuid.New().String(), "test")
	if err == nil {
		t.Errorf("Expected error, got nil")
	}

	_db := db.DB
	defer func() {
		_db = db.Conn()
		db.DB = _db
	}()

	_db.Close()
	err = ChangeEmailByUuid(uuid.New().String(), "test")
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
}

func TestChangeNameById_Success(t *testing.T) {
	_db := db.DB
	uid := uuid.New()
	defer func() {
		_db.Exec("DELETE FROM users WHERE user_id = 1")
		db.DB = _db
	}()

	_db.Exec("INSERT INTO users (user_id, user_uuid, nama, email, role, password) VALUES (1, $1, 'test', 'test', 'test','akjfkjdfkskdf')", uid)

	err := ChangeNameById(1, "test")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}

func TestChangeNameById_Fail(t *testing.T) {
	err := ChangeNameById(1, "test")
	if err == nil {
		t.Errorf("Expected error, got nil")
	}

	_db := db.DB
	defer func() {
		_db = db.Conn()
		db.DB = _db
	}()

	_db.Close()
	err = ChangeNameById(1, "test")
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
}

func TestChangeNameByUuid_Success(t *testing.T) {
	_db := db.DB
	uid := uuid.New()
	defer func() {
		_db.Exec("DELETE FROM users WHERE user_id = 1")
		db.DB = _db
	}()

	_db.Exec("INSERT INTO users (user_id, user_uuid, nama, email, role, password) VALUES (1, $1, 'test', 'test', 'test','akjfkjdfkskdf')", uid)

	err := ChangeNameByUuid(uid.String(), "test")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}

func TestChangeNameByUuid_Fail(t *testing.T) {
	err := ChangeNameByUuid(uuid.New().String(), "test")
	if err == nil {
		t.Errorf("Expected error, got nil")
	}

	_db := db.DB
	defer func() {
		_db = db.Conn()
		db.DB = _db
	}()

	_db.Close()
	err = ChangeNameByUuid(uuid.New().String(), "test")
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
}

func TestChangeRoleById_Success(t *testing.T) {
	_db := db.DB
	uid := uuid.New()
	defer func() {
		_db.Exec("DELETE FROM users WHERE user_id = 1")
		db.DB = _db
	}()

	_db.Exec("INSERT INTO users (user_id, user_uuid, nama, email, role, password) VALUES (1, $1, 'test', 'test', 'test','akjfkjdfkskdf')", uid)

	err := ChangeRoleById(1, "test")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}

func TestChangeRoleById_Fail(t *testing.T) {
	err := ChangeRoleById(1, "test")
	if err == nil {
		t.Errorf("Expected error, got nil")
	}

	_db := db.DB
	defer func() {
		_db = db.Conn()
		db.DB = _db
	}()

	_db.Close()
	err = ChangeRoleById(1, "test")
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
}

func TestChangeRoleByUuid_Success(t *testing.T) {
	_db := db.DB
	uid := uuid.New()
	defer func() {
		_db.Exec("DELETE FROM users WHERE user_id = 1")
		db.DB = _db
	}()

	_db.Exec("INSERT INTO users (user_id, user_uuid, nama, email, role, password) VALUES (1, $1, 'test', 'test', 'test','akjfkjdfkskdf')", uid)

	err := ChangeRoleByUuid(uid.String(), "test")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}

func TestChangeRoleByUuid_Fail(t *testing.T) {
	err := ChangeRoleByUuid(uuid.New().String(), "test")
	if err == nil {
		t.Errorf("Expected error, got nil")
	}

	_db := db.DB
	defer func() {
		_db = db.Conn()
		db.DB = _db
	}()

	_db.Close()
	err = ChangeRoleByUuid(uuid.New().String(), "test")
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
}

func TestChangePasswordById_Success(t *testing.T) {
	_db := db.DB
	uid := uuid.New()
	defer func() {
		_db.Exec("DELETE FROM users WHERE user_id = 1")
		db.DB = _db
	}()

	_db.Exec("INSERT INTO users (user_id, user_uuid, nama, email, role, password) VALUES (1, $1, 'test', 'test', 'test','akjfkjdfkskdf')", uid)

	err := ChangePasswordById(1, "test")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}

func TestChangePasswordById_Fail(t *testing.T) {
	err := ChangePasswordById(1, "test")
	if err == nil {
		t.Errorf("Expected error, got nil")
	}

	_db := db.DB
	defer func() {
		_db = db.Conn()
		db.DB = _db
	}()

	_db.Close()
	err = ChangePasswordById(1, "test")
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
}

func TestChangePasswordByUuid_Success(t *testing.T) {
	_db := db.DB
	uid := uuid.New()
	defer func() {
		_db.Exec("DELETE FROM users WHERE user_id = 1")
		db.DB = _db
	}()

	_db.Exec("INSERT INTO users (user_id, user_uuid, nama, email, role, password) VALUES (1, $1, 'test', 'test', 'test','akjfkjdfkskdf')", uid)

	err := ChangePasswordByUuid(uid.String(), "test")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}

func TestChangePasswordByUuid_Fail(t *testing.T) {
	err := ChangePasswordByUuid(uuid.New().String(), "test")
	if err == nil {
		t.Errorf("Expected error, got nil")
	}

	_db := db.DB
	defer func() {
		_db = db.Conn()
		db.DB = _db
	}()

	_db.Close()
	err = ChangePasswordByUuid(uuid.New().String(), "test")
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
}
