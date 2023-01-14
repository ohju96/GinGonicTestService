package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"sample/ent"
	"sample/model"
)

func CreateUser(r *gin.Context, client *ent.Client) (*ent.User, error) {

	var request model.User

	if err := r.ShouldBindJSON(&request); err != nil {
		r.JSON(http.StatusBadRequest, gin.H{
			"ERROR": "Binding error",
		})
	}

	fmt.Println(request.Age)
	fmt.Println(request.Password)
	fmt.Println(request.Name)
	fmt.Println(request.Email)

	user, err := client.User.Create().
		SetName(request.Name).
		SetAge(request.Age).
		SetEmail(request.Email).
		SetPassword(request.Password).
		Save(r.Request.Context())

	if err != nil {
		return nil, fmt.Errorf("failed creating user: %w", err)
	}

	return user, nil
}
