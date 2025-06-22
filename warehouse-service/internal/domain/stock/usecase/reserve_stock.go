package usecase

import (
	"context"
	"errors"
	"warehouse-service/package/structs"
)


func (u *UsecaseStock) ReserveStock(ctx context.Context, req *structs.RequestReserveStock) error{
	totalStock, err := u.repo.GetAvailableStockByProductIdWarehouseId(ctx, req.ProductID, req.WarehouseID)
	if err != nil {
		return err
	}
	if totalStock < req.Quantity {
		return errors.New("Insufficient stock in source warehouse")
	}
	reverseStock := req.NewReservedStock()
	return u.repo.ReserveStock(ctx, reverseStock)
}