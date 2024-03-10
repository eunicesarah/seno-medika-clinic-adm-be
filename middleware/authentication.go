package middleware

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"seno-medika.com/config/db"
	"seno-medika.com/model/person"
)

func Authentication(r *gin.Context) {
	tokenString, err := r.Cookie("token")

	if err != nil {
		r.AbortWithStatus(http.StatusUnauthorized)
	}

	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("JWT SECRET")), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// check the expiration time
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			r.AbortWithStatus(http.StatusUnauthorized)
		}

		var user person.User
		row := db.DB.QueryRow("SELECT user_id, user_uuid, nama, password, email, role FROM users WHERE user_uuid = $1", claims["user_uuid"]).Scan(&user.UserID, &user.UserUUID, &user.Nama, &user.Password, &user.Email, &user.Role)
		if row != nil {
			r.AbortWithStatus(http.StatusUnauthorized)
		}

		r.Set("user", user)
		r.Next()
	} else {
		r.AbortWithStatus(http.StatusUnauthorized)
	}

}