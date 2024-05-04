package apotek

import (
	"github.com/gin-gonic/gin"
	"seno-medika.com/controller/role"
)

func ApotekRouter(r *gin.Engine) {
	r.GET("/apotek", role.GetApotek)
}
