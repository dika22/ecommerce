package repository

import (
	"context"
	"order-service/package/structs"
)

func (r OrderRepository) StoreOrderItem(ctx context.Context, orderItem *structs.OrderItem) error{
	return r.db.Create(orderItem).Error
}

func (r OrderRepository) StoreOrderItems(ctx context.Context, orderItems []structs.OrderItem) error {
	return r.db.Create(&orderItems).Error
}