package usecase

import (
	"context"
	"shop-service/internal/domain/shop/repository"
	"shop-service/package/structs"
)

type IShop interface{
	GetAll(ctx context.Context) ([]*structs.Shop, error)
	Create(ctx context.Context, req *structs.RequestCreateShop) error
}

type ShopeUsecase struct{
	repo repository.ShopRepository
}


func NewShopUsecase(repo repository.ShopRepository) IShop  {
	return &ShopeUsecase{
		repo: repo,
	}
}
