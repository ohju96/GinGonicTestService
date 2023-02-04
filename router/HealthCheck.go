package router

import (
	"github.com/gin-gonic/gin"
	"user_service/controller"
)

func HealthCheckRouter(r *gin.Engine) {
	router := r.Group("/")
	{
		router.GET("/", controller.HealthCheck)
	}
}
