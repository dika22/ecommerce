package repository

import "gorm.io/gorm"


type CategoryProductRepository struct{
	db *gorm.DB
}

func NewCategoryProductRepository(g *gorm.DB) CategoryProductRepository {
	return CategoryProductRepository{
		db: g,
	}
}