package router

import "github.com/gin-gonic/gin"

func InitRouter(r *gin.Engine) {
	UserRouter(r) // 유저 라우터
}
