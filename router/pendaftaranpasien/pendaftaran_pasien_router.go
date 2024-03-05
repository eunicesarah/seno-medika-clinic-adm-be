package pendaftaranpasien

import (
	"github.com/gin-gonic/gin"
	"seno-medika.com/controller/pendaftaranpasien"
)

func PendaftaranPasienRouter(r *gin.Engine) {
	r.POST("/pendaftaranpasien", pendaftaranpasien.AddPasien)
	r.GET("/pendaftaranpasien", pendaftaranpasien.GetPasien)
	r.PUT("/pendaftaranpasien", pendaftaranpasien.UpdatePasien)
	r.DELETE("/pendaftaranpasien", pendaftaranpasien.DeletePasien)
}
