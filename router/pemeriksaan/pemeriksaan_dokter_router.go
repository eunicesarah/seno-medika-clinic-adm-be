package pemeriksaan

import (
	"github.com/gin-gonic/gin"
	"seno-medika.com/controller/pemeriksaan"
)

func PemeriksaanDokterRouter(r *gin.Engine) {
	r.POST("/pemeriksaan_dokter", pemeriksaan.AddPemeriksaanDokter)
	r.GET("/pemeriksaan_dokter", pemeriksaan.GetPemeriksaanDokter)
	r.DELETE("/pemeriksaan_dokter", pemeriksaan.DeletePemeriksaanDokter)
	r.PATCH("/pemeriksaan_dokter", pemeriksaan.PatchPemeriksaanDokter)
}
