package superadmin

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"seno-medika.com/config/db"
	"testing"
)

func TestDeleteUserById_Success(t *testing.T) {
	_db := db.DB

	defer func() {
		db.DB = _db
	}()

	_db.Query("INSERT INTO users (user_id, user_uuid, nama, email, password, role) VALUES (1, $1, 'test', '', '', '')", uuid.New())

	err := DeleteUserById(1)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}

func TestDeleteUserById_Fail(t *testing.T) {
	_db := db.DB

	defer func() {
		db.DB = _db
	}()

	err := DeleteUserById(1)
	require.Error(t, err)
}

func TestDeleteUserByUuid_Success(t *testing.T) {
	_db := db.DB

	defer func() {
		db.DB = _db
	}()

	uid := uuid.New()
	_db.Query("INSERT INTO users (user_id, user_uuid, nama, email, password, role) VALUES (1, $1, 'test', '', '', '')", uid)

	err := DeleteUserByUuid(uid.String())
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}

func TestDeleteUserByUuid_Fail(t *testing.T) {
	_db := db.DB

	defer func() {
		db.DB = _db
	}()

	err := DeleteUserByUuid(uuid.New().String())
	require.Error(t, err)
}
