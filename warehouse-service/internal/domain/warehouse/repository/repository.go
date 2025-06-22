package repository

import (
	"context"

	"gorm.io/gorm"
)

type Repository interface{
	AddWarehouse(ctx context.Context) error
	SetWarehouseStatus(ctx context.Context, IdWarehouse int64) error
	GetAll(ctx context.Context, dest interface{}) error
}

type WarehouseRepository struct{
	db *gorm.DB
}

func NewRepositoryWarehouse(g *gorm.DB) WarehouseRepository {
	return WarehouseRepository{
		db: g,
	}
}