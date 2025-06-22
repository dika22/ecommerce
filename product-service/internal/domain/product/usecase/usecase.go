package usecase

import (
	"context"
	"product-service/internal/domain/product/repository"
	"product-service/package/structs"
)

type IProduct interface{
	GetAll(ctx context.Context) ([]*structs.Product, error)
}

type ProductUsecase struct{
	repo repository.ProductRepository
}


func NewProductUsecase(repo repository.ProductRepository) IProduct  {
	return &ProductUsecase{
		repo: repo,
	}
}
