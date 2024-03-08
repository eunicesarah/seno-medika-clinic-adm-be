package pasien

import (
	"github.com/gin-gonic/gin"
	"seno-medika.com/controller/pasien"
)

func PasienRouter(r *gin.Engine) {
	r.POST("/pasien", pasien.AddPasien)
	r.GET("/pasien", pasien.GetPasien)
	r.PUT("/pasien", pasien.UpdatePasien)
	r.DELETE("/pasien", pasien.DeletePasien)
}
