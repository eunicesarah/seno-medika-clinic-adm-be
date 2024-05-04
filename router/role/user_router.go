package role

import (
	"github.com/gin-gonic/gin"
	"seno-medika.com/controller/role"
)

func UserRouter(r *gin.Engine) {

	r.GET("/user", role.GetUser)
	r.POST("/user", role.AddUser)
	r.PUT("/user", role.PutUser)
	r.PATCH("/user", role.PatchUser)
	r.DELETE("/user", role.DeleteUser)
}
