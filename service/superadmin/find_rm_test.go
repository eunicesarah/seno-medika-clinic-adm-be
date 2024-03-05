package superadmin

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"seno-medika.com/config/db"
	"testing"
)

func TestFindAll_Success(t *testing.T) {
	_db := db.DB
	defer func() {
		_db.Exec("DELETE FROM users WHERE user_id = 1")
		db.DB = _db
	}()

	uid := uuid.New()
	_db.Exec("INSERT INTO users (user_id, user_uuid, nama, email, role) VALUES (1, $1, 'test', 'test', 'test')", uid)
	vals, _ := _db.Exec("SELECT * FROM users")

	val, err := FindAll()
	sum, er := vals.RowsAffected()
	require.NoError(t, err)
	require.NoError(t, er)
	require.Equal(t, sum, int64(len(val)))
}

func TestFindAll_Fail(t *testing.T) {
	_db := db.DB
	defer func() {
		db.DB = _db
	}()

	_db.Close()

	_, err := FindAll()
	require.Error(t, err)
}

//func TestFindById_Success(t *testing.T) {
//	_db := db.DB
//	defer func() {
//		_db.Exec("DELETE FROM users WHERE user_id = 2")
//		db.DB = _db
//	}()
//
//	_db.Exec("INSERT INTO users (user_id, user_uuid, nama, email, role, password) VALUES ($2, $1, 'test', 'test', 'test','akjfkjdfkskdf')", uuid.New(), 2)
//
//	val, err := FindById(2)
//	t.Log(val)
//	require.NoError(t, err)
//}
