package service

import (
	"context"
	"order-service/model"
)

type OrderServiceInterface interface {
	CreateOrder(ctx context.Context, input CreateOrderInput) (*model.Order, error)
	DeleteOrder(ctx context.Context, input DeleteOrderInput) error
	UpdateOrder(ctx context.Context, input UpdateOrderInput) (*model.Order, error)
	ListOrders(ctx context.Context, input ListOrderInput) ([]*model.Order, error)
	GetOrder(ctx context.Context, input GetOrderInput) (*model.Order, error)
}
