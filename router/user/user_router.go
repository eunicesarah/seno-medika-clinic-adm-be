package user

import (
	"github.com/gin-gonic/gin"
	"seno-medika.com/controller/superadmin"
)

func UserRouter(r *gin.Engine) {
	r.GET("/user", superadmin.GetUser)
	r.POST("/user", superadmin.AddUser)
	r.PUT("/user", superadmin.PutUser)
	r.PATCH("/user", superadmin.PatchUser)
	r.DELETE("/user", superadmin.DeleteUser)
}
