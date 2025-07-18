package usecase

import (
	"context"
	"warehouse-service/internal/domain/warehouse/repository"
	rabbitmq "warehouse-service/package/rabbit-mq"
	"warehouse-service/package/structs"
)

type IWarehouse interface{
	AddWarehouse(ctx context.Context, req *structs.RequestAddWarehouse) error
	SetWarehouseStatus(ctx context.Context, IdWarehouse int64) error
	GetAll(ctx context.Context) ([]*structs.Warehouse, error)
}

type UsecaseWarehouse struct{
	repo repository.WarehouseRepository
	mqClient *rabbitmq.RabbitMQClient
}


func NewUsecaseWarehouse(repo repository.WarehouseRepository, mqClient *rabbitmq.RabbitMQClient) IWarehouse  {
	return &UsecaseWarehouse{
		repo: repo,
		mqClient: mqClient,
	}
}
