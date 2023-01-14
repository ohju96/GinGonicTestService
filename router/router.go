package router

import "github.com/gin-gonic/gin"

func InitRouters(r *gin.Engine) {

	UserRouter(r)

}
