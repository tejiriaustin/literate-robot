package repository

import (
	"context"

	"github.com/tejiriaustin/literate-robot/core/database"
	"github.com/tejiriaustin/literate-robot/core/repository"

	orderModel "order-service/model"
)

type (
	OrderServiceRepository struct {
		OrderRepo *repository.Repository[orderModel.Order]
	}

	OrderRepositoryInterface[T orderModel.Order] interface {
		Create(ctx context.Context, data T) (T, error)
		FindOne(ctx context.Context, queryFilter *repository.Query) (T, error)
		FindMany(ctx context.Context, queryFilter *repository.Query) ([]T, error)
		DeleteMany(ctx context.Context, queryFilter *repository.Query) error
		Update(ctx context.Context, dataObject T) (T, error)
		Select(ctx context.Context, target interface{}, query string, args ...interface{}) error
	}
)

func NewOrderServiceRepository(dbClient *database.Client) *OrderServiceRepository {
	return &OrderServiceRepository{
		OrderRepo: repository.NewRepository[orderModel.Order](dbClient),
	}
}
