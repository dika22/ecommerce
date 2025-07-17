package usecase

import (
	"context"
	"product-service/package/structs"
)

func (u *CategoryProductUsecase) GetAll(ctx context.Context) ([]*structs.CategoryProduct, error) {
	dest := []*structs.CategoryProduct{}
	if err := u.repo.GetAll(ctx, &dest); err != nil {return nil, err}
	return dest, nil
}