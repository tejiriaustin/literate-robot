package service

import (
	"context"
	"user-service/model"
)

type UserService struct{}

var _ UserServiceInterface = (*UserService)(nil)

func NewUserService() UserServiceInterface {
	return &UserService{}
}

func (u *UserService) CreateUser(ctx context.Context, user *model.User) error {
	return nil
}

func (u *UserService) GetUserDetails(ctx context.Context) (*model.User, error) {
	return nil, nil
}

func (u *UserService) UpdateUserDetails(ctx context.Context, user *model.User) error {
	return nil
}

func (u *UserService) DeleteUser(ctx context.Context, user *model.User) error {
	return nil
}
