package usecase

import (
	"context"
	"product-service/package/structs"
)

func (u *CategoryProductUsecase) Create(ctx context.Context, req *structs.RequestCreateCategoryProduct) error{
	categoryProduct := req.NewCategoryProduct()
	if err := u.repo.Store(ctx, categoryProduct); err != nil {
		return err
	}
	return  nil
}