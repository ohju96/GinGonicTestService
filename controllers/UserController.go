package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sample/models"
	"sample/services/impl"
)

func CreateUser(router *gin.Context) {

	// user 구조체 가져와서 요청값을 바인딩 시킨다.
	var user models.User
	router.BindJSON(&user)

	// 구조체에 바인딩 된 값을 Service로 보내 서비스 로직 처리
	serviceImpl := impl.UserServiceImpl{}
	createUser, err := serviceImpl.CreateUser(&user, nil)

	// create에서 err발생 시 에러 처리
	if err != nil {
		router.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	// 정상 반환된 값을 처리 및 리턴
	router.JSON(http.StatusOK, createUser)
	return
}
