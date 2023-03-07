package service

import "sample/app/repository"

type UserService interface {
}

type userService struct {
	userRepository repository.UserRepository
}
