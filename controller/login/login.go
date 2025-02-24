package login

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"seno-medika.com/model/common"
	"seno-medika.com/query/login"
)

type loginInput struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Login(r *gin.Context) {
	var userInput loginInput

	if err := r.ShouldBindJSON(&userInput); err != nil {
		r.JSON(http.StatusBadRequest, common.Response{
			Message:    err.Error(),
			Status:     "Bad Request",
			StatusCode: http.StatusBadRequest,
			Data:       nil,
		})
		return
	}

	token, err := login.LoginCheck(userInput.Email, userInput.Password)

	if err != nil {
		r.JSON(http.StatusBadRequest, gin.H{"error": "username or password is incorrect."})
		return
	}

	r.SetSameSite(http.SameSiteLaxMode)
	r.SetCookie("token", token, 3600*24, "", "", false, true)

	r.JSON(http.StatusOK, gin.H{"token": token})
}

func Validate(r *gin.Context) {
	user, _ := r.Get("user")
	r.JSON(http.StatusOK, gin.H{"message": user})
}
