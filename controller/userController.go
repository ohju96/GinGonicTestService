package controller

import (
	"github.com/gin-gonic/gin"
	"user-service/config"
	"user-service/service"
)

func CreateUser(r *gin.Context) {
	service.CreateUser(r, config.DB)
}

func ReadUser(r *gin.Context) {
	service.ReadUser(r, config.DB)

}

func UpdateUser(r *gin.Context) {
	service.UpdateUser(r, config.DB)
}

func DeleteUser(r *gin.Context) {
	service.DeleteUser(r, config.DB)
}
