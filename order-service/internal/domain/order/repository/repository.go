package repository

import (
	"context"
	"order-service/package/structs"

	"gorm.io/gorm"
)

type Repository interface{
	GetAll(ctx context.Context) error
	Create(ctx context.Context, order *structs.Order) error
	GetOrderByOrderId(ctx context.Context, orderID int64, dest interface{}) error
	GetOrderItemsByOrderId(ctx context.Context, orderID int64, dest interface{}) error
}

type OrderRepository struct{
	db *gorm.DB
}

func (r OrderRepository) GetAll(ctx context.Context, dest interface{}) error {
	return r.db.Find(dest).Error
}

func NewOrderRepository(g *gorm.DB) OrderRepository {
	return OrderRepository{
		db: g,
	}
}