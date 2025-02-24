package role

import (
	"github.com/gin-gonic/gin"
	"seno-medika.com/controller/role"
)

func DokterRouter(r *gin.Engine) {
	r.POST("/dokter", role.AddDokter)
	r.GET("/dokter", role.GetDokter)
	r.PATCH("/dokter", role.PatchDokter)
	r.DELETE("/dokter", role.DeleteDokter)
	r.POST("/dokter/tindakan", role.AddTindakan)
	r.GET("/dokter/tindakan", role.GetTindakan)
	r.PUT("/dokter/tindakan", role.PutTindakan)
	r.DELETE("/dokter/tindakan", role.DeleteTindakan)
}
