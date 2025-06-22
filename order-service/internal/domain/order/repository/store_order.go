package repository

import (
	"context"
	"order-service/package/structs"
)

func (r OrderRepository) CreateOrder(ctx context.Context, order structs.Order) (orderID int64, err error) {
	err = r.db.Create(&order).Error
	if err != nil {
		return 0, err
	}
	return order.ID, nil
}