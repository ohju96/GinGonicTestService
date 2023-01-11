package routers

import "github.com/gin-gonic/gin"

type MyInfo struct {
	Name string
	Age  rune
}

func PingTest(r *gin.Context) MyInfo {

	myInfo := MyInfo{"오주현", 28}

	return myInfo

}
