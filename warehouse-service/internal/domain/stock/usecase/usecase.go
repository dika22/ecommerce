package usecase

import (
	"context"
	"warehouse-service/internal/domain/stock/repository"
	"warehouse-service/package/structs"
)

type IStock interface{
	AddStock(ctx context.Context, req *structs.RequestAddStock) error
	TransferStockProduct(ctx context.Context, req *structs.RequestTransferStockProduct) error
	GetAvailableStock(ctx context.Context, productID string) (structs.ResponseTotalStock, error)
	ReserveStock(ctx context.Context, req *structs.RequestReserveStock) error
	ReleaseStock(ctx context.Context, req *structs.RequestReleaseStock) error
}

type UsecaseStock struct{
	repo repository.StokRepository
}


func NewUsecaseStock(repo repository.StokRepository) IStock  {
	return &UsecaseStock{
		repo: repo,
	}
}
