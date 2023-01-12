package services

import "sample/models"

type UserService interface {

	// GetUsers 유저 목록 가져오기
	GetUsers() []models.User

	// GetUserById 유저 아이디로 유저 정보 가져오기
	GetUserById(id int) (models.User, error)

	// CreateUser 유저 생성
	CreateUser(user models.User) (models.User, error)

	// UpdateUser 유저 정보 업데이트
	UpdateUser(user models.User) error

	// DeleteUser 유저 정보 삭제
	DeleteUser(id int) error
}
