package router

import "github.com/gin-gonic/gin"

func UserRouter(r *gin.Engine) {
	users := r.Group("/api/users")
	{
		users.POST("/")   // 유저 생성
		users.GET("/")    // 유저 조회
		users.PUT("/")    // 유저 수정
		users.DELETE("/") // 유저 삭제
	}
}
