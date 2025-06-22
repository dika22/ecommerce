package repository

import (
	"context"
	"warehouse-service/package/structs"
)

func (r StokRepository) ReserveStock(ctx context.Context, reverseStock structs.ReservedStock) error{
	return r.db.Create(&reverseStock).Error
}