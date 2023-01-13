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

func ReadUser(router *gin.Context) {

	// Param 으로 받은 id 값 바인딩
	id := router.Param("id")

	// service로 값을 보내어 로직 처리
	serviceImpl := impl.UserServiceImpl{}
	readUser, err := serviceImpl.ReadUser(id)

	// error 처리
	if err != nil {
		router.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
	}

	// Json 구조로 담아 리턴 진행
	router.JSON(http.StatusOK, readUser)
	return
}

func UpdateUser(router *gin.Context) {

	// Param 으로 받은 id 값 바인딩
	id := router.Param("id")

	// service로 값을 보내어 로직 처리
	serviceImpl := impl.UserServiceImpl{}
	updateUser, err := serviceImpl.UpdateUser(id)

	// error 처리
	if err != nil {
		router.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
	}

	// Json 구조로 담아 리턴 진행
	router.JSON(http.StatusOK, updateUser)
	return
}

func DeleteUser(router *gin.Context) {

	// Param 으로 받은 id 값 바인딩
	id := router.Param("id")

	// service로 값을 보내어 로직 처리
	serviceImpl := impl.UserServiceImpl{}
	err := serviceImpl.DeleteUser(id)

	// error 처리
	if err != nil {
		router.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
	}

	// Json 구조로 담아 리턴 진행
	router.JSON(http.StatusOK, gin.H{"status": "success"})
	return
}
