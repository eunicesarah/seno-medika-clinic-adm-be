package antrianpasien

import (
	"github.com/gin-gonic/gin"
	"seno-medika.com/controller/antrian"
)

func AntrianPasienRouter(r *gin.Engine) {
	r.POST("/antrianpasien", antrian.AddAntrianOffline)
	r.POST("/antrianpasienonline", antrian.AddAntrianOnline)
	r.GET("/antrianpasien", antrian.GetAllAntrian)
	r.DELETE("/antrianpasien", antrian.DeleteAntrian)
}
