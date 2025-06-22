package repository

import (
	"context"
	"order-service/package/structs"
)

func (r OrderRepository) UpdateOrderStatus(ctx context.Context, orderID int64) error {
	return r.db.Model(&structs.Order{}).Where("id = ?", orderID).Preload("orders").UpdateColumn("status", 3).Error 
}
