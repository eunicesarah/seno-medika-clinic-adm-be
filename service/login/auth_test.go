package login

import (
	"errors"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
	"seno-medika.com/config/db"
	"seno-medika.com/model/person"
)

func TestVerifyPassword_Fail(t *testing.T) {
	err := VerifyPassword("test", "test")
	require.Error(t, err)
}

func TestVerifyPassword_Success(t *testing.T) {
	password := "test"
	//make hashed with bcrypt
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	err := VerifyPassword("test", string(hashedPassword))
	require.NoError(t, err)
}

func TestGenerateToken_Success(t *testing.T) {
	user := person.User{
		UserID:   1,
		UserUUID: uuid.New(),
		Nama:     "test",
		Email:    "test",
		Role:     "test",
	}
	_, err := GenerateToken(user)
	require.NoError(t, err)
}

type MockJwtToken struct {
	mock.Mock
}

func (m *MockJwtToken) SignedString(claims jwt.Claims) (string, error) {
	args := m.Called(claims)
	return args.String(0), args.Error(1)
}

func GenerateTokenWithCustomSigning(user person.User, signingFunc func(claims jwt.Claims) (string, error)) (string, error) {
	claims := jwt.MapClaims{
		"user_id":   user.UserID,
		"user_uuid": user.UserUUID,
		"nama":      user.Nama,
		"email":     user.Email,
		"role":      user.Role,
		"exp":       time.Now().Add(time.Hour * 24).Unix(), // expiring in 24 hours
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := signingFunc(token.Claims)
	if err != nil {
		return "", fmt.Errorf("failed to generate token")
	}

	return tokenString, nil
}

func TestGenerateToken_Failure(t *testing.T) {
	originalSecret := os.Getenv("JWT_SECRET")
	os.Setenv("JWT_SECRET", "")
	defer func() {
		os.Setenv("JWT_SECRET", originalSecret)
	}()

	mockToken := new(MockJwtToken)
	mockToken.On("SignedString", mock.Anything).Return("", errors.New("mocked error"))

	user := person.User{
		UserID:   1,
		UserUUID: uuid.New(),
		Nama:     "John Doe",
		Email:    "john@example.com",
		Role:     "admin",
	}

	tokenString, err := GenerateTokenWithCustomSigning(user, mockToken.SignedString)

	assert.Empty(t, tokenString)
	assert.Error(t, err)

	mockToken.AssertExpectations(t)
}

func TestLoginCheck_Success(t *testing.T) {
	_db := db.DB
	defer func() {
		_db = db.Conn()
		db.DB = _db
	}()
	uid := uuid.New()
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("akjfkjdfkskdf"), bcrypt.DefaultCost)
	_db.Exec("INSERT INTO users (user_id, user_uuid, nama, email, role, password) VALUES (1, $1, 'test', 'test', 'test', $2)", uid, string(hashedPassword))
	fmt.Println("uid", uid)
	token, err := LoginCheck("test", "akjfkjdfkskdf")
	require.NoError(t, err)
	require.NotEmpty(t, token)
}

func TestLoginCheck_Fail(t *testing.T) {
	_db := db.DB
	defer func() {
		_db = db.Conn()
		db.DB = _db
	}()
	token, err := LoginCheck("test", "123")
	require.Error(t, err)
	require.Empty(t, token)
}
