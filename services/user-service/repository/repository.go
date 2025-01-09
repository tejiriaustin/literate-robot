package repository

import (
	"context"

	"github.com/tejiriaustin/literate-robot/core/database"
	"github.com/tejiriaustin/literate-robot/core/repository"

	userModel "user-service/model"
)

type (
	UserServiceRepository struct {
		UserRepo        *repository.Repository[userModel.User]
		UserProfileRepo *repository.Repository[userModel.UserProfile]
	}

	UserRepositoryInterface[T userModel.User] interface {
		Create(ctx context.Context, data T) (T, error)
		FindOne(ctx context.Context, queryFilter *repository.Query) (T, error)
		FindMany(ctx context.Context, queryFilter *repository.Query) ([]T, error)
		DeleteMany(ctx context.Context, queryFilter *repository.Query) error
		Update(ctx context.Context, dataObject T) (T, error)
		Select(ctx context.Context, target interface{}, query string, args ...interface{}) error
	}
)

func NewUserServiceRepository(dbClient *database.Client) *UserServiceRepository {
	return &UserServiceRepository{
		UserRepo:        repository.NewRepository[userModel.User](dbClient),
		UserProfileRepo: repository.NewRepository[userModel.UserProfile](dbClient),
	}
}
