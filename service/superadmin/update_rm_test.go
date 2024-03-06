package superadmin

import (
	"github.com/google/uuid"
	"seno-medika.com/config/db"
	"seno-medika.com/model/person"
	"testing"
)

func TestUpdateUserById_Success(t *testing.T) {
	_db := db.DB
	defer func() {
		_db.Exec("DELETE FROM users WHERE user_id = 1")
		db.DB = _db
	}()

	uid := uuid.New()
	_db.Exec("INSERT INTO users (user_id, user_uuid, nama, email, role, password) VALUES (1, $1, 'test', 'test', 'test','akjfkjdfkskdf')", uid)

	err := UpdateUserById(1, person.User{
		Nama:  "test",
		Email: "test",
		Role:  "test",
	})

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}

func TestUpdateUserById_Fail(t *testing.T) {
	err := UpdateUserById(1, person.User{
		Nama:  "test",
		Email: "test",
		Role:  "test",
	})

	if err == nil {
		t.Errorf("Expected error, got nil")
	}
}

func TestUpdateUserByUuid_Success(t *testing.T) {
	_db := db.DB
	defer func() {
		_db.Exec("DELETE FROM users WHERE user_id = 1")
		db.DB = _db
	}()

	uid := uuid.New()
	_db.Exec("INSERT INTO users (user_id, user_uuid, nama, email, role, password) VALUES (1, $1, 'test', 'test', 'test','akjfkjdfkskdf')", uid)

	err := UpdateUserByUuid(uid.String(), person.User{
		Nama:  "test",
		Email: "test",
		Role:  "test",
	})

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}

func TestUpdateUserByUuid_Fail(t *testing.T) {
	err := UpdateUserByUuid(uuid.New().String(), person.User{
		Nama:  "test",
		Email: "test",
		Role:  "test",
	})

	if err == nil {
		t.Errorf("Expected error, got nil")
	}
}
