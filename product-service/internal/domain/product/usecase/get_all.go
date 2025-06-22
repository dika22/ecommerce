package usecase

import (
	"context"
	"product-service/package/structs"
)

func (u *ProductUsecase) GetAll(ctx context.Context) ([]*structs.Product, error) {
	dest := []*structs.Product{}
	if err := u.repo.GetAll(ctx, &dest); err != nil {return nil, err}
	return dest, nil
}