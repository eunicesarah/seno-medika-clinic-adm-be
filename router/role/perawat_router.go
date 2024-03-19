package role

import (
	"github.com/gin-gonic/gin"
	"seno-medika.com/controller/role"
)

func PerawatRouter(r *gin.Engine) {
	r.POST("/perawat", role.AddPerawat)
	r.GET("/perawat", role.GetPerawat)
	r.PATCH("/perawat", role.PatchPerawat)
	r.DELETE("/perawat", role.DeletePerawat)
}
