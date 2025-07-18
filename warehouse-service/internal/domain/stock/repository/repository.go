package repository

import (
	"context"
	"warehouse-service/package/structs"

	"gorm.io/gorm"
)

type Repository interface{
	StoreStock(ctx context.Context) error
	GetAvailableStockByProductId(ctx context.Context, productID, totalStock int64) error 
	ReserveStock(ctx context.Context, req *structs.RequestReserveStock) error
}

type StokRepository struct{
	db *gorm.DB
}

func NewRepositoryStock(g *gorm.DB) StokRepository {
	return StokRepository{
		db: g,
	}
}