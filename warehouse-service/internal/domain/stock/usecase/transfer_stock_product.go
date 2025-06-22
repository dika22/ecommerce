package usecase

import (
	"context"
	"warehouse-service/package/structs"
)

func (u *UsecaseStock) TransferStockProduct(ctx context.Context, req *structs.RequestTransferStockProduct) error  {
	return u.repo.TransferStockProduct(ctx, req)
}