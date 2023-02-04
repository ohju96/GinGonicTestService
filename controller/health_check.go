package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func HealthCheck(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "health check successfully !!",
		"name":    "이지혜가 세상에서 제일 이쁘다 !",
	})
	fmt.Println("health check successfully !!")
}
