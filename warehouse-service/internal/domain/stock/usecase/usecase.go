package usecase

import (
	"context"
	"warehouse-service/internal/domain/stock/repository"
	rabbitmq "warehouse-service/package/rabbit-mq"
	"warehouse-service/package/structs"
)

type IStock interface{
	AddStock(ctx context.Context, req *structs.RequestAddStock) error
	TransferStockProduct(ctx context.Context, req *structs.RequestTransferStockProduct) error
	GetAvailableStock(ctx context.Context, productID string) (structs.ResponseTotalStock, error)
	ReserveStock(ctx context.Context, req *structs.RequestReserveStock) error
	ReleaseStock(ctx context.Context, req *structs.RequestReleaseStock) error
	BatchStock(ctx context.Context, req *structs.RequestBatchStock) (structs.ResponseBatchStock, error)
}

type UsecaseStock struct{
	repo repository.StokRepository
	mqClient *rabbitmq.RabbitMQClient
}


func NewUsecaseStock(repo repository.StokRepository, 	mqClient *rabbitmq.RabbitMQClient) IStock  {
	return &UsecaseStock{
		repo: repo,
		mqClient: mqClient,
	}
}
