package repository

import "gorm.io/gorm"


type SellerRepository struct {
	db *gorm.DB
}

func NewSellerRepository(g *gorm.DB) SellerRepository {
	return SellerRepository{
		db: g,
	}
}