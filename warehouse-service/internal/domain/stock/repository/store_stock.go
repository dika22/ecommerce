package repository

import (
	"context"
	"warehouse-service/package/structs"
)

func (r StokRepository) StoreStock(ctx context.Context, stock *structs.Stock) error{
	return r.db.Create(&stock).Error
}