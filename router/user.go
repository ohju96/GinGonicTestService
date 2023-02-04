package router

import (
	"github.com/gin-gonic/gin"
	"user_service/controller"
)

func UserRouter(r *gin.Engine) {
	router := r.Group("/api/user")
	{
		router.POST("/", controller.CreateUser)
		router.GET("/", controller.GetUser)
		router.PUT("/", controller.UpdateUser)
		router.DELETE("/", controller.DeleteUser)
	}
}
