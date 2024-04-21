package kasir

import (
	"github.com/gin-gonic/gin"
	"seno-medika.com/controller/kasir"
)

func KasirRouter(r *gin.Engine) {
	r.POST("/kasir", kasir.AddNota)
	r.GET("/kasir", kasir.GetNota)
	// r.DELETE("/kasir", kasir.DeleteKasir)
	// r.PATCH("/kasir", kasir.PatchKasir)
}