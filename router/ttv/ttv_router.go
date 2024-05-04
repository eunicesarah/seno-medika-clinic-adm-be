package ttv

import (
	"github.com/gin-gonic/gin"
	"seno-medika.com/controller/pemeriksaan"
)

func TTVRouter(r *gin.Engine) {
	r.GET("/ttv", pemeriksaan.FindTTV)
	r.POST("/ttv", pemeriksaan.AddTTV)
	r.PATCH("/ttv", pemeriksaan.PatchTTV)
	r.DELETE("/ttv", pemeriksaan.DeleteTTV)
}
