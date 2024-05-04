package router

import (
	"github.com/gin-gonic/gin"
	"seno-medika.com/router/antrian"
	"seno-medika.com/router/apotek"
	"seno-medika.com/router/login"
	"seno-medika.com/router/obat"
	"seno-medika.com/router/pasien"
	"seno-medika.com/router/role"
	"seno-medika.com/router/ttv"
)

func MainRouter(r *gin.Engine) {
	role.UserRouter(r)
	login.LoginRouter(r)
	antrian.AntrianRouter(r)
	pasien.PasienRouter(r)
	role.PerawatRouter(r)
	role.DokterRouter(r)
	role.ApotekerRouter(r)
	ttv.TTVRouter(r)
	obat.ObatRouter(r)
	role.KasirRouter(r)
	apotek.ApotekRouter(r)
}
