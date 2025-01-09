package service

import (
	"context"
	"user-service/model"
)

type UserServiceInterface interface {
	CreateUser(ctx context.Context, user *model.User) error
	GetUserDetails(ctx context.Context) (*model.User, error)
	UpdateUserDetails(ctx context.Context, user *model.User) error
	DeleteUser(ctx context.Context, user *model.User) error
}
