package usecase

import (
	"context"
	"product-service/internal/domain/category-product/repository"
	"product-service/package/structs"
)

type CategoryProductUsecase struct {
	repo repository.CategoryProductRepository
}

type ICategoryProduct interface {
	GetAll(ctx context.Context) ([]*structs.CategoryProduct, error)
	Create(ctx context.Context, req *structs.RequestCreateCategoryProduct) error
	Update(ctx context.Context, req *structs.RequestUpdateCategoryProduct) error
	Delete(ctx context.Context, id int64) error
}

func NewCategoryProductUsecase() ICategoryProduct {
	return &CategoryProductUsecase{}
	
}