package router

import (
	"github.com/gin-gonic/gin"
	ctl "sample/controller"
)

func UserRouter(r *gin.Engine) {
	user := r.Group("/api/user")
	{
		user.POST("/", ctl.CreateUser)
		user.POST("/login", ctl.LoginUser)
		user.GET("/", ctl.GetUser)
		user.PUT("/", ctl.UpdateUser)
		user.DELETE("/", ctl.DeleteUser)
	}
}
