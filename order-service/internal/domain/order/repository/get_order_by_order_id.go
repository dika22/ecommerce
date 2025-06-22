package repository

import "context"

func (r OrderRepository) GetOrderByOrderId(ctx context.Context, orderID int64, dest interface{}) error {
	return r.db.Where("id = ?", orderID).Find(&dest).Error 
}
