package repository

import (
	"context"
	"shop-service/package/structs"

	"gorm.io/gorm"
)

type Repository interface{
	GetAll(ctx context.Context) error
	Store(ctx context.Context, payload structs.Shop) error
}

type ShopRepository struct{
	db *gorm.DB
}

func (r ShopRepository) GetAll(ctx context.Context, dest interface{}) error {
	return r.db.Find(dest).Error
}

func NewShopRepository(g *gorm.DB) ShopRepository {
	return ShopRepository{
		db: g,
	}
}