package role

import (
	"github.com/gin-gonic/gin"
	"seno-medika.com/controller/role"
)

func KasirRouter(r *gin.Engine) {
	r.POST("/kasir", role.AddNota)
	r.GET("/kasir", role.GetNota)
	// r.DELETE("/kasir", kasir.DeleteKasir)
	r.PATCH("/kasir", role.PatchNota)
}
