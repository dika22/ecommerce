package repository

import (
	"context"
	"shop-service/package/structs"
)

func (r ShopRepository) Store(ctx context.Context, payload structs.Shop) error  {
	return r.db.Create(&payload).Error
}