package main

import (
	"fmt"
	gin "github.com/gin-gonic/gin"
	"sample/routers"
)

func main() {

	// 진 라우터 생성
	router := gin.Default()

	// 정적 페이지 템플릿 경로 지정
	//router.LoadHTMLGlob("templates/**")

	// 라우터 묶음 적용
	fmt.Println("INIT 라우터 실행")
	routers.InitRouter(router)

	// 서버 실행
	router.Run()
}
