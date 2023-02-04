package router

import (
	"github.com/gin-gonic/gin"
	"user_service/config"
)

func InitRouter(r *gin.Engine, c *config.Config) {

	// Init HealthCheck Router
	HealthCheckRouter(r)

	// Init User Routers
	UserRouter(r)
}
