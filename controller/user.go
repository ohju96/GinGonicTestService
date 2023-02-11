package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"user_service/dto"
	"user_service/service"
)

type UserController interface {
	CreateUser(g *gin.Context)
}

type Controller struct {
	service service.UserService
}

func NewController(s service.UserService) UserController {
	return &Controller{service: s}
}

func (controller *Controller) CreateUser(g *gin.Context) {

	ctx := g.Request.Context()

	var req dto.CreateUserReq
	if err := g.Bind(&req); err != nil {
		g.JSON(http.StatusBadRequest, "err")
		return
	}

	response, _ := controller.service.CreateUser(ctx, req)

	g.JSON(http.StatusOK, response)
	return
}

func GetUser(*gin.Context) {

}

func UpdateUser(*gin.Context) {

}

func DeleteUser(*gin.Context) {

}
