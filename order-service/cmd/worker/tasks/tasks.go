package tasks

import (
	"context"
	"order-service/internal/domain/order/repository"
	"order-service/package/config"

	"github.com/hibiken/asynq"

	http_client "order-service/package/http_client"
)

const TypeReleaseStock = "order:release_stock"

type AsyncTask struct{
	conf *config.Config
	repo repository.OrderRepository
	workerClient *asynq.Client
	http_clients http_client.HTTPClients
}

type Tasks interface {
	StartStockReleaseJob(ctx context.Context, task *asynq.Task) error
}

func NewAsynqTask(conf *config.Config, repo repository.OrderRepository,  http_clients http_client.HTTPClients, workClient *asynq.Client) Tasks {
	return AsyncTask{
		conf: conf,
		repo: repo,
		workerClient: workClient,
		http_clients: http_clients,
	}
}