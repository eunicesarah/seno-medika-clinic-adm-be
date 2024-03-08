package antrian

import (
	"github.com/gin-gonic/gin"
	"seno-medika.com/controller/antrian"
)

func AntrianRouter(r *gin.Engine) {
	r.POST("/antrian", antrian.AddAntrian)
	r.GET("/antrian", antrian.GetAntrian)
	r.DELETE("/antrian", antrian.DeleteAntrian)
}
