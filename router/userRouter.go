package router

import (
	"github.com/gin-gonic/gin"
	"sample/controller"
)

func UserRouter(r *gin.Engine) {
	user := r.Group("/api/user")
	{
		user.POST("/create", controller.CreateUser)
	}
}
