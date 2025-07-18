package repository

import (
	"auth-service/package/structs"
	"context"
)

func (r SellerRepository) UpdateSeller(ctx context.Context, payload structs.Seller) error {
	return r.db.Table("sellers").Where("id = ?", payload.ID).Updates(&payload).Error
}