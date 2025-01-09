package service

import (
	"context"
	"order-service/model"
)

type (
	OrderService struct{}

	CreateOrderInput struct{}
	DeleteOrderInput struct{}
	UpdateOrderInput struct{}
	ListOrderInput   struct{}
	GetOrderInput    struct{}
)

var _ OrderServiceInterface = (*OrderService)(nil)

func NewOrderService() *OrderService {
	return &OrderService{}
}

func (o *OrderService) CreateOrder(ctx context.Context, input CreateOrderInput) (*model.Order, error) {
	return nil, nil
}

func (o *OrderService) DeleteOrder(ctx context.Context, input DeleteOrderInput) error {
	return nil
}

func (o *OrderService) UpdateOrder(ctx context.Context, input UpdateOrderInput) (*model.Order, error) {
	return nil, nil
}

func (o *OrderService) ListOrders(ctx context.Context, input ListOrderInput) ([]*model.Order, error) {
	return nil, nil
}

func (o *OrderService) GetOrder(ctx context.Context, input GetOrderInput) (*model.Order, error) {
	return nil, nil
}
