package service

import (
	"context"
	"user_service/dto"
	"user_service/repository"
)

type UserService interface {
	CreateUser(ctx context.Context, request dto.CreateUserReq) (*dto.CreateUserRes, error)
}

type Service struct {
	repository repository.UserRepository
}

func NewService(repository repository.UserRepository) UserService {
	return &Service{repository: repository}
}

func (s *Service) CreateUser(ctx context.Context, request dto.CreateUserReq) (*dto.CreateUserRes, error) {

	user, _ := s.repository.CreateUser(ctx, request)

	res := dto.CreateUserRes{
		UserID:   user.ID,
		UserName: user.UserName,
	}
	return &res, nil
}
