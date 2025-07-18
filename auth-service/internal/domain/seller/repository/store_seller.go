package repository

import (
	"auth-service/package/structs"
	"context"
)


func (r SellerRepository) StoreSeller(ctx context.Context, payload structs.Seller) error {
	return  r.db.Table("sellers").Create(&payload).Error
}