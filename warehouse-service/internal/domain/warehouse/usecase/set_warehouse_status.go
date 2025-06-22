package usecase

import (
	"context"
)

func (u *UsecaseWarehouse) SetWarehouseStatus(ctx context.Context, IdWarehouse int64) error{
	return u.repo.SetWarehouseStatus(ctx, IdWarehouse)
}