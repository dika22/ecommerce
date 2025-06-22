package repository

import (
	"context"

	"gorm.io/gorm"
)

type Repository interface{
	GetAll(ctx context.Context) error
}

type ProductRepository struct{
	db *gorm.DB
}

func (r ProductRepository) GetAll(ctx context.Context, dest interface{}) error {
	return r.db.Find(dest).Error
}

func NewProductRepository(g *gorm.DB) ProductRepository {
	return ProductRepository{
		db: g,
	}
}