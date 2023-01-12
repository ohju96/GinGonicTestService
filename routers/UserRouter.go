package routers

import (
	"github.com/gin-gonic/gin"
	ctl "sample/controllers"
)

type User struct {
	Name string
	Age  rune
}

func UserRouter(router *gin.Engine) {
	//user라는 라우터 그룹 생성
	user := router.Group("/api/user")
	{
		user.POST("/create", ctl.CreateUser) //user 라우터 그룹에 Post 형식으로 /create api 추가
	}
}
