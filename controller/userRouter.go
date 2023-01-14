package controller

import (
	"github.com/gin-gonic/gin"
	"sample/config"
	"sample/service"
)

func CreateUser(r *gin.Context) {
	service.CreateUser(r, config.DB)
}
