package usecase

import (
	"context"
	"warehouse-service/package/structs"
)


func (u *UsecaseWarehouse) AddWarehouse(ctx context.Context, req *structs.RequestAddWarehouse) error  {
	warehouse := req.NewWarehouse()
	if err := u.repo.AddWarehouse(ctx, warehouse); err != nil {
		return err
	}
	return nil
}