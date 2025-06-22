package usecase

import (
	"context"
	"warehouse-service/package/structs"
)

func (u *UsecaseWarehouse) GetAll(ctx context.Context) ([]*structs.Warehouse, error) {
	dest := []*structs.Warehouse{}
	if err := u.repo.GetAll(ctx, &dest); err != nil{
		return nil, err
	}
	return dest, nil
}