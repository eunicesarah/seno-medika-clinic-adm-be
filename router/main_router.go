package router

import (
	"github.com/gin-gonic/gin"
	"seno-medika.com/router/antrian"
	"seno-medika.com/router/pasien"
	"seno-medika.com/router/user"
)

func MainRouter(r *gin.Engine) {
	user.UserRouter(r)
	antrian.AntrianRouter(r)
	pasien.PasienRouter(r)
}
