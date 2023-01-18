package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"user-service/ent"
	"user-service/model"
)

func CreateUser(r *gin.Context, db *ent.Client) (*ent.User, error) {

	var request model.CreateUserRequest

	if err := r.ShouldBindJSON(&request); err != nil {
		r.JSON(http.StatusBadRequest, gin.H{"ERROR": "Binding error !!"})
	}

	user, err := db.User.Create().SetName(request.Name).SetAge(request.Age).SetCode(request.Code).Save(r)

	if err != nil {
		return nil, fmt.Errorf("유저 생성 실패: %w", err)
	}

	return user, nil

}

func ReadUser(r *gin.Context, db *ent.Client) (*ent.User, error) {
	user, err := db.User.Get(r, 1)

	if err != nil {
		return nil, fmt.Errorf("유저가 없습니다. : %w", err)
	}

	return user, nil

}

func UpdateUser(r *gin.Context, db *ent.Client) {

}

func DeleteUser(r *gin.Context, db *ent.Client) {

}
