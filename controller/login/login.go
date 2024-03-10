package login

import (
	"net/http"
	"seno-medika.com/model/common"
	"github.com/gin-gonic/gin"
	"seno-medika.com/service/login"
)

type loginInput struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Login(r *gin.Context) {
	var userInput loginInput

	if err := r.ShouldBindJSON(&userInput); err != nil {
		r.JSON(http.StatusBadRequest, common.Response{
			Message:     "Bad Request",
			Status:      "Bad Request",
			StatusCode: http.StatusBadRequest,
			Data:        nil,
		})
		return
	}

	token, err := login.LoginCheck(userInput.Email, userInput.Password)

	if err != nil {
		r.JSON(http.StatusUnauthorized, common.Response{
			Message:     "Unauthorized",
			Status:      "Unauthorized",
			StatusCode: http.StatusUnauthorized,
			Data:        nil,
		})
		return
	}

	// Set HttpOnly to true for enhanced security
	r.SetSameSite(http.SameSiteLaxMode)
	r.SetCookie("token", token, 3600*24, "", "", false, true)

	r.JSON(http.StatusOK, gin.H{"token": token})
}

func Validate(r *gin.Context) {
	user, _ := r.Get("user")
	r.JSON(http.StatusOK, gin.H{"message": user})
}