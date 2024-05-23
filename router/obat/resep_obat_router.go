package obat

import (
	"github.com/gin-gonic/gin"
	"seno-medika.com/controller/obat"
)

func ResepRouter(r *gin.Engine) {
	r.POST("/resep", obat.AddResep)
	r.GET("/resep", obat.GetResep)
	r.PUT("/resep", obat.PutResep)
	r.DELETE("/resep", obat.DeleteResep)
}
