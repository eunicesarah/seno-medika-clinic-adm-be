package superadmin

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"seno-medika.com/config/db"
	"seno-medika.com/model/person"
	"testing"
)

func TestFindById_Fail(t *testing.T) {
	val, err := FindById(1)
	require.Error(t, err)
	require.Equal(t, person.UserWithoutPassword{}, val)
}

func TestFindById_Success(t *testing.T) {
	_db := db.DB
	defer func() {
		_db.Exec("DELETE FROM users WHERE user_id = 2")
		db.DB = _db
	}()

	_db.Exec("INSERT INTO users (user_id, user_uuid, nama, email, role, password) VALUES ($2, $1, 'test', 'test', 'test','akjfkjdfkskdf')", uuid.New(), 2)

	val, err := FindById(2)
	t.Log(val)
	require.NoError(t, err)
}

func TestFindByUuid_Fail(t *testing.T) {
	val, err := FindByUuid(uuid.New().String())
	require.Error(t, err)
	require.Equal(t, person.UserWithoutPassword{}, val)
}

func TestFindByUuid_Success(t *testing.T) {
	_db := db.DB
	defer func() {
		_db.Exec("DELETE FROM users WHERE user_id = 3")
		db.DB = _db
	}()

	uid := uuid.New()
	_db.Exec("INSERT INTO users (user_id, user_uuid, nama, email, role, password) VALUES (3, $1, 'test', 'test', 'test','akjfkjdfkskdf')", uid)

	val, err := FindByUuid(uid.String())
	require.NoError(t, err)
	require.Equal(t, uid, val.UserUUID)
}

func TestFindByName_Fail(t *testing.T) {
	val, err := FindByName("test")
	require.NoError(t, err)
	require.Equal(t, []person.UserWithoutPassword(nil), val)
}

func TestFindByName_Success(t *testing.T) {
	_db := db.DB
	defer func() {
		_db.Exec("DELETE FROM users WHERE user_id = 4")
		db.DB = _db
	}()

	uid := uuid.New()
	_db.Exec("INSERT INTO users (user_id, user_uuid, nama, email, role, password) VALUES (4, $1, 'test', 'test', 'test','akjfkjdfkskdf')", uid)

	val, err := FindByName("test")
	require.NoError(t, err)
	require.Equal(t, "test", val[0].Nama)
}

func TestFindByEmail_Fail(t *testing.T) {
	val, err := FindByEmail("test")
	require.NoError(t, err)
	require.Equal(t, []person.UserWithoutPassword(nil), val)
}

func TestFindByEmail_Success(t *testing.T) {
	_db := db.DB
	defer func() {
		_db.Exec("DELETE FROM users WHERE user_id = 5")
		db.DB = _db
	}()

	uid := uuid.New()
	_db.Exec("INSERT INTO users (user_id, user_uuid, nama, email, role, password) VALUES (5, $1, 'test', 'test', 'test','akjfkjdfkskdf')", uid)

	val, err := FindByEmail("test")
	require.NoError(t, err)
	require.Equal(t, "test", val[0].Email)
}

func TestFindByRole_Fail(t *testing.T) {
	val, err := FindByRole("test")
	require.NoError(t, err)
	require.Equal(t, []person.UserWithoutPassword(nil), val)
}

func TestFindByRole_Success(t *testing.T) {
	_db := db.DB
	defer func() {
		_db.Exec("DELETE FROM users WHERE user_id = 6")
		db.DB = _db
	}()

	uid := uuid.New()
	_db.Exec("INSERT INTO users (user_id, user_uuid, nama, email, role, password) VALUES (6, $1, 'test', 'test', 'test','akjfkjdfkskdf')", uid)

	val, err := FindByRole("test")
	require.NoError(t, err)
	require.Equal(t, "test", val[0].Role)
}

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

	val, err := FindAll()
	require.Error(t, err)
	require.Equal(t, []person.UserWithoutPassword(nil), val)
}
