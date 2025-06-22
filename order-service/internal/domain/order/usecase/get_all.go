package usecase

import (
	"context"
	"order-service/package/structs"
)

func (u *OrderUsecase) GetAll(ctx context.Context) ([]*structs.Order, error) {
	dest := []*structs.Order{}
	if err := u.repo.GetAll(ctx, &dest); err != nil {return nil, err}
	return dest, nil
}