package impl

import (
	"fmt"
	"sample/models"
)

type UserServiceImpl struct {
	user models.User
}

// 유저 생성
func (s *UserServiceImpl) CreateUser(user *models.User, err error) (*models.User, error) {

	fmt.Printf("############## \n [회원 가입] \n Name : %s \n ID : %d \n ##############\n®", user.Name, user.ID)

	if err != nil {
		fmt.Errorf("오류 : %s", err)
	}

	return user, err
}
