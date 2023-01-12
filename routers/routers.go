package routers

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func InitRouter(router *gin.Engine) {

	fmt.Println("UserRouter 그룹 실행")
	UserRouter(router)
}
