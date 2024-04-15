package obat

import (
	"github.com/gin-gonic/gin"
	"seno-medika.com/controller/obat"
)

func ObatRouter(r *gin.Engine) {
	r.POST("/obat", obat.AddObat)
	r.GET("/obat", obat.GetObat)
	r.PUT("/obat", obat.UpdateObat)
	r.DELETE("/obat", obat.DeleteObat)
}
