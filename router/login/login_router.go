package login

import (
	"github.com/gin-gonic/gin"
	"seno-medika.com/controller/login"
	"seno-medika.com/middleware"
)

func LoginRouter(r *gin.Engine) {
	r.POST("/login", login.Login)
	r.GET("/validate", middleware.Authentication, login.Validate)
}
