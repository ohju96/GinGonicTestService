package impl

import (
	"fmt"
	"sample/models"
	"strconv"
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

func (u *UserServiceImpl) ReadUser(id string) (*models.User, error) {

	// Param으로 요청한 id 값을 int로 형변환
	requestId, _ := strconv.Atoi(id)

	// user 구조체를 활용해 user 생성
	user := models.User{"Ohju", requestId}
	return &user, nil
}

func (u *UserServiceImpl) UpdateUser(id string) (*models.User, error) {

	requestId, _ := strconv.Atoi(id)

	user := models.User{"ohju", 0}
	user.ID = requestId

	return &user, nil
}

func (u *UserServiceImpl) DeleteUser(id string) error {

	fmt.Printf("####### \n [회원 삭제] \n 삭제 할 아이디 : %s \n ###### \n", id)

	return nil
}
