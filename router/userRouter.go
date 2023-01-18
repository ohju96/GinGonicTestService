package router

import (
	"github.com/gin-gonic/gin"
	ctl "user-service/controller"
)

func UserRouters(r *gin.Engine) {
	user := r.Group("/api/user")
	{
		user.POST("/create", ctl.CreateUser)
		user.GET("/read", ctl.ReadUser)
		user.PUT("/update", ctl.UpdateUser)
		user.DELETE("/delete", ctl.DeleteUser)
	}
}
