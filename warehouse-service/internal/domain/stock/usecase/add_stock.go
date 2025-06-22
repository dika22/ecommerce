package usecase

import (
	"context"
	"warehouse-service/package/structs"
)


func (u *UsecaseStock) AddStock(ctx context.Context, req *structs.RequestAddStock) error  {
	stock := req.NewStock()
	if err := u.repo.StoreStock(ctx, &stock); err != nil {
		return err
	}
	return nil
}