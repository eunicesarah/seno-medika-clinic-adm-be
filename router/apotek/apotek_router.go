package apotek

import (
	"github.com/gin-gonic/gin"
	"seno-medika.com/controller/apotek"
)

func ApotekRouter(r *gin.Engine) {
	r.GET("/apotek", apotek.GetApotek)
}