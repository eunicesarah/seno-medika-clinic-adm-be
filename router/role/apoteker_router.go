package role

import (
	"github.com/gin-gonic/gin"
	"seno-medika.com/controller/role"
)

func ApotekerRouter(r *gin.Engine) {
	r.POST("/apoteker", role.AddApoteker)
	r.GET("/apoteker", role.GetApoteker)
	r.PATCH("/apoteker", role.PatchApoteker)
	r.DELETE("/apoteker", role.DeleteApoteker)
}
