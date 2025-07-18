package usecase

import (
	"context"
	"order-service/internal/domain/order/repository"
	rabbitmq "order-service/package/rabbit-mq"
	"order-service/package/structs"
	"sync"

	http_client "order-service/package/http_client"

	"github.com/hibiken/asynq"
)

type IOrder interface{
	GetAll(ctx context.Context) ([]*structs.Order, error)
	CreateOrder(ctx context.Context, order *structs.RequestCreateOrder) error
}

type OrderUsecase struct{
	repo         repository.OrderRepository
	workerClient *asynq.Client
	http_clients http_client.HTTPClients
	mu           sync.Mutex
	mqClient     *rabbitmq.RabbitMQClient
}


func NewOrderUsecase(repo repository.OrderRepository, http_clients http_client.HTTPClients, workerClient *asynq.Client, mqClient *rabbitmq.RabbitMQClient) IOrder  {
	return &OrderUsecase{
		repo: repo,
		http_clients : http_clients,
		workerClient: workerClient,
		mqClient: mqClient,
	}
}
