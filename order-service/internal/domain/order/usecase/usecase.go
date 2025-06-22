package usecase

import (
	"context"
	"order-service/internal/domain/order/repository"
	"order-service/package/structs"

	"github.com/hibiken/asynq"
)

type IOrder interface{
	GetAll(ctx context.Context) ([]*structs.Order, error)
	CreateOrder(ctx context.Context, order *structs.RequestCreateOrder) error
}

type OrderUsecase struct{
	repo repository.OrderRepository
	workerClient *asynq.Client
}


func NewOrderUsecase(repo repository.OrderRepository, workerClient *asynq.Client) IOrder  {
	return &OrderUsecase{
		repo: repo,
		workerClient: workerClient,
	}
}
