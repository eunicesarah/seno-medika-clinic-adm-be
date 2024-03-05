package user

import (
	"github.com/gin-gonic/gin"
	"seno-medika.com/controller/superadmin"
)

func UserRouter(r *gin.Engine) {
	r.GET("/user", superadmin.GetUser)
	r.POST("/user", superadmin.AddUser)
	r.PATCH("/user", superadmin.UpdateUser)
	r.PUT("/user", superadmin.PutUser)
	r.DELETE("/user", superadmin.DeleteUser)
}
