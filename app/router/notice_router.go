package router

import "github.com/gin-gonic/gin"

func NoticeRouter(r *gin.Engine) {
	notice := r.Group("/api/notices")
	{
		notice.POST("/")   // 공지 생성
		notice.GET("/")    // 공지 조회
		notice.PUT("/")    // 공지 수정
		notice.DELETE("/") // 공지 삭제
	}
}
