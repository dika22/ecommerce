package repository

import "context"

func (r OrderRepository) GetOrderItemsByOrderId(ctx context.Context, orderID int64, dest interface{}) error {
	return r.db.Table("order_items").Where("order_id = ?", orderID).Find(dest).Error 
}