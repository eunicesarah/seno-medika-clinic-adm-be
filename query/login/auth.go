package login

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"seno-medika.com/config/db"
	"seno-medika.com/model/person"
)

func VerifyPassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func GenerateToken(user person.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":   user.UserID,
		"user_uuid": user.UserUUID,
		"nama":      user.Nama,
		"email":     user.Email,
		"role":      user.Role,
		"exp":       time.Now().Add(time.Hour * 24).Unix(), // expiring in 24 hours
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", fmt.Errorf("failed to generate token")

	}

	return tokenString, nil
}

func LoginCheck(email, password string) (string, error) {
	var user person.User

	row := db.DB.QueryRow("SELECT user_id, user_uuid, nama, password, email, role FROM users WHERE email = $1", email).Scan(&user.UserID, &user.UserUUID, &user.Nama, &user.Password, &user.Email, &user.Role)
	if row != nil {
		return "", fmt.Errorf("invalid email or password")
	}

	err := VerifyPassword(password, user.Password)
	if err != nil {
		return "", fmt.Errorf("invalid email or password")
	}

	token, err := GenerateToken(user)

	if err != nil {
		return "", err
	}

	return token, nil
}
