package router

import (
	"github.com/gin-gonic/gin"
	"user_service/config"
	"user_service/controller"
	"user_service/repository"
	"user_service/service"
)

func UserRouter(r *gin.Engine) {
	router := r.Group("/api/user")
	{
		r := repository.NewRepository(config.DB)
		s := service.NewService(r)
		c := controller.NewController(s)

		router.POST("/", c.CreateUser)
		/*	router.GET("/", c.GetUser)
			router.PUT("/", c.UpdateUser)
			router.DELETE("/", controller.DeleteUser)*/
	}
}
