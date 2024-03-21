package ttv

import (
	"github.com/gin-gonic/gin"
	"seno-medika.com/controller/ttv"
)

func TTVRouter(r *gin.Engine) {
	r.GET("/ttv", ttv.FindTTV)
	r.POST("/ttv", ttv.AddTTV)
	r.PATCH("/ttv/:id", ttv.PatchTTV)
	r.DELETE("/ttv/:id", ttv.DeleteTTV)
}
