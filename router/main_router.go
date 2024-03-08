package router

import (
	"github.com/gin-gonic/gin"
	"seno-medika.com/router/login"
	"seno-medika.com/router/user"
)

func MainRouter(r *gin.Engine) {
	user.UserRouter(r)
	login.LoginRouter(r)
}
