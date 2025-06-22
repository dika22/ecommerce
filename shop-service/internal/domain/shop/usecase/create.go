package usecase

import (
	"context"
	"shop-service/package/structs"
)

func (u *ShopeUsecase) Create(ctx context.Context, req *structs.RequestCreateShop) error {
	shop := req.NewShop()
	return u.repo.Store(ctx, shop)
}