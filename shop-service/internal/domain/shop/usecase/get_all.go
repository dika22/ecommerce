package usecase

import (
	"context"
	"shop-service/package/structs"
)

func (u *ShopeUsecase) GetAll(ctx context.Context) ([]*structs.Shop, error) {
	dest := []*structs.Shop{}
	if err := u.repo.GetAll(ctx, &dest); err != nil {return nil, err}
	return dest, nil
}