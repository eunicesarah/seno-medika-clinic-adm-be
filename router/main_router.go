package router

import (
	"github.com/gin-gonic/gin"
	"seno-medika.com/router/antrian"
	"seno-medika.com/router/login"
	"seno-medika.com/router/pasien"
	"seno-medika.com/router/role"
	"seno-medika.com/router/ttv"
	"seno-medika.com/router/user"
	"seno-medika.com/router/obat"
	"seno-medika.com/router/kasir"
)

func MainRouter(r *gin.Engine) {
	user.UserRouter(r)
	login.LoginRouter(r)
	antrian.AntrianRouter(r)
	pasien.PasienRouter(r)
	role.PerawatRouter(r)
	role.DokterRouter(r)
	role.ApotekerRouter(r)
	ttv.TTVRouter(r)
	obat.ObatRouter(r)
	kasir.KasirRouter(r)
}
