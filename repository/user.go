package repository

import (
	"context"
	"user_service/dto"
	"user_service/ent"
)

type UserRepository interface {
	CreateUser(ctx context.Context, request dto.CreateUserReq) (*ent.User, error)
}

type repository struct {
	client *ent.Client
}

func NewRepository(client *ent.Client) UserRepository {
	return &repository{client: client}
}

func (r repository) CreateUser(ctx context.Context, request dto.CreateUserReq) (*ent.User, error) {
	return r.client.User.Create().
		SetUID(request.Uid).
		SetUserName(request.UserName).
		SetUserType("USER").
		SetAge(request.Age).
		SetEmail(request.Email).
		SetPassword(request.Password).
		Save(ctx)
}
